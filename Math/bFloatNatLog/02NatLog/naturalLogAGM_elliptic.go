package naturalLogCalcs

import (
	"math"
	"math/big"
)

// Computes AGM(1, k\') and elliptic integral K(k)

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

	ePrefix := "EllipticContext.ComputeAGMForElliptic()"

	two := new(big.Float).SetPrec(ctx.Prec).SetFloat64(2.0)

	epsilon := new(big.Float).SetPrec(ctx.Prec).SetMantExp(big.NewFloat(1.0), -int(ctx.Prec)+4)

	a := new(big.Float).SetPrec(ctx.Prec).SetFloat64(1.0)
	g := new(big.Float).SetPrec(ctx.Prec).Copy(ctx.KPrime)

	for i := 0; i < 256; i++ {

		// a_{n+1} = (a + g) / 2
		aNext := new(big.Float).SetPrec(ctx.Prec).Add(a, g)

		aNext.Quo(aNext, two)

		// g_{n+1} = sqrt(a * g)
		prod := new(big.Float).SetPrec(ctx.Prec).Mul(a, g)

		gNext := new(big.Float).SetPrec(ctx.Prec).Sqrt(prod)

		// convergence check
		diff := new(big.Float).SetPrec(ctx.Prec).Sub(aNext, gNext)

		if diff.Abs(diff).Cmp(epsilon) < 0 {
			ctx.AGM.Copy(aNext)
			return nil
		}

		a = aNext
		g = gNext
	}

	return &FuncReturnError{
		ErrPrefix:  ePrefix,
		ReturnFunc: "",
		ErrContext: "",
		ErrMessage: "Error: Unexpected Result!\n" +
			"AGM(1,k') did not converge\n\n",
	}
}

func (ctx *EllipticContext) ComputeEllipticK() (*big.Float, error) {

	ePrefix := "EllipticContext.ComputeEllipticK()"

	if ctx.AGM.Sign() == 0 {
		return nil,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Unexpected Result!\n" +
					"AGM not computed\n\n",
			}
	}

	pi := new(big.Float).SetPrec(ctx.Prec).SetFloat64(math.Pi)

	two := new(big.Float).SetPrec(ctx.Prec).SetFloat64(2.0)

	denom := new(big.Float).SetPrec(ctx.Prec).Mul(two, ctx.AGM)

	K := new(big.Float).SetPrec(ctx.Prec).Quo(pi, denom)

	return K, nil
}
