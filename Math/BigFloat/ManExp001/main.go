package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {

	baseNumStr := "91.35743"

	MantissaExponent04(baseNumStr)

}

func MantissaExponent01() {

	funcName := "MantissaExponent01()"

	breakStr := strings.Repeat("=", 70)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	tFloat01 := new(big.Float).
		SetMode(big.AwayFromZero).
		SetFloat64(1234.567890)

	mantissa := big.NewFloat(0)

	exponent := tFloat01.MantExp(mantissa)

	fmt.Printf("\nExtraction Test #1-a\n"+
		"tFloat01 = %v\n\n"+
		"mantissa = %v\n\n"+
		"mantissa precision = %v\n\n"+
		"exponent = %v\n\n",
		tFloat01.Text('f', -1),
		mantissa.Text('f', -1),
		mantissa.Prec(),
		exponent)

	tFloat02 := big.NewFloat(0)

	tFloat02.SetMantExp(mantissa, exponent)

	mantissaStr := mantissa.Text('f', -1)

	mantissaStr = mantissaStr[2:]

	lenMantissaStr := len(mantissaStr)

	fmt.Printf("\nBuild Test #2-a\n"+
		"              tFloat02 = %v\n\n"+
		"              mantissa = %v\n\n"+
		"    mantissa precision = %v\n\n"+
		"       mantissa string = %v\n\n"+
		"Length mantissa string = %v\n\n"+
		"               exponent = %v\n\n",
		tFloat02.Text('f', -1),
		mantissa.Text('f', -1),
		mantissa.Prec(),
		mantissaStr,
		lenMantissaStr,
		exponent)

	mantissaBInt,
		ok := big.NewInt(0).SetString(
		mantissaStr, 10)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error: mantissaBInt=SetString(mantissaStr)\n"+
			"SetString Failed!\n"+
			"mantissaStr = %v\n",
			funcName,
			mantissaStr)

		return
	}

	bIntExponent := new(big.Int).SetInt64(2)
	bIntPower := new(big.Int).SetInt64(int64(exponent))
	bIntExponent.Exp(bIntExponent, bIntPower, nil)
	bIntValue := new(big.Int).SetInt64(0)
	bIntValue.Mul(mantissaBInt, bIntExponent)

	fmt.Printf("\n\nTest # 3-a\n"+
		" mantissaBInt = %v\n"+
		"BInt Exponent = %v\n"+
		"   BInt Value = %v\n\n",
		mantissaBInt.Text(10),
		bIntExponent.Text(10),
		bIntValue.Text(10))

	mantissaOfMantissa := big.NewFloat(0)

	mantissaOfMantissaExponent :=
		mantissa.MantExp(mantissaOfMantissa)

	fmt.Printf("\nMantissa of Mantissa Test #4-a\n"+
		"Original tFloat01 = %v\n\n"+
		"mantissaOfMantissa = %v\n\n"+
		"mantissaOfMantissa precision = %v\n\n"+
		"mantissaOfMantissa exponent = %v\n\n",
		tFloat01.Text('f', -1),
		mantissaOfMantissa.Text('f', -1),
		mantissaOfMantissa.Prec(),
		mantissaOfMantissaExponent)

	fmt.Printf("\n\n%v\n"+
		"   Successful Completion!\n"+
		"Function: %v\n%v\n\n",
		breakStr,
		funcName,
		breakStr)

}

