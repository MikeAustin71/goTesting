package naturalLogCalcs

import (
	"fmt"
	"math/big"
)

type naturalLogTaylor struct{}

// lnTaylorDirect computes ln(x) using Taylor series with
// mantissa/exponent normalization.
//
//	x = m * 2^k, m in [1,2)
//	ln(x) = ln(m) + k * ln(2)
func (nl *naturalLogTaylor) lnTaylorDirect(x *big.Float, prec uint) (*big.Float, error) {

	if x.Sign() <= 0 {
		return nil, fmt.Errorf("lnTaylorDirect: x must be > 0")
	}

	workPrec := prec + 64
	nlShared := new(naturalLogShared)

	xWork := nlShared.newFloat(workPrec)
	xWork.Set(x)

	// Normalize mantissa into [1,2) and get exponent k.
	nlAGMMech := new(naturalLogAGMMechanics)
	m, k := nlAGMMech.normalizeMantissa(xWork, workPrec)

	// ln(m) via Taylor core.
	nlTaylorMech := new(naturalLogTaylorMechanics)
	lnmWork := nlTaylorMech.lnTaylorCore(m, workPrec)

	// k * ln(2)
	nlSharedMech := new(naturalLogSharedMechanics)
	ln2 := nlSharedMech.getLn2(workPrec)

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
