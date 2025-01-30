package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	ParseFloat6403()
}

func InitializeTest01() {

	funcName := "InitializeTest01"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	var float64Num float64

	float64Num = -50.125

	numberStr := strconv.FormatFloat(
		float64Num, 'f', -1, 64)

	fmt.Printf("Original Num Str: %v\n",
		numberStr)

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}

func ParseFloat6401() {

	funcName := "ParseFloat6401"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	var float64Num float64
	var err error

	float64Num,
		err = strconv.ParseFloat("-50.125", 64)

	if err != nil {

		fmt.Printf("%v\n"+
			"Error return from  strconv.ParseFloat(\"-50.125\", 64)\n"+
			"Error: \n%v\n",
			funcName,
			err.Error())

		return
	}

	numberStr := strconv.FormatFloat(
		float64Num, 'f', -1, 64)

	fmt.Printf("Parsed Number String Value: %v\n",
		numberStr)

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}

func ParseFloat6402() {

	funcName := "ParseFloat6402"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	var float64Num float64
	var err error

	nativeNumStr := "12345678.123456789125"

	float64Num,
		err = strconv.ParseFloat(nativeNumStr, 64)

	if err != nil {

		fmt.Printf("%v\n"+
			"Error return from  strconv.ParseFloat(nativeNumStr, 64)\n"+
			"nativeNumStr = %v\n"+
			"Error: \n%v\n",
			funcName,
			nativeNumStr,
			err.Error())

		return
	}

	numberStr := strconv.FormatFloat(
		float64Num, 'f', -1, 64)

	fmt.Printf("Parsed Number String Value: %v\n",
		numberStr)

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}

// ParseFloat6403
// ----------------------------------------------------------------
//
//	Be advised that the capacity of a type 'float64' is
//	approximately 15 to 17 digits including integer and
//	fractional digits.
//
//	Source:
//	https://en.wikipedia.org/wiki/Floating-point_arithmetic#IEEE_754:_floating_point_in_modern_computers
//
//	Type	Sign  Exponent	Significand	  Total   Exponent	 Bits	     Number of
//							   field       bits     bias    precision	decimal digits
//	         ----  --------	-----------   -----   --------  ---------	--------------
//	Double	  1		 11			52			64	    1023	   53		   ~15.9
func ParseFloat6403() {

	funcName := "ParseFloat6403"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	initialNativeNumStr := "-1."

	var isNegativeNum bool

	if initialNativeNumStr[0] == '-' {

		isNegativeNum = true

	} else {

		isNegativeNum = false
	}

	nativeNumStr := []rune(initialNativeNumStr)

	nextDigit := '1'

	foundLastGoodDigit := false

	var float64Num float64
	var err error
	var initialNumStr, finalNumStr,
		lastGoodNumStr string

	for foundLastGoodDigit == false {

		lastGoodNumStr = string(nativeNumStr)

		nativeNumStr = append(
			nativeNumStr,
			nextDigit)

		nextDigit += 1

		if nextDigit > '9' {
			nextDigit = '1'
		}

		initialNumStr = string(nativeNumStr)

		float64Num,
			err = strconv.ParseFloat(
			initialNumStr, 64)

		if err != nil {

			fmt.Printf("%v\n"+
				"Error return from  strconv.ParseFloat(nativeNumStr, 64)\n"+
				"nativeNumStr = %v\n"+
				"Error: \n%v\n",
				funcName,
				nativeNumStr,
				err.Error())

			return
		}

		finalNumStr = strconv.FormatFloat(
			float64Num, 'f', -1, 64)

		if initialNumStr != finalNumStr {

			numStrElements :=
				strings.Split(lastGoodNumStr, ".")

			lenIntegerDigits := len(numStrElements[0])
			lenFracDigits := len(numStrElements[1])
			totalNumericDigits := lenIntegerDigits +
				lenFracDigits

			if isNegativeNum {
				lenIntegerDigits--
				totalNumericDigits--
			}

			fmt.Printf("\n%v\n"+
				"Found Max Limit float64\n"+
				"    initialNumStr = %v\n"+
				"      finalNumStr = %v\n"+
				"----------------------\n"+
				"   lastGoodNumStr = %v\n"+
				"   Integer Digits = %v\n"+
				"Fractional Digits = %v\n"+
				"     Total Digits = %v\n",
				funcName,
				initialNumStr,
				finalNumStr,
				lastGoodNumStr,
				lenIntegerDigits,
				lenFracDigits,
				totalNumericDigits)

			foundLastGoodDigit = true
		}

	}

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

	return
}
