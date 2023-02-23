package model

import "time"

type User struct {
	Id        int64     `json:"id"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
