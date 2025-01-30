package main

import "fmt"

// This example ranges over a slice of ints
// and returns only the index i.

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range s {
		fmt.Println("index=", i)
	}

}

/* Output
$ go run main.go
index= 0
index= 1
index= 2
index= 3
index= 4
index= 5
index= 6
index= 7
index= 8
index= 9
*/
