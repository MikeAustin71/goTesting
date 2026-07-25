package naturalLogCalcs

import "math/big"

// Computes AGM(1, k') and elliptic integral K(k)

type EllipticContext struct {
	K      *big.Float
	KPrime *big.Float
	AGM    *big.Float
	Prec   uint
}

func NewEllipticContext(mod *ModulusContext) *EllipticContext {
	return &EllipticContext{
		K:      new(big.Float).Copy(mod.K),
		KPrime: new(big.Float).Copy(mod.KPrime),
		AGM:    new(big.Float).SetPrec(mod.Prec),
		Prec:   mod.Prec,
	}
}

func (ctx *EllipticContext) ComputeAGMForElliptic() error {
	// TODO: AGM(1, k')
	return nil
}

func (ctx *EllipticContext) ComputeEllipticK() (*big.Float, error) {
	// TODO: K(k) = π / (2 * AGM(1, k'))
	K := new(big.Float).SetPrec(ctx.Prec)
	return K, nil
}
