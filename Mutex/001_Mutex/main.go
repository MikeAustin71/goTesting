package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently across multiple goroutines.
type SafeCounter struct {
	mu    sync.Mutex // Protects the count field below. Zero value is ready to use.
	count int
}

// Increment safely increases the counter by 1.
// A pointer receiver (*SafeCounter) is required so the struct (and its mutex) is not copied.
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock() // Guarantees the lock is released when the method returns.

	c.count++
}

// Value safely reads and returns the current count.
// Reads must also acquire the lock to prevent a data race with concurrent writes.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

func main() {
	// No explicit constructor is needed because the zero value of sync.Mutex is unlocked.
	var counter SafeCounter
	var wg sync.WaitGroup

	const workers = 1000
	wg.Add(workers)

	// Launch 1000 concurrent goroutines that all modify the same SafeCounter.
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Final count: %d\n", counter.Value())
}
