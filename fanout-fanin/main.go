package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// process simulates work on input
func process(input int) int {
	time.Sleep(1 * time.Millisecond)
	return input * 2
}

// fanOut starts N worker goroutines that read from the jobs channel
func fanOut(ctx context.Context, jobs <-chan int, numWorkers int) []<-chan int {
	var outs []<-chan int
	for i := 0; i < numWorkers; i++ {
		out := make(chan int)
		outs = append(outs, out)

		go func(out chan<- int) {
			defer close(out)
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-jobs:
					if !ok {
						return
					}
					result := process(job)
					select {
					case <-ctx.Done():
						return
					case out <- result:
					}
				}
			}
		}(out)
	}
	return outs
}

// fanIn merges multiple result channels into a single channel
func fanIn(ctx context.Context, channels []<-chan int) <-chan int {
	var wg sync.WaitGroup
	merged := make(chan int)

	output := func(ch <-chan int) {
		defer wg.Done()
		for v := range ch {
			select {
			case <-ctx.Done():
				return
			case merged <- v:
			}
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	jobs := make(chan int)
	go func() {
		defer close(jobs)
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
	}()

	// Fan-out: 3 workers read from jobs
	workers := fanOut(ctx, jobs, 3)

	// Fan-in: merge results from all workers
	results := fanIn(ctx, workers)

	for r := range results {
		fmt.Printf("Result: %d\n", r)
	}
}
