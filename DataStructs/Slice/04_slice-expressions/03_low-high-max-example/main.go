package main

import "fmt"

/*
	For a slice, the full expression is
	formatted as a[low:high:max]. 'low' is
	the starting index. The length is equal to
	'high' minus 'low'. The capacity is equal
	to 'max' minus 'low'.
*/

func main() {
	s := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	a := s[1:8:20]
	fmt.Println("a=", a)
	fmt.Println("Length of a =", len(a))
	fmt.Println("Capacity of a =", cap(a))
}

/*	Output
	$ go run main.go
	a= [2 3 4 5 6 7 8]
	Length of a = 7
	Capacity of a = 19
*/
