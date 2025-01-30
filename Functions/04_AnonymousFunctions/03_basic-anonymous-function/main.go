package main

import "fmt"

/*
	An Anonymous function is a function which has no name.
	This is a self-executing anonymous function which executes
	by means of the trailing () parentheses.
*/

func main() {

	func() {
		fmt.Println("Hello")
	}() // In order to execute this anonymous function
	// you need the following () parentheses

}

/*	Output
	$ go run main.go
	Hello
*/
