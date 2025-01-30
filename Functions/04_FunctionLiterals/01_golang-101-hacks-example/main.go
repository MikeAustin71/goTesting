/*
Taken from Nanxiao golang-101-hacks.
https://nanxiao.gitbooks.io/golang-101-hacks/content/posts/functional-literals.html

In Golang, lambda functions, anonymous functions and inline functions are known
as function literals.

A function literal just represents an anonymous function. You can assign
a functional literal to a variable. In this example, the anonymous function
is invoked indirectly, by means of the variable.

*/

package main

import (
	"fmt"
)

func main() {
	f := func() { fmt.Println("Hello, 中国！") }
	f()
}

/*	Output
	$ go run main.go
	Hello, 中国！
*/
