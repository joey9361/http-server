package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

const queueName = "myqueue"

type Task struct {
	ID 		string
	Payload	string
	CreatedAt time.Time
}

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()
	
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		Worker(rdb)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Process(rdb)
	}()

	wg.Wait()
}

func Process(rdb *redis.Client) {
	for i := 0; i < 5; i++ {
		task := Task{
			ID: "he",
			Payload: "somebody once told me",
			CreatedAt: time.Now(),
		}

		jsonBody, err := json.Marshal(task)
		if err != nil {
			fmt.Println("Couldnt marshal task")
			continue
		}

		err = rdb.LPush(ctx, queueName, jsonBody).Err()
		if err != nil {
			fmt.Println("Error pushing to redis")
			continue
		}


	}
}

func Worker(rdb *redis.Client) {
	for {
		result, err := rdb.BRPop(ctx, 5, queueName).Result()
		if err != nil {
			fmt.Println("Failed in popping task from redis queue")
			continue
		}
		// [key, value]
		respBody := result[1]

		var task Task

		if err := json.Unmarshal([]byte(respBody), &task); err != nil {
			fmt.Println("Failed to unmarshal json body")
			continue
		}

		fmt.Println("Succesfully unmarshaled task from redis")

		fmt.Printf("[Worker] Processing: %s -> %s (Lag: %v)\n", 
			task.ID, task.Payload, time.Since(task.CreatedAt))
		// Simulate task being executed
		time.Sleep(2 * time.Second)
	}
}

