package main

import "fmt"

func main() {
	A := 141
	fmt.Println("A =", A)

	B := 0
	B = 1
	B = B << 3
	fmt.Println("B =", B)
	B = ^B
	fmt.Println("^B =", B)

	C := A & B
	fmt.Println("A & B =", C)
	/* Output
		$ go run main.go
		A = 141
		B = 8
		A & B = 8
	 */
}
