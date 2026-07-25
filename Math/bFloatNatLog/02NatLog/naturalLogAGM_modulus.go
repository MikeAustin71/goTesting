package naturalLogCalcs

import "math/big"

// Computes k = (m-1)/(m+1) and k' = sqrt(1 - k^2)

type ModulusContext struct {
	Mantissa *big.Float
	K        *big.Float
	KPrime   *big.Float
	Prec     uint
}

func NewModulusContext(m *big.Float, prec uint) *ModulusContext {
	return &ModulusContext{
		Mantissa: new(big.Float).Copy(m),
		Prec:     prec,
		K:        new(big.Float).SetPrec(prec),
		KPrime:   new(big.Float).SetPrec(prec),
	}
}

func (ctx *ModulusContext) ComputeModulus() error {
	// TODO: implement:
	// k  = (m - 1) / (m + 1)
	// k' = sqrt(1 - k^2)
	return nil
}