func MantissaExponent02() {

	funcName := "MantissaExponent02()"

	breakStr := strings.Repeat("=", 70)
	breakStr2 := strings.Repeat("-", 60)
	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	tFloat01 := new(big.Float).
		SetMode(big.AwayFromZero).
		SetFloat64(1234.567890)

	mantissa := big.NewFloat(0)

	exponent := tFloat01.MantExp(mantissa)

	fmt.Printf("\n\n%v\n\tOriginal Value and Mantissa\n"+
		"tFloat01  = %v\n"+
		" mantissa = %v\n"+
		"precision = %v\n"+
		" exponent = %v\n"+
		"%v\n\n",
		breakStr2,
		tFloat01.Text('f', -1),
		mantissa.Text('f', -1),
		mantissa.Prec(),
		exponent,
		breakStr2)

	mantissa.SetMantExp(mantissa, int(mantissa.MinPrec()))

	newMantissa := big.NewFloat(0)

	newExponent := mantissa.MantExp(newMantissa)

	newMantissa.SetMantExp(newMantissa, int(newMantissa.MinPrec()))

	newMantissaBigInt,
		accuracy := newMantissa.Int(nil)

	if accuracy != big.Exact {

		fmt.Printf("\n%v\n"+
			"Error: newMantissaBigInt = newMantissa.Int(nil)\n"+
			"The accuracy returned is NOT 'exact'!\n",
			funcName)

		return
	}

	fmt.Printf("\n\n%v\n\tMinimum Precision Mantissa\n"+
		"new mantissa = %v\n"+
		"new precison = %v\n"+
		"new exponent = %v\n"+
		"new Mantissa Big Int = %v\n"+
		"%v\n\n",
		breakStr2,
		newMantissa.Text('f', -1),
		newMantissa.Prec(),
		newExponent,
		newMantissaBigInt.Text(10),
		breakStr2)

	bigIntExponentValue := big.NewInt(2)

	bigIntExponentValue.Exp(
		bigIntExponentValue,
		big.NewInt(int64(newExponent)), nil)

	newBigIntValue := big.NewInt(0).Mul(
		newMantissaBigInt,
		bigIntExponentValue)

	fmt.Printf("\n\n%v\n\tBig Int Mantissa\n"+
		"  newMantissaBigInt = %v\n"+
		"       new Mantissa = %v\n"+
		"  New Big Int Value = %v\n"+
		"%v\n\n",
		breakStr2,
		newMantissaBigInt.Text(10),
		newMantissa.Text('f', -1),
		newBigIntValue.Text(10),
		breakStr2)

	fmt.Printf("\n\n%v\n"+
		"   Successful Completion!\n"+
		"Function: %v\n%v\n\n",
		breakStr,
		funcName,
		breakStr)
}

