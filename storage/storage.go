package storage

import (
	"sync"

	"th/GoDoIt/models"
)

type TodoStorage struct {
	mutex sync.RWMutex
	items map[string]*models.Todo
}

func NewTodoStorage() *TodoStorage {
	return &TodoStorage{
		items: make(map[string]*models.Todo),
	}
}

func (storage *TodoStorage) getAll() []*models.Todo {
	storage.mutex.RLock()
	defer storage.mutex.RUnlock()

	result := make([]*models.Todo, len(storage.items))

	for _, todo := range storage.items {
		result = append(result, todo)
	}

	return result
}

func (storage *TodoStorage) getById(id string) *models.Todo {
	storage.mutex.RLock()
	defer storage.mutex.RUnlock()

	return storage.items[id]
}

func (storage *TodoStorage) add(newTodo *models.Todo) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	storage.items[newTodo.ID] = newTodo
}

func (storage *TodoStorage) delete(id string) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	delete(storage.items, id)
}
