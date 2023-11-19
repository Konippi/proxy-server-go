package redis

import "github.com/redis/go-redis/v9"

type RedisClient interface {
}

type redisClient struct {
	uc redis.UniversalClient
}

func NewRedisClient(uc redis.UniversalClient) RedisClient {
	return &redisClient{
		uc: uc,
	}
}

// client := redis.NewClient(&redis.Options{
// 	Addr:     "localhost:6379",
// 	Password: "",
// 	DB:       0,
// })
