package naturalLogCalcs

// *************************************
// ** xt_sqrtBigFloat_02_stress_test.go **
// ** 013Fix — additional sqrt tests   **
// *************************************

import (
	"math/big"
	"testing"
)

func Test_sqrtBigFloat_Stress_LargeInteger_01(t *testing.T) {

	ePrefix := "Test_sqrtBigFloat_Stress_LargeInteger_01"

	// A large integer (about 100 digits)
	xValStr := "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"

	expectedSqrRootStr := "35136418288201442531112223816998829391748408772394.003368165480068665646116465388194949300644461930322838575068520197376180363492490461026150613869585877503446674821327283055649366670554432971536973836028670337769014815299736954980670699179332797666669842006106642089406890613345171254295684975439711433834909426081517364187130685312361793419724046891466366618034056616205344790030023880250438487117909397632832145238569442102323153076202543130529385706339271542640467015299440287878081532413018341063346067051241422286782070747298072500592760750827061092850779809211190919933516381474457787101136988364737622928842329702550577296867327529312758542657890814874261417974017646869977027929965921865237970452441866270642043733026855259058636928953088739484085385686592428436644891459077032704319097273974174888959813896091592715063152587752593600631834901289640602887367898409383257304667370045948184397677575294458654895867058253836170806228997570007598065934714222703347355836483434828614228020283017203652"

	bFloatHlpr := new(BigFloatHelper)

	sqrtCalc := new(BigFloatMath)

	integerDigits, expectedDecimalDigits := bFloatHlpr.CountDigits(expectedSqrRootStr, '.')

	// calculationDecimalDigits := expectedDecimalDigits + 50
	calculationDecimalDigits := expectedDecimalDigits

	precBits := bFloatHlpr.ComputeBigFloatPrecisionBits(uint(calculationDecimalDigits), 1)

	xVal, ok := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(precBits).
		SetString(xValStr)

	if !ok {

		t.Fatalf("%v\nFailed SetString(%v)\n",
			ePrefix, xValStr)
	}

	sqrtVal, err := sqrtCalc.SqrtBigFloat(xVal, uint(integerDigits), precBits)

	if err != nil {
		t.Fatalf("%v\nSqrtBigFloat error: %v\n", ePrefix, err)
	}

	actualStr := sqrtVal.Text('f', expectedDecimalDigits)

	if expectedSqrRootStr != actualStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualStr = '%v'\n"+
			"  Actual actualStr = '%v'\n\n",
			ePrefix, expectedSqrRootStr, actualStr)
	}

	// Verify round-trip: (sqrtVal)^2 ≈ xVal
	ySquared := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(precBits)

	ySquared.Mul(sqrtVal, sqrtVal)

	// Compare as strings with limited decimal digits
	xStrNorm := xVal.Text('f', 0)
	ySquaredStr := ySquared.Text('f', 0)

	if xStrNorm != ySquaredStr {
		t.Errorf("%v\nRound-trip sqrt(x)^2 mismatch\nx:        %v\ny^2:      %v\n",
			ePrefix, xStrNorm, ySquaredStr)
	}
}

func Test_sqrtBigFloat_Stress_LargeInteger_02(t *testing.T) {

	ePrefix := "Test_sqrtBigFloat_Stress_LargeInteger_02"

	// Even larger integer (about 400 digits)
	xStr := "9876543210987654321098765432109876543210987654321098765432109876543210987654321098765432109876543210" +
		"12345678901234567890123456789012345678901234567890123456789012345678901234567890"

	bFloatHlpr := new(BigFloatHelper)

	intDigits, _ := bFloatHlpr.CountDigits(xStr, '.')

	decDigits := intDigits + 20

	precBits := bFloatHlpr.ComputeBigFloatPrecisionBits(uint(decDigits), 2)

	xVal, ok := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(precBits).
		SetString(xStr)

	if !ok {
		t.Fatalf("%v\nSetString FAILED for xStr: %v\n", ePrefix, xStr)
	}

	sqrtVal, err := new(BigFloatMath).SqrtBigFloat(xVal, uint(intDigits), precBits)
	if err != nil {
		t.Fatalf("%v\nSqrtBigFloat error: %v\n", ePrefix, err)
	}

	ySquared := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(precBits)

	ySquared.Mul(sqrtVal, sqrtVal)

	xStrNorm := xVal.Text('f', 0)
	ySquaredStr := ySquared.Text('f', 0)

	if xStrNorm != ySquaredStr {
		t.Errorf("%v\nRound-trip sqrt(x)^2 mismatch\nx:        %v\ny^2:      %v\n",
			ePrefix, xStrNorm, ySquaredStr)
	}
}
