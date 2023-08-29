package models

import "time"

type User struct {
	Id          string    `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	LineToken   string    `json:"lineToken"`
	DisplayName string    `json:"displayName"`
	PictureUrl  string    `json:"pictureUrl"`
	UserId      string    `json:"userId"`
	CreatedData time.Time `json:"-"`
}