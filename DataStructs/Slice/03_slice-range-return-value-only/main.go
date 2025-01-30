package main

import "fmt"

// This example ranges over a slice of ints
// and returns only the value v. Note the use
// of the Blank Identifier underscore (_)

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range s {
		fmt.Println("value=", v)
	}

}

/*	Output
	$ go run main.go
	value= 1
	value= 2
	value= 3
	value= 4
	value= 5
	value= 6
	value= 7
	value= 8
	value= 9
	value= 10
*/
