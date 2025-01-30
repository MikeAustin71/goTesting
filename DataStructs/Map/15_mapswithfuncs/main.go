package main

import (
	"fmt"
	"sync"
)

type NumStrSetup struct {
	Id      int
	Country string
}

type CountryFmt struct {
	lock *sync.Mutex
}

func (cntryFmt *CountryFmt) France() NumStrSetup {

	if cntryFmt.lock == nil {
		cntryFmt.lock = new(sync.Mutex)
	}

	cntryFmt.lock.Lock()

	defer cntryFmt.lock.Unlock()

	newSet := NumStrSetup{}

	newSet.Id = 4
	newSet.Country = "France"

	return newSet
}

func (cntryFmt *CountryFmt) Germany() NumStrSetup {

	if cntryFmt.lock == nil {
		cntryFmt.lock = new(sync.Mutex)
	}

	cntryFmt.lock.Lock()

	defer cntryFmt.lock.Unlock()

	newSet := NumStrSetup{}

	newSet.Id = 3
	newSet.Country = "Germany"

	return newSet
}

func (cntryFmt *CountryFmt) UnitedKingdom() NumStrSetup {

	if cntryFmt.lock == nil {
		cntryFmt.lock = new(sync.Mutex)
	}

	cntryFmt.lock.Lock()

	defer cntryFmt.lock.Unlock()

	newSet := NumStrSetup{}

	newSet.Id = 2
	newSet.Country = "United Kingdom"

	return newSet
}

func (cntryFmt *CountryFmt) UnitedStates() NumStrSetup {

	if cntryFmt.lock == nil {
		cntryFmt.lock = new(sync.Mutex)
	}

	cntryFmt.lock.Lock()

	defer cntryFmt.lock.Unlock()

	newSet := NumStrSetup{}

	newSet.Id = 1
	newSet.Country = "United States"

	return newSet
}

func (cntryFmt CountryFmt) Ptr() *CountryFmt {

	if cntryFmt.lock == nil {
		cntryFmt.lock = new(sync.Mutex)
	}

	newCFmt := new(CountryFmt)

	return newCFmt
}

var mapNumSet = map[int]NumStrSetup{
	1: CountryFmt{}.Ptr().UnitedStates(),
	2: CountryFmt{}.Ptr().UnitedKingdom(),
	3: CountryFmt{}.Ptr().Germany(),
	4: CountryFmt{}.Ptr().France(),
}

//
// This works for value receivers.
//var mapNumSet = map[int] NumStrSetup {
//	1 : CountryFmt{}.UnitedStates(),
//	2 : CountryFmt{}.UnitedKingdom(),
//	3 : CountryFmt{}.Germany(),
//	4 : CountryFmt{}.France(),
//}

func main() {

	fmt.Println()
	fmt.Println("Main - Maps With Functions")
	fmt.Println("------------------------------------------------------")

	for i := 1; i < 5; i++ {
		country := mapNumSet[i]

		fmt.Printf("Country Id: %v  Country Name: %v\n",
			country.Id, country.Country)

	}

	fmt.Println("------------------------------------------------------")

}
