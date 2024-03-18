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
	validator "th/GoDoIt/validators"

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

func TestElementNotFound(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/todo/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	storage := storage.New()
	handler := handler.New(storage)

	if assert.NoError(t, handler.GetById(context)) {
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	}
}

func TestAddTodoItem(t *testing.T) {
	exampleJson := `{"id":"111","title":"test","description":"example description","due_date":"0001-01-01T00:00:00Z"}`

	e := echo.New()
	e.Validator = validator.New()

	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(exampleJson))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/todo")

	storage := storage.New()
	handler := handler.New(storage)

	assert.Empty(t, storage.GetAll())

	if assert.NoError(t, handler.AddOrUpdateTodo(context)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
		assert.Equal(t, exampleJson, strings.TrimSuffix(recorder.Body.String(), "\n"))
	}
}

func TestTodoItemCanBeUpdated(t *testing.T) {
	updatedData := `{"id":"222","title":"test","description":"example description","due_date":"0001-01-01T00:00:00Z"}`

	e := echo.New()
	e.Validator = validator.New()

	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(updatedData))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/todo")

	storage := storage.New()

	var existingItem models.Todo
	existingItem.ID = "222"
	existingItem.Title = "new title"
	existingItem.Description = "new desc"
	existingItem.DueDate = time.Now()

	storage.Add(&existingItem)

	handler := handler.New(storage)

	assert.Equal(t, 1, len(storage.GetAll()))

	if assert.NoError(t, handler.AddOrUpdateTodo(context)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
		assert.Equal(t, updatedData, strings.TrimSuffix(recorder.Body.String(), "\n"))
	}
}

func TestItemCanBeDeleted(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodDelete, "/", nil)

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

	storage := storage.New()
	storage.Add(&exampleTodoItem)
	handler := handler.New(storage)

	if assert.NoError(t, handler.Remove(context)) {
		assert.Equal(t, http.StatusNoContent, recorder.Code)
		assert.Equal(t, "null", strings.TrimSuffix(recorder.Body.String(), "\n"))
	}
}

func TestValidation(t *testing.T) {
	exampleJsonWithInvalidData := `{"id":"invalid id here","title":"test","description":"example description","due_date":"0001-01-01T00:00:00Z"}`
	expectedError := `{"message":"Key: 'Todo.ID' Error:Field validation for 'ID' failed on the 'number' tag"}`

	e := echo.New()
	e.Validator = validator.New()

	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(exampleJsonWithInvalidData))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/todo")

	storage := storage.New()
	handler := handler.New(storage)

	assert.Empty(t, storage.GetAll())

	if assert.NoError(t, handler.AddOrUpdateTodo(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Equal(t, expectedError, strings.TrimSuffix(recorder.Body.String(), "\n"))
	}
}
