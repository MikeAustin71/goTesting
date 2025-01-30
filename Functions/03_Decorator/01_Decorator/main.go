package main

import "fmt"

/*	This example was taken  from:
	Golang Go - YouTube - Go Golang function decorator
	https://www.youtube.com/watch?v=JlRrKyLl8gU

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

	// These two statements work and produce the expected result.
	fun1 := decorator(myfunc)

	fun1("Hi")
	/* This line Prints:
	Start
	Hi
	End
	*/

}

/*	Output
	$ go run main.go
	==decorator== 0x4012d0
	Start
	Hi
	End
*/
