package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	bigIntLeftShift()
}

// bigIntLeftShift
//
// Takes two numbers, left shifts the bits of the
// first operand, the second operand decides the
// number of places to shift. Or in other words
// left shifting an integer “x” with an integer
// “y” denoted as ‘(x<<y)’ is equivalent to
// multiplying x with 2^y (2 raised to power y).
func bigIntLeftShift() {

	funcName := "bigIntLeftShift"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	bigIntNum := big.NewInt(2)

	fmt.Printf("Original Value:         %v\n",
		bigIntNum)

	leftShiftBits := uint(1)

	fmt.Printf("Left Shift Value:       %v\n",
		leftShiftBits)

	bigIntNum.Lsh(bigIntNum, leftShiftBits)

	fmt.Printf("After Lsh Value:        %v\n",
		bigIntNum)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}
