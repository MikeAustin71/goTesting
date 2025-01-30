package main

import "fmt"

// This example demonstrates the use of the var keyword
// in declaring multiple variables and initializing those
// variables to different values.
func main() {
	var b, c int = 1, 2
	fmt.Printf("b type= %T - value= %v\n", b, b) // int - 1
	fmt.Printf("c type= %T - value= %v\n", c, c) // int - 2

}

/*	Output
	$ go run main.go
	b type= int - value= 1
	c type= int - value= 2
*/
