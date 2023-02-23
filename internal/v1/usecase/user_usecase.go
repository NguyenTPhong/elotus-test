package usecase

import (
	"elotus/internal/v1/entity"
	"elotus/internal/v1/repository"
)

type UserUseCase interface {
	CreateUser(user *entity.CreateUserRequest) (entity.CreateUserResponse, error)
	Login(req *entity.LoginRequest) (*entity.LoginResponse, error)
}

type userUseCase struct {
	repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{userRepository}
}

func (u *userUseCase) CreateUser(user *entity.CreateUserRequest) (entity.CreateUserResponse, error) {
	panic("implement me")
}

func (u *userUseCase) Login(req *entity.LoginRequest) (*entity.LoginResponse, error) {
	panic("implement me")
}
