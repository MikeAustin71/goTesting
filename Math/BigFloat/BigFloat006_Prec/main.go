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

	InitializeBigFloatFirstPlace()
}

func PositiveBigFloat() {

	funcName := "PositiveBigFloat"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	floatStrVal := "12345.678901"

	var bigFloatVal *big.Float
	var ok bool

	bigFloatVal,
		ok = big.NewFloat(0).SetString(floatStrVal)

	if !ok {
		fmt.Printf("Error: big.NewFloat(0).SetString(floatStrVal)")

		return
	}

	fmt.Printf("Big Float String Value:         '%v'\n",
		floatStrVal)

	fmt.Printf(breakStr + "\n\n")

	initialPrec := bigFloatVal.Prec()

	fmt.Printf("Big Float Intial Precision:     '%v'\n",
		initialPrec)

	minimumPrec := bigFloatVal.MinPrec()

	fmt.Printf("Big Float Intial Min Precision: '%v'\n",
		minimumPrec)

	fmt.Printf("Big Float Intial Mode: '%v'\n",
		bigFloatVal.Mode().String())

	fmt.Printf("Big Float Intial Accuracy:      '%v'\n",
		bigFloatVal.Acc().String())

	fmt.Printf("Big Float Numeric Value g -1:   '%v'\n",
		bigFloatVal.Text('g', -1))

	fmt.Printf("Big Float Numeric Value f -1:   '%v'\n",
		bigFloatVal.Text('f', -1))

	fmt.Printf("Big Float Val f 6:              '%v'\n",
		bigFloatVal.Text('f', 6))

	fmt.Printf("Big Float Val f 7:              '%v'\n",
		bigFloatVal.Text('f', 7))

	fmt.Printf("Big Float Val f 11g:            '%.11g'\n",
		bigFloatVal)

	fmt.Printf("Big Float Val f  20 Prec :      '%v'\n",
		bigFloatVal.Text('f', 20))

	fmt.Printf(breakStr + "\n\n")

	fmt.Printf("Setting Precision to Minimum Precision\n"+
		"Minimum Precision = '%v'\n",
		minimumPrec)

	fmt.Printf("Setting Rounding to Away From Zero\n"+
		"Mode = '%v'\n",
		big.AwayFromZero.String())

	bigFloatVal.SetPrec(minimumPrec)
	bigFloatVal.SetMode(big.AwayFromZero)

	fmt.Printf("Big Float MinimumPrec Accuracy: '%v'\n",
		bigFloatVal.Acc().String())

	fmt.Printf("Big Float Rounding Mode:        '%v'\n",
		bigFloatVal.Mode().String())

	fmt.Printf("Big Float Numeric Value g -1:   '%v'\n",
		bigFloatVal.Text('g', -1))

	fmt.Printf("Big Float Val Min Prec f -1:    '%v'\n",
		bigFloatVal.Text('f', -1))

	fmt.Printf("Big Float Val f 6:              '%v'\n",
		bigFloatVal.Text('f', 6))

	fmt.Printf("Big Float Val f 7:              '%v'\n",
		bigFloatVal.Text('f', 7))

	fmt.Printf("Big Float Val  f 11g:           '%.11g'\n",
		bigFloatVal)

	fmt.Printf("Big Float Val f  20 Prec :      '%v'\n",
		bigFloatVal.Text('f', 20))

	fmt.Printf(breakStr + "\n\n")

	fmt.Printf("Setting Precison to 4096\n")
	fmt.Printf("Re-Setting Mode to 'Away From Zero'\n")

	bigFloatVal.SetPrec(4096)
	bigFloatVal.SetMode(big.AwayFromZero)

	fmt.Printf("Big Float 4096 Prec Accuracy:   '%v'\n",
		bigFloatVal.Acc().String())

	fmt.Printf("Big Float Rounding Mode:        '%v'\n",
		bigFloatVal.Mode().String())

	fmt.Printf("Big Float Val   4096 Prec f -1: '%v'\n",
		bigFloatVal.Text('f', -1))

	fmt.Printf("Big Float Val f 6:              '%v'\n",
		bigFloatVal.Text('f', 6))

	fmt.Printf("Big Float Val f 7:              '%v'\n",
		bigFloatVal.Text('f', 7))

	fmt.Printf("Big Float Val 4096 Prec f 11g:  '%.11g'\n",
		bigFloatVal)

	fmt.Printf("Big Float Val f  20 Prec :      '%v'\n",
		bigFloatVal.Text('f', 20))

	fmt.Printf("\n\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")
}

