package naturalLogCalcs

import (
	"fmt"
	"math/big"
)

type naturalLogAGM struct{}

// lnAGMDirect computes ln(x) using the Borwein AGM logarithm algorithm.
//
// It assumes x > 0 and returns a result rounded to 'prec' bits.
func (nl *naturalLogAGM) lnAGMDirect(x *big.Float, prec uint) (*big.Float, error) {

	if x.Sign() <= 0 {
		return nil, fmt.Errorf("lnAGMDirect: x must be > 0")
	}

	// Working precision with safety margin.
	workPrec := prec + 64

	// Core Borwein AGM log computation.
	nlAGMMech := new(naturalLogAGMMechanics)
	lnxWork := nlAGMMech.lnAGMBorweinCore(x, workPrec)

	// Round to requested precision.
	nlShared := new(naturalLogShared)
	lnx := nlShared.newFloat(prec)
	lnx.Set(lnxWork)

	return lnx, nil
}

// lnAGMMantissa computes ln(x) using mantissa/exponent normalization
// combined with the Borwein AGM logarithm algorithm.
//
//	x = m * 2^k, m in [1,2)
//	ln(x) = ln(m) + k * ln(2)
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

	// ln(m) via Borwein AGM core.
	lnmWork := nlAGMMech.lnAGMBorweinCore(m, workPrec)

	// k * ln(2) using full-precision ln(2).
	nlSharedMech := new(naturalLogSharedMechanics)
	ln2 := nlSharedMech.getLn2(workPrec) // ln2Full

	kInt := big.NewInt(int64(k))
	kFloat := nlShared.newFloat(workPrec)
	kFloat.SetInt(kInt)

	kLn2 := nlShared.newFloat(workPrec)
	kLn2.Mul(kFloat, ln2)

	// ln(x) = ln(m) + k * ln(2)
	lnxWork := nlShared.newFloat(workPrec)
	lnxWork.Add(lnmWork, kLn2)

	// Round to requested precision.
	lnx := nlShared.newFloat(prec)
	lnx.Set(lnxWork)

	return lnx, nil
}
