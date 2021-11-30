package redis_management

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisManagement struct {
	redisClient *redis.Client
}

func NewRedisManagement(redisClient *redis.Client) RedisManagement{
	return RedisManagement{
		redisClient: redisClient,
	}
}

func (rm *RedisManagement) Read(key string) string{
	return rm.redisClient.Get(key).Val()
}

func (rm *RedisManagement) Store(key string, value string, duration time.Duration){
	rm.redisClient.Set(key,value, duration)
}