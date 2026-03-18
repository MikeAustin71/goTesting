package naturalLogCalcs

import (
	"math/big"
	"testing"
)

// 010Fix — AGM validation using verified Taylor expected values
func Test_Natural_Logarithm_AGM_01(t *testing.T) {

	tests := []struct {
		xStr     string
		expected string
	}{
		// Same values used in Taylor tests
		{"3237", "8.082402253926244350924204901779"},
		{"245", "5.50125821054472698481146482011254709879977081343224009883142819134933871321179630020545087526892266478608969143341784472317150974468024988593153280779150727936782857370573281241801155421927627388228010260330798331306088066067384276824977246856395981"},
		{"1.0001", "0.00009999500033330833533316668095113106348206440107107551266129432164491607407171907733994721288860975"},
		{"1.0000001", "0.00000009999999500000033333330833333533333316666668095237970238106349205349206440115431782107551337480"},
		{"0.125", "-2.079441541679835928251696364375"},
		{"2", "0.693147180559945309417232121458"},
		{"4", "1.386294361119890618834464242916"},
		{"8", "2.079441541679835928251696364375"},
	}

	agm := new(naturalLogAGM)
	bFloatHlpr := new(BigFloatHelper)

	for _, tc := range tests {

		// Determine required decimal digits
		_, decDigits := bFloatHlpr.CountDigits(tc.expected, '.')
		decDigits += 3

		precBits := bFloatHlpr.ComputeBigFloatPrecisionBits(uint(decDigits), 1)

		xVal, ok := new(big.Float).
			SetMode(big.AwayFromZero).
			SetPrec(precBits).
			SetString(tc.xStr)

		if !ok {
			t.Fatalf("SetString failed for %v", tc.xStr)
		}

		actual, err := agm.lnAGMDirect(xVal, precBits)
		if err != nil {
			t.Fatalf("AGM error for ln(%v): %v", tc.xStr, err)
		}

		actualStr := actual.Text('f', decDigits-3)

		if actualStr != tc.expected {
			t.Errorf("ln(%v) mismatch\nExpected: %v\nActual:   %v\n",
				tc.xStr, tc.expected, actualStr)
		}
	}
}