func InitializeBigFloatSecondPlace() {

	funcName := "InitializeBigFloatSecondPlace()"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	numValueStr := "123456.78901234567890"
	fmt.Printf("t2 Example With Set To Minimum Precision.\n\n")

	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.
	var t2 big.Float
	var ok bool
	_,
		ok = t2.SetString(numValueStr)

	if !ok {
		fmt.Printf("Error: t.SetString(numValueStr) Failed!\n")

		return
	}

	t2.SetMode(big.AwayFromZero)
	t2InitialVal := t2.Text('f', -1)

	fmt.Printf("Values Immediately After Initialization\n"+
		"    Starting Value Num Str: '%v'\n"+
		"\n ---- Before Minimum Precision ----\n"+
		"            t2 Intial Mode: '%v'\n"+
		"        t2 Intial Accuracy: '%v'\n"+
		"      t3 Initial Precision: '%v'\n"+
		"t2 Initial Value (Prec=-1): '%v'\n",
		numValueStr,
		t2.Mode().String(),
		t2.Acc().String(),
		t2.Prec(),
		t2InitialVal)

	t2.SetPrec(t2.MinPrec())

	t2InitialVal = t2.Text('f', -1)

	fmt.Printf(
		"\n ----- After Minimum Precision -----\n"+
			"      t2 Post Minimum Mode: '%v'\n"+
			"  t2 Post Minimum Accuracy: '%v'\n"+
			" t2 Post Minimum Precision: '%v'\n"+
			"     t2 Post Minimum Value: '%v'\n\n",
		t2.Mode().String(),
		t2.Acc().String(),
		t2.Prec(),
		t2InitialVal)

	// Precision as used here is the number decimal places
	fmt.Printf("\n ----- After Minimum Precision2 -----\n"+
		"t2 Numeric Value Text f 14: '%v'\n"+
		"t2 Numeric Value Printf .14f: '%.14f'\n",
		&t2,
		&t2)

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}

func BigFloatFromPureNumStr(
	pureNumberValueStr string,
	numOfExtraDigitsPrecisionBuffer int64) BigFloatDto {

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
			bigFloatDto.EstimatedPrecisionBits.NumOfExtraDigitsBuffer)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"Estimated Required Precision Bits\n"+
			"Error returned by calculateRequiredPrecision8()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return bigFloatDto
	}

	bigFloatDto.Value.SetInt64(0).
		SetPrec(bigFloatDto.EstimatedPrecisionBits.
			EstimatedNumberOfPrecisionBits).
		SetMode(big.AwayFromZero)

	var ok bool
	_,
		ok = bigFloatDto.Value.SetString(pureNumberValueStr)

	if !ok {
		fmt.Printf("Error: bigFloatDto.Value.SetString(numValueStr) Failed!\n")

		return bigFloatDto
	}

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

	return bigFloatDto
}

