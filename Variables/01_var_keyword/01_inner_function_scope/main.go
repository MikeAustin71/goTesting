package main

import "fmt"

// variable with package level scope
var x int

func main() {
	x = 2
	fmt.Println("x initialized to ", x) // 2
	// function block begins
	var y int
	fmt.Println("y initialzied to", y)

	{
		// inner block begins
		x = 3
		fmt.Println("value of x changed to", x) // 3
		var z int

		z = x + y			 // z = 3 + 0
		fmt.Println("z =", z) // 3
		// inner block ends
	}

	// z no longer exists
	y = x
	fmt.Println("y = x, y =", y) // 3

	// function block ends
}

/*	Output
	$ go run main.go
	x initialized to  2
	y initialzied to 0
	value of x changed to 3
	z = 3
	y = x, y = 3
 */
