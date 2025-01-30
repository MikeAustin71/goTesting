package main

import "fmt"

/*
	The object of this example is to test an original value A = 773
	to determine if all of Bit #'s 7, 5 and 2 were set. This test
	will be accomplished by AND'ing the original value A by the
	Bit Mask B = 164.
*/

func main() {
	A := 773
	fmt.Println("Original Value A =", A)
	B := 164
	fmt.Println("Bit Mask B =", B)
	Result := A & B
	fmt.Println("Result = A & B =", Result)
	fmt.Println("In Value 'A', Are All of Bits 7, 5 and 2 Set? Answer:", B == Result)
	// If Result == Bit Mask it signals True, Bit #'s 7, 5 and 2
	// were set equal to one.
	//
	// If Result == 0 it signals False, none of the bits,
	// Bit #'s 7, 5 and 2 were set equal to one.
	//
	// If Result > 0 AND Result < Bit Mask, it signals that
	// at least some, but not all of the sampled bits were
	// set equal to one.

	/*	Output
		$ go run main.go
		Original Value A = 773
		Bit Mask B = 164
		Result = A & B = 4
		In Value 'A', Are All of Bits 7, 5 and 2 Set? Answer: false
	*/

}
