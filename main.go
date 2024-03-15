package main

import (
	handler "th/GoDoIt/handlers"
	"th/GoDoIt/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	storage := storage.New()
	requestsHandler := handler.New(storage)

	e.GET("/todo", requestsHandler.GetAllTodos)
	e.POST("/todo/add", requestsHandler.AddTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