// InitializeBigFloatFirstPlace
//
// Best approach so far for transferring Big Float
// values for processing.
func InitializeBigFloatFirstPlace() {

	funcName := "InitializeBigFloatFirstPlace()"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	numValueStr := "123456.789012345678901"

	var err error

	initialBigFloatDto :=
		BigFloatFromPureNumStr(
			numValueStr,
			10)

	//t1.SetMode(big.AwayFromZero)

	initialVal := initialBigFloatDto.Value.Text('f', -1)

	t1Value := BigFloatFromPureNumStr(
		initialVal,
		10)

	fmt.Printf("Values Immediately After Initialization\n"+
		"\t\t------ t1 Values ------\n"+
		"       Starting Value Num Str: '%v'\n"+
		"    Starting Value Int Digits: '%v'\n"+
		"   Starting Value Frac Digits: '%v'\n"+
		"    Starting Val Total Digits: '%v'\n"+
		"NumStr Est Required Precision: '%v'\n"+
		"\n ---- Before Minimum Precision ----\n"+
		"               t1 Intial Mode: '%v'\n"+
		"      Starting Value Accuracy: '%v'\n"+
		"           t1 Intial Accuracy: '%v'\n"+
		"         t1 Initial Precision: '%v'\n"+
		"    t1 Est Required Precision: '%v'\n"+
		"   t1 Initial Value (Prec=-1): '%v'\n"+
		"  t1 Initial Value Int Digits: '%v'\n"+
		" t1 Initial Value Frac Digits: '%v'\n"+
		"t1 Initial Value Total Digits: '%v'\n",
		numValueStr,
		initialBigFloatDto.EstimatedPrecisionBits.NumIntegerDigits,
		initialBigFloatDto.EstimatedPrecisionBits.NumFractionalDigits,
		initialBigFloatDto.GetTotalNumericDigits(),
		initialBigFloatDto.EstimatedPrecisionBits.EstimatedNumberOfPrecisionBits,
		t1Value.Value.Mode().String(),
		initialBigFloatDto.Value.Acc().String(),
		t1Value.Value.Acc(),
		t1Value.Value.Prec(),
		t1Value.EstimatedPrecisionBits.EstimatedNumberOfPrecisionBits,
		t1Value.Value.Text('f', -1),
		t1Value.EstimatedPrecisionBits.NumIntegerDigits,
		t1Value.EstimatedPrecisionBits.NumFractionalDigits,
		t1Value.GetTotalNumericDigits())

	t1Value.Value.SetPrec(t1Value.Value.MinPrec())

	initialVal = t1Value.Value.Text('f', -1)

	t1bValue := BigFloatFromPureNumStr(
		initialVal,
		10)

	fmt.Printf(
		"\n ----- After Minimum Precision -----\n"+
			"        t1 Post Minimum Mode: '%v'\n"+
			"    t1 Post Minimum Accuracy: '%v'\n"+
			"   t1 Post Minimum Precision: '%v'\n"+
			"   t1 Est Required Precision: '%v'\n"+
			"  t1 Post Min Value ('f' -1): '%v'\n"+
			"  t1 Post Min Val Int Digits: '%v'\n"+
			" t1 Post Min Val Frac Digits: '%v'\n"+
			"t1 Post Min Val Total Digits: '%v'\n"+
			"\n\n",
		t1bValue.Value.Mode().String(),
		t1bValue.Value.Acc().String(),
		t1bValue.Value.Prec(),
		t1bValue.EstimatedPrecisionBits.EstimatedNumberOfPrecisionBits,
		initialVal,
		t1bValue.ActualNumStrComponents.NumIntegerDigits,
		t1bValue.ActualNumStrComponents.NumFractionalDigits,
		t1bValue.GetTotalNumericDigits())

	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.
	var t2 big.Float

	//t2.Copy(&t1)
	t2.SetMode(big.AwayFromZero)
	t2.SetPrec(t1bValue.EstimatedPrecisionBits.EstimatedNumberOfPrecisionBits)
	t2.Set(&t1Value.Value)

	t2InitialVal := t2.Text('f', -1)

	var pureNumStrStats02a PureNumberStrComponents

	pureNumStrStats02a,
		err = breakNumStrToComponents(
		numValueStr)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"Call #1"+
			"pureNumStrStats02a\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	var estimatedRequiredPrecision uint

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision8(
			pureNumStrStats02a.NumIntegerDigits,
			pureNumStrStats02a.NumFractionalDigits,
			0)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision8()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf("Values Immediately After Initialization\n"+
		"\t\t------ t2 Values ------\n"+
		"    Starting Value Num Str: '%v'\n"+
		" Starting Value Int Digits: '%v'\n"+
		"Starting Value Frac Digits: '%v'\n"+
		" Starting Val Total Digits: '%v'\n"+
		"\n ---- Before Minimum Precision ----\n"+
		"            t2 Intial Mode: '%v'\n"+
		"        t2 Intial Accuracy: '%v'\n"+
		"      t2 Initial Precision: '%v'\n"+
		" t1 Est Required Precision: '%v'\n"+
		"t2 Initial Value (Prec=-1): '%v'\n",
		t2InitialVal,
		pureNumStrStats02a.NumIntegerDigits,
		pureNumStrStats02a.NumFractionalDigits,
		pureNumStrStats02a.NumIntegerDigits+
			pureNumStrStats02a.NumFractionalDigits,
		t2.Mode().String(),
		t2.Acc().String(),
		t2.Prec(),
		estimatedRequiredPrecision,
		t2InitialVal)

	t2.SetPrec(t2.MinPrec())

	t2InitialVal = t2.Text('f', -1)

	var pureNumStrStats02b PureNumberStrComponents

	pureNumStrStats02b,
		err = breakNumStrToComponents(
		t2InitialVal)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats02b\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision8(
			pureNumStrStats02b.NumIntegerDigits,
			pureNumStrStats02b.NumFractionalDigits,
			0)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision8()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf(
		"\n ----- After Minimum Precision -----\n"+
			"        t2 Post Minimum Mode: '%v'\n"+
			"    t2 Post Minimum Accuracy: '%v'\n"+
			"   t2 Post Minimum Precision: '%v'\n"+
			"   t2 Est Required Precision: '%v'\n"+
			"       t2 Post Minimum Value: '%v'\n"+
			"  t2 Post Min Val Int Digits: '%v'\n"+
			" t2 Post Min Val Frac Digits: '%v'\n"+
			"t2 Post Min Val Total Digits: '%v'\n"+
			"\n\n",
		t2.Mode().String(),
		t2.Acc().String(),
		t2.Prec(),
		estimatedRequiredPrecision,
		t2InitialVal,
		pureNumStrStats02b.NumIntegerDigits,
		pureNumStrStats02b.NumFractionalDigits,
		pureNumStrStats02b.NumIntegerDigits+
			pureNumStrStats02b.NumFractionalDigits)

	//printfFmt1 :=
	//	fmt.Sprintf(".%vf: '\\%."+"%v"+"f'\n",
	//		pureNumStrStats02b.NumFractionalDigits,
	//		pureNumStrStats02b.NumFractionalDigits)

	/*	printFmt2 :=
		fmt.Sprintf(printfFmt1,
			&t2)
	*/

	// Precision as used here is the number decimal places
	/*	fmt.Printf("\n ----- After Minimum Precision2 -----\n"+
		"t2 Numeric Value Text f -1: '%v'\n"+
		"t2 Numeric Value Printf printfFmt1",
		t2.Text('f', -1),
		pureNumStrStats02b.NumFractionalDigits,
		&t2)
	*/

	fmt.Printf("\n ----- After Minimum Precision2 -----\n"+
		"                     t2 Everything f -1: '%v'\n\n"+
		"  t2 All Current Fractional Digits f %v: '%v'\n"+
		"t2 Original Number Fractional Digits %v: '%v'\n ",
		t2.Text('f', -1),
		pureNumStrStats02b.NumFractionalDigits,
		t2.Text('f', int(pureNumStrStats02b.NumFractionalDigits)),
		initialBigFloatDto.ActualNumStrComponents.NumFractionalDigits,
		t2.Text('f',
			int(initialBigFloatDto.ActualNumStrComponents.NumFractionalDigits)))

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}

