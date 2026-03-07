package naturalLogCalcs

import (
	"fmt"
	"math/big"
)

type NaturalLogDispatcher struct{}

// Compute chooses the production algorithm for ln(x).
// Currently: Taylor is the production path.
// AGM is marked experimental and will return an error.
func (nlogDispatcher *NaturalLogDispatcher) Compute(x *big.Float, prec uint) (*big.Float, error) {

	if x.Sign() <= 0 {
		return nil, fmt.Errorf("lnDispatcher: x must be > 0")
	}

	// Production path: Taylor
	return new(naturalLogTaylor).lnTaylorDirect(x, prec)
}
