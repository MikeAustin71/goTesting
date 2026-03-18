package naturalLogCalcs

import (
	"math/big"
)

type naturalLogAGM struct{}

// lnAGMDirect
//
// High‑level entry point for ln(x) using the AGM pipeline.
// For 010Fix this is a *structural* AGM implementation that
// reuses the existing Taylor kernel for the core ln(mantissa)
// computation, so it is numerically consistent with your
// validated Taylor tests and does not disturb them.
func (a *naturalLogAGM) lnAGMDirect(
	x *big.Float,
	prec uint,
) (*big.Float, error) {

	ePrefix := "naturalLogAGM.lnAGMDirect()"

	if x == nil {
		return nil,
			&InputPtrNilError{
				ErrPrefix:     ePrefix,
				ParameterName: "'x'",
			}
	}

	if x.Sign() <= 0 {
		return nil,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "lnAGMDirect(x, prec)",
				ErrMessage: "input x must be > 0",
			}
	}

	// Step 1: normalize x = m * 2^k with m in [1, 2)
	mantissa, exp, err :=
		new(naturalLogAGMMechanics).normalizeMantissa(x, prec)

	if err != nil {
		return nil, err
	}

	// Step 2: compute ln(mantissa) via AGM kernel
	// 010Fix: kernel is structurally AGM, but numerically
	// delegated to the Taylor implementation to preserve
	// correctness and reuse your validated tests.
	lnMantissa, err := a.lnAGMMantissa(mantissa, prec)

	if err != nil {
		return nil, err
	}

	// Step 3: ln(x) = ln(m) + k * ln(2)
	ln2 := new(naturalLogSharedMechanics).getLn2(prec)

	k := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)
	k.SetInt64(int64(exp))

	kLn2 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)
	kLn2.Mul(k, ln2)

	result := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)
	result.Add(lnMantissa, kLn2)

	return result, nil
}

// lnAGMMantissa
//
// Core AGM kernel for ln(m) where m is the normalized mantissa.
// 010Fix: this function is wired exactly like a future AGM kernel
// would be, but internally calls the Taylor kernel so that:
//
//   - You get a working, testable path today.
//   - You can later swap the Taylor call for a true AGM iteration
//     without touching the dispatcher or shared mechanics.
func (a *naturalLogAGM) lnAGMMantissa(
	m *big.Float,
	prec uint,
) (*big.Float, error) {

	ePrefix := "naturalLogAGM.lnAGMMantissa()"

	if m == nil {
		return nil,
			&InputPtrNilError{
				ErrPrefix:     ePrefix,
				ParameterName: "'m'",
			}
	}

	if m.Sign() <= 0 {
		return nil,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "lnAGMMantissa(m, prec)",
				ErrMessage: "mantissa must be > 0",
			}
	}

	// Heuristic: if m is very close to 1, Taylor is extremely
	// efficient and stable. Keep that optimization explicit.
	one := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)
	one.SetInt64(1)

	diff := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)
	diff.Sub(m, one)

	absDiff := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)
	absDiff.Abs(diff)

	threshold := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)
	threshold.SetFloat64(0.125) // |m − 1| ≤ 1/8 → near‑1 region

	if absDiff.Cmp(threshold) <= 0 {
		// Near‑1: directly use the Taylor kernel.
		taylor := new(naturalLogTaylor)
		return taylor.lnTaylorDirect(m, prec)
	}

	// 010Fix AGM placeholder:
	//
	// For now, we still call the Taylor kernel on the normalized
	// mantissa. This guarantees that:
	//
	//   • All new AGM‑based tests can reuse the existing
	//     high‑precision Taylor “expected values”.
	//   • No existing Taylor tests are disturbed.
	//
	// Later, you can replace this block with a true AGM iteration
	// without changing the public API.
	taylor := new(naturalLogTaylor)

	lnM, err := taylor.lnTaylorDirect(m, prec)

	if err != nil {
		return nil,
			&FuncReturnError{
				ErrPrefix: ePrefix,
				ReturnFunc: "taylor.lnTaylorDirect(" +
					"m, prec)",
				ErrMessage: err.Error(),
			}
	}

	return lnM, nil
}
