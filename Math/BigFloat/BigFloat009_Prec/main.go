package main

import (
	"fmt"
	"math/big"
	"strings"
)

// BigFloatPrecisionDto
// Used in computing a precision bits value for
// a specific big.Float value.
type BigFloatPrecisionDto struct {
	NumIntegerDigits int64

	NumFractionalDigits int64

	NumOfExtraDigitsBuffer int64

	EstimatedNumberOfPrecisionBits uint
}

type BigFloatDto struct {
	Value                  big.Float
	ActualNumStrComponents PureNumberStrComponents
	EstimatedPrecisionBits BigFloatPrecisionDto
}

func (bigFloatDto *BigFloatDto) GetTotalNumericDigits() int64 {

	return bigFloatDto.ActualNumStrComponents.NumIntegerDigits +
		bigFloatDto.ActualNumStrComponents.NumFractionalDigits

}

func (bigFloatDto *BigFloatDto) GetEstPrecisionTotalNumDigits() int64 {

	return bigFloatDto.EstimatedPrecisionBits.NumIntegerDigits +
		bigFloatDto.EstimatedPrecisionBits.NumFractionalDigits +
		bigFloatDto.EstimatedPrecisionBits.NumOfExtraDigitsBuffer

}

type PureNumberStrComponents struct {
	NumberSign int
	// -1 == negative
	//  0 == zero
	//  1 == positive

	NumberType int
	//	0 == invalid
	//	1 == integer number
	//	2 == floating point number

	NumIntegerDigits int64

	NumFractionalDigits int64

	AbsoluteValueNumStr string

	AllIntegerDigitsNumStr string
}

func main() {

	testPrecSmallNums01("32.75819",
		10,
		0,
		big.AwayFromZero,
		false,
		true)

}

func testPrecSmallNums01(numStr string,
	numOfExtraDigits int64,
	precisionBitsOverride uint,
	roundingMode big.RoundingMode,
	turnOffMinPrec bool,
	turnOnMessages bool) {

	funcName := "testPrecSmallNums01()"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	fmt.Printf("\n\n%v\nInitializtion\n"+
		"           Original numStr = %v\n"+
		"    Number of Extra Digits = %v\n"+
		"   Precision Bits Override = %v\n"+
		"             Rounding Mode = %v\n"+
		"Turn Off Minimum Precision = %v\n",
		funcName,
		numStr,
		numOfExtraDigits,
		precisionBitsOverride,
		roundingMode,
		turnOffMinPrec)

	bFloatDto := BigFloatFromPureNumStr(
		numStr,
		numOfExtraDigits,
		precisionBitsOverride,
		roundingMode,
		turnOffMinPrec,
		turnOnMessages)

	finalBFloatValStr := bFloatDto.Value.Text('f', -1)

	fmt.Printf("\n%v\nPost Configuration\n"+
		"     bFloatDto.Value = %v\n"+
		"bFloatDto.Value.Prec = %v\n"+
		"bFloatDto.Value.Mode = %v\n"+
		"bFloatDto.Value.Acc  = %v\n",
		funcName,
		finalBFloatValStr,
		bFloatDto.Value.Prec(),
		bFloatDto.Value.Mode(),
		bFloatDto.Value.Acc())

	if finalBFloatValStr != numStr {
		fmt.Printf("\n\n%v\n"+
			"Calculation Error!!\n"+
			"'numStr' DOES NOT MATCH bFloatDto.Value!!\n"+
			"         numStr = %v\n"+
			"bFloatDto.Value = %v\n\n\n",
			funcName,
			numStr,
			finalBFloatValStr)
	} else {
		fmt.Printf("\n\n%v\n"+
			"       !!!! SUCCESS !!!!\n"+
			"'numStr' MATCHES bFloatDto.Value!!\n"+
			"         numStr = %v\n"+
			"bFloatDto.Value = %v\n\n\n",
			funcName,
			numStr,
			finalBFloatValStr)
	}
}

