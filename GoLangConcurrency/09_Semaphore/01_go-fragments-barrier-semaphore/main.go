/*
	Go Fragments barrier Semaphore Example
	http://www.gofragments.net/client/blog/concurrency/2015/10/23/barrierSemaphore/index.html
===============================================================================
Synchronization Study 6: barrier.
3 concurrent processes: A(a1, a1), B(b1, b2) and C(c1, c2).
Each has 2 tasks.
How to execute them in that order:
a1 -> b1 -> c1 -> then the remaining tasks.
The generalization of 'rendezvous', to n processes, requires to implement the semaphore Pattern named:
'barrier'.
A barrier is a way of synchronizing a 'fixed' number N of processes.
Each process calls the barrier.Acquire() operation and blocks until all N processes do the same.
At that point, all the processes should be free to continue.
To be noted:
1- 2 named semaphores, 'barrier' and 'mutex' 2- their initial state is different:
0 (blocking) and 1 (available) 3- the numberOfProcess is known:
numberOfThreads here.
4- the current state is known too, currentValue here.
===============================================================================
Note: you will need the package:
$ go get github.com/robryk/semaphore
Reference: https://github.com/robryk/semaphore
*/

package main

import (
	"fmt"
	"sync"

	"github.com/robryk/semaphore"
)

var (
	barrier                       *semaphore.Semaphore
	mutex                         *semaphore.Semaphore // a binary semaphore used as a Mutex
	currentValue, numberOfThreads int
)

// Comment all the line with 'aBarrier()'. Observe the difference
func process(name string, barrier *semaphore.Semaphore) {
	fmt.Printf("Process %s - Task %d\n", name, 1)

	//aBarrier()
	fmt.Printf("Process %s - Task %d\n", name, 2)
}

// Let's dispatch 3 concurrent processes
func dispatch(wg *sync.WaitGroup, f func()) {
	wg.Add(1)
	go func() {
		f()
		wg.Done()
	}()
}

// a re-usable barrier
func aBarrier() {
	// updating the currentValue must be done safely
	mutex.Acquire(1)
	// as long as the numberOfThreads is not reached we must be locked
	currentValue--
	// when numberOfThreads is reached, we drain the semaphore 'barrier'
	if currentValue == 0 {
		for i := 0; i < numberOfThreads; i++ {
			barrier.Release(1)
		}
		currentValue = numberOfThreads
	}
	// currentValue is made available again
	mutex.Release(1)

	// the following statements, together in that order, define a 'turnstile'
	barrier.Acquire(1) // initial state being 0, in wait until a release is made
	barrier.Release(1) // allow the next thread to be unblocked
}
func main() {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	// all threads are 'blocked' at the 'barrier' until the last one has arrived
	numberOfThreads = 3 // value of N for this example
	currentValue = numberOfThreads
	// The 2 semaphores have different initial state
	barrier = semaphore.New(0) // blocking
	mutex = semaphore.New(1)   // non blocking
	for i := 0; i < numberOfThreads; i++ {
		name := string('A' + i)
		dispatch(wg, func() { process(name, barrier) })
	}
}

/* Output:
Process A - Task 1
Process B - Task 1
Process C - Task 1
Process C - Task 2
Process A - Task 2
Process B - Task 2
When aBarrier() line is commented in all 3 processes, we get:
Process C - Task 1
Process C - Task 2
Process A - Task 1
Process A - Task 2
Process B - Task 1
Process B - Task 2
*/
