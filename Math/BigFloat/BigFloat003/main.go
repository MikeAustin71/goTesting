package main

import (
	"fmt"
	"math/big"
)

func main() {


	var newFloat *big.Float
	var b int
	var err error


	floatStr := "1.6666666666666666666666666666666666666667"
	newFloat,
	b,
	err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n" +
			"Error='%v'\n", err.Error())
		return
	}


	fmt.Println()
	fmt.Println("Testing big.Float Parse Function!")
	fmt.Println("---------------------------------")
	fmt.Printf("         Expected value:    %v\n",
		floatStr)
	fmt.Printf("  Actual newFloat value: %45.40f\n",
		newFloat)
	fmt.Printf("  Returned Value of 'b': %v\n", b)

}
