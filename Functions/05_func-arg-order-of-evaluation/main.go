/*
Taken from Kuree Gitbooks - The Go Programming Language Report
Evaluation Order of Function Parameters in Go
https://kuree.gitbooks.io/the-go-programming-language-report/content/15/text.html

This example is designed to demonstrate that the Go Language evaluates
function parameters from left to right.

*/

package main

import "fmt"

func arg1() int {
	fmt.Println("arg1 gets called")
	return 1
}

func arg2() int {
	fmt.Println("arg2 gets called")
	return 2
}

func arg3() int {
	fmt.Println("arg3 gets called")
	return 3
}

func foo(a int, b int, c int) {
	return
}

func main() {
	foo(arg1(), arg2(), arg3())
}

/*	Output
	$ go run main.go
	arg1 gets called
	arg2 gets called
	arg3 gets called
*/
