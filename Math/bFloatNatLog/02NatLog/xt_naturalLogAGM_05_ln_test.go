package naturalLogCalcs

import (
	"math/big"
	"testing"
)

func Test_LnAGM_Mantissa_05_01(t *testing.T) {
	m := big.NewFloat(1.5)
	prec := uint(2048)

	ctx := NewLnAGMContext(m, prec)
	_, err := ctx.LnAGMMantissa()
	if err != nil {
		t.Fatalf("LnAGMMantissa failed: %v", err)
	}
}
