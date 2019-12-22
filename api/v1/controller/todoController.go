package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"main/model"
)

func ListTodos(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("ListTodos")
	// parse request parameter
	err := request.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	// check session is valid
	isValid, err := isSessionValid(request)
	if isValid {
		// get todo list
		fmt.Println("session is valid")
		userID := request.Form.Get("user_id")
		todos, err := model.ListTasks(userID)
		if err != nil {
			return
		}
		// return todos
		writer.Header().Set("Content-Type", "application/json")
		output, err := json.MarshalIndent(&todos, "", "\t\t")
		if err != nil {
			fmt.Println(err)
		}
		writer.Write(output)
	}

}

func AddTodo(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("AddTodo")
	// parse request parameter
	err := request.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	// check session is valid
	isValid, err := isSessionValid(request)
	if isValid {
		fmt.Println("session is valid")
		// add todo to task
		userID, _ := strconv.Atoi(request.Form.Get("user_id"))
		todo := model.Todo{
			Title:    request.Form.Get("title"),
			Body:     request.Form.Get("body"),
			Deadline: request.Form.Get("deadline"),
			UserID:   userID,
		}
		if err := todo.CreateTask(); err != nil {
			fmt.Println("err ", err)
		}
		// return status code
		writer.Header().Set("Content-Type", "application/json")
		output, err := json.MarshalIndent(&todo, "", "\t\t")
		if err != nil {
			fmt.Println(err)
		}
		writer.Write(output)
	}

}

func UpdateTodo(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("UpdateTodo")
	// parse request parameter
	err := request.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	// check session is valid
	isValid, err := isSessionValid(request)
	if isValid {
		fmt.Println("session is valid")
		// update todo task
		todoID, _ := strconv.Atoi(request.Form.Get("todo_id"))
		userID, _ := strconv.Atoi(request.Form.Get("user_id"))
		todo := model.Todo{
			ID:       todoID,
			Title:    request.Form.Get("title"),
			Body:     request.Form.Get("body"),
			Deadline: request.Form.Get("deadline"),
			UserID:   userID,
		}
		if err := todo.UpdateTask(); err != nil {
			fmt.Println("err ", err)
		}
		// return status code
		writer.Header().Set("Content-Type", "application/json")
		output, err := json.MarshalIndent(&todo, "", "\t\t")
		if err != nil {
			fmt.Println(err)
		}
		writer.Write(output)
	}
}

func DeleteTodo(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("DeleteTodo")
	// parse request parameter
	err := request.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	// check session is valid
	isValid, err := isSessionValid(request)
	if isValid {
		// delete todo
		fmt.Println("session is valid")
		todoID, _ := strconv.Atoi(request.Form.Get("todo_id"))
		todo := model.Todo{
			ID: todoID,
		}
		if err := todo.DeleteTask(); err != nil {
			fmt.Println("err ", err)
		}

		// return status code
	}
}
