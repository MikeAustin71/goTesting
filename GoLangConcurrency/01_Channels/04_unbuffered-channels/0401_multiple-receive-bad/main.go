package _401_multiple_receive_bad

import (
	"fmt"
	//"time"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

func main()  {
	done:= make(chan bool)
	exit:= make(chan bool)
	msg := make(chan string)
	amITheSolution := false
	for i:=1; i < 11; i++ {
		wg.Add(1)
		if i == 5 {
				amITheSolution = true
		} else {
				amITheSolution = false
		}

		go doStuff(msg, done, exit, amITheSolution, i )

	}

	wg.Wait()

	for j:= 1; j <11 ; j++ {
		m:= <-exit
		fmt.Println("Exit Equal: ", m, "  Item: ", j)
	}

	fmt.Println("End of Program")
}



func doStuff(msg chan <-  string, done chan bool, exit chan bool, amITheSolution bool, item int) {
		for i:=0; i< 1000000; i++ {
			select{

			case <-done:
				fmt.Println("Found Done - Exiting function: ", item)
				exit <- true
				wg.Done()
				runtime.Gosched()
				return
			default:
				//time.Sleep(20 * time.Millisecond)

				if amITheSolution && i == 150000 {
					done <- true
					close(done)
					msg <- "Found 150000"
					close(msg)
					wg.Done()
					runtime.Gosched()
					return
				}

			}

			wg.Done()
			runtime.Gosched()
			return

		}






}
