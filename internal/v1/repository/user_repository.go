package repository

import (
	"elotus/internal/v1/repository/model"
	"gorm.io/gorm"
)

type UserFilter struct {
	Username string
	Id       int64
}

type UserRepository interface {
	CreateUser(user *model.User) error
	FindUsers(filter UserFilter) ([]*model.User, error)
}

type UserRepositoryImpl struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &UserRepositoryImpl{database: database}
}

func (u *UserRepositoryImpl) CreateUser(user *model.User) error {
	return u.database.Create(user).Error
}

func (u *UserRepositoryImpl) FindUsers(filter UserFilter) ([]*model.User, error) {
	var users []*model.User
	query := u.database.Model(&model.User{})
	if filter.Username != "" {
		query = query.Where("username = ?", filter.Username)
	}
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	err := query.Find(&users).Error
	return users, err
}
