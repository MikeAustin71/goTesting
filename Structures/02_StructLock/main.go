package main

import (
	"fmt"
	"sync"
)

type stringDto struct {
	Str1 string
	Str2 string

	lock *sync.Mutex
}

func main() {

	strDto := newStringDto()

	if strDto.lock == nil {
		fmt.Printf("strDto.lock is nil")
		return
	}

	strDto.lock.Lock()

	fmt.Printf("Successfully called strDto.lock.Lock()\n")

	strDto.lock.Unlock()

	fmt.Printf("Successfully called strDto.lock.Unlock()\n")

}

func newStringDto() stringDto {

	strDto := stringDto{}

	// new sync.Mutex is created as 'Unlocked'
	strDto.lock = new(sync.Mutex)

	return strDto
}
