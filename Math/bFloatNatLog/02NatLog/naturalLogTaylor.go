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
//
// Special-case: If x is extremely close to 1, skip normalization
// to avoid rounding error amplification.
//
// 007Fix (near-1 bypass) retained.
func (nl *naturalLogTaylor) lnTaylorDirect(x *big.Float, prec uint) (*big.Float, error) {

	if x.Sign() <= 0 {
		return nil, fmt.Errorf("lnTaylorDirect: x must be > 0")
	}

	// Working precision with safety margin.
	workPrec := prec + 64

	// 008Fix01 - Calculate working decimal digits of precision.
	workingDecimalDigitsOfPrecision := new(BigFloatHelper).ComputeBigFloatPrecisionBits(workPrec, 1)

	nlShared := new(naturalLogShared)

	xWork := nlShared.newFloat(workPrec)
	xWork.Set(x)

	// ------------------------------------------------------------
	// Near-1 bypass (007Fix)
	// ------------------------------------------------------------
	one := nlShared.newFloat(workPrec)
	one.SetFloat64(1.0)

	delta := nlShared.newFloat(workPrec)
	delta.Sub(xWork, one)

	absDelta := nlShared.newFloat(workPrec)
	absDelta.Abs(delta)

	// Threshold for "near 1" — values with |x - 1| < 1e-3
	// are handled without mantissa normalization.
	threshold := nlShared.newFloat(workPrec)
	threshold.SetFloat64(1e-3)

	var m *big.Float
	var k int

	if absDelta.Cmp(threshold) < 0 {
		// Skip normalization entirely for near-1 values.
		m = nlShared.newFloat(workPrec)
		m.Set(xWork)
		k = 0
	} else {
		// Normal path: mantissa/exponent normalization.
		nlAGMMech := new(naturalLogAGMMechanics)
		m, k = nlAGMMech.normalizeMantissa(xWork, workPrec)
	}

	// ln(m) via Taylor core (now decimal-aware stopping rule).
	nlTaylorMech := new(naturalLogTaylorMechanics)
	lnmWork := nlTaylorMech.lnTaylorCore(m, workPrec)

	// k * ln(2)
	nlSharedMech := new(naturalLogSharedMechanics)
	//ln2 := nlSharedMech.getLn2(workPrec)

	// 008Fix01 - Calculate number of digits needed for
	// Natural Log of 2 constant value.
	uintNumOfLn2DecimalDigits := uint(DEFAULT_BASE_NUM_DECIMAL_DIGITS)

	if workingDecimalDigitsOfPrecision > 5000 {
		uintNumOfLn2DecimalDigits = 10000
	} else if workingDecimalDigitsOfPrecision > 1000 {
		uintNumOfLn2DecimalDigits = 4000
	}

	ln2 := nlSharedMech.getLn2(uintNumOfLn2DecimalDigits)

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
