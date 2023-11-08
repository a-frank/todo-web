package todos

import (
	"database/sql"
	"fmt"
)

var todoState = map[uint32]Todo{
	1: {Id: 1, Todo: "Something", Done: false},
	2: {Id: 2, Todo: "Nothing", Done: true},
}

var DbConnection *sql.DB

func GetTodos() ([]Todo, error) {
	rows, err := DbConnection.Query("SELECT * FROM Todo ORDER BY id")

	if err != nil {
		fmt.Printf("Can't get todos, %s", err.Error())
		return nil, err
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Todo, &todo.Done)
		if err != nil {
			fmt.Printf("Can't scan a single todo, %s", err.Error())
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func AddNewTodo(todoText string) (*Todo, error) {
	var id int
	err := DbConnection.QueryRow("INSERT INTO Todo (todo, done) VALUES($1, $2) RETURNING id", todoText, false).Scan(&id)
	if err != nil {
		fmt.Printf("Can't insert todo into DB: %s", err.Error())
		return nil, err
	}

	todo := Todo{
		Id:   uint32(id),
		Todo: todoText,
		Done: false,
	}
	return &todo, nil
}

func ToggleDone(todoId uint32) (*Todo, error) {
	row := DbConnection.QueryRow("SELECT * FROM Todo Where id = $1", todoId)
	var todo Todo
	err := row.Scan(&todo.Id, &todo.Todo, &todo.Done)
	if err != nil {
		fmt.Printf("Can't get done state of todo %d: %s", todoId, err.Error())
		return nil, err
	}

	todo.Done = !todo.Done

	_, err = DbConnection.Exec("UPDATE Todo SET done = $1 Where id = $2", todo.Done, todoId)
	if err != nil {
		fmt.Printf("Can't update done state of todo %d: %s", todoId, err.Error())
		return nil, err
	}

	return &todo, nil
}

func DeleteTodo(todoId uint32) ([]Todo, error) {
	_, err := DbConnection.Exec("DELETE FROM Todo WHERE id = $1", todoId)
	if err != nil {
		fmt.Printf("Can't delete todo %d: %s", todoId, err.Error())
		return nil, err
	}
	return GetTodos()
}
