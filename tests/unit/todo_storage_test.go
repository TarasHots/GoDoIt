package storage

import (
	"testing"
	"time"

	"th/GoDoIt/models"
	"th/GoDoIt/storage"

	"github.com/stretchr/testify/assert"
)

func TestEmptyStorage(t *testing.T) {
	storage := storage.New()

	assert.Empty(t, storage.GetAll(), "Newly created storage should be empty")
}

func TestAddingElement(t *testing.T) {
	storage := storage.New()

	var exampleTodo models.Todo
	exampleTodo.ID = "1"
	exampleTodo.Title = "Test"
	exampleTodo.Description = "Test description"
	exampleTodo.DueDate = time.Now()

	storage.Add(&exampleTodo)

	assert.NotEmpty(t, storage.GetAll(), "Storage must not be empty")
	assert.Equal(t, 1, len(storage.GetAll()), "Storage must have exactly one element")
}

func TestCanGetElementById(t *testing.T) {
	storage := storage.New()

	exampleId := "1"

	var exampleTodo models.Todo
	exampleTodo.ID = exampleId
	exampleTodo.Title = "Test"
	exampleTodo.Description = "Test description"
	exampleTodo.DueDate = time.Now()

	storage.Add(&exampleTodo)

	assert.NotEmpty(t, storage.GetAll(), "Storage must not be empty")
	assert.Equal(t, &exampleTodo, storage.GetById(exampleId), "Element is different")
}

func TestCanUpdateElement(t *testing.T) {
	storage := storage.New()

	exampleId := "1"

	var beforeChanges models.Todo
	beforeChanges.ID = exampleId
	beforeChanges.Title = "Test"
	beforeChanges.Description = "Test description"
	beforeChanges.DueDate = time.Now()

	storage.Add(&beforeChanges)
	assert.NotEmpty(t, storage.GetAll(), "Storage must not be empty")

	var afterChanges models.Todo
	afterChanges.ID = exampleId
	afterChanges.Title = "Some new title"
	afterChanges.Description = "Completely new desc"
	afterChanges.DueDate = time.Now().Add(time.Duration(time.Duration.Hours(2)))

	storage.Add(&afterChanges)

	assert.Equal(t, 1, len(storage.GetAll()), "Storage must have exactly one element")
	assert.Equal(t, &afterChanges, storage.GetById(exampleId), "Updated element should have all new data")
	assert.NotEqual(t, beforeChanges, afterChanges)
}

func TestCanDeleteElement(t *testing.T) {
	storage := storage.New()

	exampleId := "1"

	var beforeChanges models.Todo
	beforeChanges.ID = exampleId
	beforeChanges.Title = "Test"
	beforeChanges.Description = "Test description"
	beforeChanges.DueDate = time.Now()

	storage.Add(&beforeChanges)
	assert.NotEmpty(t, storage.GetAll(), "Storage must not be empty")

	storage.Delete(exampleId)

	assert.Empty(t, storage.GetAll(), "Storage should be empty if all elements were removed")
}
