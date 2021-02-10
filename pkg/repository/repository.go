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
}
type TodoList interface {
}
type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
