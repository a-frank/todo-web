package test

import (
	"github.com/a-frank/web-dev/todos"
	"golang.org/x/exp/maps"
)

type MemStore struct {
	Todos map[uint32]todos.Todo
}

func NewMemStore() *MemStore {
	return &MemStore{
		Todos: make(map[uint32]todos.Todo),
	}
}

func (t *MemStore) GetAll() ([]todos.Todo, error) {
	return maps.Values(t.Todos), nil
}

func (t *MemStore) Add(todoText string) (*todos.Todo, error) {
	count := len(maps.Keys(t.Todos))
	id := uint32(count + 1)
	todo := todos.Todo{
		Id:   id,
		Todo: todoText,
		Done: false,
	}
	t.Todos[id] = todo
	return &todo, nil
}
func (t *MemStore) Delete(todoId uint32) ([]todos.Todo, error) {
	maps.DeleteFunc(t.Todos, func(id uint32, todo todos.Todo) bool {
		return id == todoId
	})
	return t.GetAll()
}
func (t *MemStore) ToggleDone(todoId uint32) (*todos.Todo, error) {
	todo := t.Todos[todoId]
	todo.Done = !todo.Done
	t.Todos[todoId] = todo
	return &todo, nil
}
func (t *MemStore) Close() {}
