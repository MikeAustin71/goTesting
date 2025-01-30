package main

import "fmt"

// This example ranges over a slice of ints
// and returns both the index and the value.

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, value := range s {
		fmt.Println("index=", i, " value=", value)
	}

}

/* Output
$ go run main.go
index= 0  value= 1
index= 1  value= 2
index= 2  value= 3
index= 3  value= 4
index= 4  value= 5
index= 5  value= 6
index= 6  value= 7
index= 7  value= 8
index= 8  value= 9
index= 9  value= 10
*/
