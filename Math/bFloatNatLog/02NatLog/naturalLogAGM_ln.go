package naturalLogCalcs

import (
	"math/big"
)

// Full AGM ln(m) pipeline

type LnAGMContext struct {
	Mantissa     *big.Float
	Prec         uint
	Modulus      *ModulusContext
	Elliptic     *EllipticContext
	AGMSequenceA []*big.Float
	AGMSequenceG []*big.Float
}

func NewLnAGMContext(m *big.Float, prec uint) *LnAGMContext {
	ctx := &LnAGMContext{
		Mantissa: new(big.Float).Copy(m),
		Prec:     prec,
	}
	ctx.Modulus = NewModulusContext(ctx.Mantissa, prec)
	ctx.Elliptic = NewEllipticContext(ctx.Modulus)
	return ctx
}

func (ctx *LnAGMContext) buildAGMSequences() error {

	ePrefix := "LnAGMContext.buildAGMSequences()"

	two := new(big.Float).SetPrec(ctx.Prec).SetFloat64(2.0)

	epsilon := new(big.Float).SetPrec(ctx.Prec).SetMantExp(big.NewFloat(1.0), -int(ctx.Prec)+4)

	a := new(big.Float).SetPrec(ctx.Prec).SetFloat64(1.0)

	g := new(big.Float).SetPrec(ctx.Prec).Copy(ctx.Modulus.KPrime)

	ctx.AGMSequenceA = []*big.Float{new(big.Float).Copy(a)}

	ctx.AGMSequenceG = []*big.Float{new(big.Float).Copy(g)}

	for i := 0; i < 256; i++ {

		aNext := new(big.Float).SetPrec(ctx.Prec).Add(a, g)

		aNext.Quo(aNext, two)

		prod := new(big.Float).SetPrec(ctx.Prec).Mul(a, g)

		gNext := new(big.Float).SetPrec(ctx.Prec).Sqrt(prod)

		ctx.AGMSequenceA = append(ctx.AGMSequenceA, new(big.Float).Copy(aNext))

		ctx.AGMSequenceG = append(ctx.AGMSequenceG, new(big.Float).Copy(gNext))

		diff := new(big.Float).SetPrec(ctx.Prec).Sub(aNext, gNext)

		if diff.Abs(diff).Cmp(epsilon) < 0 {
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
			"AGM sequence did not converge\n\n",
	}

}

func (ctx *LnAGMContext) computeCorrectionSeries() (*big.Float, error) {

	S := new(big.Float).SetPrec(ctx.Prec)

	for n := 0; n < len(ctx.AGMSequenceA); n++ {

		a := ctx.AGMSequenceA[n]

		g := ctx.AGMSequenceG[n]

		a2 := new(big.Float).SetPrec(ctx.Prec).Mul(a, a)

		g2 := new(big.Float).SetPrec(ctx.Prec).Mul(g, g)

		c := new(big.Float).SetPrec(ctx.Prec).Sub(a2, g2)

		scale := new(big.Float).SetPrec(ctx.Prec).SetInt(new(big.Int).Lsh(big.NewInt(1), uint(n+1)))

		term := new(big.Float).SetPrec(ctx.Prec).Mul(scale, c)

		S.Add(S, term)
	}

	return S, nil
}

func (ctx *LnAGMContext) LnAGMMantissa() (*big.Float, error) {

	ePrefix := "LnAGMContext.LnAGMMantissa()"

	if err := ctx.Modulus.ComputeModulus(); err != nil {

		return nil,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "Error returned by ctx.Modulus.ComputeModulus()\n",
				ErrMessage: err.Error() + "\n\n",
			}
	}

	if err := ctx.Elliptic.ComputeAGMForElliptic(); err != nil {

		return nil,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "Error returned by ctx.Elliptic.ComputeAGMForElliptic()\n",
				ErrMessage: err.Error() + "\n\n",
			}
	}

	K, err := ctx.Elliptic.ComputeEllipticK()

	if err != nil {

		return nil,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "Error returned by ctx.Elliptic.ComputeEllipticK()\n",
				ErrMessage: err.Error() + "\n\n",
			}
	}

	if err := ctx.buildAGMSequences(); err != nil {

		return nil,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "Error returned by ctx.buildAGMSequences()\n",
				ErrMessage: err.Error() + "\n\n",
			}
	}

	S, err := ctx.computeCorrectionSeries()

	if err != nil {

		return nil,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "Error returned by ctx.computeCorrectionSeries()\n",
				ErrMessage: err.Error() + "\n\n",
			}
	}

	lnM := new(big.Float).SetPrec(ctx.Prec)

	lnM.Sub(K, S)

	return lnM, nil
}
