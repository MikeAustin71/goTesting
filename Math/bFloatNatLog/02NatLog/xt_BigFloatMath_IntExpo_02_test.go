package naturalLogCalcs

// *************************************
// ** xt_BigFloatMath_Expo_02_test.go **
//      013Fix — IntExpoBigFloat tests
// *************************************
// These tests are designed to evaluate
// exponent calculation expected values
// to 1,000+ decimal places.

import (
	"math/big"
	"testing"
)

func Test_BigFloatMath_IntExpo_ManyDigits_01(t *testing.T) {

	ePrefix := "Test_ExpoBigFloat_1000_01"

	// baseStr to 57 decimal digits of accuracy
	baseStr := "3.716586431412345678901234567896665321489219683361042189632"

	exponentValue := int64(8)

	// Expected Result to 456 decimal digits of accuracy
	expectedResultStr := "36404.401562225975591741663072973189677987488996396044626112641492651053286615144796058628951198923359976623755686783423428496969116290496797290496395603454115934731971114958321403448741582304231510133842124829429416753443259489442443197222136207898977414246517466800582599164548256625911734795116932334044587952739317391815896874646908544842081850235941724431811478937457601741575847137432859848470036907221552484215297644767344006306721500401036591964855730176"

	bFloatHlpr := new(BigFloatHelper)

	var workingPrecisionBits uint

	var sourceCalculatedDecimalDigits, sourceBaseDecimalDigits, sourceExpectedResultDecimalDigits int

	_, sourceBaseDecimalDigits = bFloatHlpr.CountDigits(baseStr, '.')

	_, sourceExpectedResultDecimalDigits = bFloatHlpr.CountDigits(expectedResultStr, '.')

	sourceCalculatedDecimalDigits = sourceExpectedResultDecimalDigits + 3

	workingPrecisionBits = bFloatHlpr.ComputeBigFloatPrecisionBits(uint(sourceCalculatedDecimalDigits), 1)

	baseBigFloat, isOk := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(workingPrecisionBits).
		SetString(baseStr)

	if !isOk {
		t.Errorf("%v\n"+
			"Failed to set base string as big.Float!\n"+
			"baseStr= '%v'\n\n",
			ePrefix, baseStr)
		return
	}

	actualResult, err := new(BigFloatMath).IntExpoBigFloat(baseBigFloat, exponentValue, workingPrecisionBits)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" actualResult, err := new(BigFloatMath).\n"+
			"IntExpoBigFloat(baseBigFloat, exponentValue, workingPrecisionBits)\n"+
			"multiplierStr= '%v'\n"+
			"baseBigFloat= '%v'\n"+
			"workingPrecisionBits= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseBigFloat.Text('f', sourceBaseDecimalDigits),
			exponentValue,
			workingPrecisionBits,
			err.Error())

		return
	}

	actualResultStr := actualResult.Text('f', sourceExpectedResultDecimalDigits)

	if expectedResultStr != actualResultStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedResult != actualResultStr\n"+
			"Expected actualResultStr = '%v'\n"+
			"  Actual actualResultStr = '%v'\n\n",
			ePrefix, expectedResultStr, actualResultStr)
	}

	return
}
