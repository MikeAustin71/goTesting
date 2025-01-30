/*
Taken from Nanxiao golang-101-hacks.
https://nanxiao.gitbooks.io/golang-101-hacks/content/posts/functional-literals.html

In Golang, lambda functions, anonymous functions and inline functions are known
as function literals.

A function literal just represents an anonymous function. This example
demonstrates that a function literal is also a closure. In this case, the
function inherits the 'i' parameter from the parent function.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 3; i++ {

		// Function literal is also a closure,
		// so it can access the variables of its
		// surrounding function. Thanks to closure,
		// 'i' is inherited from the parent function.
		// No parameter required.
		go func() { fmt.Println(i) }()

		time.Sleep(time.Second * 2) // give the goroutine time to finish.
	}

}

/*	Output
	$ go run main.go
	1
	2
	3
*/
