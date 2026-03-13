package naturalLogCalcs

import (
	"math/big"
	"testing"
)

func Test_Natural_Logarithm_Dispatcher_Range(t *testing.T) {

	ePrefix := "Test_Natural_Logarithm_Dispatcher_Range()"

	tests := []struct {
		xStr     string
		expected string
	}{
		// Near 1
		// {"1.0001", "0.000099995000333308335416654166250"},
		{"1.0001", "0.00009999500033330833533316668095113106348206440107107551266129432164491607407171907733994721288860975"},
		//                          12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901
		//                                   1         2         3         4         5         6         7         8         9         0
		//                                                                                                                             0
		//                                                                                                                             1
		//               Calculator   0.00009999500033330833533316668095113106348206440107107551266129432164491607407171907733994721288860975
		// Original Local Copilot     0.000099995000333308335333166680951
		//      Bing Copilot          0.000099995000333308335333166680555341171633562180100586805488327095926381493498941289678
		//                    Google  0.00009999500033329732302199351671461307129289868903350109720552546654971833095803199109004110324709458890

		// Small
		{"0.125", "-2.079441541679835928251696364375"},
		//                       −2.0794415416798359282516963643745
		// Large
		{"3237", "8.082402253926244350924204901779"},

		// Powers of 2
		{"2", "0.693147180559945309417232121458"},
		{"4", "1.386294361119890618834464242916"},
		{"8", "2.079441541679835928251696364375"},
		//                   2.079441541679835928251696364374 5
		// Mid-range
		// {"17.5", "2.862200880929468980535438056161"},
		{"17.5", "2.86220088092946837028887995521119080"},
	}

	dispatch := new(NaturalLogDispatcher)

	bFloatHlpr := new(BigFloatHelper)

	var workingPrecision uint

	var sourceExpectedDecimalDigits, sourceInputDecimalDigits, sourceCalculatedDecimalDigits int

	for idx, tc := range tests {

		tc.expected = bFloatHlpr.CleanNumberString(tc.expected)

		_, sourceExpectedDecimalDigits = bFloatHlpr.CountDigits(tc.expected, '.')

		_, sourceInputDecimalDigits = bFloatHlpr.CountDigits(tc.xStr, '.')

		if sourceInputDecimalDigits > sourceExpectedDecimalDigits {
			sourceCalculatedDecimalDigits = sourceInputDecimalDigits
		} else {
			sourceCalculatedDecimalDigits = sourceExpectedDecimalDigits
		}

		if sourceCalculatedDecimalDigits == 0 {
			sourceCalculatedDecimalDigits = 3
		} else {
			sourceCalculatedDecimalDigits += 3
		}

		workingPrecision = bFloatHlpr.ComputeBigFloatPrecisionBits(uint(sourceCalculatedDecimalDigits), 1)

		bFloatXValue, ok := new(big.Float).
			SetMode(big.AwayFromZero).
			SetPrec(workingPrecision).
			SetString(tc.xStr)

		if !ok {
			t.Errorf("%v\n"+
				"FAILED: SetString(%v)\n"+
				"Structure Index: %v\n",
				ePrefix, tc.xStr, idx)
			continue
		}

		actual, err := dispatch.Compute(bFloatXValue, workingPrecision)

		if err != nil {
			t.Errorf("%v\n"+
				"Error computing ln(%v): %v\n",
				ePrefix,
				tc.xStr,
				err)
			continue
		}

		actualStr := actual.Text('f', sourceExpectedDecimalDigits)

		if actualStr != tc.expected {
			t.Errorf("%v\n"+
				"ln(%v) mismatch\n"+
				"Expected: %v\n"+
				"  Actual: %v\n"+
				"sourceCalculatedDecimalDigits = %v\n"+
				"Big Float Precision Bits - workingPrecision = %v\n",
				ePrefix,
				tc.xStr,
				tc.expected,
				actualStr,
				sourceCalculatedDecimalDigits,
				workingPrecision)
		}

	} // End of 'for' loop

}
