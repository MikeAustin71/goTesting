package main

import "fmt"

/*
	This is a standard switch statement. In
	Go, there is no default fall through from
	one case to the next. If fall through behavior
	is required it must be added manually.
*/

func main() {

	x := 2

	switch x {

	case 1:
		fmt.Println("x is equal to 1")
	case 2:
		fmt.Println("x is equal to 2")
	case 3:
		fmt.Println("x is equal to 3")
		return
	case 4:
		fmt.Println("x is equal to 4")
	case 5:
		fmt.Println("x is equal to 5")
	}

	fmt.Println("End of Program - Execution Terminated!")
}

/*	Output
	x is equal to 2
	End of Program - Execution Terminated!
*/
