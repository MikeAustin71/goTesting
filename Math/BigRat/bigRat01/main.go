package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {

	BigRatToNumStr04()

}

func BigRatToNumStr01() {

	funcName := "BigRatToNumStr01"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	newRat := big.NewRat(1, 3)

	floatStr := newRat.FloatString(2000)

	fmt.Printf("\n%v\n"+
		"'newRat' FloatString\n"+
		"newRat = %v\n",
		funcName,
		floatStr)

}

func BigRatToNumStr02() {

	funcName := "BigRatToNumStr02"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	newRat := big.NewRat(1, 3)

	bigFloatNum := new(big.Float)

	bigFloatNum.SetRat(newRat)

	floatStr := bigFloatNum.Text('f', -1)

	fmt.Printf("\n%v\n"+
		"'newRat' FloatString\n"+
		"Original New Rat = %v\n"+
		"newRat = %v\n",
		funcName,
		newRat.RatString(),
		floatStr)

}

func BigRatToNumStr03() {

	funcName := "BigRatToNumStr03"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	newRat := big.NewRat(50, 100)

	floatStr := newRat.FloatString(2000)

	fmt.Printf("\n%v\n"+
		"'newRat' FloatString\n"+
		"newRat = %v\n",
		funcName,
		floatStr)

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}

func BigRatToNumStr04() {

	funcName := "BigRatToNumStr04"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)
	fmt.Printf("    Converting a number string\n" +
		"    to a BigRat Number.\n")

	fmt.Printf(breakStr + "\n\n")

	originalNumStr := "250.75"
	fmt.Printf("    Original Number String: %v\n\n",
		originalNumStr)

	fmt.Printf(breakStr + "\n\n")

	compositeStr := "25075"

	exponent := big.NewInt(int64(2))

	denominator := big.NewInt(10)

	denominator.Exp(denominator, exponent, nil)

	numerator := big.NewInt(0)

	var ok bool

	_,
		ok = numerator.
		SetString(
			compositeStr,
			10)

	if !ok {

		fmt.Printf("%v\n"+
			"Error Converting 'nativeNumStr' to 'numerator' (*big.Int)!\n"+
			"The following integerDigits string generated an error.\n"+
			"compositeStr = '%v'\n",
			funcName,
			compositeStr)

		return
	}

	bigRatNum := new(big.Rat)

	bigRatNum.SetFrac(numerator, denominator)

	fmt.Printf("%v\n"+
		"bigRatNum =\t   %v\n"+
		"Original Num Str = %v\n",
		funcName,
		bigRatNum.FloatString(2),
		originalNumStr)

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}
