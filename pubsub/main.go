package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Message represents a message in the system
type Message struct {
	Topic   string
	Content interface{}
}

// PubSub represents a publish/subscribe system
type PubSub struct {
	mu          sync.RWMutex
	subscribers map[string][]chan Message
}

// NewPubSub creates a new PubSub instance
func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan Message),
	}
}

// Subscribe subscribes to a topic
func (ps *PubSub) Subscribe(topic string) chan Message {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan Message, 1)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)
	return ch
}

// Publish publishes a message to a topic
func (ps *PubSub) Publish(topic string, content interface{}) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	msg := Message{
		Topic:   topic,
		Content: content,
	}

	for _, ch := range ps.subscribers[topic] {
		select {
		case ch <- msg:
		default:
			// Skip if subscriber is not ready
		}
	}
}

// Unsubscribe removes a subscription
func (ps *PubSub) Unsubscribe(topic string, ch chan Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subscribers := ps.subscribers[topic]
	for i, subscriber := range subscribers {
		if subscriber == ch {
			ps.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
			close(ch)
			return
		}
	}
}

func main() {
	// Create a new PubSub instance
	ps := NewPubSub()

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Subscribe to topics
	newsCh := ps.Subscribe("news")
	weatherCh := ps.Subscribe("weather")

	// Start subscribers
	var wg sync.WaitGroup
	wg.Add(2)

	// News subscriber
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-newsCh:
				fmt.Printf("News: %v\n", msg.Content)
			}
		}
	}()

	// Weather subscriber
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-weatherCh:
				fmt.Printf("Weather: %v\n", msg.Content)
			}
		}
	}()

	// Publish some messages
	ps.Publish("news", "Breaking: Go 1.21 released!")
	ps.Publish("weather", "Sunny, 25°C")
	ps.Publish("news", "Dinamo wins championship")
	ps.Publish("weather", "Rain expected tomorrow")

	// Wait for subscribers to finish
	wg.Wait()
}
