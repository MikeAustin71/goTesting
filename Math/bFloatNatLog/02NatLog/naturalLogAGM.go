package naturalLogCalcs

import (
	"fmt"
	"math/big"
)

type NaturalLogAGM struct{}

// lnAGMDirect computes ln(x) using the direct AGM formula:
//
//	ln(x) = π / (2 * AGM(1, 4/x)) - ln(4)
//
// It assumes x > 0 and returns a result rounded to 'prec' bits.
func (nl *NaturalLogAGM) lnAGMDirect(x *big.Float, prec uint) (*big.Float, error) {

	if x.Sign() <= 0 {
		return nil, fmt.Errorf("lnAGM_Direct: x must be > 0")
	}

	if x.Sign() <= 0 {
		return nil, fmt.Errorf("lnAGM_Direct: x must be > 0")
	}

	// Working precision with safety margin.
	workPrec := prec + 64

	nlShared := new(NaturalLogShared)

	xWork := nlShared.newFloat(workPrec)
	xWork.Set(x)

	// t = 4 / x
	four := nlShared.newFloat(workPrec)
	four.SetFloat64(4.0)

	t := nlShared.newFloat(workPrec)
	t.Quo(four, xWork)

	// AGM(1, t)
	one := nlShared.newFloat(workPrec)
	one.SetFloat64(1.0)

	nlAGMMech := new(NaturalLogAGMMechanics)

	A := nlAGMMech.agm(one, t, workPrec)

	// Constants: π and ln(2)
	nlSharedMech := new(NaturalLogSharedMechanics)
	pi := nlSharedMech.getPi(workPrec)
	ln2 := nlSharedMech.getLn2(workPrec)

	// term1 = π / (2 * A)
	two := nlShared.newFloat(workPrec)
	two.SetFloat64(2.0)

	denom := nlShared.newFloat(workPrec)
	denom.Mul(two, A)

	term1 := nlShared.newFloat(workPrec)
	term1.Quo(pi, denom)

	// ln(4) = 2 * ln(2)
	ln4 := nlShared.newFloat(workPrec)
	ln4.Mul(two, ln2)

	// ln(x) = term1 - ln(4)
	lnxWork := nlShared.newFloat(workPrec)
	lnxWork.Sub(term1, ln4)

	// Round to requested precision.
	lnx := nlShared.newFloat(prec)
	lnx.Set(lnxWork)

	return lnx, nil
}

// lnAGMMantissa computes ln(x) using mantissa/exponent normalization:
//
//	x = m * 2^k, m in [1,2)
//	ln(x) = π / (2 * AGM(1, 4/m)) - ln(4) + k * ln(2)
//
// It assumes x > 0 and returns a result rounded to 'prec' bits.
func (nl *NaturalLogAGM) lnAGMMantissa(x *big.Float, prec uint) (*big.Float, error) {

	if x.Sign() <= 0 {
		return nil, fmt.Errorf("lnAGMMantissa: x must be > 0")
	}

	workPrec := prec + 64

	nlShared := new(NaturalLogShared)

	xWork := nlShared.newFloat(workPrec)
	xWork.Set(x)

	// Normalize mantissa into [1,2) and get exponent k.
	nlAGMMech := new(NaturalLogAGMMechanics)

	m, k := nlAGMMech.normalizeMantissa(xWork, workPrec)

	// t = 4 / m
	four := nlShared.newFloat(workPrec)
	four.SetFloat64(4.0)

	t := nlShared.newFloat(workPrec)
	t.Quo(four, m)

	// AGM(1, t)
	one := nlShared.newFloat(workPrec)
	one.SetFloat64(1.0)

	A := nlAGMMech.agm(one, t, workPrec)

	// Constants: π and ln(2)
	nlShareMech := new(NaturalLogSharedMechanics)

	pi := nlShareMech.getPi(workPrec)
	ln2 := nlShareMech.getLn2(workPrec)

	// term1 = π / (2 * A)
	two := nlShared.newFloat(workPrec)
	two.SetFloat64(2.0)

	denom := nlShared.newFloat(workPrec)
	denom.Mul(two, A)

	term1 := nlShared.newFloat(workPrec)
	term1.Quo(pi, denom)

	// ln(4) = 2 * ln(2)
	ln4 := nlShared.newFloat(workPrec)
	ln4.Mul(two, ln2)

	// ln(m) = term1 - ln(4)
	lnm := nlShared.newFloat(workPrec)
	lnm.Sub(term1, ln4)

	// k * ln(2)
	kInt := big.NewInt(int64(k))
	kFloat := nlShared.newFloat(workPrec)
	kFloat.SetInt(kInt)

	kLn2 := nlShared.newFloat(workPrec)
	kLn2.Mul(kFloat, ln2)

	// ln(x) = ln(m) + k * ln(2)
	lnxWork := nlShared.newFloat(workPrec)
	lnxWork.Add(lnm, kLn2)

	// Round to requested precision.
	lnx := nlShared.newFloat(prec)
	lnx.Set(lnxWork)

	return lnx, nil
}
