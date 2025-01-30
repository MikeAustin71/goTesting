package main

import "fmt"

/*
	This code taken from Socket Loop
	Golang builtin.copy() function example.
	https://www.socketloop.com/references/golang-builtin-copy-function-example

	package builtin
	The copy built-in function copies elements from a source slice into a destination slice.
*/

func main() {
	source := []string{"Hi", "Hello", "Hey"}

	fmt.Println("Source : ", source[:])

	destination := make([]string, 5)

	fmt.Println("Destination before copy :", destination[:])

	num := copy(destination, source[:])

	fmt.Println("Destination AFTER copy :", destination[:])

	fmt.Println("Element(s) copied : ", num)
}

/*	Output
	$ go run main.go
	Source :  [Hi Hello Hey]
	Destination before copy : [    ]
	Destination AFTER copy : [Hi Hello Hey  ]
	Element(s) copied :  3
*/
