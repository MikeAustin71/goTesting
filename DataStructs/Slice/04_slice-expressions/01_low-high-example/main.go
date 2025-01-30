package main

import "fmt"

/*
	For a slice, the primary expression is
	a[low:high]. The output of a[low:high]
	starts as index 'low' and continues for
	a length of high minus low.
*/

func main() {

	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:4]
	fmt.Println(s)
}
/* 	Output
	$ go run main.go
	[2 3 4]
*/
