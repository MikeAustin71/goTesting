/*
In Go, functions may return multiple values. This code
demonstrates the ability to ignore specified function
return values.

'returnTwoVars' returns a string and an int. When called
from function 'main' below the int return value is ignored
using the blank identifier symbol, the underscore character (_).
*/

package main

import (
	"fmt"
)

func returnTwoVars() (string, int) {
	return "Hello World", 7
}

func main() {
	a, _ := returnTwoVars()
	fmt.Println("a=", a)
	fmt.Println(`Second Value ignored using Blank Identifier"_"`)
}

/*	Output
	$ go run main.go
	a= Hello World
	Second Value ignored using Blank Identifier"_"
*/
