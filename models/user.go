package models

type User struct {
	Name  string
	Email string
}

type userPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Tag   string `default:"user" json:"tag"`
}

func NewUserPayload(name string, email string) userPayload {
	return userPayload{
		Name:  name,
		Email: email,
		Tag:   "user",
	}
}
