package repository

import (
	"elotus/internal/v1/repository/model"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

const (
	RedisAuthTokenKey = "user_authenticate_token"
)

type AuthCacheRepository interface {
	CacheAuthSession(userId int64, session *model.Session) error
	GetAuthSessionFromCache(userId int64) (*model.Session, error)
}

type AuthCacheRepositoryImpl struct {
	redis *redis.Client
}

func NewAuthCacheRepository(redis *redis.Client) AuthCacheRepository {
	return &AuthCacheRepositoryImpl{redis: redis}
}

func (u *AuthCacheRepositoryImpl) CacheAuthSession(userId int64, session *model.Session) error {
	return u.redis.HSet(RedisAuthTokenKey, fmt.Sprint(userId), session).Err()
}

func (u *AuthCacheRepositoryImpl) GetAuthSessionFromCache(userId int64) (*model.Session, error) {
	res, err := u.redis.HGet(RedisAuthTokenKey, fmt.Sprint(userId)).Result()
	if err != nil {
		return nil, err
	}

	var session model.Session
	if err = json.Unmarshal([]byte(res), &session); err != nil {
		return nil, err
	}
	return &session, nil
}
