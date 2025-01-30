package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {

	ExponentPositive01()
}

func ExponentPositive01() {

	funcName := "ExponentPositive01"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	var result, base, exponent *big.Int

	base = big.NewInt(33)
	exponent = big.NewInt(10)

	result = big.NewInt(0)
	result.Exp(base, exponent, nil)

	// Expected result: 1,531,578,985,264,449

	fmt.Printf("Expected Result: %v\n",
		"1,531,578,985,264,449")

	fmt.Printf("Actual Result: %v\n",
		result.Text(10))

	fmt.Printf("\n" + breakStr + "\n")

	fmt.Printf("  Successful Completion!\n" +
		"Function: " +
		funcName + "\n")

	fmt.Printf(breakStr + "\n")

}
