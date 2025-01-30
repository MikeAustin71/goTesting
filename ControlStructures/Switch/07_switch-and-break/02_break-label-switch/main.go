package main

import "fmt"

/*
	This example demonstrates the use a break
	statement inside a swtich block. The break
	statement uses a label to exit the for loop.
*/

func main() {

	var i = 0

EXITLOOP:

	for { // forever loop

		i++

		switch i {

		case 5:
			fmt.Println("Breaking out of loop!")
			break EXITLOOP // Break to Label EXITLOOP

		default:
			fmt.Println("i =", i)

		} // End of switch statement

	} // End of for loop

	fmt.Println("End of Program - Execution Terminated!")
	fmt.Println("Ending Value of i =", i)
}

/*	Output
	$ go run main.go
	i = 1
	i = 2
	i = 3
	i = 4
	Breaking out of loop!
	End of Program - Execution Terminated!
	Ending Value of i = 5

*/
