package main

import "fmt"

/*
	This example demonstrates use of the break statement
	inside a switch statement. break is not required for
	switch case statements, but there are situations where
	break can prove useful. Also, notice the use of continue
	statement within the switch.
*/

func main() {

	for i := 0; i < 10; i++ {

		switch i {

		case 7:
			fmt.Println("i=7 - break out of switch")
			break // break out of switch statement

		default:
			fmt.Println("i=", i)
			continue

		}

		fmt.Println("2nd break - breaking out of for loop")
		break // break out of for loop

	}

	fmt.Println("End of Program - Execution Terminated!")

}

/*	Output
	$ go run main.go
	i= 0
	i= 1
	i= 2
	i= 3
	i= 4
	i= 5
	i= 6
	i=7 - break out of switch
	2nd break - breaking out of for loop
	End of Program - Execution Terminated!
*/
