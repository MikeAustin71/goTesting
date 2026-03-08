package naturalLogCalcs

import (
	"math/big"
	"testing"
)

func Test_Natural_Logarithm_Taylor_01(t *testing.T) {

	ePrefix := "Test_Natural_Logarithm_Taylor_01()"

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

	actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, prec)

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, prec)\n"+
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

func Test_Natural_Logarithm_Taylor_02(t *testing.T) {

	ePrefix := "Test_Natural_Logarithm_Taylor_02()"

	xValueStr := "245"

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
	expectedResult := "5." +
		"5012582105447269848114648201125470987997708134322400988314281913493387132117963002054508752689226647" +
		"8608969143341784472317150974468024988593153280779150727936782857370573281241801155421927627388228010" +
		"260330798331306088066067384276824977246856395981"

	prec := uint(7000)

	actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, prec)

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResult, err := new(naturalLogTaylor).lnTaylorDirect(xValue, prec)\n"+
			"xValue= '%v'\n"+
			"prec= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			xValue.Text('f', 30),
			prec,
			err.Error())

		return
	}

	bFloatHlpr := new(BigFloatHelper)

	_, decDigits := bFloatHlpr.CountDigits(expectedResult, '.')

	actualResultStr := actualResult.Text('f', decDigits)

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
