package model

import (
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

type Todo struct {
	ID        int
	UUID      string
	Title     string
	Body      string
	Deadline  string
	UserID    int
	CreatedAt time.Time
}
