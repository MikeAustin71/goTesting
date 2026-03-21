package naturalLogCalcs

import (
	"math/big"
	"testing"
)

func Test_sqrtBigFloat_Basic_01(t *testing.T) {

	ePrefix := "Test_sqrtBigFloat_Basic_01()"

	bFloatHlpr := new(BigFloatHelper)
	sqrtCalc := new(BigFloatMath)

	type testCase struct {
		xStr     string
		expected string
	}

	tests := []testCase{
		// Exact squares
		{"0", "0"},
		{"1", "1"},
		{"4", "2"},
		{"9", "3"},
		// Irrationals with known high‑precision values
		// sqrt(2) to 50 decimal digits
		{"2", "1.41421356237309504880168872420969807856967187537695"},
		// sqrt(5) to 50 decimal digits
		{"5", "2.23606797749978969640917366873127623544061835961153"},
	}

	for idx, tc := range tests {

		tc.expected = bFloatHlpr.CleanNumberString(tc.expected)

		_, decDigits := bFloatHlpr.CountDigits(tc.expected, '.')
		if decDigits == 0 {
			decDigits = 3
		} else {
			decDigits += 3
		}

		precBits := bFloatHlpr.ComputeBigFloatPrecisionBits(uint(decDigits), 1)

		xVal, ok := new(big.Float).
			SetMode(big.AwayFromZero).
			SetPrec(precBits).
			SetString(tc.xStr)

		if !ok {
			t.Fatalf("%v\nFailed SetString(%v) at index %v\n",
				ePrefix, tc.xStr, idx)
		}

		actual, err := sqrtCalc.SqrtBigFloat(xVal, precBits)

		if err != nil {
			t.Fatalf("%v\nError computing sqrt(%v): %v\n",
				ePrefix, tc.xStr, err)
		}

		// Format with the same number of decimal digits as expected
		_, expectedDecDigits := bFloatHlpr.CountDigits(tc.expected, '.')

		actualStr := actual.Text('f', expectedDecDigits)

		if actualStr != tc.expected {
			t.Errorf("%v\n"+
				"sqrt(%v) mismatch\n"+
				"Expected: %v\n"+
				"  Actual: %v\n"+
				"Index: %v\n",
				ePrefix,
				tc.xStr,
				tc.expected,
				actualStr,
				idx)
		}
	}
}

func Test_sqrtBigFloat_NegativeInput_01(t *testing.T) {

	ePrefix := "Test_sqrtBigFloat_NegativeInput_01()"

	sqrtCalc := new(BigFloatMath)

	xVal := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(128).
		SetFloat64(-1.0)

	_, err := sqrtCalc.SqrtBigFloat(xVal, 128)

	if err == nil {
		t.Errorf("%v\nExpected error for negative input, but got nil\n", ePrefix)
	}
}
func Test_sqrtBigFloat_1000Digits_01(t *testing.T) {

	ePrefix := "Test_sqrtBigFloat_1000Digits_01()"

	bFloatHlpr := new(BigFloatHelper)

	sqrtCalc := new(BigFloatMath)

	// This expected result was calculated with the Python math library
	// 'mpmath'. The expected result is accurate to 1,001 digits.

	expectedResultStr := "1.4142135623730950488016887242096980785696718753769480731766797379907324784621070388503875343276415727350138462309122970249248360558507372126441214970999358314132226659275055927557999505011527820605714701095599716059702745345968620147285174186408891986095523292304843087143214508397626036279952514079896872533965463318088296406206152583523950547457502877599617298355752203375318570113543746034084988471603868999706990048150305440277903164542478230684929369186215805784631115966687130130156185689872372352885092648612494977154218334204285686060146824720771435854874155657069677653720226485447015858801620758474922657226002085584466521458398893944370926591800311388246468157082630100594858704003186480342194897278290641045072636881313739855256117322040245091227700226941127573627280495738108967504018369868368450725799364729060762996941380475654823728997180326802474420629269124859052181004459842150591120249441341728531478105803603371077309182869314710171111683916581726889419758716582152128229518488472"

	xValStr := "2"

	expectedResultStr = bFloatHlpr.CleanNumberString(expectedResultStr)

	_, expectedDecimalDigits := bFloatHlpr.CountDigits(expectedResultStr, '.')

	calculationDecimalDigits := expectedDecimalDigits + 3

	precBits := bFloatHlpr.ComputeBigFloatPrecisionBits(uint(calculationDecimalDigits), 1)

	xVal, ok := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(precBits).
		SetString(xValStr)

	if !ok {

		t.Fatalf("%v\nFailed SetString(%v)\n",
			ePrefix, xValStr)
	}

	actual, err := sqrtCalc.SqrtBigFloat(xVal, precBits)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actual, err := sqrtCalc.SqrtBigFloat(xVal, precBits)\n"+
			"xVal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, xVal.Text('f', expectedDecimalDigits), err.Error())
		return
	}

	actualStr := actual.Text('f', expectedDecimalDigits)

	if expectedResultStr != actualStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualStr = '%v'\n"+
			"  Actual actualStr = '%v'\n\n",
			ePrefix, expectedResultStr, actualStr)
	}

	return
}

func Test_sqrtBigFloat_1000Digits_02(t *testing.T) {

	ePrefix := "Test_sqrtBigFloat_1000Digits_02()"

	bFloatHlpr := new(BigFloatHelper)

	sqrtCalc := new(BigFloatMath)

	// This expected result was calculated with the Python math library
	// 'mpmath'.

	// Expected Result is accurate to 1,001 digits
	expectedResultStr := "7.57079916521366983223589085940257446093186571087576262685458548752028368016292044019371264038358920159219431761161343478341722312033808711481028398850277324769575361147269676708111552122984173571111308561260276604176186688312421272624906985236007856343457027209594345853191903088613329191598829786009384812594492163086529847163238059259557813382046954511807942863853618661059947354962259472589318452892101429949845505718000085584629676441594678307027175960644997958868103853369904702926209798520646627225516720310891369277123731009261291699411300769428682112936645769406537648989991531536118282970906451299539350711154626401948385741987154491365143603000111970573140921812178193505059779374948372157801988393863726238171055428165453865542459428672357466719509632783571223291092825357687050664310570894884387773882526623858389275004702862537818107093816155802801440007778257703236497293896895992296702495445078264904311305347550183142486829913501983294261435370110365552309566556730026915497566854977636"

	// Square Root of '57.317'
	xValStr := "57.317"

	expectedResultStr = bFloatHlpr.CleanNumberString(expectedResultStr)

	_, expectedDecimalDigits := bFloatHlpr.CountDigits(expectedResultStr, '.')

	calculationDecimalDigits := expectedDecimalDigits + 3

	precBits := bFloatHlpr.ComputeBigFloatPrecisionBits(uint(calculationDecimalDigits), 1)

	xVal, ok := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(precBits).
		SetString(xValStr)

	if !ok {

		t.Fatalf("%v\nFailed SetString(%v)\n",
			ePrefix, xValStr)
	}

	actual, err := sqrtCalc.SqrtBigFloat(xVal, precBits)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actual, err := sqrtCalc.SqrtBigFloat(xVal, precBits)\n"+
			"xVal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, xVal.Text('f', expectedDecimalDigits), err.Error())
		return
	}

	actualStr := actual.Text('f', expectedDecimalDigits)

	if expectedResultStr != actualStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualStr = '%v'\n"+
			"  Actual actualStr = '%v'\n\n",
			ePrefix, expectedResultStr, actualStr)
	}

	return
}
