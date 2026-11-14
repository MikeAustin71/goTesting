package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

// Clone safely creates an independent copy of SafeCounter.
// Returning *SafeCounter ensures the new mutex is never copied by value.
func (c *SafeCounter) Clone() *SafeCounter {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Only copy the protected data (count).
	// By omitting 'mu', the new instance gets a fresh, unlocked sync.Mutex.
	return &SafeCounter{
		count: c.count,
	}
}

func main() {
	original := &SafeCounter{}
	original.Increment()
	original.Increment() // original is now 2

	// Create a safe, independent copy while original might be used concurrently
	cloned := original.Clone()

	var wg sync.WaitGroup
	wg.Add(2)

	// Modify original concurrently
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			original.Increment()
		}
	}()

	// Modify cloned concurrently
	go func() {
		defer wg.Done()
		cloned.Increment()
	}()

	wg.Wait()

	fmt.Printf("Original count: %d\n", original.Value()) // 5
	fmt.Printf("Cloned count:   %d\n", cloned.Value())   // 3
}
