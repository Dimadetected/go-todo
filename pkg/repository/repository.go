package repository

import (
	"github.com/Dimadetected/todo-app"
	"github.com/jmoiron/sqlx"
)

//docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres  команда для запуска базы в докер

//migrate create -ext sql -dir ./schema -seq init - создать миграцию
//migrate -path ./schema -database postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable up - миграции запустить

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}
type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, id int) (todo.TodoList, error)
	Delete(userId, id int) error
	Update(userId, id int, input todo.UpdateListInput) error
}
type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, id int) error
	Update(userId, id int, input todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
