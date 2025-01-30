package main

import "fmt"

/*	This example was taken  from:
	Golang Go - YouTube - Go Golang Function Decorator Example
	https://www.youtube.com/watch?v=eGC0xu6Ohac

	A decorator is just a function that takes another function
	In most cases they return a function that is a modified version
	of the function they are wrapping.
*/
func decorator(f func(s string)) func(s string) {

	return func(s string) {
		fmt.Println("Start")
		f(s)
		fmt.Println("End")
	}
}

func myfunc(s string) {
	fmt.Println(s)
}

func main() {

	// Prints address of wrapper function
	// ==decorator== 0x4012d0
	fmt.Println("==decorator==", decorator(myfunc))

	// fun1 := decorator(myfunc)
	// fun1("Hi")

	// We can replace the two commented lines
	// above with the following line.
	decorator(myfunc)("Hello")

	/* This line Prints:
	Start
	Hello
	End
	*/
}

/*	Output
	$ go run main.go
	==decorator== 0x4012d0
	Start
	Hello
	End
*/