func MantissaExponent03(
	baseNumStr string) {

	funcName := "MantissaExponent03()"

	breakStr := strings.Repeat("=", 70)

	breakStr2 := strings.Repeat("-", 60)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	tFloat01,
		ok := new(big.Float).
		SetMode(big.AwayFromZero).
		SetString(baseNumStr)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error: tFloat01=SetString(baseNumStr)\n"+
			"SetString Failed!\n"+
			"baseNumStr = %v\n",
			funcName,
			baseNumStr)

		return
	}

	mantissa := big.NewFloat(0)

	exponent := tFloat01.MantExp(mantissa)

	tFloat02 := big.NewFloat(0.0)

	tFloat02.Copy(tFloat01)

	tFloat02.SetPrec(tFloat02.MinPrec())

	mantissa2 := big.NewFloat(0.0)

	exponent2 := tFloat02.MantExp(mantissa2)

	fmt.Printf("\n%v\n"+
		"\tExtraction Test #1-a\n"+
		"           tFloat01 = %v\n"+
		"           mantissa = %v\n"+
		" mantissa precision = %v\n"+
		"           exponent = %v\n"+
		"-----------------------\n"+
		"            tFloat2 = %v\n"+
		"          mantissa2 = %v\n"+
		"mantissa2 precision = %v\n"+
		"          exponent2 = %v\n"+
		"%v\n\n",
		breakStr2,
		tFloat01.Text('f', -1),
		mantissa.Text('f', -1),
		mantissa.Prec(),
		exponent,
		tFloat02.Text('f', -1),
		mantissa2.Text('f', -1),
		mantissa2.Prec(),
		exponent2,
		breakStr2)

	fracDigitsMantissaStr :=
		mantissa.Text('f', -1)

	startIndex := 2

	if fracDigitsMantissaStr[0] == '-' {
		startIndex = 3
	}

	fracDigitsMantissaStr = fracDigitsMantissaStr[startIndex:]

	lenMantissaFracDigits := len(fracDigitsMantissaStr)

	bIntMantissa,
		ok := big.NewInt(0).SetString(
		fracDigitsMantissaStr, 10)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error: bIntMantissa=SetString(fracDigitsMantissaStr)\n"+
			"SetString Failed!\n"+
			"fracDigitsMantissaStr = %v\n",
			funcName,
			fracDigitsMantissaStr)

		return
	}

	bIntExponent := big.NewInt(int64(exponent))

	bIntExponentVal := big.NewInt(2)

	bIntExponentVal.Exp(
		bIntExponentVal,
		bIntExponent,
		nil)

	bIntNumberVal := big.NewInt(0).Mul(
		bIntMantissa,
		bIntExponentVal)

	bIntNumValStr :=
		bIntNumberVal.Text(10)

	bIntNumValStrTotalLen := len(bIntNumValStr)

	intDigitsLen := bIntNumValStrTotalLen -
		lenMantissaFracDigits

	floatingPointStr :=
		bIntNumValStr[0:intDigitsLen] +
			"." +
			bIntNumValStr[intDigitsLen:]

	fmt.Printf("\nMantissa-1\n%v\n"+
		"         bIntMantissa = %v\n"+
		"      bIntExponentVal = %v\n"+
		"        bIntNumberVal = %v\n"+
		"bIntNumValStrTotalLen = %v\n"+
		" Mantissa Frac Digits = %v\n"+
		"Calculated Float Str  = %v\n"+
		"%v\n\n",
		breakStr2,
		bIntMantissa.Text(10),
		bIntExponentVal.Text(10),
		bIntNumberVal.Text(10),
		bIntNumValStrTotalLen,
		lenMantissaFracDigits,
		floatingPointStr,
		breakStr2)

	fracDigitsMantissa2Str :=
		mantissa2.Text('f', -1)

	startIndex = 2

	if fracDigitsMantissa2Str[0] == '-' {
		startIndex = 3
	}

	fracDigitsMantissa2Str = fracDigitsMantissa2Str[startIndex:]

	lenMantissa2FracDigits := len(fracDigitsMantissa2Str)

	bIntMantissa2,
		ok := big.NewInt(0).SetString(
		fracDigitsMantissa2Str, 10)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error: bIntMantissa2=SetString(fracDigitsMantissa2Str)\n"+
			"SetString Failed!\n"+
			"fracDigitsMantissa2tr = %v\n",
			funcName,
			fracDigitsMantissa2Str)

		return
	}

	bIntExponent2 := big.NewInt(int64(exponent2))

	bIntExponent2Val := big.NewInt(2)

	bIntExponent2Val.Exp(
		bIntExponent2Val,
		bIntExponent2,
		nil)

	bIntNumber2Val := big.NewInt(0).Mul(
		bIntMantissa2,
		bIntExponent2Val)

	bIntNumVal2Str :=
		bIntNumber2Val.Text(10)

	bIntNumVal2StrTotalLen := len(bIntNumVal2Str)

	intDigits2Len := bIntNumVal2StrTotalLen -
		lenMantissa2FracDigits

	floatingPoint2Str :=
		bIntNumVal2Str[0:intDigits2Len] +
			"." +
			bIntNumVal2Str[intDigits2Len:]
	//
	//bIntExponent2Val := big.NewInt(2)
	//
	//bIntExponent2Val.Exp(
	//	bIntExponent2Val,
	//	bIntExponent2,
	//	nil)
	//
	//bIntNumber2Val := big.NewInt(0).Mul(
	//	bIntMantissa2,
	//	bIntExponent2Val)
	//
	//bIntNumVal2Str =
	//	bIntNumber2Val.Text(10)
	//
	//bIntNumVal2StrTotalLen = len(bIntNumValStr)
	//
	//intDigits2Len = bIntNumValStrTotalLen -
	//	lenMantissaFracDigits

	fmt.Printf("\nMantissa-2\n%v\n"+
		"              tFloat2 = %v\n"+
		"            mantissa2 = %v\n"+
		"  mantissa2 precision = %v\n"+
		"   mantissa2 exponent = %v\n"+
		"         bIntMantissa2 = %v\n"+
		"      bIntExponent2Val = %v\n"+
		"        bIntNumber2Val = %v\n"+
		"bIntNumVal2StrTotalLen = %v\n"+
		" Mantissa2 Frac Digits = %v\n"+
		"Calculated Float Str2  = %v\n"+
		"%v\n\n",
		breakStr2,
		tFloat02.Text('f', -1),
		mantissa2.Text('f', -1),
		mantissa2.Prec(),
		exponent2,
		bIntMantissa2.Text(10),
		bIntExponent2Val.Text(10),
		bIntNumber2Val.Text(10),
		bIntNumVal2StrTotalLen,
		lenMantissa2FracDigits,
		floatingPoint2Str,
		breakStr2)

	fmt.Printf("\n\n%v\n"+
		"   Successful Completion!\n"+
		"Function: %v\n%v\n\n",
		breakStr,
		funcName,
		breakStr)

	return
}

