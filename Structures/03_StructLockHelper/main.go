package main

import (
	"fmt"
	"sync"
)

type mutexHelper struct {
	lock *sync.Mutex
}

func (mtxHelper *mutexHelper) New() *mutexHelper {

	newMtxHelper := mutexHelper{}

	newMtxHelper.lock = new(sync.Mutex)

	return &newMtxHelper
}

func (mtxHelper *mutexHelper) Lock() {

	if mtxHelper.lock == nil {
		mtxHelper.lock = new(sync.Mutex)
	}

	mtxHelper.lock.Lock()
}

func (mtxHelper *mutexHelper) Unlock() {

	if mtxHelper.lock == nil {
		mtxHelper.lock = new(sync.Mutex)
		return
	}

	mtxHelper.lock.Unlock()
}

type stringDto struct {
	Str1 string
	Str2 string

	lock mutexHelper
}

func main() {

	strDto := newStringDto2()

	/*
	if strDto.lock == nil {
		fmt.Printf("strDto.lock is nil")
		return
	}
*/

	strDto.lock.Lock()

	fmt.Printf("Successfully called strDto.lock.Lock()\n")

	strDto.lock.Unlock()

	fmt.Printf("Successfully called strDto.lock.Unlock()\n")

	strDto.lock.Lock()

	fmt.Printf("Successfully called strDto.lock.Lock() Round-2\n")

	strDto.lock.Unlock()

	fmt.Printf("Successfully called strDto.lock.Unlock() Round-2\n")

}

func newStringDto2() *stringDto {

	strDto := stringDto{}
	// new sync.Mutex is created as 'Unlocked'

	return &strDto
}
