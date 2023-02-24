package repository

import (
	"elotus/internal/v1/repository/model"

	"gorm.io/gorm"
)

type UserFilter struct {
	Username string
	Id       int64
}

type AuthRepository interface {
	CreateUser(user *model.User) error
	FindFirstUser(filter UserFilter) (*model.User, error)
}

type AuthRepositoryImpl struct {
	database *gorm.DB
}

func NewAuthRepository(database *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{database: database}
}

func (u *AuthRepositoryImpl) CreateUser(user *model.User) error {
	return u.database.Create(user).Error
}

func (u *AuthRepositoryImpl) FindFirstUser(filter UserFilter) (*model.User, error) {
	var user model.User

	query := u.database.Model(&model.User{})

	if filter.Username != "" {
		query = query.Where("username = ?", filter.Username)
	}
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}

	err := query.First(&user).Error
	return &user, err
}