// CopyPreInitializeBigFloat03
func CopyPreInitializeBigFloat03() {

	funcName := "CopyPreInitializeBigFloat03()"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	numValueStr := "123456.789012345678901"

	var pureNumStrStats01a PureNumberStrComponents
	var err error

	pureNumStrStats01a,
		err = breakNumStrToComponents(
		numValueStr)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	var estimatedRequiredPrecisionNumStr uint

	estimatedRequiredPrecisionNumStr,
		err =
		calculateRequiredPrecision8(
			pureNumStrStats01a.NumIntegerDigits,
			pureNumStrStats01a.NumFractionalDigits,
			10)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"estimatedRequiredPrecisionNumStr\n"+
			"Error returned by calculateRequiredPrecision8()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	//t1 := big.Float{}
	t1 := big.NewFloat(0).
		SetPrec(estimatedRequiredPrecisionNumStr).
		SetMode(big.AwayFromZero)

	var ok bool
	_,
		ok = t1.SetString(numValueStr)

	if !ok {
		fmt.Printf("Error: t1.SetString(numValueStr) Failed!\n")

		return
	}

	floatAccuracy := t1.Acc()

	//t1.SetMode(big.AwayFromZero)

	t1InitialVal := t1.Text('f', -1)

	var estimatedRequiredPrecision uint

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision8(
			pureNumStrStats01a.NumIntegerDigits,
			pureNumStrStats01a.NumFractionalDigits,
			2)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision8()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	var purNumStrStats01a2 PureNumberStrComponents

	purNumStrStats01a2,
		err = breakNumStrToComponents(
		t1InitialVal)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"purNumStrStats01a2\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf("Values Immediately After Initialization\n"+
		"\t\t------ t1 Values ------\n"+
		"       Starting Value Num Str: '%v'\n"+
		"    Starting Value Int Digits: '%v'\n"+
		"   Starting Value Frac Digits: '%v'\n"+
		"    Starting Val Total Digits: '%v'\n"+
		"NumStr Est Required Precision: '%v'\n"+
		"\n ---- Before Minimum Precision ----\n"+
		"               t1 Intial Mode: '%v'\n"+
		"      Starting Value Accuracy: '%v'\n"+
		"           t1 Intial Accuracy: '%v'\n"+
		"         t1 Initial Precision: '%v'\n"+
		"    t1 Est Required Precision: '%v'\n"+
		"   t1 Initial Value (Prec=-1): '%v'\n"+
		"      t1 Initial Value Digits: '%v'\n"+
		" t1 Initial Value Frac Digits: '%v'\n"+
		"t1 Initial Value Total Digits: '%v'\n",
		numValueStr,
		pureNumStrStats01a.NumIntegerDigits,
		pureNumStrStats01a.NumFractionalDigits,
		pureNumStrStats01a.NumIntegerDigits+
			pureNumStrStats01a.NumFractionalDigits,
		estimatedRequiredPrecisionNumStr,
		t1.Mode().String(),
		floatAccuracy,
		t1.Acc().String(),
		t1.Prec(),
		estimatedRequiredPrecision,
		t1InitialVal,
		purNumStrStats01a2.NumIntegerDigits,
		purNumStrStats01a2.NumFractionalDigits,
		purNumStrStats01a2.NumIntegerDigits+
			purNumStrStats01a2.NumFractionalDigits)

	t1.SetPrec(t1.MinPrec())

	t1InitialVal = t1.Text('f', -1)

	var pureNumStrStats01b PureNumberStrComponents

	pureNumStrStats01b,
		err = breakNumStrToComponents(
		t1InitialVal)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01b\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision8(
			pureNumStrStats01b.NumIntegerDigits,
			pureNumStrStats01b.NumFractionalDigits,
			0)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision8()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf(
		"\n ----- After Minimum Precision -----\n"+
			"        t1 Post Minimum Mode: '%v'\n"+
			"    t1 Post Minimum Accuracy: '%v'\n"+
			"   t1 Post Minimum Precision: '%v'\n"+
			"   t1 Est Required Precision: '%v'\n"+
			"       t1 Post Minimum Value: '%v'\n"+
			"  t1 Post Min Val Int Digits: '%v'\n"+
			" t1 Post Min Val Frac Digits: '%v'\n"+
			"t1 Post Min Val Total Digits: '%v'\n"+
			"\n\n",
		t1.Mode().String(),
		t1.Acc().String(),
		t1.Prec(),
		estimatedRequiredPrecision,
		t1InitialVal,
		pureNumStrStats01b.NumIntegerDigits,
		pureNumStrStats01b.NumFractionalDigits,
		pureNumStrStats01b.NumIntegerDigits+
			pureNumStrStats01b.NumFractionalDigits)

	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.
	var t2 big.Float

	//t2.Copy(&t1)
	t2.SetMode(big.AwayFromZero)
	t2.SetPrec(estimatedRequiredPrecisionNumStr)
	t2.Set(t1)

	t2InitialVal := t2.Text('f', -1)

	var pureNumStrStats02a PureNumberStrComponents

	pureNumStrStats02a,
		err = breakNumStrToComponents(
		numValueStr)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"Call #1"+
			"pureNumStrStats02a\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision8(
			pureNumStrStats02a.NumIntegerDigits,
			pureNumStrStats02a.NumFractionalDigits,
			0)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision8()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf("Values Immediately After Initialization\n"+
		"\t\t------ t2 Values ------\n"+
		"    Starting Value Num Str: '%v'\n"+
		" Starting Value Int Digits: '%v'\n"+
		"Starting Value Frac Digits: '%v'\n"+
		" Starting Val Total Digits: '%v'\n"+
		"\n ---- Before Minimum Precision ----\n"+
		"            t2 Intial Mode: '%v'\n"+
		"        t2 Intial Accuracy: '%v'\n"+
		"      t2 Initial Precision: '%v'\n"+
		" t1 Est Required Precision: '%v'\n"+
		"t2 Initial Value (Prec=-1): '%v'\n",
		t2InitialVal,
		pureNumStrStats02a.NumIntegerDigits,
		pureNumStrStats02a.NumFractionalDigits,
		pureNumStrStats02a.NumIntegerDigits+
			pureNumStrStats02a.NumFractionalDigits,
		t2.Mode().String(),
		t2.Acc().String(),
		t2.Prec(),
		estimatedRequiredPrecision,
		t2InitialVal)

	t2.SetPrec(t2.MinPrec())

	t2InitialVal = t2.Text('f', -1)

	var pureNumStrStats02b PureNumberStrComponents

	pureNumStrStats02b,
		err = breakNumStrToComponents(
		t2InitialVal)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats02b\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision8(
			pureNumStrStats02b.NumIntegerDigits,
			pureNumStrStats02b.NumFractionalDigits,
			0)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision8()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf(
		"\n ----- After Minimum Precision -----\n"+
			"        t2 Post Minimum Mode: '%v'\n"+
			"    t2 Post Minimum Accuracy: '%v'\n"+
			"   t2 Post Minimum Precision: '%v'\n"+
			"   t2 Est Required Precision: '%v'\n"+
			"       t2 Post Minimum Value: '%v'\n"+
			"  t2 Post Min Val Int Digits: '%v'\n"+
			" t2 Post Min Val Frac Digits: '%v'\n"+
			"t2 Post Min Val Total Digits: '%v'\n"+
			"\n\n",
		t2.Mode().String(),
		t2.Acc().String(),
		t2.Prec(),
		estimatedRequiredPrecision,
		t2InitialVal,
		pureNumStrStats02b.NumIntegerDigits,
		pureNumStrStats02b.NumFractionalDigits,
		pureNumStrStats02b.NumIntegerDigits+
			pureNumStrStats02b.NumFractionalDigits)

	//printfFmt1 :=
	//	fmt.Sprintf(".%vf: '\\%."+"%v"+"f'\n",
	//		pureNumStrStats02b.NumFractionalDigits,
	//		pureNumStrStats02b.NumFractionalDigits)

	/*	printFmt2 :=
		fmt.Sprintf(printfFmt1,
			&t2)
	*/

	// Precision as used here is the number decimal places
	/*	fmt.Printf("\n ----- After Minimum Precision2 -----\n"+
		"t2 Numeric Value Text f -1: '%v'\n"+
		"t2 Numeric Value Printf printfFmt1",
		t2.Text('f', -1),
		pureNumStrStats02b.NumFractionalDigits,
		&t2)
	*/

	fmt.Printf("\n ----- After Minimum Precision2 -----\n"+
		"                     t2 Everything f -1: '%v'\n\n"+
		"  t2 All Current Fractional Digits f %v: '%v'\n"+
		"t2 Original Number Fractional Digits %v: '%v'\n ",
		t2.Text('f', -1),
		pureNumStrStats02b.NumFractionalDigits,
		t2.Text('f', int(pureNumStrStats02b.NumFractionalDigits)),
		pureNumStrStats01a.NumFractionalDigits,
		t2.Text('f', int(pureNumStrStats01a.NumFractionalDigits)))

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}

