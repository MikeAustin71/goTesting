package main

import (
	"fmt"
	"math/big"
	"strings"

	ePref "github.com/MikeAustin71/errpref"
	"github.com/mikeaustin71/Math/NthRoot/NewtonExamples/NewtonBuild03/bigFloatPower"
)

func main() {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"Math/NthRoot/NewtonExamples/NewtonBuild03/main.go",
		"")

	if err != nil {
		return
	}

	// 0.37885748409213036900649240796424989244463341114874
	expectedNumStr := "0.3788574840921303690064"

	base := big.NewFloat(1.382)

	powerInt64 := int64(-3)

	maxInternalPrecisionUint := uint(4096)

	roundingMode := big.AwayFromZero

	breakStr := strings.Repeat("=", 65)

	subBreakStr := strings.Repeat("-", 65)

	setupStr := "Preparing to call: raiseToPowerBySquaring()"

	fmt.Printf("\n\n %v\n"+
		" Function: %v\n"+
		" %v\n"+
		"               base value = %v\n"+
		"            base Accuracy = %v\n"+
		"           base Precision = %v\n"+
		"               powerInt64 = %v\n"+
		" maxInternalPrecisionUint = %v\n"+
		"             roundingMode = %v\n"+
		" %v\n"+
		" %v\n"+
		" %v \n\n",
		breakStr,
		ePrefix,
		subBreakStr,
		base.Text('f', -1),
		base.Acc(),
		base.Prec(),
		powerInt64,
		maxInternalPrecisionUint,
		roundingMode.String(),
		subBreakStr,
		setupStr,
		breakStr)

	result, err := new(bigFloatPower.BigFloatPower).RaiseToPower(
		base, powerInt64, maxInternalPrecisionUint, roundingMode)

	if err != nil {
		fmt.Printf(" %v\n"+
			" Error returned by:\n"+
			" result, err := new(bigFloatRaiseToPower).raiseToPowerBySquaring(n"+
			"   base, powerInt64, maxInternalPrecisionUint, roundingMode)\n"+
			" base= '%v'\n"+
			" powerInt64= '%v'\n"+
			" maxInternalPrecisionUint= '%v'\n"+
			" roundingMode= '%v'\n"+
			" Error='%v'\n\n",
			ePrefix,
			base.Text('f', -1),
			powerInt64,
			maxInternalPrecisionUint,
			roundingMode.String(),
			err.Error())
		return

	}

	setupStr = "Successful Return from: raiseToPowerBySquaring()"

	resultRoundingMode := result.Mode()

	resultAccuracy := result.Acc()

	resultPrecisionUint := result.Prec()

	fmt.Printf(" %v\n"+
		" Function: %v\n"+
		" %v\n"+
		" Results: raiseToPowerBySquaring()\n"+
		"      Expected Result: %v\n"+
		"        Actual Result: %v\n"+
		" Actual Rounding Mode: %v\n"+
		"      Actual Accuracy: %v\n"+
		"     Actual Precision: %v\n"+
		"           base value: %v\n"+
		"        base Accuracy: %v\n"+
		"       base Precision: %v\n"+
		" %v\n"+
		" %v\n"+
		" %v\n\n",
		breakStr,
		ePrefix.String(),
		subBreakStr,
		expectedNumStr,
		result.Text('f', -1),
		resultRoundingMode.String(),
		resultAccuracy.String(),
		resultPrecisionUint,
		base.Text('f', -1),
		base.Acc(),
		base.Prec(),
		subBreakStr,
		setupStr,
		breakStr)

	return
}
