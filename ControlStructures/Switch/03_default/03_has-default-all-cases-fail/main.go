package main

import "fmt"

/*
This example demonstrates the most common use of the
default case. In Go the 'default' case is NOT added
automatically - it must be added manually.

If all cases fail and default case is present, the code
under the default case will be executed as shown below.
*/

func main() {
	x := 5

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
	default:
		fmt.Println("default - no cases evaluated as true!")
		return
	}

	fmt.Println("Completed switch #1 - no cases triggered!")
}

/*	Output
	$ go run main.go
	default - no cases evaluated as true!
*/
