package naturalLogCalcs

import (
	"math/big"
	"testing"
)

func Test_Elliptic_K_04(t *testing.T) {
	m := big.NewFloat(1.5)
	prec := uint(1024)

	mod := NewModulusContext(m, prec)
	_ = mod.ComputeModulus()

	ell := NewEllipticContext(mod)
	_ = ell.ComputeAGMForElliptic()

	_, err := ell.ComputeEllipticK()
	if err != nil {
		t.Fatalf("ComputeEllipticK failed: %v", err)
	}
}
