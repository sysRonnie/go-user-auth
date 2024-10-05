package types

import "time"

type User struct {
	USER_ID       int       `json:"USER_ID"`
	USERNAME string    `json:"USERNAME"`
	PASSWORD  string    `json:"PASSWORD"`
	EMAIL     string    `json:"EMAIL"`
	CREATED_AT time.Time `json:"CREATED_AT"`
}