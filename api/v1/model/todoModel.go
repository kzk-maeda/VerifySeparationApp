package model

import (
	"time"
)

func ListTasks(userID string) (todos []Todo, err error) {
	sql, err := readSQLFile("model/sql/select_todos.sql")
	if err != nil {
		return
	}
	rows, err := Db.Query(sql, userID)
	if err != nil {
		return
	}
	for rows.Next() {
		task := Todo{}
		if err = rows.Scan(&task.ID, &task.UUID, &task.Title, &task.Body, &task.Deadline, &task.UserID, &task.CreatedAt); err != nil {
			return
		}
		todos = append(todos, task)
	}
	rows.Close()
	return

}

func GetTask() {

}

func (todo *Todo) CreateTask() (err error) {
	sql, err := readSQLFile("model/sql/insert_todo.sql")
	stmt, err := Db.Prepare(sql)
	if err != nil {
		print(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), todo.Title, todo.Body, todo.Deadline, todo.UserID, time.Now()).
		Scan(&todo.ID, &todo.UUID, &todo.Title, &todo.Body, &todo.Deadline, &todo.UserID, &todo.CreatedAt)
	return

}

func (todo *Todo) UpdateTask() (err error) {
	sql, err := readSQLFile("model/sql/update_todo.sql")
	stmt, err := Db.Prepare(sql)
	if err != nil {
		print(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(todo.ID, todo.Title, todo.Body, todo.Deadline, time.Now()).
		Scan(&todo.ID, &todo.Title, &todo.Body, &todo.Deadline)
	return
}

func DeleteTask() {

}
