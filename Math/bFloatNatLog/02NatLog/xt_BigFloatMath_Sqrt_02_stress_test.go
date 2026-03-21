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

	// A large integer (about 200 digits)
	xStr := "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"

	bFloatHlpr := new(BigFloatHelper)

	// We want enough digits to represent x and its sqrt comfortably
	intDigits, _ := bFloatHlpr.CountDigits(xStr, '.')
	decDigits := intDigits + 10

	precBits := bFloatHlpr.ComputeBigFloatPrecisionBits(uint(decDigits), 2)

	xVal, ok := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(precBits).
		SetString(xStr)

	if !ok {
		t.Fatalf("%v\nSetString FAILED for xStr: %v\n", ePrefix, xStr)
	}

	sqrtVal, err := new(BigFloatMath).SqrtBigFloat(xVal, precBits)
	if err != nil {
		t.Fatalf("%v\nSqrtBigFloat error: %v\n", ePrefix, err)
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
	decDigits := intDigits + 10

	precBits := bFloatHlpr.ComputeBigFloatPrecisionBits(uint(decDigits), 2)

	xVal, ok := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(precBits).
		SetString(xStr)

	if !ok {
		t.Fatalf("%v\nSetString FAILED for xStr: %v\n", ePrefix, xStr)
	}

	sqrtVal, err := new(BigFloatMath).SqrtBigFloat(xVal, precBits)
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
