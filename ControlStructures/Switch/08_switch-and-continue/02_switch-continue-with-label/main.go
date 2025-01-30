package main

import "fmt"

/*
	This example employs a continue statement
	with label inside of a switch block. The
	switch statement is configured inside of
	two for loops - an inner and an outer for
	loop.
*/

func main() {

OuterLoop:
	for i := 0; i < 3; i++ {

		for j := 0; j < 7; j++ {

			switch j {

			case 2:
				continue OuterLoop
			default:
				fmt.Println("i=", i, "  j=", j)

			}// End of switch block
		} // End of j Loop
	} // End of i Loop

	fmt.Println("End of Program - Execution Terminated!")
}

/* Output
	$ go run main.go
	i= 0   j= 0
	i= 0   j= 1
	i= 1   j= 0
	i= 1   j= 1
	i= 2   j= 0
	i= 2   j= 1
	End of Program - Execution Terminated!
*/
