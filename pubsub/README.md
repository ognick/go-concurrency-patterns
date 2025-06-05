# Publish/Subscribe Pattern

The Publish/Subscribe pattern enables decoupled communication between components using channels and goroutines.

## Key Concepts

- Decoupled communication using channels
- Topic-based message routing
- Asynchronous message delivery
- Synchronization with sync.RWMutex

## Implementation Details

- Channel-based message delivery
- Topic management with map
- Proper cleanup of channels and goroutines

## Use Cases

- Event distribution between goroutines
- Asynchronous communication between program components
- Event handling in different parts of the application
- Data transfer between independent modules
- Internal state change notifications

## Example

[main.go](./main.go) demonstrates a simple pub/sub implementation with:
- Topic-based message routing
- Multiple subscribers per topic
- Asynchronous message delivery
- Proper cleanup

## ASCII Diagram

```
Publisher 1 ----[msg1]----> Topic A ----[msg1]----> Subscriber 1 (goroutine)
Publisher 2 ----[msg2]----> |                       Subscriber 2 (goroutine)
Publisher 3 ----[msg3]----> |            
                            |            
Publisher 4 ----[msg4]----> Topic B ----[msg4]----> Subscriber 3 (goroutine)
                            |           
Publisher 5 ----[msg5]----> |           
                            |           
                       Topic Registry
                     (map[string][]chan)
                            |
                            v
                        sync.RWMutex
                            |
                            v
                         Context
                      (cancellation)
```

## Potential Pitfalls

- Channel leaks
- Goroutine leaks
- Deadlocks
- Memory issues

## Best Practices

- Use context for cancellation
- Close channels properly
- Monitor goroutine counts
- Handle errors appropriately

[Back to main README](../README.md) 