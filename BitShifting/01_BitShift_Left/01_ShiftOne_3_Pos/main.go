package main

import "fmt"

func main() {
	A := 0
	A = 1
	fmt.Println("A=", A)
	A = A << 3
	fmt.Println("A << 3 =", A)
	/* Output
		$ go run main.go
		A= 1
		A << 3 = 8
		Shifting a value of 1 three bit positions
		to the left will yield a value of 8.
	*/
}
