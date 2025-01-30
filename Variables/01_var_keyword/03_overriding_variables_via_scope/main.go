package main

import "fmt"

/*
	This example demonstrates the use of scope in
	overriding a variable with the same name. Notice
	that x has a different type and value depending
	on whether it is inside or outside of the inner
	function block.
*/

func main() {
	var x = 5
	fmt.Println("Here x is an integer =", x) //5
	fmt.Printf("x is of type %T\n", x)       // integer
	{
		// begin inner function block
		var x bool
		x = true
		fmt.Println("x is now a boolean =", x) // true
		fmt.Printf("x is now of type %T\n", x) // bool
		// end of inner function block
	}
	fmt.Println("Outside of the Inner Function Block")
	fmt.Println("x is an integer = ", x) // 5
	fmt.Printf("x is of type %T\n", x)   // integer
}

/*	Output
	$ go run main.go
	Here x is an integer = 5
	x is of type int
	x is now a boolean = true
	x is now of type bool
	Outside of the Inner Function Block
	x is an integer =  5
	x is of type int
*/
