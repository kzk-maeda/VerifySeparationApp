package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"main/model"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Login")
	// parse request parameter
	err := request.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	// search user by email
	email := request.Form.Get("email")
	password := request.Form.Get("password")

	user, err := model.GetUserByEmail(email)
	if err != nil {
		return
	}

	// create session if password is valid
	var session model.Session
	if user.Password == model.Encrypt(password) {
		session, err = user.CreateSession()
		if err != nil {
			return
		}
	} else {
		fmt.Println("Password Invalid")
	}

	// return session
	writer.Header().Set("Content-Type", "application/json")
	output, err := json.MarshalIndent(&session, "", "\t\t")
	if err != nil {
		fmt.Println(err)
	}
	writer.Write(output)
}

func Logout(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Logout")

}

func Signup(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Signin")
	// parse requset parameter
	err := request.ParseForm()
	if err != nil {
		return
	}

	// create user to user_table
	user := model.User{
		Name:     request.Form.Get("name"),
		Email:    request.Form.Get("email"),
		Password: request.Form.Get("password"),
	}
	fmt.Println("DEBUG : ", request.Form.Get("name"), request.Form.Get("email"), request.Form.Get("password"))
	if err := user.CreateUser(); err != nil {
		fmt.Println("err ", err)
	}
	// return status code
	writer.Header().Set("Content-Type", "application/json")
	output, err := json.MarshalIndent(&user, "", "\t\t")
	if err != nil {
		fmt.Println(err)
	}
	writer.Write(output)

}
