package model

import "time"

type Session struct {
	Id         int64     `json:"id"`
	Username   string    `json:"username"`
	LoggedInAt time.Time `json:"logged_in_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}
