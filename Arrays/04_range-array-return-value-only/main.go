package main

import "fmt"

/*
	This example ranges over an array and
	extracts only the value v.
*/

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range nums {
		fmt.Println("value =", v)
	}
}

/*	Output
	$ go run main.go
	value = 1
	value = 2
	value = 3
	value = 4
	value = 5
	value = 6
	value = 7
	value = 8
	value = 9
	value = 10
*/
