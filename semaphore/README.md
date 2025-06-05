# Semaphore Pattern

The Semaphore pattern controls concurrent access to resources using a counting semaphore implemented with channels.

## Key Concepts

- Resource access control
- Permit counting
- Channel-based implementation
- Synchronization

## Implementation Details

- Channel-based permit management
- Context for timeout control
- Cleanup management
- Error handling

## Use Cases

- Rate limiting
- Connection pooling
- Resource throttling
- Concurrent access control
- Task limiting

## Example

[main.go](./main.go) demonstrates a simple semaphore implementation with:
- Permit acquisition and release
- Timeout handling
- Resource cleanup
- Error handling

## ASCII Diagram

```
Request 1 ----[acquire]----> Semaphore ----[permit]----> Resource 1
Request 2 ----[acquire]---->    |              |         Resource 2
Request 3 ----[acquire]---->    |              |         Resource 3
Request 4 ----[acquire]---->    |              |
Request 5 ----[acquire]---->    |              |
                                |              |
                                v              v
                            [permits]      [release]
                            (channel)      (channel)
```

## Potential Pitfalls

- Channel leaks
- Goroutine leaks
- Deadlocks
- Memory issues

## Best Practices

- Use context for cancellation
- Close channels properly
- Monitor permit count
- Handle errors appropriately

[Back to main README](../README.md) 