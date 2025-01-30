package main

import "fmt"

func main() {
	A := 8
	fmt.Println("A =", A)
	A = ^A
	fmt.Println("A = ^8 =", A)

	/* Output
		$ go run main.go
		A = 8
		A = ^8 = -9
	 */

}
