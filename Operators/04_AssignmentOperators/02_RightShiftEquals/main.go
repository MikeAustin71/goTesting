package main

import "fmt"

func main() {
	// Right Shift Assignment Operator >>=
	// C >>= A is equivalent to C = C >> A
	// or in this case C = 16 >> 2
	// Note: Shift Count 'A' must be unsigned integer
	C := 16
	A := uint16(2)
	C >>= A
	fmt.Println("C = 16 and A = unint16(2).")
	fmt.Println("C >>= A is equal to ", C)
}
