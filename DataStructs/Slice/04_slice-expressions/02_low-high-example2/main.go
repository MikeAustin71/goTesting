package main

import "fmt"

/*
	For a slice, the primary expression is
	a[low:high]. The output of a[low:high]
	starts as index 'low' and continues for
	a length of high minus low.
*/

func main() {
	s := [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	a := s[1:4]
	fmt.Println("a=", a)
	b := s[5:10]
	fmt.Println("b=", b)
	c := s[14:15]
	fmt.Println("c=", c)
	// d := s[15:14] // Generates invalid slice index: 15 > 14 error!
	// fmt.Println("d=",d)
}

/*	Output
	$ go run main.go
	a= [2 3 4]
	b= [6 7 8 9 10]
	c= [15]
*/
