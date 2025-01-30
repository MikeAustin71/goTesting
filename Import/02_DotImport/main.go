package main

/*
	This example uses the dot import
	declaration to import the 'math'
	package. By using the dot syntax,
	all exported fields and methods in
	the math package can be accessed directly.
	There is no need for the usual prefix, 'math.'
	That is, we no longer reference Pi as math.Pi.
	Instead, we can access it directly using the
	syntax, 'Pi'.

	Note: golint recommends against using the dot
	import statement.
*/

import (
	"fmt"
	. "math"
)

func main() {
	fmt.Println("PI is ", Pi)
	var n float64 = 144
	s := Sqrt(n)
	fmt.Println("The Square Root of 144 is", s)
}

/*	Output
	PI is  3.141592653589793
	The Square Root of 144 is 12
*/
