package main

import "fmt"

/*
	An Anonymous function is a function which has no name.
	This is a self-executing anonymous function which executes
	by means of the trailing () parentheses.

	Note: This anonymous function accepts parameters.
*/

func main() {

	var msg string = "Hello World"

	func(message string) {
		fmt.Println(message)
	}(msg) // In order to execute this anonymous function
	// you need the following () parentheses
	// Note that the varialbe 'msg' is passed to this
	// anonymous function.

}

/*	Output
	$ go run main.go
	Hello World
*/
