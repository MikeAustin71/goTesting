package naturalLogCalcs

import (
	"math/big"
	"testing"
)

// 007Fix
func Test_Natural_Logarithm_Taylor_NearOne(t *testing.T) {

	tests := []struct {
		xStr     string
		expected string
	}{
		// Very close to 1
		{"1.0000001", "0.000000099999995000000333333341666666708"},
		{"1.000001", "0.000000999999500000333333416666708333208"},
		{"1.00001", "0.000009999950000333341666708333208333208"},
		{"1.0001", "0.000099995000333308335333166680951"},
		{"1.001", "0.000999500333083423333416708333208"},
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
