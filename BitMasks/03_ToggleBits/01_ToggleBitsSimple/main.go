package main

import "fmt"

/*
	This example demonstrates the 'toggle bit'
	technique using The Bitwise Exclusive OR Operator
	XOR (^)
*/
func main() {
	A := 141
	fmt.Println("Original Value A =", A)
	B := 40
	fmt.Println("Bit Map B =", B)
	fmt.Println("Result A ^ B =", A^B) // Yields 165
	// Bit Numbers 5 and 3 have been toggled in the
	// final value (165).

	/*	Output
		$ go run main.go
		Original Value A = 141
		Bit Map B = 40
		Result A ^ B = 165
	*/

}
