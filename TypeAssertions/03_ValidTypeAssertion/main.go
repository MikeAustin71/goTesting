package main

import "fmt"

// Adapted from a Tour of Go
// This example demonstrates valid
// type assertions.

func main() {

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

}

/*	Output
	$ go run main.go
	hello
	hello true
	0 false
*/