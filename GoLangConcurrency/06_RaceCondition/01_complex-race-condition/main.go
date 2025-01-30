// Source Code: Ardan Labs - Data Race - Example1.go
// https://github.com/ardanlabs/gotraining/blob/master/topics/data_race/advanced/example1/example1.go
// Sample program to show a more complicated race condition using
// an interface value. This produces a read to an interface value after
// a partial write.
// **************** THIS CODE WILL NOT EXECUTE !!!! ***************************
package main

import "fmt"

// Speaker allows for speaking behavior.
type Speaker interface {
	Speak() bool
}

// Ben is a person who can speak.
type Ben struct {
	name string
}

// Speak allows Ben to say hello. It returns false if the method is
// called through the interface value after a partial write.
func (b *Ben) Speak() bool {
	if b.name != "Ben" {
		fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
		return false
	}

	return true
}

// Jerry is a person who can speak.
type Jerry struct {
	name string
}

// Speak allows Jerry to say hello. It returns false if the method is
// called through the interface value after a partial write.
func (j *Jerry) Speak() bool {
	if j.name != "Jerry" {
		fmt.Printf("Jerry says, \"Hello my name is %s\"\n", j.name)
		return false
	}

	return true
}

// main is the entry point for all Go programs.
func main() {
	// Create values of type Ben and Jerry.
	ben := Ben{"Ben"}
	jerry := Jerry{"Jerry"}
	// Assign the pointer to the Ben value to the interface value.
	person := Speaker(&ben)

	// Have a goroutine constantly assign the pointer of
	// the Ben value to the interface.
	go func() {
		for {
			person = &ben
		}
	}()

	// Have a goroutine constantly assign the pointer of
	// the Jerry value to the interface.
	go func() {
		for {
			person = &jerry
		}
	}()
	// Keep calling the Speak method against the interface
	// value until we have a race condition.
	for {
		if !person.Speak() {
			break
		}
	}
}

/*	Output
	$ go run -race main.go
	...
	Ben says, "Hello my name is Jerry"
	Found 2 data race(s)
	exit status 66
*/
