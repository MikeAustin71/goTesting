package naturalLogCalcs

import (
	"fmt"
	"math/big"
)

type naturalLogAGM struct{}

// lnAGMDirect computes ln(x) using the direct AGM formula:
//
//	ln(x) = π / (2 * AGM(1, 4/x))
//
// It assumes x > 0 and returns a result rounded to 'prec' bits.
func (nl *naturalLogAGM) lnAGMDirect(x *big.Float, prec uint) (*big.Float, error) {

	if x.Sign() <= 0 {
		return nil, fmt.Errorf("lnAGM_Direct: x must be > 0")
	}

	// Working precision with safety margin.
	workPrec := prec + 64

	nlShared := new(naturalLogShared)

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

	nlAGMMech := new(naturalLogAGMMechanics)

	A := nlAGMMech.agm(one, t, workPrec)

	// Constants: π (full precision)
	nlSharedMech := new(naturalLogSharedMechanics)

	pi := nlSharedMech.getPi(workPrec) // returns piFull

	// term1 = π / (2 * A)
	two := nlShared.newFloat(workPrec)
	two.SetFloat64(2.0)

	denom := nlShared.newFloat(workPrec)
	denom.Mul(two, A)

	term1 := nlShared.newFloat(workPrec)
	term1.Quo(pi, denom)

	// ln(x) = term1
	lnxWork := nlShared.newFloat(workPrec)
	lnxWork.Set(term1)

	// Round to requested precision.
	lnx := nlShared.newFloat(prec)
	lnx.Set(lnxWork)

	return lnx, nil
}

// lnAGMMantissa computes ln(x) using mantissa/exponent normalization:
//
//	x = m * 2^k, m in [1,2)
//	ln(x) = π / (2 * AGM(1, 4/m)) + k * ln(2)
//
// It assumes x > 0 and returns a result rounded to 'prec' bits.
func (nl *naturalLogAGM) lnAGMMantissa(x *big.Float, prec uint) (*big.Float, error) {

	if x.Sign() <= 0 {
		return nil, fmt.Errorf("lnAGMMantissa: x must be > 0")
	}

	workPrec := prec + 64

	nlShared := new(naturalLogShared)

	xWork := nlShared.newFloat(workPrec)
	xWork.Set(x)

	// Normalize mantissa into [1,2) and get exponent k.
	nlAGMMech := new(naturalLogAGMMechanics)

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

	// Constants: π and ln(2) (full precision)
	nlShareMech := new(naturalLogSharedMechanics)

	pi := nlShareMech.getPi(workPrec)   // returns piFull
	ln2 := nlShareMech.getLn2(workPrec) // returns ln2Full

	// term1 = π / (2 * A)
	two := nlShared.newFloat(workPrec)
	two.SetFloat64(2.0)

	denom := nlShared.newFloat(workPrec)
	denom.Mul(two, A)

	term1 := nlShared.newFloat(workPrec)
	term1.Quo(pi, denom)

	// ln(m) = term1
	lnm := nlShared.newFloat(workPrec)
	lnm.Set(term1)

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
