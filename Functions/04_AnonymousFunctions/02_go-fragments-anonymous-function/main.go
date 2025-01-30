/* Taken from Go Fragments
http://www.gofragments.net/client/blog/fundamentals/2015/09/09/closureFuncAnExample/index.html
*/

/* gofragments.net, example: 'closureFuncAnExample.go' */

/*
Closures, briefly, are functions that carry along with them variables declared within lexical scope of the defined function.
So, what does this mean? The variable 'foo' is in scope for the function that we declare and return from a Func().
When we return our function, foo stays in scope, even though the function it was declared in has returned.
*/

package main

import "fmt"

type generator func() int

func aFunc() generator {
	foo := 0

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