func BigFloatFromPureNumStr(
	pureNumberValueStr string,
	numOfExtraDigitsPrecisionBuffer int64,
	precisionBitsOverride uint,
	roundingMode big.RoundingMode,
	turnOffMinPrec bool,
	turnOnMessages bool) BigFloatDto {

	funcName := "BigFloatFromPureNumStr"
	var err error
	var bigFloatDto BigFloatDto

	bigFloatDto.ActualNumStrComponents,
		err = breakNumStrToComponents(
		pureNumberValueStr)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return bigFloatDto
	}

	var precisionBitsSpec uint

	if precisionBitsOverride == 0 {

		bigFloatDto.EstimatedPrecisionBits.NumIntegerDigits =
			bigFloatDto.ActualNumStrComponents.NumIntegerDigits

		bigFloatDto.EstimatedPrecisionBits.NumFractionalDigits =
			bigFloatDto.ActualNumStrComponents.NumFractionalDigits

		bigFloatDto.EstimatedPrecisionBits.NumOfExtraDigitsBuffer =
			numOfExtraDigitsPrecisionBuffer

		bigFloatDto.EstimatedPrecisionBits.EstimatedNumberOfPrecisionBits,
			err =
			calculateRequiredPrecision8(
				bigFloatDto.EstimatedPrecisionBits.NumIntegerDigits,
				bigFloatDto.EstimatedPrecisionBits.NumFractionalDigits,
				bigFloatDto.EstimatedPrecisionBits.NumOfExtraDigitsBuffer,
				turnOnMessages)

		if err != nil {

			fmt.Printf("\n\n%v\n"+
				"Estimated Required Precision Bits\n"+
				"Error returned by calculateRequiredPrecision8()\n"+
				"Error = %v\n\n",
				funcName,
				err.Error())

			return bigFloatDto
		}

		precisionBitsSpec =
			bigFloatDto.EstimatedPrecisionBits.EstimatedNumberOfPrecisionBits

	} else {

		bigFloatDto.EstimatedPrecisionBits.
			EstimatedNumberOfPrecisionBits =
			precisionBitsOverride

		precisionBitsSpec = precisionBitsOverride
	}

	if turnOnMessages {

		fmt.Printf("\n\n%v\n"+
			"Precision Bits Specification Before SetPrec()\n"+
			"precisionBitsSpec = %v\n\n",
			funcName,
			precisionBitsSpec)

	}

	/*	bigFloatDto.Value.SetInt64(0).
		SetPrec(precisionBitsSpec).
		SetMode(roundingMode)
	*/

	bigFloatDto.Value.
		SetPrec(precisionBitsSpec).
		SetMode(roundingMode)

	if turnOnMessages {

		fmt.Printf("%v\n"+
			"Actual Precision Bits After SetPrec()\n"+
			"precisionBits = %v\n\n",
			funcName,
			bigFloatDto.Value.Prec())
	}

	var ok bool
	_,
		ok = bigFloatDto.Value.SetString(pureNumberValueStr)

	if !ok {
		fmt.Printf("Error: bigFloatDto.Value.SetString(numValueStr) Failed!\n")

		return bigFloatDto
	}

	/*	if bigFloatDto.Value.Acc() != big.Exact {

			fmt.Printf("\n\n%v\n"+
				"Accuracy Error After SetString()\n"+
				"An exact floating pointing number value could NOT\n"+
				"be calculated from the Pure Number Value string,\n"+
				"'pureNumberValueStr',\n"+
				"pureNumberValueStr = %v\n",
				funcName,
				pureNumberValueStr)

			return BigFloatDto{}

		}
	*/
	if turnOnMessages {

		fmt.Printf("%v\n"+
			"Actual Precision Bits After SetString()\n"+
			"and Before turnOffMinPrec.\n"+
			"precisionBits = %v\n\n",
			funcName,
			bigFloatDto.Value.Prec())

	}

	if turnOffMinPrec == false {

		if !bigFloatDto.Value.IsInt() {

			bigFloatDto.Value.SetPrec(bigFloatDto.Value.MinPrec())

			if bigFloatDto.Value.Acc() != big.Exact {

				fmt.Printf("\n\n%v\n"+
					"Accuracy Error\n"+
					"An exact floating pointing number value could NOT\n"+
					"be calculated from the Pure Number Value string,\n"+
					"'pureNumberValueStr',\n"+
					"pureNumberValueStr = %v\n",
					funcName,
					pureNumberValueStr)

				return BigFloatDto{}

			}

		}

	}

	if turnOnMessages {
		fmt.Printf("%v\n"+
			"Actual Precision Bits After turnOffMinPrec.\n"+
			"precisionBits = %v\n\n",
			funcName,
			bigFloatDto.Value.Prec())

	}

	if turnOnMessages {

		fmt.Printf("%v\n"+
			"Actual Precision Bits Final Function Exit Status.\n"+
			"precisionBits = %v\n\n",
			funcName,
			bigFloatDto.Value.Prec())

	}

	return bigFloatDto
}

