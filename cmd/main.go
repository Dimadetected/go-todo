package main

import (
	"github.com/Dimadetected/todo-app"
	"github.com/Dimadetected/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occuered while running http server: %s", err.Error())
	}
}
