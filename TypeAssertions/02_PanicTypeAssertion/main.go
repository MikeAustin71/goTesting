package main

import "fmt"

// Taken from a Tour of Go

func main() {

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)


	f = i.(float64) // panic
	fmt.Println(f)

}
