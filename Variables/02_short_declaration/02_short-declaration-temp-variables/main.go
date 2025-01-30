package main

import "fmt"

func main() {
	// example of short variable declaration
	// for local temporary variables.
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

/*	Output
	$ go run main.go
	0
	1
	2
	3
	4
	5
	6
	7
	8
	9
	10
*/
