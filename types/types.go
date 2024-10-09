package types

import "time"

type User struct {
	USER_ID       int       `json:"USER_ID"`
	USERNAME string    `json:"USERNAME"`
	PASSWORD  string    `json:"PASSWORD"`
	EMAIL     string    `json:"EMAIL"`
	CREATED_AT time.Time `json:"CREATED_AT"`
}


type RegisterUserPayload struct {
	USERNAME string `json:"username" validate:"required,max=12"`
	EMAIL     string `json:"email" validate:"required,email"`
	PASSWORD  string `json:"password" validate:"required,min=3,max=130"`
}

type UserStore interface {
	GetUserByUsername(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}