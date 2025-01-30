package main

import (
	"fmt"
	"math"
)

func main() {
	A := 13 << 4
	fmt.Println("13 << 4 =", A)
	twoToPower4 := int(math.Pow(2.0, 4.0))
	B := 13 * twoToPower4
	fmt.Println("13 X (2 to power of 4) =", B)

	/* Output
		13 << 4 = 208
		13 X (2 to power of 4) = 208
	 */
}
