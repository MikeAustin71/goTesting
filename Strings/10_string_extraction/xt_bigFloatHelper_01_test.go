package Strings_string_extraction10

import "testing"

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
