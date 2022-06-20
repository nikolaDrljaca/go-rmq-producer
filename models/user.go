package models

type User struct {
	Name string
	Email string
}

type userPayload struct {
	Name string
	Email string
	Tag string `default:"user"`
}

func NewUserPayload(name string, email string) userPayload {
	return userPayload{
		Name: name,
		Email: email,
		Tag: "user",
	}
}