func MantissaExponent04(
	baseNumStr string) {

	funcName := "MantissaExponent04()"

	breakStr := strings.Repeat("=", 70)

	breakStr2 := strings.Repeat("-", 60)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	var ok bool

	tFloat02 := big.NewFloat(0.0)

	_,
		ok = tFloat02.
		SetString(baseNumStr)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error: tFloat01=SetString(baseNumStr)\n"+
			"SetString Failed!\n"+
			"baseNumStr = %v\n",
			funcName,
			baseNumStr)

		return
	}

	tFloat02.SetPrec(tFloat02.MinPrec())

	mantissa := big.NewFloat(0)

	exponent := tFloat02.MantExp(mantissa)

	fmt.Printf("\n%v\n"+
		"\tExtraction Test #1-a\n"+
		"            tFloat2 = %v\n"+
		"           mantissa = %v\n"+
		" mantissa precision = %v\n"+
		"           exponent = %v\n"+
		"%v\n\n",
		breakStr2,
		tFloat02.Text('f', -1),
		mantissa.Text('f', -1),
		mantissa.Prec(),
		exponent,
		breakStr2)

	fracDigitsMantissaStr :=
		mantissa.Text('f', -1)

	startIndex := 2

	if fracDigitsMantissaStr[0] == '-' {
		startIndex = 3
	}

	fracDigitsMantissaStr = fracDigitsMantissaStr[startIndex:]

	lenMantissaFracDigits := len(fracDigitsMantissaStr)

	bIntMantissa,
		ok := big.NewInt(0).SetString(
		fracDigitsMantissaStr, 10)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error: bIntMantissa=SetString(fracDigitsMantissaStr)\n"+
			"SetString Failed!\n"+
			"fracDigitsMantissaStr = %v\n",
			funcName,
			fracDigitsMantissaStr)

		return
	}

	bIntExponent := big.NewInt(int64(exponent))

	bIntExponentVal := big.NewInt(2)

	bIntExponentVal.Exp(
		bIntExponentVal,
		bIntExponent,
		nil)

	bIntNumberVal := big.NewInt(0).Mul(
		bIntMantissa,
		bIntExponentVal)

	bIntNumValStr :=
		bIntNumberVal.Text(10)

	bIntNumValStrTotalLen := len(bIntNumValStr)

	intDigitsLen := bIntNumValStrTotalLen -
		lenMantissaFracDigits

	floatingPointStr :=
		bIntNumValStr[0:intDigitsLen] +
			"." +
			bIntNumValStr[intDigitsLen:]

	fmt.Printf("\nMantissa-1\n%v\n"+
		"         bIntMantissa = %v\n"+
		"      bIntExponentVal = %v\n"+
		"        bIntNumberVal = %v\n"+
		"bIntNumValStrTotalLen = %v\n"+
		" Mantissa Frac Digits = %v\n"+
		"Calculated Float Str  = %v\n"+
		"%v\n\n",
		breakStr2,
		bIntMantissa.Text(10),
		bIntExponentVal.Text(10),
		bIntNumberVal.Text(10),
		bIntNumValStrTotalLen,
		lenMantissaFracDigits,
		floatingPointStr,
		breakStr2)

	return
}

func GetPrecisionFactor() (
	*big.Float,
	error) {

	funcName := "GetPrecisionFactor"
	var err error
	precToDigitsFactor := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	// 3.3219789132197891321978913219789

	_,
		ok := precToDigitsFactor.SetString("3.3219789132197891321978913219789")

	if !ok {
		err = fmt.Errorf("\n%v\n"+
			"expectedNthRoot.SetString(\"3.3219789132197891321978913219789\") FAILED!\n",
			funcName)

	}

	return precToDigitsFactor, err
}
