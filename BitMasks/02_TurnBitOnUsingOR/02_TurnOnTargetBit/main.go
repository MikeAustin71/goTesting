package main

import "fmt"

/*
	This example is designed to Turn On Bit # 2
	using the Logical OR (|) Operator.
*/
func main() {

	A := 97
	fmt.Println("A =", A)
	// Create the Bit Mask
	// Start with a value of one
	B := 1

	// Shift value of 1 left
	// two bit positions.
	B = B << 2
	fmt.Println("B = ", B) // B = 4

	// OR the original value
	// and the Bit Mask
	result := A | B

	fmt.Println("A | B =", result)
	//	The result is a value of 101.
	//  Bit # 2 has been turned on.

	/*	Output
		$ go run main.go
			A = 97
			B =  4
			A & B = 101
	*/
}
