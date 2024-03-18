package main

import (
	"net/http"

	handler "th/GoDoIt/handlers"
	"th/GoDoIt/storage"
	validator "th/GoDoIt/validators"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Validator = validator.New()

	e.Use(middleware.Recover())

	storage := storage.New()
	requestsHandler := handler.New(storage)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	e.GET("/todo", requestsHandler.GetAllTodos)
	e.GET("/todo/:id", requestsHandler.GetById)
	e.POST("/todo/add", requestsHandler.AddOrUpdateTodo)
	e.DELETE("/todo/:id", requestsHandler.Remove)

	e.Logger.Fatal(e.Start(":8080"))
}
