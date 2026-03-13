package naturalLogCalcs

import (
	"math/big"
	"testing"
)

func Test_Natural_Logarithm_Dispatcher_Range(t *testing.T) {

	tests := []struct {
		xStr     string
		expected string
	}{
		// Near 1
		// {"1.0001", "0.000099995000333308335416654166250"},
		{"1.0001", "0.000099995000333308335333166681"},
		//               Calculator   0.00009999500033330833533316668095113106348206440107107551266129432164491607407171907733994721288860975
		// Original Local Copilot     0.000099995000333308335333166680951
		//      Bing Copilot          0.000099995000333308335333166680555341171633562180100586805488327095926381493498941289678
		//                    Google  0.00009999500033329732302199351671461307129289868903350109720552546654971833095803199109004110324709458890
		//                              12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901
		//                                       1         2         3         4         5         6         7         8         9         0
		//                                                                                                                                 0
		//                                                                                                                                 1

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

	var workingPrecision uint

	var sourceDecimalDigits int

	dispatch := new(NaturalLogDispatcher)

	bFloatHlpr := new(BigFloatHelper)

	var decDigits int

	for _, tc := range tests {

		_, sourceDecimalDigits = bFloatHlpr.CountDigits(tc.expected, '.')

		if sourceDecimalDigits == 0 {
			sourceDecimalDigits = 3
		} else {
			sourceDecimalDigits += 3
		}

		workingPrecision = bFloatHlpr.ComputeBigFloatPrecisionBits(uint(sourceDecimalDigits), 1)

		x, ok := new(big.Float).
			SetMode(big.AwayFromZero).
			SetPrec(workingPrecision).
			SetString(tc.xStr)

		if !ok {
			t.Errorf("FAILED: SetString(%v)", tc.xStr)
			continue
		}

		actual, err := dispatch.Compute(x, workingPrecision)

		if err != nil {
			t.Errorf("Error computing ln(%v): %v", tc.xStr, err)
			continue
		}

		_, decDigits = bFloatHlpr.CountDigits(tc.expected, '.')

		//actualStr := actual.Text('f', 30)

		actualStr := actual.Text('f', decDigits)

		if actualStr != tc.expected {
			t.Errorf("\nln(%v) mismatch\nExpected: %v\nActual:   %v\n",
				tc.xStr, tc.expected, actualStr)
		}
	}
}
