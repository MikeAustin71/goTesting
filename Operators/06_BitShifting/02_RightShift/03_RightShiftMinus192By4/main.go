package main

import (
	"fmt"
	"math"
)

func main() {
	A := -192 >> 4
	fmt.Println("-192 >> 4 =", A)
	twoToPwrOf4 := int(math.Pow(2.0, 4.0))
	B := -192 / twoToPwrOf4
	fmt.Println("-192 / 2 to Power of 4 ignore remainder =", B)

	/*
		The Right Shift Operation
		-192 >> 4
		Equivalent Binary Division
		A = -191 >> 4
		-192 / (2 to the power of 4)
		= -192 /16
		= -12 and ignore any remainder
	*/
}
