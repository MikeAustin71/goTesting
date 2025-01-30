package main

import "fmt"
// Uses the AND NOT Operator to
// turn off or clear specific bits
// in 'A'.  The expected result of
// 141 &^ 12 = 129
func main() {
	A := 141
	B := 12

	result := A &^ B

	fmt.Println("141 &^ 12 =", result)
}
