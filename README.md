# in-memory-cache

A minimal in-memory key-value cache written in Go.

This package provides a simple API for storing, retrieving, and deleting values in memory using string keys.

## Features

- Simple key-value storage
- O(1) access using `map`
- Idiomatic Go API
- Zero external dependencies
- Suitable for small services, prototypes, and learning purposes

## Installation

```bash
go get github.com/artfourteen/in-memory-cache
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	cache "github.com/artfourteen/in-memory-cache"
)

func main() {
	c := cache.New()

	// Store value
	c.Set("token", "YOUR_TOKEN_GOES_HERE")

	// Read value
	token, ok := c.Get("token")
	if !ok {
		log.Fatal("token not found in cache")
	}

	fmt.Println(token)

	// Delete value
	c.Delete("token")

	// Try to read again
	_, ok = c.Get("token")
	if !ok {
		fmt.Println("token not found")
	}
}
```

## API

### `func New() *Cache`

Creates and returns a new cache instance.

### `func (c *Cache) Set(key string, value interface{})`

Stores a value in the cache by key.

### `func (c *Cache) Get(key string) (interface{}, bool)`

Retrieves a value by key.

- Returns the value and `true` if the key exists
- Returns `nil` and `false` if the key does not exist

### `func (c *Cache) Delete(key string)`

Removes a value from the cache by key.

## Limitations

- **Not thread-safe**  
  Concurrent access requires synchronization (`sync.Mutex` or `sync.RWMutex`).

- **No expiration (TTL)**  
  Values live in memory until explicitly deleted.

- **No eviction policy**  
  Cache grows until the program ends or values are removed.

## When to use

- Small services and internal tools
- Prototypes and MVPs
- Learning Go and understanding in-memory data structures

## When not to use

- High-concurrency applications
- Large-scale caching
- Distributed systems

In those cases, consider Redis or a cache with eviction and TTL support.

## License

MIT
