// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"sync"
// 	"time"

// 	"://github.com"
// )

// const queueName = "task_queue"

// // Task represents the structured payload we send through the queue
// type Task struct {
// 	ID        string    `json:"id"`
// 	Payload   string    `json:"payload"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// var ctx = context.Background()

// func main() {
// 	// 1. Initialize Redis client
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 	})
// 	defer rdb.Close()

// 	var wg sync.WaitGroup

// 	// 2. Start the Consumer (Worker) loop in a goroutine
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		worker(rdb)
// 	}()

// 	// 3. Start the Producer to simulate incoming tasks
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		producer(rdb)
// 	}()

// 	// Wait for goroutines to finish (in production, use OS signals for graceful shutdown)
// 	wg.Wait()
// }

// // Producer sends work to the queue using LPUSH
// func producer(rdb *redis.Client) {
// 	for i := 1; i <= 5; i++ {
// 		task := Task{
// 			ID:        fmt.Sprintf("task-%d", i),
// 			Payload:   fmt.Sprintf("Process video segment #%d", i),
// 			CreatedAt: time.Now(),
// 		}

// 		// Serialize to JSON string
// 		taskBytes, err := json.Marshal(task)
// 		if err != nil {
// 			log.Printf("[Producer] Failed to marshal task: %v", err)
// 			continue
// 		}

// 		// LPUSH to add to the head of the list
// 		err = rdb.LPush(ctx, queueName, taskBytes).Err()
// 		if err != nil {
// 			log.Printf("[Producer] LPUSH error: %v", err)
// 		} else {
// 			fmt.Printf("[Producer] Enqueued: %s\n", task.ID)
// 		}

// 		// Sleep briefly between producing items
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

// // Consumer waits for work from the queue using BRPOP
// func worker(rdb *redis.Client) {
// 	fmt.Println("[Worker] Started and waiting for tasks...")

// 	for {
// 		// BRPOP blocks indefinitely (timeout = 0) until an item is available.
// 		// It returns a slice: [key_name, popped_value]
// 		results, err := rdb.BRPop(ctx, 0, queueName).Result()
// 		if err != nil {
// 			log.Printf("[Worker] BRPOP error: %v", err)
// 			time.Sleep(1 * time.Second) // Prevent tight crash loop on network error
// 			continue
// 		}

// 		// Extract the actual message string (index 1)
// 		// index 0 contains the key name ("task_queue")
// 		taskData := results[1]

// 		// Deserialize JSON back into our struct
// 		var task Task
// 		if err := json.Unmarshal([]byte(taskData), &task); err != nil {
// 			log.Printf("[Worker] Failed to unmarshal task: %v", err)
// 			continue
// 		}

// 		// Process the task
// 		fmt.Printf("[Worker] Processing: %s -> %s (Lag: %v)\n", 
// 			task.ID, task.Payload, time.Since(task.CreatedAt))
		
// 		// Simulate heavy processing work
// 		time.Sleep(1 * time.Second)
// 	}
// }
