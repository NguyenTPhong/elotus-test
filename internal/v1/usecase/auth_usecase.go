package usecase

import (
	"context"
	_const "elotus/const"
	"elotus/global"
	"elotus/internal/v1/entity"
	"elotus/internal/v1/repository"
	"elotus/internal/v1/repository/model"
	"elotus/package/logger"
	"elotus/package/sha"
	"fmt"
	"time"

	"github.com/google/uuid"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthUseCase interface {
	CreateUser(ctx context.Context, req *entity.CreateUserRequest) (*entity.CreateUserResponse, error)
	Login(ctx context.Context, req *entity.LoginRequest, expiredIn int64) (*entity.LoginResponse, error)
}

type authUseCase struct {
	authRepo      repository.AuthRepository
	authCacheRepo repository.AuthCacheRepository
}

func NewUserUseCase(userRepository repository.AuthRepository, authCacheRepo repository.AuthCacheRepository) AuthUseCase {
	return &authUseCase{
		authRepo:      userRepository,
		authCacheRepo: authCacheRepo,
	}
}

func (u *authUseCase) CreateUser(ctx context.Context, req *entity.CreateUserRequest) (*entity.CreateUserResponse, error) {
	// find user by username
	user, err := u.authRepo.FindFirstUser(repository.UserFilter{
		Username: req.Username,
	})

	// process error
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Logger.Error("login error", zap.Error(err), logger.TraceID(ctx))
		return nil, fmt.Errorf(_const.InternalServerErr)
	}

	// check user exist
	if user != nil {
		return nil, fmt.Errorf(_const.UserAlreadyExist)
	}

	// create user
	// gen salt
	salt := uuid.New().String()
	user = &model.User{
		Username: req.Username,
		Salt:     salt,
		Password: sha.Decode(req.Password, salt),
	}

	// save db
	err = u.authRepo.CreateUser(user)
	if err != nil {
		global.Logger.Error("login error", zap.Error(err), logger.TraceID(ctx))
		return nil, fmt.Errorf(_const.InternalServerErr)
	}

	// return response
	return &entity.CreateUserResponse{
		Id: user.Id,
	}, nil

}

func (u *authUseCase) Login(ctx context.Context, req *entity.LoginRequest, expiredIn int64) (*entity.LoginResponse, error) {

	// find user by username
	user, err := u.authRepo.FindFirstUser(repository.UserFilter{
		Username: req.Username,
	})

	// process error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf(_const.UserNotFound)
		}
		global.Logger.Error("login error", zap.Error(err), logger.TraceID(ctx))
		return nil, fmt.Errorf(_const.InternalServerErr)
	}
	// check password
	sendHashedPass := sha.Decode(req.Password, user.Salt)
	if sendHashedPass != user.Password {
		return nil, fmt.Errorf(_const.UserWrongPassword)
	}

	session := &model.Session{
		Id:         user.Id,
		UserName:   user.Username,
		LoggedInAt: time.Now(),
		ExpiredAt:  time.Now().Add(time.Duration(expiredIn) * time.Minute),
	}

	// generate token
	token, expiredAt, err := u.generateAuthToken(session)
	if err != nil {
		global.Logger.Error("generate token error", zap.Error(err), logger.TraceID(ctx))
		return nil, fmt.Errorf(_const.InternalServerErr)
	}

	// save session to redis
	err = u.authCacheRepo.CacheAuthSession(user.Id, session)
	if err != nil {
		global.Logger.Error("save session to redis error", zap.Error(err), logger.TraceID(ctx))
		return nil, fmt.Errorf(_const.InternalServerErr)
	}

	return &entity.LoginResponse{
		Token:    token,
		ExpireAt: expiredAt,
	}, nil
}

func (u *authUseCase) generateAuthToken(session *model.Session) (string, time.Time, error) {
	panic("implement me")
}
