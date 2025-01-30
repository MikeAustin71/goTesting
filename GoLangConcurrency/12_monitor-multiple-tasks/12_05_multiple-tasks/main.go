package main

import (
	"fmt"
	"runtime"
	"sync"

	"time"
)

type MessageDto struct {
	Threshold   int
	SearchCount int
	CountMsg    string
	TargetMsg   string
	IsSuccess   bool
}

func main() {

	var wg sync.WaitGroup

	 target := 432
	// target := 9000000
	cycleCnt := 11

	msg := make(chan MessageDto)
	done := make(chan bool)

	wg.Add(cycleCnt)



	for i := 100; i <= cycleCnt*100; i += 100 {
		go search(msg, done, &wg, i, target)
	}

	go waitForIt(&wg, msg)
	isSuccess := false
	for m := range msg {

		if m.IsSuccess && !isSuccess {
			isSuccess = true
			done <- true
			close(done)
			runtime.Gosched()
			fmt.Println(m.TargetMsg + " " + m.CountMsg)
		} else {
			fmt.Println(m.CountMsg)
		}

	}

}

func waitForIt(wg *sync.WaitGroup, msg chan MessageDto) {
	wg.Wait()
	close(msg)
}

func search(msg chan<- MessageDto, done chan bool, wg *sync.WaitGroup, threshold int, target int) {

	max := threshold + 99
	dto := MessageDto{}
	dto.Threshold = threshold
	for i := threshold; i < max; i++ {
		dto.SearchCount++
		select {
		case <-done:
			dto.CountMsg = fmt.Sprintf("Done Exit - Threshold: %v  Count: %v", threshold, dto.SearchCount)
			msg <- dto
			wg.Done()
			return
		default:
			if i == target {
				dto.TargetMsg = fmt.Sprintf("Target located: %v Threshold: %v", target, threshold)
				dto.IsSuccess = true
				dto.CountMsg = fmt.Sprintf("Target Exit - Threshold: %v  Count: %v", threshold, dto.SearchCount)
				msg <- dto
				wg.Done()
				runtime.Gosched()
				return
			}

			runtime.Gosched()

		}
	}


	for {
		select {
		case <-done:
			dto.CountMsg = fmt.Sprintf(" Final - Threshold: %v  Count: %v", threshold, dto.SearchCount)
			msg <- dto
			wg.Done()
			return
		default:
			runtime.Gosched()
			time.Sleep(50 * time.Microsecond)
		}
	}


	/*
	dto.CountMsg = fmt.Sprintf("Func Exit - Threshold: %v  Count: %v", threshold, dto.SearchCount)
	msg <- dto
	wg.Done()
	//time.Sleep(50 * time.Microsecond)
	return
	*/

}
