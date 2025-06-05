package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work
type Task struct {
	ID       int
	Duration time.Duration
}

// Worker processes tasks from the jobs channel
func Worker(ctx context.Context, id int, jobs <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("Worker %d processing task %d\n", id, job.ID)
			time.Sleep(job.Duration)
			select {
			case results <- fmt.Sprintf("Task %d completed by worker %d", job.ID, id):
			case <-ctx.Done():
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 5

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Create channels for jobs and results
	jobs := make(chan Task, numJobs)
	results := make(chan string, numJobs)

	// Create WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go Worker(ctx, w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		select {
		case jobs <- Task{
			ID:       j,
			Duration: time.Duration(j) * 1 * time.Millisecond,
		}:
		case <-ctx.Done():
			fmt.Println("Context cancelled while sending jobs")
			return
		}
	}
	close(jobs)

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for {
		select {
		case result, ok := <-results:
			if !ok {
				return
			}
			fmt.Println(result)
		case <-ctx.Done():
			fmt.Println("Context cancelled while collecting results")
			return
		}
	}
}
