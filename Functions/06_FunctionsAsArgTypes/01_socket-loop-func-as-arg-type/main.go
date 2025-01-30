/*
Taken from Socket Loop Golang : Function as an argument type example
https://socketloop.com/tutorials/golang-function-as-an-argument-type-example

This example demonstrates the function as argument type in that the function
literal 'strToInt' is passed as a parameter to the function 'OuterFunc'.
*/

package main

import (
	"fmt"
	"strconv"
)

var a string = "1"

// OuterFunc receives receives two parameters,
// a function parameter and an integer parameter.
func OuterFunc(strToInt func(s string) int, b int) string {
	c := strToInt(a) + b
	return strconv.Itoa(c)
}

func main() {

	strToInt := func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	}

	result := OuterFunc(strToInt, 2)
	fmt.Println(result)                 // Result = 3
	fmt.Println(OuterFunc(strToInt, 4)) // Result = 5
}

/*	Output
	$ go run main.go
	3
	5
*/
