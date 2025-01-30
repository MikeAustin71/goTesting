package main

import (
	"fmt"
	"runtime"
)

type MessageDto struct {
	Threshold   int
	SearchCount int
	CountMsg    string
	TargetMsg   string
	IsSuccess   bool
}

func main() {


	//  target := 432
	 target := 9000000
	cycleCnt := 11

	msg := make(chan MessageDto)
	done := make(chan bool)

	for i := 100; i <= cycleCnt*100; i += 100 {
		go search(msg, done,  i, target)
	}

	isSuccess := false
	cnt := 0
	for m := range msg {
		cnt++
		if m.IsSuccess && !isSuccess {
			isSuccess = true
			done <- true
			close(done)
			runtime.Gosched()
			fmt.Println(m.TargetMsg + " " + m.CountMsg)
			if cnt == 11 {
				close(msg)
			}
		} else {
			fmt.Println(m.CountMsg)
			if cnt==11{
				close(msg)
			}
		}

	}

}

/*
func waitForIt(wg *sync.WaitGroup, msg chan MessageDto) {
	wg.Wait()
	close(msg)
}
*/


func search(msg chan<- MessageDto, done chan bool, threshold int, target int) {

	max := threshold + 99
	dto := MessageDto{}
	dto.Threshold = threshold
	for i := threshold; i < max; i++ {
		dto.SearchCount++
		select {
		case <-done:
			dto.CountMsg = fmt.Sprintf("Done Exit - Threshold: %v  Count: %v", threshold, dto.SearchCount)
			msg <- dto
			return
		default:
			if i == target {
				dto.TargetMsg = fmt.Sprintf("Target located: %v Threshold: %v", target, threshold)
				dto.IsSuccess = true
				dto.CountMsg = fmt.Sprintf("Target Exit - Threshold: %v  Count: %v", threshold, dto.SearchCount)
				msg <- dto
				// runtime.Gosched()
				return
			}

			runtime.Gosched()

		}
	}

/*
	for {
		select {
		case <-done:
			dto.CountMsg = fmt.Sprintf(" Final - Threshold: %v  Count: %v", threshold, dto.SearchCount)
			msg <- dto
			return
		default:
			time.Sleep(50 * time.Microsecond)
		}
	}

*/

	dto.CountMsg = fmt.Sprintf("Func Exit - Threshold: %v  Count: %v", threshold, dto.SearchCount)
	msg <- dto
//time.Sleep(50 * time.Microsecond)
	return

}
