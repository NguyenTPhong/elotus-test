package repository

import (
	"elotus/internal/v1/repository/model"

	"gorm.io/gorm"
)

type UploadFileRepository interface {
	Create(file *model.UploadedFile) error
}

type UploadFileRepositoryImpl struct {
	db *gorm.DB
}

func NewUploadFileRepository(db *gorm.DB) UploadFileRepository {
	return &UploadFileRepositoryImpl{db: db}
}

func (u *UploadFileRepositoryImpl) Create(file *model.UploadedFile) error {
	return u.db.Create(file).Error
}
