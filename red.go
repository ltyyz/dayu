package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

const (
	RedisAddr     = "redis.lt:6379"
	RedisPassword = "********"
)

var RedisClient *redis.Client

func ConnectToRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Password: RedisPassword,
		Addr:     RedisAddr,
		DB:       0,
	})
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Client connection ping test:", pong)
}
