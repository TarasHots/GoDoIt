package handler

import (
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
	if todo := handler.store.GetById(c.Param("id")); todo != nil {
		return c.JSON(http.StatusOK, todo)
	}

	return c.JSON(http.StatusNotFound, nil)
}

func (handler *TodoHandler) AddOrUpdateTodo(c echo.Context) error {
	newTodo := new(models.Todo)

	if err := c.Bind(newTodo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(newTodo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	handler.store.Add(newTodo)
	return c.JSON(http.StatusCreated, newTodo)
}

func (handler *TodoHandler) Remove(c echo.Context) error {
	if todo := handler.store.GetById(c.Param("id")); todo != nil {
		handler.store.Delete(todo.ID)

		return c.JSON(http.StatusNoContent, nil)
	}

	return c.JSON(http.StatusNotFound, nil)
}
