package main

import "fmt"

// This is a simple switch example.

func main() {
	var x = 2

	switch x {

	case 1:
		fmt.Println("x is equal to 1")
	case 2:
		fmt.Println("x is equal to 2")
	case 3:
		fmt.Println("x is equal to 3")
	case 4:
		fmt.Println("x is equal to 4")
	case 5:
		fmt.Println("x is equal to 5")
	}

	fmt.Println("End of Program - Execution Terminated!")

}

/*	Output
	$ go run main.go
	x is equal to 2
	End of Program - Execution Terminated!
*/
