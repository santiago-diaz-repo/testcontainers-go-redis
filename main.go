package main

import (
	"fmt"
	"github.com/go-redis/redis"
	redis_management "testcontainers-go-redis/redis-management"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	redisMgnt := redis_management.NewRedisManagement(client)
	redisMgnt.Store("testGolang", "golang", 1*time.Second)
	fmt.Println(redisMgnt.Read("testGolang"))
}
