package main

import "fmt"

/*
	This example adapted from byte-sized Tuts:
	https://gist.github.com/honkskillet/bd1f72223dd8e06b5ce6#golang-12---closure
*/

func myCounter() func() {
	theCount := 0
	increment := func() {
		theCount++
		fmt.Println("The count is", theCount)
	}
	return increment
}

func main() {
	var incFunc func()
	incFunc = myCounter()
	for i := 0; i < 5; i++ {
		incFunc() //use () to execute increment
	}
}

/*	Output
	$ go run main.go
	The count is 1
	The count is 2
	The count is 3
	The count is 4
	The count is 5
*/
