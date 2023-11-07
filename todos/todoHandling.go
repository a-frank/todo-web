package todos

import (
	"fmt"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var todoState = map[uint32]Todo{
	1: {Id: 1, Todo: "Something", Done: false},
	2: {Id: 2, Todo: "Nothing", Done: true},
}

func GetTodos() []Todo {
	values := maps.Values(todoState)
	slices.SortStableFunc(values, func(a, b Todo) int {
		return int(a.Id) - int(b.Id)
	})
	return values
}

func AddNewTodo(todoText string) Todo {
	numTodos := len(todoState)

	id := uint32(numTodos) + 1
	todo := Todo{
		Id:   id,
		Todo: todoText,
		Done: false,
	}
	todoState[id] = todo
	return todo
}

func ToggleDone(todoId uint32) (Todo, error) {
	todo, found := todoState[todoId]
	if !found {
		return Todo{}, fmt.Errorf("todo %d not found", todoId)
	}

	todo.Done = !todo.Done
	todoState[todoId] = todo

	return todo, nil
}
