package main

import (
	"fmt"
	"math"
)

func main() {
	A := -12 << 4
	fmt.Println("-12 << 4 =", A)
	TwoTo4 := int(math.Pow(2.0, 4.0))
	B := -12 * TwoTo4
	fmt.Println("-12 X 2 to Power of 4 = ", B)

	/* Ouput
		-12 << 4 = -192
		-12 X 2 to Power of 4 =  -192
	*/

}
