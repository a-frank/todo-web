package test

import (
	"errors"
	"github.com/a-frank/web-dev/todos"
)

type MemStore struct {
	todos []todos.Todo
}

func (t *MemStore) GetAll() ([]todos.Todo, error) {
	return t.todos, nil
}

func (t *MemStore) Add(todoText string) (*todos.Todo, error) {
	count := len(t.todos)
	todo := todos.Todo{
		Id:   uint32(count + 1),
		Todo: todoText,
		Done: false,
	}
	t.todos = append(t.todos, todo)
	return &todo, errors.New("not supported")
}
func (t *MemStore) Delete(todoId uint32) ([]todos.Todo, error) {
	return nil, errors.New("not supported")
}
func (t *MemStore) ToggleDone(todoId uint32) (*todos.Todo, error) {
	return nil, errors.New("not supported")
}
func (t *MemStore) Close() {}
