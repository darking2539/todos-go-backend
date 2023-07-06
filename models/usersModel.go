package models

import "time"

type User struct {
	Id          string    `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	CreatedData time.Time `json:"-"`
}
