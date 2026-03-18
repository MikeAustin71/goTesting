package naturalLogCalcs

// 010Fix - naturalLogAGM_mechanics.go

import "math/big"

type naturalLogAGMMechanics struct{}

// normalizeMantissa
//
// Rewrites x as:
//
//	x = m · 2^k
//
// with mantissa m constrained to the interval [1, 2).
// This is the standard floating‑point style normalization
// and is a natural fit for an AGM‑based ln(x) pipeline:
//
//	ln(x) = ln(m) + k · ln(2)
func (mec *naturalLogAGMMechanics) normalizeMantissa(
	x *big.Float,
	prec uint,
) (*big.Float, int, error) {

	ePrefix := "naturalLogAGMMechanics.normalizeMantissa()"

	if x == nil {
		return nil, 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix,
				ParameterName: "'x'",
			}
	}

	if x.Sign() <= 0 {
		return nil, 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "normalizeMantissa(x, prec)",
				ErrMessage: "input x must be > 0",
			}
	}

	// Work on a copy to avoid mutating the caller’s value.
	tmp := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)
	tmp.Copy(x)

	mant := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	// MantExp returns mantissa in [0.5, 1) and exponent e such that:
	//
	//   tmp = mant · 2^e
	exp := tmp.MantExp(mant)

	one := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)
	one.SetInt64(1)

	two := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)
	two.SetInt64(2)

	// Adjust to [1, 2) if necessary.
	if mant.Cmp(one) < 0 {
		mant.Mul(mant, two) // mant *= 2
		exp--               // compensate exponent
	}

	return mant, exp, nil
}
