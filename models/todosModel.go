package models

import "time"

type Todo struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Complete    bool      `json:"complete"`
	Username    string    `json:"username"`
	CreatedData time.Time `json:"-"`
}
