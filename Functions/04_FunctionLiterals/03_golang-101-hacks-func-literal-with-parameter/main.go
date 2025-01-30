/*
Taken from Nanxiao golang-101-hacks.
https://nanxiao.gitbooks.io/golang-101-hacks/content/posts/functional-literals.html

In Golang, lambda functions, anonymous functions and inline functions are known
as function literals.

A function literal just represents an anonymous function. In this example,
a parameter 'i' is passed to the function literal.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 3; i++ {

		go func(i int) { fmt.Println(i) }(i) 	// parameter 'i' is passed to
												// function literal.

		time.Sleep(time.Second * 2) // give the goroutine time to finish.
	}

}

/*	Output
	$ go run main.go
	1
	2
	3
*/
