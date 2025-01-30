package main

import "fmt"

/*  This example shows a switch statement with NO default
case where one case succeeds.
*/

func main() {
	x := 1

	switch x {
	case 1:
		fmt.Println("x=1")
		return
	case 2:
		fmt.Println("x=2")
		return
	case 3:
		fmt.Println("x=3")
		return
	}

	fmt.Println("Completed switch #1 - no cases triggered!")
}

/*	Output
	$ go run main.go
	x=1
*/
