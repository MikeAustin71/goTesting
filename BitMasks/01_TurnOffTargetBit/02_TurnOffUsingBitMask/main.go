package main

import "fmt"

/*
	This example is designed to turn off bit # 3 in a target
	integer using the Logical Operator AND (&)
*/
const (
	maskBit0 = -1<<iota - 1
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
	fmt.Println("----------------------")
	A := 141
	fmt.Println("Original Value A =", A)
	fmt.Println("Bit Mask maskBit3 =", maskBit3)
	fmt.Println("A & maskBit3 =", A&maskBit3)
	// 141 & -9 = 133. The result 133 means
	// that bit # 3 has been turned off.

	/* Output
	$ go run main.go
	maskBit0 = -2
	maskBit1 = -3
	maskBit2 = -5
	maskBit3 = -9
	maskBit4 = -17
	maskBit5 = -33
	maskBit6 = -65
	maskBit7 = -129
	maskBit8 = -257
	maskBit9 = -513
	maskBit10 = -1025
	maskBit11 = -2049
	maskBit12 = -4097
	maskBit13 = -8193
	maskBit14 = -16385
	maskBit15 = -32769
	----------------------
	Original Value A = 141
	Bit Mask maskBit3 = -9
	A & maskBit3 = 133
	*/

}
