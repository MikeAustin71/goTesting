package naturalLogCalcs

import (
	"testing"
)

func Test_CountDigits_01(t *testing.T) {

	ePrefix := "Test_CountDigits_01"

	tests := []struct {
		numString         string
		numStringName     string
		decimalSeparator  rune
		expectedIntDigits int
		expectedDecDigits int
	}{
		{eulersConstNumStr, "eulersConstNumStr", '.', 1, 49999},
		{pi20kDigitStr, "pi20kDigitStr", '.', 1, 19999},
		{natLog2Str20kDigits, "natLog2Str20kDigits", '.', 1, 20001},
		{"123.45", "Test String #1", '.', 3, 2},
		{"123", "Test String #2", '.', 3, 0},
		{".45", "Test String #3", '.', 0, 2},
		{"0.1234567890", "Test String #3", '.', 1, 10},
	}

	for _, tt := range tests {

		gotIntDigits, gotDecDigits := new(BigFloatHelper).CountDigits(tt.numString, tt.decimalSeparator)

		if gotIntDigits != tt.expectedIntDigits {
			t.Errorf("%v\n"+
				"Error: Integer Digits Read Out is INCORRECT!\n"+
				"Test String Name: '%v'\n"+
				"Expected gotIntDigits = '%v'\n"+
				"  Actual gotIntDigits = '%v'\n\n",
				ePrefix, tt.numStringName, tt.expectedIntDigits, gotIntDigits)
		}

		if gotDecDigits != tt.expectedDecDigits {
			t.Errorf("%v\n"+
				"Error: Decimal Digits Read Out is INCORRECT!\n"+
				"Test String Name: '%v'\n"+
				"Expected fixedDecNumStr2 = '%v'\n"+
				"  Actual fixedDecNumStr2 = '%v'\n\n",
				ePrefix, tt.numStringName, tt.expectedDecDigits, gotDecDigits)
		}

	}

	return
}

func Test_CleanNumberString_01(t *testing.T) {

	ePrefix := "Test_CleanNumberString_01"

	tests := []struct {
		DirtyString string
		CleanString string
	}{
		{"123.456", "123.456"},
		{"123 456", "123456"},
		{"123,456.1233", "123456.1233"},
		{"-123456.1233", "-123456.1233"},
		{"+123456.1233", "+123456.1233"},
		{"$123456.1233", "123456.1233"},
		{"123456789", "123456789"},
		{"123456789.1234568", "123456789.1234568"},
		{"XXX123456789.1234568XXX", "123456789.1234568"},
		{"@@@123456789.1234568@@@", "123456789.1234568"},
		{"♣123456789.1234568♣", "123456789.1234568"},
	}

	var testResultStr string

	bFloatHlpr := new(BigFloatHelper)

	for idx, tCase := range tests {

		testResultStr = bFloatHlpr.CleanNumberString(tCase.DirtyString)

		if testResultStr != tCase.CleanString {
			t.Errorf("%v\n"+
				"Error: CleanString is INCORRECT!\n"+
				"Test Case Index: %v\n"+
				"Expected CleanString = '%v'\n"+
				"  Actual CleanString = '%v'\n"+
				"Test Index = %v\n",
				ePrefix, idx, tCase.CleanString, testResultStr, idx)
		}

	}
}

func Test_ComputeBigFloatPrecisionBits_01(t *testing.T) {

	uintNumOfDecimalDigits := uint(100)

	expectedPrecisionResult := uint(397)

	bFloatHlpr := new(BigFloatHelper)

	actualPrecisionResult := bFloatHlpr.ComputeBigFloatPrecisionBits(uintNumOfDecimalDigits, 0)

	if actualPrecisionResult != expectedPrecisionResult {
		t.Errorf("Error: Precision Bits is INCORRECT!\n"+
			"Expected Precision Bits = '%v'\n"+
			"  Actual Precision Bits = '%v'\n",
			expectedPrecisionResult, actualPrecisionResult)
	}

	return
}

func Test_ComputeBigFloatDecimalDigits_01(t *testing.T) {

	inputPrecisionResult := uint(397)

	expectedNumOfDecimalDigits := uint(120)

	bFloatHlpr := new(BigFloatHelper)

	actualNumOfDecimalDigits := bFloatHlpr.ComputeBigFloatDecimalDigits(inputPrecisionResult, 0)

	if actualNumOfDecimalDigits != expectedNumOfDecimalDigits {
		t.Errorf("Error: Decimal Digits is INCORRECT!\n"+
			"Expected Decimal Digits = '%v'\n"+
			"  Actual Decimal Digits = '%v'\n",
			expectedNumOfDecimalDigits, actualNumOfDecimalDigits)
	}

	return

}
