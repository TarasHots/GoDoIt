package handler

import (
	"encoding/json"
	"net/http"

	"th/GoDoIt/models"
	"th/GoDoIt/storage"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	store *storage.TodoStorage
}

func New(storage *storage.TodoStorage) *TodoHandler {
	return &TodoHandler{store: storage}
}

func (handler *TodoHandler) GetAllTodos(c echo.Context) error {
	todos := handler.store.GetAll()

	return c.JSON(http.StatusOK, todos)
}

func (handler *TodoHandler) GetById(c echo.Context) error {
	id := c.QueryParam("id")

	return c.JSON(http.StatusOK, json.NewEncoder(c.Response().Writer).Encode(handler.store.GetById(id)))
}

func (handler *TodoHandler) AddOrUpdateTodo(c echo.Context) error {
	var newTodo models.Todo

	if err := json.NewDecoder(c.Request().Body).Decode(&newTodo); err != nil {
		return c.JSON(http.StatusBadRequest, newTodo)
	}

	handler.store.Add(&newTodo)
	return c.JSON(http.StatusCreated, newTodo)
}
