// Sample program to show how to use a mutex to define critical
// sections of code that need synchronous access.
// Sorry, don't know where this code originated.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// counter is a variable incremented by all goroutines.
var counter int

// mutex is used to define a critical section of code.
var mutex sync.Mutex

// main is the entry point for the application.
func main() {
	// Number of goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines.
	for i := 0; i < grs; i++ {
		go func() {
			incCounter()
			wg.Done()
		}()
	}
	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

// incCounter increments the package level Counter variable
// using the Mutex to synchronize and provide safe access.
func incCounter() {
	for count := 0; count < 2; count++ {
		// Only allow one goroutine through this
		// critical section at a time.
		mutex.Lock()
		{
			// Capture the value of counter.
			value := counter
			// Yield the thread and be placed back in queue.
			runtime.Gosched()
			// Increment our local value of counter.
			value++
			// Store the value back into counter.
			counter = value
		}

		mutex.Unlock()
		// Release the lock and allow any
		// waiting goroutine through.
	}
}


/*	Output
	$ go run -race main.go
	Final Counter: 4
*/
