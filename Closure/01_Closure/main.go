package main

import "fmt"

/*
	This closure example adapted from:
	http://www.gofragments.net/client/blog/fundamentals/2015/09/09/closureFuncAnExample/index.html
*/
type generator func() int

func aFunc() generator {
	foo := 0
	// the returned function of type generator is a closure, it encompasses the
	// var foo.
	return func() int {
		foo++
		return foo
	}
}

func main() {
	bar := aFunc()
	fmt.Println(bar())     // 1
	fmt.Println(bar())     // 2
	fmt.Println(bar() - 3) // 0
}

/*	Output
	$ go run main.go
	1
	2
	0
*/
