package naturalLogCalcs

import (
	"math/big"
)

// && = and
// Computes k = (m-1)/(m+1) && k' = sqrt(1 - k^2)

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

	ePrefix := "ModulusContext.ComputeModulus()"

	one := new(big.Float).SetPrec(ctx.Prec).SetFloat64(1.0)

	// k = (m - 1) / (m + 1)
	num := new(big.Float).SetPrec(ctx.Prec).Sub(ctx.Mantissa, one)

	den := new(big.Float).SetPrec(ctx.Prec).Add(ctx.Mantissa, one)

	if den.Cmp(new(big.Float).SetFloat64(0)) == 0 {

		return &FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Unexpected Result!\n" +
					"modulus computation: denominator == zero\n\n",
			}
	}

	ctx.K.Quo(num, den)

	// k' = sqrt(1 - k^2)
	k2 := new(big.Float).SetPrec(ctx.Prec).Mul(ctx.K, ctx.K)

	oneMinusK2 := new(big.Float).SetPrec(ctx.Prec).Sub(one, k2)

	if oneMinusK2.Sign() < 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix,
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "Error: Unexpected Result!\n" +
				"modulus computation: 1 - k^2 < 0\n\n",
		}
	}

	// sqrt(1 - k^2)
	ctx.KPrime.Sqrt(oneMinusK2)

	return nil
}