func CopyPreInitializeBigFloat02() {

	funcName := "CopyPreInitializeBigFloat02()"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	numValueStr := "123456.789012345678901"

	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.
	var t1 big.Float

	var pureNumStrStats01a PureNumberStrComponents
	var err error

	pureNumStrStats01a,
		err = breakNumStrToComponents(
		numValueStr)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	var estimatedRequiredPrecision uint

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision2(
			pureNumStrStats01a.NumIntegerDigits,
			pureNumStrStats01a.NumFractionalDigits,
			0)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision2()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	//t1.SetPrec(estimatedRequiredPrecision)

	// t1.SetMode(big.AwayFromZero)

	var ok bool
	_,
		ok = t1.SetString(numValueStr)

	if !ok {
		fmt.Printf("Error: t1.SetString(numValueStr) Failed!\n")

		return
	}

	floatAccuracy := t1.Acc()

	//t1.SetMode(big.AwayFromZero)

	t1InitialVal := t1.Text('f', -1)

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision2(
			pureNumStrStats01a.NumIntegerDigits,
			pureNumStrStats01a.NumFractionalDigits,
			2)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision2()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	var purNumStrStats01a2 PureNumberStrComponents

	purNumStrStats01a2,
		err = breakNumStrToComponents(
		t1InitialVal)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"purNumStrStats01a2\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf("Values Immediately After Initialization\n"+
		"\t\t------ t1 Values ------\n"+
		"       Starting Value Num Str: '%v'\n"+
		"    Starting Value Int Digits: '%v'\n"+
		"   Starting Value Frac Digits: '%v'\n"+
		"    Starting Val Total Digits: '%v'\n"+
		"\n ---- Before Minimum Precision ----\n"+
		"               t1 Intial Mode: '%v'\n"+
		"      Starting Value Accuracy: '%v'\n"+
		"           t1 Intial Accuracy: '%v'\n"+
		"         t1 Initial Precision: '%v'\n"+
		"    t1 Est Required Precision: '%v'\n"+
		"   t1 Initial Value (Prec=-1): '%v'\n"+
		"      t1 Initial Value Digits: '%v'\n"+
		" t1 Initial Value Frac Digits: '%v'\n"+
		"t1 Initial Value Total Digits: '%v'\n",
		numValueStr,
		pureNumStrStats01a.NumIntegerDigits,
		pureNumStrStats01a.NumFractionalDigits,
		pureNumStrStats01a.NumIntegerDigits+
			pureNumStrStats01a.NumFractionalDigits,
		t1.Mode().String(),
		floatAccuracy,
		t1.Acc().String(),
		t1.Prec(),
		estimatedRequiredPrecision,
		t1InitialVal,
		purNumStrStats01a2.NumIntegerDigits,
		purNumStrStats01a2.NumFractionalDigits,
		purNumStrStats01a2.NumIntegerDigits+
			purNumStrStats01a2.NumFractionalDigits)

	t1.SetPrec(t1.MinPrec())

	t1InitialVal = t1.Text('f', -1)

	var pureNumStrStats01b PureNumberStrComponents

	pureNumStrStats01b,
		err = breakNumStrToComponents(
		t1InitialVal)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01b\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision2(
			pureNumStrStats01b.NumIntegerDigits,
			pureNumStrStats01b.NumFractionalDigits,
			0)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision2()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf(
		"\n ----- After Minimum Precision -----\n"+
			"        t1 Post Minimum Mode: '%v'\n"+
			"    t1 Post Minimum Accuracy: '%v'\n"+
			"   t1 Post Minimum Precision: '%v'\n"+
			"   t1 Est Required Precision: '%v'\n"+
			"       t1 Post Minimum Value: '%v'\n"+
			"  t1 Post Min Val Int Digits: '%v'\n"+
			" t1 Post Min Val Frac Digits: '%v'\n"+
			"t1 Post Min Val Total Digits: '%v'\n"+
			"\n\n",
		t1.Mode().String(),
		t1.Acc().String(),
		t1.Prec(),
		estimatedRequiredPrecision,
		t1InitialVal,
		pureNumStrStats01b.NumIntegerDigits,
		pureNumStrStats01b.NumFractionalDigits,
		pureNumStrStats01b.NumIntegerDigits+
			pureNumStrStats01b.NumFractionalDigits)

	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.
	var t2 big.Float

	t2.Copy(&t1)

	t2InitialVal := t2.Text('f', -1)

	var pureNumStrStats02a PureNumberStrComponents

	pureNumStrStats02a,
		err = breakNumStrToComponents(
		numValueStr)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"Call #1"+
			"pureNumStrStats02a\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision2(
			pureNumStrStats02a.NumIntegerDigits,
			pureNumStrStats02a.NumFractionalDigits,
			0)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision2()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf("Values Immediately After Initialization\n"+
		"\t\t------ t2 Values ------\n"+
		"    Starting Value Num Str: '%v'\n"+
		" Starting Value Int Digits: '%v'\n"+
		"Starting Value Frac Digits: '%v'\n"+
		" Starting Val Total Digits: '%v'\n"+
		"\n ---- Before Minimum Precision ----\n"+
		"            t2 Intial Mode: '%v'\n"+
		"        t2 Intial Accuracy: '%v'\n"+
		"      t2 Initial Precision: '%v'\n"+
		" t1 Est Required Precision: '%v'\n"+
		"t2 Initial Value (Prec=-1): '%v'\n",
		t2InitialVal,
		pureNumStrStats02a.NumIntegerDigits,
		pureNumStrStats02a.NumFractionalDigits,
		pureNumStrStats02a.NumIntegerDigits+
			pureNumStrStats02a.NumFractionalDigits,
		t2.Mode().String(),
		t2.Acc().String(),
		t2.Prec(),
		estimatedRequiredPrecision,
		t2InitialVal)

	t2.SetPrec(t2.MinPrec())

	t2InitialVal = t2.Text('f', -1)

	var pureNumStrStats02b PureNumberStrComponents

	pureNumStrStats02b,
		err = breakNumStrToComponents(
		t2InitialVal)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats02b\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision2(
			pureNumStrStats02b.NumIntegerDigits,
			pureNumStrStats02b.NumFractionalDigits,
			0)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision2()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf(
		"\n ----- After Minimum Precision -----\n"+
			"        t2 Post Minimum Mode: '%v'\n"+
			"    t2 Post Minimum Accuracy: '%v'\n"+
			"   t2 Post Minimum Precision: '%v'\n"+
			"   t2 Est Required Precision: '%v'\n"+
			"       t2 Post Minimum Value: '%v'\n"+
			"  t2 Post Min Val Int Digits: '%v'\n"+
			" t2 Post Min Val Frac Digits: '%v'\n"+
			"t2 Post Min Val Total Digits: '%v'\n"+
			"\n\n",
		t2.Mode().String(),
		t2.Acc().String(),
		t2.Prec(),
		estimatedRequiredPrecision,
		t2InitialVal,
		pureNumStrStats02b.NumIntegerDigits,
		pureNumStrStats02b.NumFractionalDigits,
		pureNumStrStats02b.NumIntegerDigits+
			pureNumStrStats02b.NumFractionalDigits)

	// Precision as used here is the number decimal places
	fmt.Printf("\n ----- After Minimum Precision2 -----\n"+
		"t2 Numeric Value Text f -1: '%v'\n"+
		"t2 Numeric Value Printf .14f: '%.14f'\n",
		t2.Text('f', -1),
		&t2)

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}