func BigFloatFromPureNumStr02(
	pureNumberValueStr string,
	numOfExtraDigitsPrecisionBuffer int64,
	precisionBitsOverride uint,
	roundingMode big.RoundingMode,
	turnOffMinPrec bool,
	turnOnMessages bool) BigFloatDto {

	funcName := "BigFloatFromPureNumStr"
	var err error
	var bigFloatDto BigFloatDto

	bigFloatDto.ActualNumStrComponents,
		err = breakNumStrToComponents(
		pureNumberValueStr)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return bigFloatDto
	}

	var precisionBitsSpec uint

	if precisionBitsOverride == 0 {

		bigFloatDto.EstimatedPrecisionBits.NumIntegerDigits =
			bigFloatDto.ActualNumStrComponents.NumIntegerDigits

		bigFloatDto.EstimatedPrecisionBits.NumFractionalDigits =
			bigFloatDto.ActualNumStrComponents.NumFractionalDigits

		bigFloatDto.EstimatedPrecisionBits.NumOfExtraDigitsBuffer =
			numOfExtraDigitsPrecisionBuffer

		bigFloatDto.EstimatedPrecisionBits.EstimatedNumberOfPrecisionBits,
			err =
			calculateRequiredPrecision8(
				bigFloatDto.EstimatedPrecisionBits.NumIntegerDigits,
				bigFloatDto.EstimatedPrecisionBits.NumFractionalDigits,
				bigFloatDto.EstimatedPrecisionBits.NumOfExtraDigitsBuffer,
				turnOnMessages)

		if err != nil {

			fmt.Printf("\n\n%v\n"+
				"Estimated Required Precision Bits\n"+
				"Error returned by calculateRequiredPrecision8()\n"+
				"Error = %v\n\n",
				funcName,
				err.Error())

			return bigFloatDto
		}

		precisionBitsSpec =
			bigFloatDto.EstimatedPrecisionBits.EstimatedNumberOfPrecisionBits

	} else {

		bigFloatDto.EstimatedPrecisionBits.
			EstimatedNumberOfPrecisionBits =
			precisionBitsOverride

		precisionBitsSpec = precisionBitsOverride
	}

	if turnOnMessages {

		fmt.Printf("\n\n%v\n"+
			"Precision Bits Specification Before SetPrec()\n"+
			"precisionBitsSpec = %v\n\n",
			funcName,
			precisionBitsSpec)

	}

	/*	bigFloatDto.Value.SetInt64(0).
		SetPrec(precisionBitsSpec).
		SetMode(roundingMode)
	*/

	bigFloatDto.Value.
		SetPrec(precisionBitsSpec).
		SetMode(roundingMode)

	if turnOnMessages {

		fmt.Printf("%v\n"+
			"Actual Precision Bits After SetPrec()\n"+
			"precisionBits = %v\n\n",
			funcName,
			bigFloatDto.Value.Prec())
	}

	var ok bool
	_,
		ok = bigFloatDto.Value.SetString(pureNumberValueStr)

	if !ok {
		fmt.Printf("Error: bigFloatDto.Value.SetString(numValueStr) Failed!\n")

		return bigFloatDto
	}

	/*	if bigFloatDto.Value.Acc() != big.Exact {

			fmt.Printf("\n\n%v\n"+
				"Accuracy Error After SetString()\n"+
				"An exact floating pointing number value could NOT\n"+
				"be calculated from the Pure Number Value string,\n"+
				"'pureNumberValueStr',\n"+
				"pureNumberValueStr = %v\n",
				funcName,
				pureNumberValueStr)

			return BigFloatDto{}

		}
	*/
	if turnOnMessages {

		fmt.Printf("%v\n"+
			"Actual Precision Bits After SetString()\n"+
			"and Before turnOffMinPrec.\n"+
			"precisionBits = %v\n\n",
			funcName,
			bigFloatDto.Value.Prec())

	}

	if turnOffMinPrec == false {

		bigFloatDto.Value.SetPrec(bigFloatDto.Value.MinPrec())

		if bigFloatDto.Value.Acc() != big.Exact {

			fmt.Printf("\n\n%v\n"+
				"Accuracy Error\n"+
				"An exact floating pointing number value could NOT\n"+
				"be calculated from the Pure Number Value string,\n"+
				"'pureNumberValueStr',\n"+
				"pureNumberValueStr = %v\n",
				funcName,
				pureNumberValueStr)

			return BigFloatDto{}

		}

	}

	if turnOnMessages {
		fmt.Printf("%v\n"+
			"Actual Precision Bits After turnOffMinPrec.\n"+
			"precisionBits = %v\n\n",
			funcName,
			bigFloatDto.Value.Prec())

	}

	if turnOnMessages {

		fmt.Printf("%v\n"+
			"Actual Precision Bits Final Function Exit Status.\n"+
			"precisionBits = %v\n\n",
			funcName,
			bigFloatDto.Value.Prec())

	}

	return bigFloatDto
}

