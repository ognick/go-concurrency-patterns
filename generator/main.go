package main

import (
	"context"
	"fmt"
	"time"
)

// Generator returns a channel that produces integers
func Generator(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- i
			time.Sleep(1 * time.Millisecond)
		}
	}()
	return ch
}

// GeneratorWithStop returns a channel that produces integers until context is cancelled
func GeneratorWithStop(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; ; i++ {
			select {
			case ch <- i:
				time.Sleep(1 * time.Millisecond)
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Example 1: Simple generator
	fmt.Println("Simple generator example:")
	numbers := Generator(5)
	for n := range numbers {
		fmt.Printf("Received: %d\n", n)
	}

	// Example 2: Generator with stop
	fmt.Println("\nGenerator with stop example:")
	numbersWithStop := GeneratorWithStop(ctx)

	// Read numbers until context is cancelled
	for i := 0; i < 5; i++ {
		fmt.Printf("Received: %d\n", <-numbersWithStop)
	}
}
