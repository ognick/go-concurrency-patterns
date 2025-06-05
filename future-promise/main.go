package main

import (
	"context"
	"fmt"
	"time"
)

// Future represents a value that will be available in the future
type Future[T any] struct {
	result chan T
	err    chan error
}

// NewFuture creates a new Future that will execute the given function
func NewFuture[T any](fn func() (T, error)) *Future[T] {
	f := &Future[T]{
		result: make(chan T, 1),
		err:    make(chan error, 1),
	}

	go func() {
		result, err := fn()
		if err != nil {
			f.err <- err
			return
		}
		f.result <- result
	}()

	return f
}

// Get waits for the result with timeout
func (f *Future[T]) Get(ctx context.Context) (T, error) {
	zeroVal := func() T {
		var v T
		return v
	}
	select {
	case result := <-f.result:
		return result, nil
	case err := <-f.err:
		return zeroVal(), err
	case <-ctx.Done():
		return zeroVal(), ctx.Err()
	}
}

func main() {
	// Create a future that simulates a long-running computation
	future := NewFuture(func() (string, error) {
		time.Sleep(100 * time.Millisecond)
		return "Computation completed", nil
	})

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Example with error
	errorFuture := NewFuture(func() (string, error) {
		time.Sleep(100 * time.Millisecond)
		return "", fmt.Errorf("computation failed")
	})
	if _, err := errorFuture.Get(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Wait for the result
	result, err := future.Get(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %v\n", result)
}
