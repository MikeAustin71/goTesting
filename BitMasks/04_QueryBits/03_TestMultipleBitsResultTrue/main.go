package main

import "fmt"

/*
	The object of this example is to test an
	original value of 933 to determine if Bit
	#'s 7, 5 and 2 were set. This test will be
	accomplished by AND'ing the original value A
	by the Bit Mask B.
*/
func main() {
	A := 933
	fmt.Println("Original Value A =", A)
	B := 164
	fmt.Println("Bit Mask B =", B)
	Result := A & B
	fmt.Println("Result = A & B =", Result)
	fmt.Println("In Value 'A', Are Bits 7, 5 and 2 Set? Answer:", B == Result)
	// If Result == Bit Mask it signals True, Bit #'s 7, 5 and 2
	// were set equal to one.

	/*	Output
		$ go run main.go
		Original Value A = 933
		Bit Mask B = 164
		Result = A & B = 164
		In Value 'A', Are Bits 7, 5 and 2 Set? Answer: true
	*/

}
