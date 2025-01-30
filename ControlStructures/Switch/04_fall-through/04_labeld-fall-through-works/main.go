package main

import "fmt"

/*
	This example demonstrates the labeled
	'fallthrough' statement. The 'fallthrough'
	statement must be the last thing in the case.
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
			goto nextCase // Goto or jump to label
		}
		fmt.Println("a=5")
	nextCase:
		fallthrough  // fallthrough must be last thing in case
	case c == 20:
		fmt.Println("c=", c) // prints c= 30
	case c == 99:
		fmt.Println("c=99")
	}

	fmt.Println("End of Program - Execution Terminated!")
}

/*	Output
	$ go run main.go
	c= 30
	End of Program - Execution Terminated!
*/
