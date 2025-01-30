package main

import "fmt"

func main() {
	a := 3
	b := 12
	c := 25

	switch {
	case a == 1 && b == 10 && c == 23:
		fmt.Println("a=1, b=10, c= 23")
	case a == 2 && b == 11 && c == 24:
		fmt.Println("a=2, b=11, c=24")
	case a == 3 && b == 12 && c == 25:
		fmt.Println("a=3, b=12, c=25")
	case a == 4 && b == 13 && c == 16:
		fmt.Println("a=4, b=13, c=16")
	default:
		fmt.Println("Values of a, b and c are unknown")
	}

	fmt.Println("End of Program - Execution Terminated!")
}

/*	Output
	$ go run main.go
	a=3, b=12, c=25
	End of Program - Execution Terminated!
*/
