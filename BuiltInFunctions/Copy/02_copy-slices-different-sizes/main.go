package main

import "fmt"

/*
	This example demonstrates the built-in function copy
	when used to copy slices where the destination length
	is less than the source length.
*/

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3)
	copy(slice2, slice1)
	fmt.Println("slice1 =", slice1)
	fmt.Println("slice2 =", slice2)
}

/*	Output
	$ go run main.go
	slice1 = [1 2 3 4 5]
	slice2 = [1 2 3]

*/
