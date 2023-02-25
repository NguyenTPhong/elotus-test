package model

import "time"

type UploadedFile struct {
	Id        int64     `json:"id" gorm:"primaryKey,autoIncrement"`
	Path      string    `json:"path"`
	MeteData  string    `json:"metadata"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
