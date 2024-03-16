package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	handler "th/GoDoIt/handlers"
	"th/GoDoIt/models"
	"th/GoDoIt/storage"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestEmptyTodoList(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/todo", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)

	storage := storage.New()
	handler := handler.New(storage)

	expectedResponseBody := "[]"

	if assert.NoError(t, handler.GetAllTodos(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, expectedResponseBody, strings.TrimSuffix(recorder.Body.String(), "\n"))
	}
}

func TestCanGetById(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/todo/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	var exampleTodoItem models.Todo
	exampleTodoItem.ID = "1"
	exampleTodoItem.Title = "test"
	exampleTodoItem.Description = "test desc"
	exampleTodoItem.DueDate = time.Now()

	expectedJson, _ := json.Marshal(exampleTodoItem)

	storage := storage.New()
	storage.Add(&exampleTodoItem)
	handler := handler.New(storage)

	if assert.NoError(t, handler.GetById(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, string(expectedJson), strings.TrimSuffix(recorder.Body.String(), "\n"))
	}
}