func breakNumStrToComponents(
	numberString string) (
	pureNumStrStats PureNumberStrComponents,
	err error) {

	funcName := ""

	lenNumberStr := len(numberString)

	if lenNumberStr == 0 {
		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'numberString'\n"+
			"is a zero length string and INVALID!\n",
			funcName)

		return pureNumStrStats, err
	}

	pureNumStrStats.NumberSign = 0
	pureNumStrStats.NumberType = 2

	if numberString[0] == '-' {

		pureNumStrStats.NumberSign = -1

		numberString = numberString[1:]

		lenNumberStr--

		if lenNumberStr == 0 {
			err = fmt.Errorf("\n\n%v\n"+
				"Error: Input parameter 'numberString'\n"+
				"is a zero length string and INVALID!\n",
				funcName)

			return pureNumStrStats, err
		}

	} else {

		pureNumStrStats.NumberSign = 1

	}

	idx := strings.Index(numberString, ".")

	if idx == -1 {

		pureNumStrStats.NumberType = 1

		pureNumStrStats.NumIntegerDigits =
			int64(lenNumberStr)

		pureNumStrStats.NumFractionalDigits =
			0

		pureNumStrStats.AbsoluteValueNumStr =
			numberString

		pureNumStrStats.AllIntegerDigitsNumStr =
			numberString

	} else {

		pureNumStrStats.NumberType = 2

		pureNumStrStats.NumIntegerDigits =
			int64(idx)

		pureNumStrStats.NumFractionalDigits =
			int64(lenNumberStr - (idx + 1))

		pureNumStrStats.AbsoluteValueNumStr =
			numberString

		pureNumStrStats.AllIntegerDigitsNumStr =
			numberString[0:idx] +
				numberString[idx+1:]
	}
	//
	//fmt.Printf("\n%v\n"+
	//	"                          numberString = %v\n"+
	//	"   pureNumStrStats.AbsoluteValueNumStr = %v\n"+
	//	"pureNumStrStats.AllIntegerDigitsNumStr = %v\n",
	//	funcName,
	//	numberString,
	//	pureNumStrStats.AbsoluteValueNumStr,
	//	pureNumStrStats.AllIntegerDigitsNumStr)

	isZeroValue := true
	lenNumberStr = len(pureNumStrStats.AbsoluteValueNumStr)

	for i := 0; i < lenNumberStr; i++ {

		if pureNumStrStats.AbsoluteValueNumStr[i] >= '0' &&
			pureNumStrStats.AbsoluteValueNumStr[i] <= '9' {

			isZeroValue = false

			break
		}
	}

	if isZeroValue {
		pureNumStrStats.NumberSign = 0
	}

	return pureNumStrStats, err
}

