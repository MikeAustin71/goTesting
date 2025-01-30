package main

import "fmt"

/*
	This switch statement includes fallthrough
	statements for cases 1-3. In Go, there is
	no default fall through from one case to the next.
	If fall through behavior is required it must be
	added manually as shown below.
*/

func main() {

	x := 1

	switch x {

	case 1:
		fmt.Println("case 1 executed")
		fallthrough
	case 2:
		fmt.Println("case 2 executed")
		fallthrough
	case 3:
		fmt.Println("case 3 executed")
		fallthrough
	case 4:
		fmt.Println("case 4 executed")
	case 5:
		fmt.Println("case 5 executed")
	}

	fmt.Println("End of Program - Execution Terminated!")
}

/*	Output
	$ go run main.go
	case 1 executed
	case 2 executed
	case 3 executed
	case 4 executed
	End of Program - Execution Terminated!
*/
