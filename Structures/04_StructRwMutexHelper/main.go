package main

import (
	"fmt"
	"sync"
)

type rWMutexHelper struct {
	lock *sync.RWMutex
}

func (rWMtxHelper *rWMutexHelper) New() *rWMutexHelper {

	newMtxHelper := rWMutexHelper{}

	newMtxHelper.lock = new(sync.RWMutex)

	return &newMtxHelper
}

func (rWMtxHelper *rWMutexHelper) Lock() {

	if rWMtxHelper.lock == nil {
		rWMtxHelper.lock = new(sync.RWMutex)
	}

	rWMtxHelper.lock.Lock()
}

func (rWMtxHelper *rWMutexHelper) RLock() {

	if rWMtxHelper.lock == nil {
		rWMtxHelper.lock = new(sync.RWMutex)
	}

	rWMtxHelper.lock.RLock()
}

func (rWMtxHelper *rWMutexHelper) RUnlock() {

	if rWMtxHelper.lock == nil {
		rWMtxHelper.lock = new(sync.RWMutex)
		return
	}

	rWMtxHelper.lock.RUnlock()
}

func (rWMtxHelper *rWMutexHelper) Unlock() {

	if rWMtxHelper.lock == nil {
		rWMtxHelper.lock = new(sync.RWMutex)
		return
	}

	rWMtxHelper.lock.Unlock()
}



type stringDto struct {
	Str1 string
	Str2 string

	lock *rWMutexHelper
}

func main() {

	strDto := newStringDto2()

	if strDto.lock == nil {
		fmt.Printf("strDto.lock is nil")
		return
	}

	strDto.lock.RLock()
	fmt.Printf("Successfully called strDto.lock.RLock() on unitialized sync.RWMutex! \n")

	strDto.lock.RUnlock()

	fmt.Printf("Successfully called strDto.lock.RUnlock()\n")


	strDto.lock.Lock()

	fmt.Printf("Successfully called strDto.lock.Lock()\n")

	strDto.lock.Unlock()

	fmt.Printf("Successfully called strDto.lock.Unlock()\n")

	strDto.lock.RLock()
	fmt.Printf("Successfully called strDto.lock.RLock() Round-2 \n")

	strDto.lock.RUnlock()

	fmt.Printf("Successfully called strDto.lock.RUnlock() Round-2\n")


	strDto.lock.Lock()

	fmt.Printf("Successfully called strDto.lock.Lock() Round-2\n")

	strDto.lock.Unlock()

	fmt.Printf("Successfully called strDto.lock.Unlock() Round-2\n")

}

func newStringDto2() stringDto {

	strDto := stringDto{}
	strDto.lock = &rWMutexHelper{}
	// new sync.Mutex is created as 'Unlocked'

	return strDto
}

