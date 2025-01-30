package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	PrecisionTest01()
}

func PrecisionTest01() {

	funcName := "PrecisionTest01()"

	// Input Parameters
	prec := uint(1024)

	factor3_3 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	factor3_3.SetString("3.3219789132197891321978913219789")

	test_Number_rat := big.NewRat(217, 3)
	//localPrecision := 50

	test_Number_float := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetRat(test_Number_rat)

	testFloatStr :=
		test_Number_float.Text('f', -1)

	lenStr := len(testFloatStr)

	foundDecimalPt := false

	var numIntDigits, numFracDigits, numTotalDigits uint64

	var digitRunes []rune

	for i := 0; i < lenStr; i++ {

		if testFloatStr[i] == '.' {
			foundDecimalPt = true

			continue
		}

		if testFloatStr[i] >= '0' &&
			testFloatStr[i] <= '9' {

			numTotalDigits++

			digitRunes =
				append(digitRunes,
					rune(testFloatStr[i]))

			if foundDecimalPt == false {
				numIntDigits++
			} else {
				numFracDigits++
			}

		}

	}

	test_Number_Float_exponent :=
		test_Number_float.MantExp(nil)

	bInt := big.NewInt(0)

	bInt.SetString(string(digitRunes), 10)

	bInt256 := big.NewInt(256)

	bInt.Quo(bInt, bInt256)

	estimatedNumOfDigits := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(int64(prec))

	estimatedNumOfDigits.Quo(estimatedNumOfDigits, factor3_3)

	estimatedTotalDigits,
		estDigitsAccuracy := estimatedNumOfDigits.Int64()

	/*
		const (
			Below Accuracy = -1
			Exact Accuracy = 0
			Above Accuracy = +1
		)

	*/
	var estDigitsAccuracyStr string

	if estDigitsAccuracy == 1 {

		estDigitsAccuracyStr = "Above Accuracy"

	} else if estDigitsAccuracy == 0 {

		estDigitsAccuracyStr = "Exact Accuracy"

	} else {
		// MUST BE -
		//  -1
		estDigitsAccuracyStr = "Below Accuracy"

	}

	estimatedIntDigits := int64(numIntDigits)
	estimatedFracDigits := int64(0)

	estimatedFracDigits =
		estimatedTotalDigits - int64(numIntDigits)

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n" + breakStr + "\n")

	fmt.Printf("Function: %v\n",
		funcName)

	fmt.Printf(breakStr)

	fmt.Printf(
		"\n"+
			"                         Precision = %v\n"+
			"            Number of Total Digits = %v\n"+
			"  Estimated Total Number of Digits = %v\n"+
			"    Total Digits Estimate Accuracy = %v\n"+
			"          Number of Integer Digits = %v\n"+
			"Estimated Number of Integer Digits = %v\n"+
			"       Number of Fractional Digits = %v\n"+
			"   Estimated Number of Frac Digits = %v\n"+
			"    Estimated Number Digits Factor = %v\n"+
			"                   Number Exponent = %v\n"+
			"                            Number = %v\n",
		prec,
		numTotalDigits,
		estimatedTotalDigits,
		estDigitsAccuracyStr,
		numIntDigits,
		estimatedIntDigits,
		numFracDigits,
		estimatedFracDigits,
		factor3_3.Text('f', 35),
		test_Number_Float_exponent,
		test_Number_float.Text('f', 80))

	/*
		,
				test_Number_float.Text('f', 80),
				test_Number_float.Text('f', -1)
		"+
				breakStr+"\n"+
				"Number = \n"+
				"%v\n\n",
				"breakStr+\n\n"+
					"Number Detail =\n%v\n\n"+
					"breakStr+\n\n"
	*/

}
