package main

import (
	"fmt"
	"time"
	// "runtime"
	"sync"
)



func main() {

	var wg sync.WaitGroup

	target := 432
	cycleCnt := 11

	msg := make(chan string)
	cnt := make(chan string)
	done := make(chan bool)

	for i := 100; i <= cycleCnt*100; i += 100 {
		wg.Add(1)
		go search(msg, done, cnt, &wg, i, target)
	}

	go waitForIt(&wg, cnt)

	for {
		select {
		case m := <-msg:
			fmt.Println("Message Received:", m)
			for n := range cnt {
				fmt.Println("Count: ", n)
			}

			return
		default:
			fmt.Println("Waiting...")
			time.Sleep(500 * time.Millisecond)
		}
	}

}

func waitForIt(wg *sync.WaitGroup, cnt chan string) {
	wg.Wait()
	close(cnt)
}

func search(ch chan<- string, done chan bool, cnt chan<- string,wg *sync.WaitGroup, threshold int, target int) {

	max := threshold + 99
	count := 0

	for i := threshold; i < max; i++ {
		count++
		select {
		case <-done:
			cnt <- fmt.Sprintf("Threshold: %v  Count: %v",threshold, count)
			wg.Done()
			return
		default:
			if i == target {
				ch <- fmt.Sprintf("Target located: %v Threshold: %v", target, threshold)
				cnt <- fmt.Sprintf("Threshold: %v  Count: %v",threshold, count)
				done <- true
				close(ch)
				close(done)
				wg.Done()
				return
			}

		}

		//runtime.Gosched()
		//time.Sleep(250 * time.Millisecond)
	}

	fmt.Println("Finished Threshold: ", threshold)

	for {
		select {
		case <-done:
			cnt <- fmt.Sprintf("Threshold: %v  Count: %v",threshold, count)
			wg.Done()
			return
		default:
			time.Sleep(250 * time.Millisecond)
		}

	}


	return
}