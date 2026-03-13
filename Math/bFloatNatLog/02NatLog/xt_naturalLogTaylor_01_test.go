package naturalLogCalcs

import (
	"math/big"
	"testing"
)

func Test_Natural_Logarithm_Taylor_01(t *testing.T) {

	ePrefix := "Test_Natural_Logarithm_Taylor_01()"

	xValueStr := "3237"

	// Fractional Digits ---------1---------2---------3
	//         Accuracy: 123456789012345678901234567890
	expectedResult := "8.082402253926244350924204901779"

	bFloatHlpr := new(BigFloatHelper)

	var workingPrecision uint

	var sourceExpectedDecimalDigits, sourceInputDecimalDigits, sourceCalculatedDecimalDigits int

	_, sourceInputDecimalDigits = bFloatHlpr.CountDigits(xValueStr, '.')

	_, sourceExpectedDecimalDigits = bFloatHlpr.CountDigits(expectedResult, '.')

	if sourceInputDecimalDigits > sourceExpectedDecimalDigits {
		sourceCalculatedDecimalDigits = sourceInputDecimalDigits
	} else {
		sourceCalculatedDecimalDigits = sourceExpectedDecimalDigits
	}

	sourceCalculatedDecimalDigits += 3

	workingPrecision = bFloatHlpr.ComputeBigFloatPrecisionBits(uint(sourceCalculatedDecimalDigits), 1)

	xValue, isOk := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(workingPrecision).
		SetString(xValueStr)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"xValue, isOk := new(big.Float)\n"+
			".SetMode(big.AwayFromZero).SetPrec(workingPrecision).SetString(xValueStr)\n"+
			"xValueStr= '%v'\n"+
			"Working Precision Bits = '%v'\n"+
			"Error= 'SetString() FAILED!'\n\n", ePrefix, workingPrecision, xValueStr)
		return
	}

	actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, workingPrecision)

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, workingPrecision)\n"+
			"xValue= '%v'\n"+
			"workingPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			xValue.Text('f', sourceInputDecimalDigits),
			workingPrecision,
			err.Error())

		return
	}

	actualResultStr := actualResult.Text('f', sourceExpectedDecimalDigits)

	if expectedResult != actualResultStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedResult != actualResultStr\n"+
			"Expected actualResultStr = '%v'\n"+
			"  Actual actualResultStr = '%v'\n\n",
			ePrefix, expectedResult, actualResultStr)
	}

	return
}

func Test_Natural_Logarithm_Taylor_02(t *testing.T) {

	ePrefix := "Test_Natural_Logarithm_Taylor_02()"

	xValueStr := "245"

	expectedResult := "5." +
		"5012582105447269848114648201125470987997708134322400988314281913493387132117963002054508752689226647" +
		"8608969143341784472317150974468024988593153280779150727936782857370573281241801155421927627388228010" +
		"260330798331306088066067384276824977246856395981"

	bFloatHlpr := new(BigFloatHelper)

	var workingPrecision uint

	var sourceExpectedDecimalDigits, sourceInputDecimalDigits, sourceCalculatedDecimalDigits int

	_, sourceInputDecimalDigits = bFloatHlpr.CountDigits(xValueStr, '.')

	_, sourceExpectedDecimalDigits = bFloatHlpr.CountDigits(expectedResult, '.')

	if sourceInputDecimalDigits > sourceExpectedDecimalDigits {
		sourceCalculatedDecimalDigits = sourceInputDecimalDigits
	} else {
		sourceCalculatedDecimalDigits = sourceExpectedDecimalDigits
	}

	sourceCalculatedDecimalDigits += 3

	workingPrecision = bFloatHlpr.ComputeBigFloatPrecisionBits(uint(sourceCalculatedDecimalDigits), 1)

	xValue, isOk := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(workingPrecision).
		SetString(xValueStr)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"xValue, isOk := new(big.Float)\n"+
			".SetMode(big.AwayFromZero).SetPrec(workingPrecision).SetString(xValueStr)\n"+
			"xValueStr= '%v'\n"+
			"Working Precision Bits = '%v'\n"+
			"Error= 'SetString() FAILED!'\n\n", ePrefix, workingPrecision, xValueStr)
		return
	}

	actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, workingPrecision)

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, workingPrecision)\n"+
			"xValue= '%v'\n"+
			"workingPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			xValue.Text('f', sourceInputDecimalDigits),
			workingPrecision,
			err.Error())

		return
	}

	actualResultStr := actualResult.Text('f', sourceExpectedDecimalDigits)

	if expectedResult != actualResultStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedResult != actualResultStr\n"+
			"Expected actualResultStr = '%v'\n"+
			"  Actual actualResultStr = '%v'\n\n",
			ePrefix, expectedResult, actualResultStr)
	}

	return
}

func Test_Natural_Logarithm_Taylor_03(t *testing.T) {

	ePrefix := "Test_Natural_Logarithm_Taylor_03()"

	xValueStr := "1.0001"

	expectedResult :=
		"0.00009999500033330833533316668095113106348206440107107551266129432164491607407171907733994721288860975"
	//   12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
	//   ---------1---------2---------3---------4---------5---------6---------7---------8---------9---------0---------0
	//                                                                                                      0         1
	//                                                                                                      1         1
	//
	//     101- Decimal Digits of Accuracy - This is Calculator Result

	bFloatHlpr := new(BigFloatHelper)

	var workingPrecision uint

	var sourceExpectedDecimalDigits, sourceInputDecimalDigits, sourceCalculatedDecimalDigits int

	_, sourceInputDecimalDigits = bFloatHlpr.CountDigits(xValueStr, '.')

	_, sourceExpectedDecimalDigits = bFloatHlpr.CountDigits(expectedResult, '.')

	if sourceInputDecimalDigits > sourceExpectedDecimalDigits {
		sourceCalculatedDecimalDigits = sourceInputDecimalDigits
	} else {
		sourceCalculatedDecimalDigits = sourceExpectedDecimalDigits
	}

	sourceCalculatedDecimalDigits += 3

	workingPrecision = bFloatHlpr.ComputeBigFloatPrecisionBits(uint(sourceCalculatedDecimalDigits), 1)

	xValue, isOk := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(workingPrecision).
		SetString(xValueStr)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"xValue, isOk := new(big.Float)\n"+
			".SetMode(big.AwayFromZero).SetPrec(workingPrecision).SetString(xValueStr)\n"+
			"xValueStr= '%v'\n"+
			"Working Precision Bits = '%v'\n"+
			"Error= 'SetString() FAILED!'\n\n", ePrefix, workingPrecision, xValueStr)
		return
	}

	actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, workingPrecision)

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, workingPrecision)\n"+
			"xValue= '%v'\n"+
			"workingPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			xValue.Text('f', sourceInputDecimalDigits),
			workingPrecision,
			err.Error())

		return
	}

	actualResultStr := actualResult.Text('f', sourceExpectedDecimalDigits)

	if expectedResult != actualResultStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedResult != actualResultStr\n"+
			"Expected actualResultStr = '%v'\n"+
			"  Actual actualResultStr = '%v'\n\n",
			ePrefix, expectedResult, actualResultStr)
	}

	return
}
