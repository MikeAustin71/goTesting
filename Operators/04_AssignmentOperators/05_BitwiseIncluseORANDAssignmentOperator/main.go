package main

import "fmt"

func main() {
	// Bitwise Inclusive OR AND Assignment Operator: |=
	// C |= 13 is the same as C = C | 13
	var C = 60
	var A = 13

	C |= A

	fmt.Println("C = 60, A = 13")
	fmt.Println("C |= A is equivalent to C = C | A")
	fmt.Println("C |= A yields C =", C)
	// C = 60 (‭0011 1100‬) and A = 13 (‭0000 1101‬)
	// C |= A will yield C = 61, which is 0011 1101. ‬‬‬‬‬‬‬‬

}
