package main

/*
Taken from Nanxiao golang-101-hacks.
https://nanxiao.gitbooks.io/golang-101-hacks/content/posts/functional-literals.html

In Golang, lambda functions, anonymous functions and inline functions are known
as function literals.

A function literal just represents an anonymous function. In this example,
the function literal is invoked directly by means of the () parentheses
located immediately after the function literal. The () parentheses causes
the function literal to execute immediately.

*/

import (
	"fmt"
)

func main() {
	func() { fmt.Println("Hello, 中国！") }()  // () triggers immediate execution
}

/*	Output
	$ go run main.go
	Hello, 中国！
*/
