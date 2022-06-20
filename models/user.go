package models

type User struct {
	Name string
	Email string
}

type UserPayload struct {
	Name string
	Email string
	Tag string `default:"user"`
}

