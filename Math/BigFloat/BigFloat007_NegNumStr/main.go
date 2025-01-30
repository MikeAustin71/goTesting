package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {

	BestInitializeBigFloat()
}

func BestInitializeBigFloat() {

	funcName := "InitializeBigFloat"

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

	t2.SetPrec(t2.MinPrec())
	t2.SetMode(big.ToNearestAway)

	fmt.Printf("t2 Intial Mode: '%v'\n",
		t2.Mode().String())

	fmt.Printf("t2 Intial Accuracy: '%v'\n",
		t2.Acc().String())

	fmt.Printf("t2 Precision: '%v'\n",
		t2.Prec())

	fmt.Printf("t2 Numeric Value f -1: '%v'\n",
		t2.Text('f', -1))

	// Precision as used here is the number decimal places
	fmt.Printf("t2 Numeric Value f 14: '%v'\n",
		t2.Text('f', 14))

	fmt.Printf("t2 Numeric Value f .14: '%.14f'",
		&t2)

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}

func InitializeNegativeBigFloat() {

	funcName := "InitializeBigFloat"

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

func InitializeBigFloat() {

	funcName := "InitializeBigFloat"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("Function: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	numValueStr := "123456.78901234567890"

	fmt.Printf("Original Numeric Value String: %v\n",
		numValueStr)

	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.
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

	fmt.Printf("t Numeric Value f -1: '%v'\n",
		t.Text('f', -1))

	fmt.Printf("\n\n" + breakStr + "\n")
	fmt.Printf("t2 Example With Set To Minimum Precision.\n\n")
	var t2 big.Float

	_,
		ok = t2.SetString(numValueStr)

	if !ok {
		fmt.Printf("Error: t.SetString(numValueStr) Failed!\n")

		return
	}

	t2.SetPrec(t2.MinPrec())
	t2.SetMode(big.ToNearestAway)

	fmt.Printf("t2 Intial Mode: '%v'\n",
		t2.Mode().String())

	fmt.Printf("t2 Intial Accuracy: '%v'\n",
		t2.Acc().String())

	fmt.Printf("t2 Precision: '%v'\n",
		t2.Prec())

	fmt.Printf("t2 Numeric Value f -1: '%v'\n",
		t2.Text('f', -1))

	// Precision as used here is the number decimal places
	fmt.Printf("t2 Numeric Value f 14: '%v'\n",
		t2.Text('f', 14))

	fmt.Printf("t2 Numeric Value f .14: '%.14f'",
		&t2)

	fmt.Printf("\n\n" + breakStr + "\n")
	fmt.Printf("t3 Example With Set To 4096 Precision.\n\n")
	var t3 big.Float

	_,
		ok = t3.SetString(numValueStr)

	if !ok {
		fmt.Printf("Error: t.SetString(numValueStr) Failed!\n")

		return
	}

	minimumPrecision := t3.MinPrec()

	fmt.Printf("Original Minimum Precision: %v\n",
		minimumPrecision)

	t3.SetPrec(4096)
	t3.SetMode(big.ToNearestAway)

	fmt.Printf("t3 Intial Mode: '%v'\n",
		t3.Mode().String())

	fmt.Printf("t3 Intial Accuracy: '%v'\n",
		t3.Acc().String())

	fmt.Printf("t3 Precision: '%v'\n",
		t3.Prec())

	fmt.Printf("t3 Numeric Value f -1: '%v'\n",
		t3.Text('f', -1))

	fmt.Printf("t3 Numeric Value f minimumPrecision: '%v'\n",
		t3.Text('f', int(minimumPrecision)))

	fmt.Printf("t3 Numeric Value f .14: '%.14f'",
		&t3)

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}
