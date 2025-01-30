package main

import (
	"fmt"
	"math"
)

func main() {
	A := 21 >> 3
	fmt.Println("21 >> 3 =", A)
	twoToPwrOf3 := int(math.Pow(2.0, 3.0))
	B := 21 / twoToPwrOf3
	fmt.Println("21 / 2 to Power of 3 ignore remainder =", B)

	/*
		The Right Shift Operation
		21 >> 3
		Equivalent Binary Division
		A = 21 >> 3
		21 / (2 to the power of 3)
		= 21 /8
		= 2 and ignore the remainder
	*/
}
