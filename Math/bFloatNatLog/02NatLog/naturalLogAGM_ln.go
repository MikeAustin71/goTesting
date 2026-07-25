package naturalLogCalcs

import "math/big"

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
	// TODO: build a_n, g_n
	return nil
}

func (ctx *LnAGMContext) computeCorrectionSeries() (*big.Float, error) {
	S := new(big.Float).SetPrec(ctx.Prec)
	// TODO: sum 2^(n+1)(a_n^2 - g_n^2)
	return S, nil
}

func (ctx *LnAGMContext) LnAGMMantissa() (*big.Float, error) {
	if err := ctx.Modulus.ComputeModulus(); err != nil {
		return nil, err
	}
	if err := ctx.Elliptic.ComputeAGMForElliptic(); err != nil {
		return nil, err
	}
	K, err := ctx.Elliptic.ComputeEllipticK()
	if err != nil {
		return nil, err
	}
	if err := ctx.buildAGMSequences(); err != nil {
		return nil, err
	}
	S, err := ctx.computeCorrectionSeries()
	if err != nil {
		return nil, err
	}
	lnM := new(big.Float).SetPrec(ctx.Prec)
	lnM.Sub(K, S)
	return lnM, nil
}
