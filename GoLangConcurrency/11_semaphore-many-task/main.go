package main

import (
	"fmt"
	"time"
)

func main()  {

	target := 432

	msg := make(chan string)
	done := make(chan bool)

	fmt.Println()

	for i:= 100 ; i <= 1100; i+=100 {
		go search(msg, done, i, target)
	}

	for {
		select {
		case m:= <-msg:
			fmt.Println("Message Recieved", m)
			// done <- true
			return
		default:
			fmt.Println("Waiting...")
			time.Sleep(500 * time.Millisecond)
		}
	}

	close(done)
}


func search(ch chan<- string, done chan bool, threshold int, target int) {

	max := threshold + 99

	for i:= threshold ; i < max; i++ {
		select {
			case <-done:
				close(ch)
				return
			default:
				if i == target {
					ch <- fmt.Sprintf("Target located: %v Threshold: %v",target, threshold )
					done <- true
					close(ch)
					return
				}
		}

		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Finished Threshold: ", threshold)
	for {
		select {
			case <-done:
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}

	return
}