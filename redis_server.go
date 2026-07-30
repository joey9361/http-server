package main

import (
	"github.com/redis/go-redis/v9"
	"context"
	"fmt"
	"log"
)

func newRedisClient() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Your Redis host and port
		Password: "",               // No password set by default
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	fmt.Println("Connected successfully! Redis responded with: ",pong)
}