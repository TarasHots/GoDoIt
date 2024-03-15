package main

import (
	"net/http"

	handler "th/GoDoIt/handlers"
	"th/GoDoIt/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	storage := storage.New()
	requestsHandler := handler.New(storage)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{}{})
	})

	e.GET("/todo", requestsHandler.GetAllTodos)
	e.GET("/todo/:id", requestsHandler.GetById)
	e.POST("/todo/add", requestsHandler.AddOrUpdateTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
