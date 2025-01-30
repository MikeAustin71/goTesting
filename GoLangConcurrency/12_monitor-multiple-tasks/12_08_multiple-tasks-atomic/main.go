package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

type MessageDto struct {
	Threshold   int
	SearchCount int
	CountMsg    string
	TargetMsg   string
	IsSuccess   bool
}

func main() {


	target := 901
	cycleCnt := 11

	msg := make(chan MessageDto)
	var done uint64 = 0
	cnt := 0
	isSuccess := false


	for i := 100; i <= cycleCnt*100; i += 100 {
		go search(msg, &done, i, target)
	}


	for m := range msg {

		cnt++
		if cnt >= cycleCnt {
			close(msg)
		}

		if m.IsSuccess && !isSuccess {
			isSuccess = true
			fmt.Println(m.TargetMsg + " " + m.CountMsg)
		} else {
			fmt.Println(m.CountMsg)
		}

	}

}



func search(msg chan<- MessageDto, done *uint64, threshold int, target int) {

	max := threshold + 99
	dto := MessageDto{}
	dto.Threshold = threshold
	var doneTest uint64

	for i := threshold; i < max; i++ {
		dto.SearchCount++
		if i == target {
			dto.TargetMsg = fmt.Sprintf("Target located: %v Threshold: %v", target, threshold)
			dto.IsSuccess = true
			dto.CountMsg = fmt.Sprintf("Target Exit - Threshold: %v  Count: %v", threshold, dto.SearchCount)
			atomic.AddUint64(done, 1)
			msg <- dto
			runtime.Gosched()
			//time.Sleep(10 * time.Millisecond)
			return
		}

		doneTest = atomic.LoadUint64(done)

		if doneTest > 0 {
			dto.CountMsg = fmt.Sprintf("Done Exit - Threshold: %v  Count: %v", threshold, dto.SearchCount)
			msg <- dto
			return
		}

		// runtime.Gosched()
		// time.Sleep(1 * time.Millisecond)

	}


	dto.CountMsg = fmt.Sprintf("Func Exit - Threshold: %v  Count: %v", threshold, dto.SearchCount)
	msg <- dto
	return
}
