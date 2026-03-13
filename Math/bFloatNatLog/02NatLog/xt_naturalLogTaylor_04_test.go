package naturalLogCalcs

import (
	"math/big"
	"testing"
)

// 008Fix01
func Test_Natural_Logarithm_Taylor_NearMinusOne_01(t *testing.T) {

	ePrefix := "Test_Natural_Logarithm_Taylor_NearMinusOne_01()"

	tests := []struct {
		xStr     string
		expected string
	}{
		//{"0.9999", "-0.000100005000333416708416708333208"}, Copilot
		{"0.9999", "-0.000100005000333358335333500014286964396835397734571075514089865762716344756611402617022495986"},
		//                           12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
		//                           ---------1---------2---------3---------4---------5---------6---------7---------8---------9---------0---------0
		//                                                                                                                              0         1
		//                                                                                                                              1         1
		//

		// {"0.99999", "-0.000010000050000333341666708333208"}, Copilot
		{"0.99999", "-0.00001000005000033333583335333350000142858392868254068254877353210755134469426136148636674376935861921"},
		//                            12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
		//                            ---------1---------2---------3---------4---------5---------6---------7---------8---------9---------0---------0
		//                                                                                                                               0         1
		//                                                                                                                               1         1
		//

		//{"0.999999", "-0.000001000000500000333333416666708"}, Copilot
		{"0.999999", "-0.00000100000050000033333358333353333350000014285726785725396835396834487742821075513382656239322905573"},
		//                             12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
		//                             ---------1---------2---------3---------4---------5---------6---------7---------8---------9---------0---------0
		//                                                                                                                                0         1
		//                                                                                                                                1         1
		//

	}

	var workingPrecision uint

	var sourceExpectedDecimalDigits, sourceInputDecimalDigits, sourceCalculatedDecimalDigits int

	taylor := new(naturalLogTaylor)

	bFloatHlpr := new(BigFloatHelper)

	for _, tc := range tests {

		_, sourceInputDecimalDigits = bFloatHlpr.CountDigits(tc.xStr, '.')

		lenExpectedStr := len(tc.expected)

		tc.expected = bFloatHlpr.CleanNumberString(tc.expected)

		lenExpectedStr = len(tc.expected)

		_, sourceExpectedDecimalDigits = bFloatHlpr.CountDigits(tc.expected, '.')

		if sourceInputDecimalDigits > sourceExpectedDecimalDigits {
			sourceCalculatedDecimalDigits = sourceInputDecimalDigits
		} else {
			sourceCalculatedDecimalDigits = sourceExpectedDecimalDigits
		}

		sourceCalculatedDecimalDigits += 3

		workingPrecision = bFloatHlpr.ComputeBigFloatPrecisionBits(uint(sourceCalculatedDecimalDigits), 1)

		bFloatXValue, ok := new(big.Float).
			SetMode(big.AwayFromZero).
			SetPrec(workingPrecision).
			SetString(tc.xStr)

		if !ok {
			t.Errorf("FAILED: SetString(%v)", tc.xStr)
			continue
		}

		actual, err := taylor.lnTaylorDirect(bFloatXValue, workingPrecision)

		if err != nil {
			t.Errorf("Error computing ln(%v): %v", tc.xStr, err)
			continue
		}

		_, decDigits := bFloatHlpr.CountDigits(tc.expected, '.')

		actualStr := actual.Text('f', decDigits)

		lenActualStr := len(actualStr)

		if actualStr != tc.expected {

			t.Errorf("\n%v\n"+
				"ln(%v) mismatch\n"+
				"Expected: %v\n"+
				"  Actual: %v\n"+
				"Length Expected: %v\n"+
				"  Length Actual: %v\n",
				ePrefix,
				tc.xStr,
				tc.expected,
				actualStr,
				lenExpectedStr,
				lenActualStr)

			return
		}
	}
}
