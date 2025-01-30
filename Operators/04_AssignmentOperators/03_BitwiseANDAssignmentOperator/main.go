package main

import "fmt"

func main() {
	// Bitwise AND (&) Assignment Operator
	// C &= 13 is the same as C = C & 13
	var C = 60
	var A = 13

	C &= A

	fmt.Println("C = 60, A = 13")
	fmt.Println("C &= A is equivalent to C = C & A")
	fmt.Println("C &= A = 12")

}
