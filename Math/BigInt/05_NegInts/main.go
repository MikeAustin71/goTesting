package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {

	convertNegToPos()
}

func convertNegToPos() {

	funcName := "convertNegToPos"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	baseNum := big.NewInt(-300)

	fmt.Printf("baseNum Before 'Neg' = %v\n",
		baseNum.Text(10))

	baseNum.Neg(baseNum)

	fmt.Printf("baseNum After 'Neg' = %v\n",
		baseNum.Text(10))

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}
