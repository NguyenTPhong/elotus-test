package redis

import "github.com/go-redis/redis"

func NewClient(host, password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
