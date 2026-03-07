package naturalLogCalcs

import "math/big"

type naturalLogShared struct{}

// newFloat creates a *big.Float with the given precision and a
// consistent rounding mode.
func (nlShared *naturalLogShared) newFloat(prec uint) *big.Float {
	return new(big.Float).SetPrec(prec).SetMode(big.AwayFromZero)
}
