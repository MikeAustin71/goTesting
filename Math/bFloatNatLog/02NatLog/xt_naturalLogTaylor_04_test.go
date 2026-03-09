package naturalLogCalcs

import (
	"math/big"
	"testing"
)

// 007Fix
func Test_Natural_Logarithm_Taylor_NearMinusOne(t *testing.T) {

	tests := []struct {
		xStr     string
		expected string
	}{
		{"0.9999", "-0.000100005000333416708416708333208"},
		{"0.99999", "-0.000010000050000333341666708333208"},
		{"0.999999", "-0.000001000000500000333333416666708"},
	}

	prec := uint(7000)
	taylor := new(naturalLogTaylor)
	bFloatHlpr := new(BigFloatHelper)

	for _, tc := range tests {

		x, ok := new(big.Float).
			SetMode(big.AwayFromZero).
			SetString(tc.xStr)

		if !ok {
			t.Errorf("FAILED: SetString(%v)", tc.xStr)
			continue
		}

		actual, err := taylor.lnTaylorDirect(x, prec)
		if err != nil {
			t.Errorf("Error computing ln(%v): %v", tc.xStr, err)
			continue
		}

		_, decDigits := bFloatHlpr.CountDigits(tc.expected, '.')
		actualStr := actual.Text('f', decDigits)

		if actualStr != tc.expected {
			t.Errorf("\nln(%v) mismatch\nExpected: %v\nActual:   %v\n",
				tc.xStr, tc.expected, actualStr)
		}
	}
}
