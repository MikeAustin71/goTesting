package main

import "fmt"

func main() {
	// This example will generate a Truth
	// Table for the AND NOT (&^) Operator

	fmt.Println("0 &^ 0 =", 0&^0) // 0
	fmt.Println("1 &^ 0 =", 1&^0) // 1
	fmt.Println("0 &^ 1 =", 0&^1) // 0
	fmt.Println("1 &^ 1 =", 1&^1) // 0

	/*	Output
		$ go run main.go
		0 &^ 0 = 0
		1 &^ 0 = 1
		0 &^ 1 = 0
		1 &^ 1 = 0
	*/
}
