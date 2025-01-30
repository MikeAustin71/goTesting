/*
	Taken form Go Fragments barrierWaitGroup Example
	http://www.gofragments.net/client/blog/concurrency/2015/10/30/barrierWaitGroup/index.html
	Fragment Description:
	Synchonization study 7: barrier using sync.WaitGroup.
	Here we simply use the sync.WaitGroup type and its methods to implement a 'barrier'. Observe the difference with:'barrierSemaphore'.
	 http://www.gofragments.net/client/blog/concurrency/2015/10/23/barrierSemaphore/index.html

The sync package proposes efficient synchronization functions and primitives.
Channels can be used as Signalling resources (semaphore), and the sync types and primitives can help implement traditional semaphores schemes (locks or condition variables).
*/

package main

import (
	"fmt"
	"sync"
)

// a sync.WaitGroup as a barrier
func process(name string, wgBarrier *sync.WaitGroup) {
	fmt.Printf("Process %s - Task %d\n", name, 1)

	wgBarrier.Done() // release one from wgBarrier which is initialized to 3
	wgBarrier.Wait() // and wait until the last process has arrived
	fmt.Printf("Process %s - Task %d\n", name, 2)
}

// Dispatching processes as goroutines (concurrently)
func dispatch(wgControl *sync.WaitGroup, f func()) {
	wgControl.Add(1)
	go func() {
		f()
		wgControl.Done()
	}()
}
func main() {
	wgControl := &sync.WaitGroup{}
	defer wgControl.Wait()
	wgBarrier := &sync.WaitGroup{}
	// Each thread will be 'blocked' at the 'barrier' until all have arrived
	numberOfThreads := 3
	wgBarrier.Add(numberOfThreads)
	for i := 0; i < numberOfThreads; i++ {
		name := string('A' + i)
		dispatch(wgControl, func() { process(name, wgBarrier) })
	}
}

/* Expected Output:
Process A - Task 1
Process B - Task 1
Process C - Task 1
Process C - Task 2
Process A - Task 2
Process B - Task 2
*/
