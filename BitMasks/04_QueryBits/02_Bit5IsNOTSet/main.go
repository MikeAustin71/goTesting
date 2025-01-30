package main

import "fmt"

/*
	The objective of this example is to query Bit #5
	in the original value A=133. The Bit Mask is
	B = 32. 'A' will be AND'd with the Bit Mask 'B' to
	determine if Bit # 5 is set (equal to one).
	If the result is non-zero, it signals that Bit #5
	is set (equal to 1). If the result is zero, it
	signals that Bit #5 is NOT set (and is equal to zero).
*/
func main() {
	A := 133
	fmt.Println("The Original Value A =", A)
	// Creating the Bit Mask by
	// Left Shifting a value of 1
	// 5-bit positions. B = 32
	B := 1 << 5
	fmt.Println("The Bit Mask B =", B)
	fmt.Println("A & B =", A&B) // Result = 0
	// A zero result signals that Bit#5 is NOT set

	/*	Output
		$ go run main.go
		The Original Value A = 133
		The Bit Mask B = 32
		A & B = 0
	*/
}
