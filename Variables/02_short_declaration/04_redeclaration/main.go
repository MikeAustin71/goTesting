package main

import "fmt"

/*
	This example demonstrates variable
	re-declaration using the short variable
	declaration operator.
*/

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