func CopyPreInitializeBigFloat01() {

	funcName := "CopyPreInitializeBigFloat01()"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	numValueStr := "123456.789012345678901"

	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.
	var t1 big.Float

	var pureNumStrStats01a PureNumberStrComponents
	var err error

	pureNumStrStats01a,
		err = breakNumStrToComponents(
		numValueStr)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	var estimatedRequiredPrecision uint

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision2(
			pureNumStrStats01a.NumIntegerDigits,
			pureNumStrStats01a.NumFractionalDigits,
			2)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision2()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	//t1.SetPrec(estimatedRequiredPrecision)

	// t1.SetMode(big.AwayFromZero)

	var ok bool
	_,
		ok = t1.SetString(numValueStr)

	if !ok {
		fmt.Printf("Error: t1.SetString(numValueStr) Failed!\n")

		return
	}

	//t1.SetMode(big.AwayFromZero)

	t1InitialVal := t1.Text('f', -1)

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision2(
			pureNumStrStats01a.NumIntegerDigits,
			pureNumStrStats01a.NumFractionalDigits,
			2)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision2()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf("Values Immediately After Initialization\n"+
		"\t\t------ t1 Values ------\n"+
		"    Starting Value Num Str: '%v'\n"+
		" Starting Value Int Digits: '%v'\n"+
		"Starting Value Frac Digits: '%v'\n"+
		" Starting Val Total Digits: '%v'\n"+
		"\n ---- Before Minimum Precision ----\n"+
		"            t1 Intial Mode: '%v'\n"+
		"        t1 Intial Accuracy: '%v'\n"+
		"      t1 Initial Precision: '%v'\n"+
		" t1 Est Required Precision: '%v'\n"+
		"t1 Initial Value (Prec=-1): '%v'\n",
		numValueStr,
		pureNumStrStats01a.NumIntegerDigits,
		pureNumStrStats01a.NumFractionalDigits,
		pureNumStrStats01a.NumIntegerDigits+
			pureNumStrStats01a.NumFractionalDigits,
		t1.Mode().String(),
		t1.Acc().String(),
		t1.Prec(),
		estimatedRequiredPrecision,
		t1InitialVal)

	//t1.SetPrec(t1.MinPrec())

	t1InitialVal = t1.Text('f', -1)

	var pureNumStrStats01b PureNumberStrComponents

	pureNumStrStats01b,
		err = breakNumStrToComponents(
		t1InitialVal)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01b\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision2(
			pureNumStrStats01b.NumIntegerDigits,
			pureNumStrStats01b.NumFractionalDigits,
			0)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision2()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf(
		"\n ----- After Minimum Precision -----\n"+
			"        t1 Post Minimum Mode: '%v'\n"+
			"    t1 Post Minimum Accuracy: '%v'\n"+
			"   t1 Post Minimum Precision: '%v'\n"+
			"   t1 Est Required Precision: '%v'\n"+
			"       t1 Post Minimum Value: '%v'\n"+
			"  t1 Post Min Val Int Digits: '%v'\n"+
			" t1 Post Min Val Frac Digits: '%v'\n"+
			"t1 Post Min Val Total Digits: '%v'\n"+
			"\n\n",
		t1.Mode().String(),
		t1.Acc().String(),
		t1.Prec(),
		estimatedRequiredPrecision,
		t1InitialVal,
		pureNumStrStats01b.NumIntegerDigits,
		pureNumStrStats01b.NumFractionalDigits,
		pureNumStrStats01b.NumIntegerDigits+
			pureNumStrStats01b.NumFractionalDigits)

	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.
	var t2 big.Float

	t2.Copy(&t1)
	t2InitialVal := t2.Text('f', -1)

	var pureNumStrStats02a PureNumberStrComponents

	pureNumStrStats02a,
		err = breakNumStrToComponents(
		numValueStr)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"Call #1"+
			"pureNumStrStats02a\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision2(
			pureNumStrStats02a.NumIntegerDigits,
			pureNumStrStats02a.NumFractionalDigits,
			0)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision2()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf("Values Immediately After Initialization\n"+
		"\t\t------ t2 Values ------\n"+
		"    Starting Value Num Str: '%v'\n"+
		" Starting Value Int Digits: '%v'\n"+
		"Starting Value Frac Digits: '%v'\n"+
		" Starting Val Total Digits: '%v'\n"+
		"\n ---- Before Minimum Precision ----\n"+
		"            t2 Intial Mode: '%v'\n"+
		"        t2 Intial Accuracy: '%v'\n"+
		"      t2 Initial Precision: '%v'\n"+
		" t1 Est Required Precision: '%v'\n"+
		"t2 Initial Value (Prec=-1): '%v'\n",
		numValueStr,
		pureNumStrStats02a.NumIntegerDigits,
		pureNumStrStats02a.NumFractionalDigits,
		pureNumStrStats02a.NumIntegerDigits+
			pureNumStrStats02a.NumFractionalDigits,
		t2.Mode().String(),
		t2.Acc().String(),
		t2.Prec(),
		estimatedRequiredPrecision,
		t2InitialVal)

	t2.SetPrec(t2.MinPrec())

	t2InitialVal = t2.Text('f', -1)

	var pureNumStrStats02b PureNumberStrComponents

	pureNumStrStats02b,
		err = breakNumStrToComponents(
		t2InitialVal)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats02b\n"+
			"Error returned by breakNumStrToComponents(numValueStr)\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	estimatedRequiredPrecision,
		err =
		calculateRequiredPrecision2(
			pureNumStrStats02a.NumIntegerDigits,
			pureNumStrStats02a.NumFractionalDigits,
			0)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"pureNumStrStats01a\n"+
			"Error returned by calculateRequiredPrecision2()\n"+
			"Error = \n%v\n\n",
			funcName,
			err.Error())

		return
	}

	fmt.Printf(
		"\n ----- After Minimum Precision -----\n"+
			"        t2 Post Minimum Mode: '%v'\n"+
			"    t2 Post Minimum Accuracy: '%v'\n"+
			"   t2 Post Minimum Precision: '%v'\n"+
			"   t2 Est Required Precision: '%v'\n"+
			"       t2 Post Minimum Value: '%v'\n"+
			"  t2 Post Min Val Int Digits: '%v'\n"+
			" t2 Post Min Val Frac Digits: '%v'\n"+
			"t2 Post Min Val Total Digits: '%v'\n"+
			"\n\n",
		t2.Mode().String(),
		t2.Acc().String(),
		t2.Prec(),
		estimatedRequiredPrecision,
		t2InitialVal,
		pureNumStrStats02b.NumIntegerDigits,
		pureNumStrStats02b.NumFractionalDigits,
		pureNumStrStats02b.NumIntegerDigits+
			pureNumStrStats02b.NumFractionalDigits)

	// Precision as used here is the number decimal places
	fmt.Printf("\n ----- After Minimum Precision2 -----\n"+
		"t2 Numeric Value Text f 14: '%v'\n"+
		"t2 Numeric Value Printf .14f: '%.14f'\n",
		&t2,
		&t2)

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}

