package main

import "fmt"

/*
	This example shows the use of continue
	statements inside of a switch block.
*/

func main() {

	for i := 0; i < 10; i++ {

		switch i {
		case 3, 4:
			continue // return to top of for loop
		case 7, 8:
			continue // return to top of for loop
		default:
			fmt.Println("i=", i)

		} // End of switch statement

	} // End of i loop

	fmt.Println("End Of Program - Execution Terminated!")
}

/*	Output
	$ go run main.go
	i= 0
	i= 1
	i= 2
	i= 5
	i= 6
	i= 9
	End Of Program - Execution Terminated!
*/
