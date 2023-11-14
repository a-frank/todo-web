package todos

import (
	"database/sql"
	"fmt"
)

type TodoStore interface {
	GetAll() ([]Todo, error)
	Add(todoText string) (*Todo, error)
	Delete(todoId uint32) ([]Todo, error)
	ToggleDone(todoId uint32) (*Todo, error)
	Close()
}

type dbHandler struct {
	connection *sql.DB
}

func NewTodoStore(host string, port int, dbname string, user string, password string, sslModeEnabled bool) (TodoStore, error) {
	var sslMode string
	if sslModeEnabled {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Printf("Can't open connection to DB, %s", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Can't ping DB, %s", err.Error())
		return nil, err
	}

	return &dbHandler{
		connection: db,
	}, nil
}

func (h *dbHandler) Close() {
	err := h.connection.Close()
	if err != nil {
		fmt.Printf("Can't close connection to DB, %s", err.Error())
	}
}

func (h *dbHandler) GetAll() ([]Todo, error) {
	rows, err := h.connection.Query("SELECT * FROM Todo ORDER BY id")

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

func (h *dbHandler) Add(todoText string) (*Todo, error) {
	var id int
	err := h.connection.QueryRow("INSERT INTO Todo (todo, done) VALUES($1, $2) RETURNING id", todoText, false).Scan(&id)
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

func (h *dbHandler) ToggleDone(todoId uint32) (*Todo, error) {
	row := h.connection.QueryRow("SELECT * FROM Todo Where id = $1", todoId)
	var todo Todo
	err := row.Scan(&todo.Id, &todo.Todo, &todo.Done)
	if err != nil {
		fmt.Printf("Can't get done state of todo %d: %s", todoId, err.Error())
		return nil, err
	}

	todo.Done = !todo.Done

	_, err = h.connection.Exec("UPDATE Todo SET done = $1 Where id = $2", todo.Done, todoId)
	if err != nil {
		fmt.Printf("Can't update done state of todo %d: %s", todoId, err.Error())
		return nil, err
	}

	return &todo, nil
}

func (h *dbHandler) Delete(todoId uint32) ([]Todo, error) {
	_, err := h.connection.Exec("DELETE FROM Todo WHERE id = $1", todoId)
	if err != nil {
		fmt.Printf("Can't delete todo %d: %s", todoId, err.Error())
		return nil, err
	}
	return h.GetAll()
}
