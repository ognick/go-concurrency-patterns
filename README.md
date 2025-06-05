# Go Concurrency Patterns

This repository contains a collection of practical concurrency patterns in Go. Each pattern includes:

- Concise explanation with key concepts
- Use cases and comparisons with related patterns
- ASCII diagram of data/control flow
- Minimal runnable Go code
- Potential pitfalls to avoid

> Before diving into these patterns, make sure you're familiar with Go's synchronization primitives described in [Concurrency in Go: Synchronization Tools](https://medium.com/@ogneslav.work/concurrency-in-go-synchronization-tools-8de05c5fb4a0).


| Name                                           | Description                                     | Use Cases                                       | Memory | CPU  |
|------------------------------------------------|-------------------------------------------------|-------------------------------------------------|--------|------|
| [Generator](./generator/README.md)             | Stream of values                                | • Infinite sequences<br>• Data streaming<br>• Lazy evaluation | Low | Low |
| [Worker Pool](./worker-pool/README.md)         | Concurrent task processing                      | • Task processing<br>• Resource management<br>• Load balancing | Medium | High |
| [Semaphore](./semaphore/README.md)             | Resource access control                         | • Rate limiting<br>• Connection pooling<br>• Resource throttling | Low | Low |
| [Future / Promise](./future-promise/README.md) | Async result handling                           | • Async operations<br>• Result caching<br>• Error handling | Low | Low |
| [Fan-out / Fan-in](./fanout-fanin/README.md)   | Parallel processing                             | • Data aggregation<br>• Load distribution<br>• Parallel computation | High | High |
| [Pub/Sub](./pubsub/README.md)                  | Event broadcasting                              | • Event handling<br>• Message broadcasting<br>• System integration | Medium | Low |

