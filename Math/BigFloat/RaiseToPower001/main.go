package main

import (
	"fmt"
	"math/big"
	"strings"
)

type BigFloatCalcStats struct {
	BaseFloatNum big.Float

	BasePrec uint

	BaseNumIntDigits int64

	BaseNumFracDigits int64

	CalcResult big.Float

	CalcResultPrec uint

	CalcResultNumIntDigits int64

	CalcResultNumFracDigits int64
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

	funcName := "Main()"

	baseNumStr := "128.773333"

	exponent := int64(11)

	expectedResult := "161467567251021767963804.20470511"

	baseNum := big.Float{}

	baseNum.SetMode(big.AwayFromZero).
		SetInt64(0)

	var ok bool

	_,
		ok = baseNum.SetString(baseNumStr)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error: baseNum.SetString(baseNumStr) FAILED!\n"+
			"baseNumStr = %v\n",
			funcName,
			baseNumStr)

		return
	}

	_,
		err := raiseToPowerInt(
		baseNum,
		exponent,
		expectedResult)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"Error Return:\n%v\n\n",
			funcName,
			err.Error())
	}

}

func raiseToPowerInt(
	base big.Float,
	exponent int64,
	expectedResult string) (
	resultsStats BigFloatCalcStats,
	err error) {

	funcName := "raiseToPowerInt()"

	breakStr := strings.Repeat("=", 70)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	var pBase *big.Float

	originalBasePrec := base.Prec()

	base.SetPrec(base.MinPrec())

	resultsStats.BaseFloatNum = base
	resultsStats.BasePrec = base.Prec()

	pBase = &base

	numStr := pBase.Text('f', -1)

	rawBaseNumStr := numStr

	var pureNumStrStats PureNumberStrComponents

	pureNumStrStats,
		err = breakNumStrToComponents(
		numStr)

	if err != nil {
		return resultsStats, err
	}

	if exponent < 0 {

		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'exponent' is INVALID!\n"+
			"'exponent' has a value less than zero.\n"+
			"exponent = %v\n",
			funcName,
			exponent)

		return resultsStats, err
	}

	resultsStats.BaseNumIntDigits =
		pureNumStrStats.NumIntegerDigits

	resultsStats.BaseNumFracDigits =
		pureNumStrStats.NumFractionalDigits

	var ok bool

	bigIntBase,
		ok := big.NewInt(0).
		SetString(
			pureNumStrStats.AllIntegerDigitsNumStr,
			10)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error: bigIntBase=SetString(AllIntegerDigitsNumStr)\n"+
			"SetString Failed!\n"+
			"AllIntegerDigitsNumStr = %v\n",
			funcName,
			pureNumStrStats.AllIntegerDigitsNumStr)

		return resultsStats, err
	}

	bigIntExponent := big.NewInt(
		exponent)

	raisedToPowerFracDigits :=
		exponent * pureNumStrStats.NumFractionalDigits

	bigIntBase.Exp(bigIntBase, bigIntExponent, nil)

	numStr =
		bigIntBase.Text(10)

	rawResultsStr := numStr

	lenNumStrInt64 := int64(len(numStr))

	if raisedToPowerFracDigits == 0 {

		resultsStats.CalcResultNumIntDigits =
			lenNumStrInt64

		resultsStats.CalcResultNumFracDigits = 0

	} else {

		resultsStats.CalcResultNumIntDigits =
			lenNumStrInt64 - raisedToPowerFracDigits

		resultsStats.CalcResultNumFracDigits =
			raisedToPowerFracDigits

		numStr =
			numStr[0:resultsStats.CalcResultNumIntDigits] +
				"." +
				numStr[resultsStats.CalcResultNumIntDigits:]

	}

	var calcPrecision uint

	calcPrecision,
		err = calculateRequiredPrecision(
		resultsStats.CalcResultNumIntDigits,
		resultsStats.CalcResultNumFracDigits,
		5)

	if err != nil {
		return resultsStats, err
	}

	_,
		ok = resultsStats.CalcResult.
		SetPrec(calcPrecision).
		SetMode(big.AwayFromZero).
		SetString(numStr)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error: CalcResult=SetString(numStr)\n"+
			"SetString Failed!\n"+
			"numStr = %v\n",
			funcName,
			numStr)

	}

	resultsStats.CalcResult.SetPrec(
		resultsStats.CalcResult.MinPrec())

	resultsStats.CalcResultPrec =
		resultsStats.CalcResult.Prec()

	fmt.Printf("\n\n\t\t%v\n"+
		"\tRaise To Exponent Results\n"+
		"                   Base = %v\n"+
		"        Raw Base NumStr = %v\n"+
		"Original Base Precision = %v\n"+
		"         Base Precision = %v\n"+
		"    Base Integer Digits = %v\n"+
		" Base Fractional Digits = %v\n"+
		"               Exponent = %v\n"+
		"     Raw Results NumStr = %v\n"+
		"     Calculation Result = %v\n"+
		"        Expected Result = %v\n"+
		"  Calc Result Precision = %v\n"+
		" Calc Result Int Digits = %v\n"+
		"Calc Result Frac Digits = %v\n\n",
		funcName,
		base.Text('f', -1),
		rawBaseNumStr,
		originalBasePrec,
		base.Prec(),
		resultsStats.BaseNumIntDigits,
		resultsStats.BaseNumFracDigits,
		exponent,
		rawResultsStr,
		resultsStats.CalcResult.Text('f', -1),
		expectedResult,
		resultsStats.CalcResult.Prec(),
		resultsStats.CalcResultNumIntDigits,
		resultsStats.CalcResultNumFracDigits)

	fmt.Printf("\n\n%v\n"+
		"   Successful Completion!\n"+
		"Function: %v\n%v\n\n",
		breakStr,
		funcName,
		breakStr)

	return resultsStats, err
}

func calculateRequiredPrecision(
	integerDigits int64,
	fractionalDigits int64,
	spareBufferDigits int64) (
	uint,
	error) {

	funcName := "calculateRequiredPrecision"

	totalDigits :=
		integerDigits +
			fractionalDigits +
			spareBufferDigits

	totalDigitsFloat := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(totalDigits)

	factorEightFloat := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(8)

	precToDigitsFactor := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	// 3.3219789132197891321978913219789

	var err error
	_,
		ok := precToDigitsFactor.SetString("3.3219789132197891321978913219789")

	if !ok {
		err = fmt.Errorf("\n%v\n"+
			"expectedNthRoot.SetString(\"3.3219789132197891321978913219789\") FAILED!\n",
			funcName)

		return 0, err
	}

	totalDigitsFloat.Mul(totalDigitsFloat, precToDigitsFactor)

	totalDigitsFloat.Quo(totalDigitsFloat, factorEightFloat)

	baseDigits,
		accuracy := totalDigitsFloat.Int64()

	if accuracy == -1 {
		baseDigits++
	}

	baseDigits = baseDigits * 8

	return uint(baseDigits), err
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

	fmt.Printf("\n%v\n"+
		"                          numberString = %v\n"+
		"   pureNumStrStats.AbsoluteValueNumStr = %v\n"+
		"pureNumStrStats.AllIntegerDigitsNumStr = %v\n",
		funcName,
		numberString,
		pureNumStrStats.AbsoluteValueNumStr,
		pureNumStrStats.AllIntegerDigitsNumStr)

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