func calculateRequiredPrecision8(

	integerDigits int64,
	fractionalDigits int64,
	spareDigitsBuffer int64,
	turnOnMessages bool) (
	uint,
	error) {

	// Precision calculated in multiples of 8

	funcName := "calculateRequiredPrecision8"

	totalDigits :=
		integerDigits +
			fractionalDigits +
			spareDigitsBuffer

	fmt.Printf("\n\n%v\n"+
		"Initial Input\n"+
		"    integerDigits = %v\n"+
		" fractionalDigits = %v\n"+
		"spareDigitsBuffer = %v\n"+
		"     Total Digits = %v\n\n",
		funcName,
		integerDigits,
		fractionalDigits,
		spareDigitsBuffer,
		totalDigits)

	totalDigitsFloat := new(big.Float).
		SetPrec(2048).
		SetMode(big.AwayFromZero).
		SetInt64(totalDigits)

	factorEightFloat := new(big.Float).
		SetPrec(512).
		SetMode(big.AwayFromZero).
		SetInt64(8)

	factorEightUint64 := uint64(8)

	precToDigitsFactor := new(big.Float).
		SetPrec(512).
		SetMode(big.AwayFromZero)

	// 3.3219789132197891321978913219789

	var err error
	_,
		ok := precToDigitsFactor.SetString("3.3219789132197891321978913219789")

	if !ok {
		err = fmt.Errorf("\n%v\n"+
			"precToDigitsFactor.SetString(\"3.3219789132197891321978913219789\") FAILED!\n",
			funcName)

		return 0, err
	}

	// precToDigitsFactor.SetPrec(precToDigitsFactor.MinPrec())

	/*	if precToDigitsFactor.Acc() != big.Exact {

			err = fmt.Errorf("\n%v\n"+
				"Accuracy Test #1\n"+
				"Error: Accuracy Test Failed!\n"+
				"precToDigitsFactor inaccurate!\n"+
				"          precToDigitsFactor = %v\n"+
				"precToDigitsFactor Precision = %v\n",
				funcName,
				precToDigitsFactor.Text('f', -1),
				precToDigitsFactor.Prec())

			return 0, err
		}
	*/
	totalDigitsFloat.Mul(totalDigitsFloat, precToDigitsFactor)

	totalDigitsFloat.Quo(totalDigitsFloat, factorEightFloat)

	baseDigits,
		accuracy := totalDigitsFloat.Uint64()

	if accuracy == -1 {
		baseDigits++
	}

	baseDigits = baseDigits * factorEightUint64

	if turnOnMessages {

		fmt.Printf("\n%v\n"+
			"Exit Status\n"+
			"    baseDigits = %v\n\n",
			funcName,
			baseDigits)

	}

	return uint(baseDigits), err
}
