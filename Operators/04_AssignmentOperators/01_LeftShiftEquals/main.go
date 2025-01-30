package main

import "fmt"

func main() {
	// C <<= A is equivalent to C = C << A
	// or in this case C = 4 << 2
	// Note: Shift Count 'A' must be unsigned integer
	C := 4
	A := uint16(2)
	C <<= A
	fmt.Println("C = 4 and A = unint16(2).")
	fmt.Println("C <<= A is equal to ", C)
}
