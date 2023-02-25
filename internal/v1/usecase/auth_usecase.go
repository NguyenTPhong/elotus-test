package usecase

import (
	"context"
	"elotus/config"
	_const "elotus/const"
	"elotus/global"
	"elotus/internal/v1/entity"
	"elotus/internal/v1/repository"
	"elotus/internal/v1/repository/model"
	"elotus/package/logger"
	"elotus/package/sha"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthUseCase interface {
	CreateUser(ctx context.Context, req *entity.CreateUserRequest) (*entity.CreateUserResponse, error)
	Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error)
	ValidateToken(token string) (*model.Session, error)
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
		Id:       user.Id,
		Username: user.Username,
	}, nil

}

func (u *authUseCase) Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error) {

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

	expiredAt := time.Now().Add(time.Duration(config.TokenLifeTime) * time.Minute)
	session := &model.Session{
		Id:         user.Id,
		Username:   user.Username,
		LoggedInAt: time.Now(),
		ExpiredAt:  expiredAt,
	}

	// generate token
	token, err := u.generateAuthToken(session)
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

func (u *authUseCase) generateAuthToken(session *model.Session) (string, error) {
	secretKey := []byte(config.JWTKey)

	// Define the claims for the token
	claims := jwt.MapClaims{
		"id":           session.Id,
		"username":     session.Username,
		"expired_at":   session.ExpiredAt.Format(time.RFC3339),
		"logged_in_at": session.LoggedInAt,
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (u *authUseCase) ValidateToken(tokenString string) (session *model.Session, err error) {
	secretKey := []byte(config.JWTKey)

	// Parse and verify the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify that the signing method is HMAC-SHA256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key for verification
		return secretKey, nil
	})
	if err != nil {
		global.Logger.Error("parse token error", zap.Error(err))
		return nil, err
	}

	// Verify that the token is valid
	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		global.Logger.Error("token is invalid", zap.Any("claims", token.Claims))
		return nil, fmt.Errorf(_const.Unauthorized)
	}

	defer func() {
		if r := recover(); r != nil {
			global.Logger.Error("parse token error", zap.Any("error", r))
			session = nil
			err = fmt.Errorf(_const.Unauthorized)
		}
	}()

	// get session from claims
	claims := token.Claims.(jwt.MapClaims)
	session = &model.Session{}
	session.Username = claims["username"].(string)
	session.Id = int64(claims["id"].(float64))
	session.ExpiredAt, _ = time.Parse(time.RFC3339, claims["expired_at"].(string))
	session.LoggedInAt, _ = time.Parse(time.RFC3339, claims["logged_in_at"].(string))

	// check the token is expired
	if session.ExpiredAt.Before(time.Now()) {
		global.Logger.Error("token is expired", zap.Any("claims", token.Claims))
		return nil, fmt.Errorf(_const.Unauthorized)
	}

	// check session in redis, in case that server force logout user
	ssInCache, err := u.authCacheRepo.GetAuthSessionFromCache(session.Id)
	if err != nil || ssInCache == nil {
		global.Logger.Error("get session from redis error", zap.Error(err))
		return nil, fmt.Errorf(_const.Unauthorized)
	}

	return

}
