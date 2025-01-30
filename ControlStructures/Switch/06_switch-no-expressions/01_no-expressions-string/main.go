package main

import "fmt"

/*
	switch statements do not require expressions.
	In this example, a switch statement will trigger
	the appropriate case without switch expression.
*/

func main() {
	x := "Some string"

	switch {
	case x == "fish":
		fmt.Println("x = fish")
	case x == "bird":
		fmt.Println("x = bird")
	case x != "Some string":
		fmt.Println("x is NOT equal to 'Some string'")
	case len(x) == 11:
		fmt.Println("x has a length of 11 x=", x)
	default:
		fmt.Println("Unknown")
	}

	fmt.Println("End of Program - Execution Terminated!")
}

/*	Output
	$ go run main.go
	x has a length of 11 x= Some string
	End of Program - Execution Terminated!
*/
