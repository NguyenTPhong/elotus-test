package model

import "time"

type User struct {
	Id        int64     `json:"id" gorm:"primary_key;auto_increment;not null;unique_index`
	Username  string    `json:"username" gorm:"not null;unique_index`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP`
}
