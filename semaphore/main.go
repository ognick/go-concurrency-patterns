package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Semaphore represents a counting semaphore
type Semaphore struct {
	sem chan struct{}
}

// NewSemaphore creates a new semaphore with the specified number of permits
func NewSemaphore(permits int) *Semaphore {
	return &Semaphore{
		sem: make(chan struct{}, permits),
	}
}

// Acquire acquires a permit, blocking until one is available
func (s *Semaphore) Acquire() func() {
	s.sem <- struct{}{}
	return s.release
}

// TryAcquire attempts to acquire a permit with timeout
func (s *Semaphore) TryAcquire(ctx context.Context) (func(), bool) {
	select {
	case s.sem <- struct{}{}:

		return s.release, true
	case <-ctx.Done():
		return func() {}, false
	}
}

// release releases a permit
func (s *Semaphore) release() {
	<-s.sem
}

func main() {
	// Create a semaphore with 2 permits
	sem := NewSemaphore(2)
	var wg sync.WaitGroup

	// Simulate 5 workers trying to access a shared resource
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Try to acquire permit with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			if release, ok := sem.TryAcquire(ctx); ok {
				release()
				fmt.Printf("Worker %d acquired permit\n", id)
				time.Sleep(1 * time.Millisecond)
				fmt.Printf("Worker %d released permit\n", id)
			} else {
				fmt.Printf("Worker %d failed to acquire permit\n", id)
			}
		}(i)
	}

	wg.Wait()
}
