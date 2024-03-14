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

func NewTodoHandler(storage *storage.TodoStorage) *TodoHandler {
	return &TodoHandler{store: storage}
}

func (handler *TodoHandler) getAllTodos(c echo.Context) error {
	todos := handler.store.getAll()

	return c.JSON(http.StatusOK, todos)
}

func (handler *TodoHandler) addTodo(c echo.Context) error {
	var newTodo models.Todo

	if err := json.NewDecoder(c.Request().Body).Decode(&newTodo); err != nil {
		return c.JSON(http.StatusBadRequest, newTodo)
	}

	handler.store.add(newTodo)
	return c.JSON(http.StatusCreated, newTodo)
}