func InitializeNegativeBigFloat() {

	funcName := "InitializeNegativeBigFloat"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("Function: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	numValueStr := "-123456.78901234567890"

	fmt.Printf("Original Numeric Value String: %v\n",
		numValueStr)

	var t big.Float
	var ok bool

	_,
		ok = t.SetString(numValueStr)

	if !ok {
		fmt.Printf("Error: t.SetString(numValueStr) Failed!\n")

		return
	}

	fmt.Printf("t Intial Mode: '%v'\n",
		t.Mode().String())

	fmt.Printf("t Intial Accuracy: '%v'\n",
		t.Acc().String())

	fmt.Printf("t Precision: '%v'\n",
		t.Prec())

	minPrecision := t.MinPrec()

	fmt.Printf("t Minimum Precision: '%v'\n",
		minPrecision)

	fmt.Printf("t Numeric Value f -1: '%v'\n",
		t.Text('f', -1))

	fmt.Printf("t Numeric Value f %v: '%v'\n",
		minPrecision,
		t.Text('f', int(minPrecision)))

	fmt.Printf("t Numeric Value f 14: '%v'\n",
		t.Text('f', 14))

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

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

func modeToString(mode big.RoundingMode) string {

	/*
		const (
			ToNearestEven RoundingMode = iota // == IEEE 754-2008 roundTiesToEven
			ToNearestAway                     // == IEEE 754-2008 roundTiesToAway
			ToZero                            // == IEEE 754-2008 roundTowardZero
			AwayFromZero                      // no IEEE 754-2008 equivalent
			ToNegativeInf                     // == IEEE 754-2008 roundTowardNegative
			ToPositiveInf                     // == IEEE 754-2008 roundTowardPositive
		)

	*/

	var modeStr string

	switch mode {

	case big.ToNearestEven:
		modeStr = "ToNearestEven"
	case big.ToNearestAway:
		modeStr = "ToNearestEven"
	case big.ToZero:
		modeStr = "ToNearestEven"
	case big.AwayFromZero:
		modeStr = "ToNearestEven"
	case big.ToNegativeInf:
		modeStr = "ToNearestEven"
	case big.ToPositiveInf:
		modeStr = "ToNearestEven"
	default:
		modeStr = "Error: mode is INVALID!\n"
	}

	return modeStr
}

func calculateRequiredPrecision2(

	integerDigits int64,
	fractionalDigits int64,
	spareBufferDigits int64) (
	uint,
	error) {

	// Precision calculated in multiples of 2
	funcName := "calculateRequiredPrecision2"

	totalDigits :=
		integerDigits +
			fractionalDigits +
			spareBufferDigits

	totalDigitsFloat := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(totalDigits)

	/*	factorEightFloat := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(8)
	*/
	factorTwoFloat := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(2)

	factorTwoUint64 := uint64(2)

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

	totalDigitsFloat.Quo(totalDigitsFloat, factorTwoFloat)

	baseDigits,
		accuracy := totalDigitsFloat.Uint64()

	if accuracy == -1 {
		baseDigits++
	}

	baseDigits = baseDigits * factorTwoUint64

	return uint(baseDigits), err
}

func calculateRequiredPrecision8(

	integerDigits int64,
	fractionalDigits int64,
	spareDigitsBuffer int64) (
	uint,
	error) {

	// Precision calculated in multiples of 8

	funcName := "calculateRequiredPrecision2"

	totalDigits :=
		integerDigits +
			fractionalDigits +
			spareDigitsBuffer

	totalDigitsFloat := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(totalDigits)

	factorEightFloat := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(8)

	factorEightUint64 := uint64(8)

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
		accuracy := totalDigitsFloat.Uint64()

	if accuracy == -1 {
		baseDigits++
	}

	baseDigits = baseDigits * factorEightUint64

	return uint(baseDigits), err
}
