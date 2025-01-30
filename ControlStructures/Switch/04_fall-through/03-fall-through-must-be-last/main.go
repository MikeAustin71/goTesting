package main

import "fmt"

/*
	This example demonstrates the wrong way to
	use a 'fallthrough' statement. The 'fallthrough'
	statement must be the last thing in the case.

	This code will issue a runtime error.
*/

func main() {
	a := 5
	b := 10
	c := 20

	switch {
	case a == 3:
		fmt.Println("a=3")
	case a == 5:
		if b == 10 {
			c = 30
			fallthrough // Error - fallthrough must be last thing in case
		}
		fmt.Println("a=5")
	case c == 20:
		fmt.Println("c=20")
	}

	fmt.Println("End of Program - Execution Terminated!")
}

/*	Output
	$ go run main.go
	# command-line-arguments
	.\main.go:16: fallthrough statement out of place

*/
