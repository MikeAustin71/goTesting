package main

import "fmt"

/*
	This example ranges over an array and
	extracts both the index (i) and the
	value (v)
*/
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range nums {
		fmt.Println("index=", i, " value=", v)
	}
}

/*	Output
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
