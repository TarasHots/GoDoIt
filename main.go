package main

import (
	handler "th/GoDoIt/handlers"
	"th/GoDoIt/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	storage := storage.NewTodoStorage()
	requestsHandler := handler.NewTodoHandler(storage)

	e.GET("/todo", requestsHandler.getAllTodos)
	e.GET("/todo/add", requestsHandler.addTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
