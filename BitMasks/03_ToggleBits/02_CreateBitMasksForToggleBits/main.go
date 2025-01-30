package main

import "fmt"

/*
	The objective of this example is to toggle the bit #'s
	3 and 5 on target integer A = 141. This operation will
	utilize the Logical Exclusive OR Operator XOR (^) and
	a bit mask consisting of all zeros except for Bit #'s
	3 and 5 which will be Turned On with a value of 1.
*/

// Step-1 Create a series of bit masks
// consisting of all zeros and a value
// of 1 in the target bit position 0-15
const (
	maskBit0 = 1 << iota
	maskBit1
	maskBit2
	maskBit3
	maskBit4
	maskBit5
	maskBit6
	maskBit7
	maskBit8
	maskBit9
	maskBit10
	maskBit11
	maskBit12
	maskBit13
	maskBit14
	maskBit15
)

func main() {
	fmt.Println("maskBit0 =", maskBit0)
	fmt.Println("maskBit1 =", maskBit1)
	fmt.Println("maskBit2 =", maskBit2)
	fmt.Println("maskBit3 =", maskBit3)
	fmt.Println("maskBit4 =", maskBit4)
	fmt.Println("maskBit5 =", maskBit5)
	fmt.Println("maskBit6 =", maskBit6)
	fmt.Println("maskBit7 =", maskBit7)
	fmt.Println("maskBit8 =", maskBit8)
	fmt.Println("maskBit9 =", maskBit9)
	fmt.Println("maskBit10 =", maskBit10)
	fmt.Println("maskBit11 =", maskBit11)
	fmt.Println("maskBit12 =", maskBit12)
	fmt.Println("maskBit13 =", maskBit13)
	fmt.Println("maskBit14 =", maskBit14)
	fmt.Println("maskBit15 =", maskBit15)
	fmt.Println("============================")
	A := 141
	fmt.Println("Original Value A =", A)

	// Step-2 Create a Bit Mask to toggle bits 3 and 5
	B := 0                        // Default all bits to zero
	B = (B | maskBit3 | maskBit5) // Turn On Bits 3 and 5

	fmt.Println("Bit Mask B =", B) // B = 40

	// Step-3 Apply the XOR (^) Operator to toggle
	// target bits.
	fmt.Println("Result  A ^ B =", A^B) // Result = 165
	// Original Value A = 141, Bit Mask B = 40
	// Result A ^ B = 165 where Bits 3 and 5 have been
	// toggled.

	/*	Output
		$ go run main.go
		maskBit0 = 1
		maskBit1 = 2
		maskBit2 = 4
		maskBit3 = 8
		maskBit4 = 16
		maskBit5 = 32
		maskBit6 = 64
		maskBit7 = 128
		maskBit8 = 256
		maskBit9 = 512
		maskBit10 = 1024
		maskBit11 = 2048
		maskBit12 = 4096
		maskBit13 = 8192
		maskBit14 = 16384
		maskBit15 = 32768
		============================
		Original Value A = 141
		Bit Mask B = 40
		Result  A ^ B = 165
	*/
}
