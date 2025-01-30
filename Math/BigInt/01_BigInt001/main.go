package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {

	bigIntNegConversion01()
}

func bigIntNegConversion01() {

	funcName := "bigIntNegConversion01"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	numValueStr := "-12345678901234567890"

	fmt.Printf("Original Number String:         %v\n",
		numValueStr)

	var ok bool
	var bigIntValue big.Int
	_,
		ok = bigIntValue.SetString(
		numValueStr,
		10)

	if !ok {

		fmt.Printf("%v\n"+
			"Error Converting 'numValueStr' to big.Int!\n"+
			"The following integerDigits string generated an error.\n"+
			"'numValueStr' = '%v'\n",
			funcName,
			numValueStr)

		return
	}

	fmt.Printf("Converted Big Int String:       %v\n",
		bigIntValue.Text(10))

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

	return
}
