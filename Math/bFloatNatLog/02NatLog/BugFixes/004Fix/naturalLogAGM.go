package naturalLogCalcs

import (
	"fmt"
	"math/big"
)

type naturalLogAGM struct{}

// lnAGMDirect is currently experimental and not suitable for production use.
func (nl *naturalLogAGM) lnAGMDirect(x *big.Float, prec uint) (*big.Float, error) {

	if x.Sign() <= 0 {
		return nil, fmt.Errorf("lnAGMDirect: x must be > 0")
	}

	return nil, fmt.Errorf("lnAGMDirect: AGM-based ln(x) is experimental and not yet implemented for production use")
}

// lnAGMMantissa is currently experimental and not suitable for production use.
func (nl *naturalLogAGM) lnAGMMantissa(x *big.Float, prec uint) (*big.Float, error) {

	if x.Sign() <= 0 {
		return nil, fmt.Errorf("lnAGMMantissa: x must be > 0")
	}

	return nil, fmt.Errorf("lnAGMMantissa: AGM-based ln(x) is experimental and not yet implemented for production use")
}
