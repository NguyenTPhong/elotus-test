package migration

import (
	"elotus/internal/v1/repository/model"

	"gorm.io/gorm"
)

func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UploadedFile{})
}

func DropTable(db *gorm.DB) {
	db.Migrator().DropTable(&model.User{})
	db.Migrator().DropTable(&model.UploadedFile{})
}
