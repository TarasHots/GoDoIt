package storage

import (
	"sync"

	"th/GoDoIt/models"
)

type TodoStorage struct {
	mutex sync.RWMutex
	items map[string]*models.Todo
}

func New() *TodoStorage {
	return &TodoStorage{
		items: make(map[string]*models.Todo),
	}
}

func (storage *TodoStorage) GetAll() []*models.Todo {
	storage.mutex.RLock()
	defer storage.mutex.RUnlock()

	result := make([]*models.Todo, 0)

	for _, todo := range storage.items {
		result = append(result, todo)
	}

	return result
}

func (storage *TodoStorage) GetById(id string) *models.Todo {
	storage.mutex.RLock()
	defer storage.mutex.RUnlock()

	return storage.items[id]
}

func (storage *TodoStorage) Add(newTodo *models.Todo) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	storage.items[newTodo.ID] = newTodo
}

func (storage *TodoStorage) Delete(id string) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	delete(storage.items, id)
}
