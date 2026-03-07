package naturalLogCalcs

import (
	"math/big"
	"testing"
)

func Test_Natural_Logarithm_AGM_Direct_01(t *testing.T) {

	ePrefix := "Test_Natural_Logarithm_Direct_AGM_01()"

	xValueStr := "3237"

	xValue, isOk := new(big.Float).
		SetMode(big.AwayFromZero).
		SetString(xValueStr)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"xValue, isOk := new(big.Float)\n"+
			".SetMode(big.AwayFromZero).SetString(xValueStr)\n"+
			"xValueStr= '%v'\n"+
			"Error= 'SetString() FAILED!'\n\n", ePrefix, xValueStr)
		return
	}

	// Fractional Digits ---------1---------2---------3
	//         Accuracy: 123456789012345678901234567890
	expectedResult := "8.082402253926244350924204901779"

	prec := uint(7000)

	actualResult, err := new(naturalLogAGM).lnAGMDirect(xValue, prec)

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResult, err := new(naturalLogAGM).lnAGMDirect(xValue, prec)\n"+
			"xValue= '%v'\n"+
			"prec= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			xValue.Text('f', 30),
			prec,
			err.Error())

		return
	}

	actualResultStr := actualResult.Text('f', 30)

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

func Test_Natural_Logarithm_AGM_Mantissa_01(t *testing.T) {

	ePrefix := "Test_Natural_Logarithm_AGM_Mantissa_01()"

	xValueStr := "3237"

	xValue, isOk := new(big.Float).
		SetMode(big.AwayFromZero).
		SetString(xValueStr)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"xValue, isOk := new(big.Float)\n"+
			".SetMode(big.AwayFromZero).SetString(xValueStr)\n"+
			"xValueStr= '%v'\n"+
			"Error= 'SetString() FAILED!'\n\n", ePrefix, xValueStr)
		return
	}

	// Fractional Digits ---------1---------2---------3
	//         Accuracy: 123456789012345678901234567890
	expectedResult := "8.082402253926244350924204901779"

	prec := uint(7000)

	actualResult, err := new(naturalLogAGM).lnAGMMantissa(xValue, prec)

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResult, err := new(naturalLogAGM).lnAGMMantissa(xValue, prec)\n"+
			"xValue= '%v'\n"+
			"prec= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			xValue.Text('f', 30),
			prec,
			err.Error())

		return
	}

	actualResultStr := actualResult.Text('f', 30)

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
