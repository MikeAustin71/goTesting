package main

/*
	This example demonstrates a variable with
	package level scope. var x declared at a
	package level can be accessed from anywhere
	in the package.
*/

import "fmt"

// variable with package scope
var x int

func main() {
	x = 3                              // x initialized to value of 3
	fmt.Println("x initialized to", x) // 3

	changeX()

	fmt.Println("After changeX() the new value of x =", x) // 5
}

func changeX() {
	x = 5

}

/*	Output
	$ go run main.go
	x initialized to 3
	After changeX() the new value of x = 5
*/