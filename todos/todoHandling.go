package todos

var todoState = []Todo{
	{Id: 1, Todo: "Something", Done: false},
	{Id: 2, Todo: "Nothing", Done: true},
}

func GetTodos() []Todo {
	return todoState
}

func AddNewTodo(todoText string) Todo {
	numTodos := len(todoState)

	todo := Todo{
		Id:   uint32(numTodos) + 1,
		Todo: todoText,
		Done: false,
	}
	todoState = append(todoState, todo)
	return todo
}
