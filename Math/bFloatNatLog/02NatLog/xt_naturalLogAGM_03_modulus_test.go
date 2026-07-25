package naturalLogCalcs

import (
	"math/big"
	"testing"
)

func Test_Modulus_Computation_03(t *testing.T) {
	m := big.NewFloat(1.5)
	prec := uint(1024)

	ctx := NewModulusContext(m, prec)
	if err := ctx.ComputeModulus(); err != nil {
		t.Fatalf("ComputeModulus failed: %v", err)
	}
}
