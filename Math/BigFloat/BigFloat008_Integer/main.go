package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {

	bigFloatInteger002()
}

func bigFloatInteger() {

	funcName := "bigFloatInteger"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	numValueStr := "123456"
	fmt.Printf("t2 Example With Set To Minimum Precision.\n\n")

	fmt.Printf("numValueStr numeric characters= %v\n",
		numValueStr)

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

	fmt.Printf("Original t2 Intial Mode: '%v'\n",
		t2.Mode().String())

	fmt.Printf("Original t2 Intial Accuracy: '%v'\n",
		t2.Acc().String())

	fmt.Printf("Original t2 Precision: '%v'\n",
		t2.Prec())

	fmt.Printf("t2 Numeric Value f -1: '%v'\n",
		t2.Text('f', -1))

	customPrec := t2.MinPrec()
	//t2.SetPrec(t2.MinPrec())
	//t2.SetPrec(customPrec)

	fmt.Printf("Seting t2 to Precision = %v\n",
		customPrec)

	t2.SetMode(big.AwayFromZero)

	fmt.Printf("t2 Intial Mode: '%v'\n",
		t2.Mode().String())

	fmt.Printf("t2 Intial Accuracy: '%v'\n",
		t2.Acc().String())

	fmt.Printf("t2 Precision: '%v'\n",
		t2.Prec())

	fmt.Printf("t2 Numeric Value f -1: '%v'\n",
		t2.Text('f', -1))

	// Precision as used here is the number decimal places
	fmt.Printf("t2 Numeric Value Text f 2: '%v'\n",
		t2.Text('f', 2))

	fmt.Printf("t2 Numeric Value Printf .2f: '%.2f'",
		&t2)

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}

func bigFloatInteger002() {

	funcName := "bigFloatInteger002"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	floatStrVal := "12345678901"

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

	fmt.Printf("Big Float Val f Zero (0) Prec:  '%v'\n",
		bigFloatVal.Text('f', 0))

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

	fmt.Printf("Big Float Val f Zero (0) Prec:  '%v'\n",
		bigFloatVal.Text('f', 0))

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

	fmt.Printf("Big Float Val f Zero (0) Prec:  '%v'\n",
		bigFloatVal.Text('f', 0))

	fmt.Printf("\n\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}
