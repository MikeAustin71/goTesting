
/*
Taken from Socket Loop - Golang : On lambda, anonymous, inline functions and function literals.
https://socketloop.com/tutorials/golang-on-lambda-anonymous-inline-functions-and-function-literals

An anonymous function is a function without a name. In Golang, lambda, anonymous
and inline functions are known as function literals. What it does is to allow
you to create a short callback function but without creating a separate named
function. In a nutshell, it allows you to specify the one-line function
"in-line" a.k.a. on the fly...

This is an example of an anonymous function (a.k.a. function literal) assigned
to a variable. A 'func expression' sets a variable equal to a function.
*/

package main

import (
	"fmt"
)

func main() {
	add := func(a ,b int) int { return a + b }  // func expression
	fmt.Println(add(1,2))
}

/* 	Output
	$ go run main.go
	3
*/