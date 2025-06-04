package db

import "github.com/redis/go-redis/v9"

func NewRedis(redisURI string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: redisURI,
	})

	return client
}
