package naturalLogCalcs

import (
	"fmt"
	"math/big"
)

type naturalLogTaylorMechanics struct{}

// lnTaylorCore computes ln(m) for m in (0, 2] using the
// atanh-style series:
//
//	t = (m - 1) / (m + 1)
//	ln(m) = 2 * (t + t^3/3 + t^5/5 + ...)
//
// It assumes m > 0 and uses workPrec bits.
//
// 008Fix (decimal-aware stopping rule, based on Option 2):
//   - Derive an effective decimal precision from workPrec.
//   - Build a decimal epsilon ~ 10^{-(d+1)}.
//   - Stop when |next term| < epsDec.
func (nlTaylorMech *naturalLogTaylorMechanics) lnTaylorCore(m *big.Float, workPrec uint) *big.Float {

	nlShared := new(naturalLogShared)

	one := nlShared.newFloat(workPrec)
	one.SetFloat64(1.0)

	two := nlShared.newFloat(workPrec)
	two.SetFloat64(2.0)

	// t = (m - 1) / (m + 1)
	num := nlShared.newFloat(workPrec)
	den := nlShared.newFloat(workPrec)

	num.Sub(m, one)
	den.Add(m, one)

	t := nlShared.newFloat(workPrec)
	t.Quo(num, den)

	// term = t
	term := nlShared.newFloat(workPrec)
	term.Set(t)

	// sum = t
	sum := nlShared.newFloat(workPrec)
	sum.Set(t)

	// t2 = t^2
	t2 := nlShared.newFloat(workPrec)
	t2.Mul(t, t)

	// k = 1, next denominator = 3, 5, 7, ...
	k := int64(1)

	// ------------------------------------------------------------
	// Decimal-aware stopping rule (Option 2)
	//
	// workPrec = prec + 64 (from caller).
	// Approximate decimal digits from the "effective" precision:
	//
	//   precBits ≈ workPrec - 64
	//   d ≈ precBits * log10(2)
	//
	// Then choose epsilon_dec ≈ 10^{-(d+1)} and stop when
	// |next term| < epsilon_dec.
	// ------------------------------------------------------------
	const log10of2 = 0.3010299956639812 // log10(2)

	precBits := int(workPrec) - 64
	if precBits < 16 {
		// Fallback for very small workPrec.
		precBits = int(workPrec)
	}

	decDigits := int(float64(precBits) * log10of2)
	if decDigits < 1 {
		decDigits = 1
	}

	// One extra decimal digit of safety.
	effDigits := decDigits + 1

	// Build epsilon_dec = 10^{-effDigits} as a decimal string.
	epsStr := fmt.Sprintf("1e-%d", effDigits)

	epsDec := nlShared.newFloat(workPrec)
	// SetString parses the decimal scientific notation exactly.
	_, ok := epsDec.SetString(epsStr)
	if !ok {
		// Fallback: if parsing fails for any reason, use a binary epsilon
		// roughly equivalent to 10^{-effDigits}.
		// 10^{-d} ≈ 2^{-d / log10(2)}.
		binExp := int(-float64(effDigits) / log10of2)
		epsDec.SetFloat64(1.0)
		epsDec.SetMantExp(epsDec, binExp)
	}

	absTerm := nlShared.newFloat(workPrec)

	for {
		// term *= t^2
		term.Mul(term, t2)

		// k += 2
		k += 2

		// term / k
		denomK := nlShared.newFloat(workPrec)
		denomK.SetInt64(k)

		tmp := nlShared.newFloat(workPrec)
		tmp.Quo(term, denomK)

		// sum += term/k
		sum.Add(sum, tmp)

		// |term/k|
		absTerm.Abs(tmp)
		if absTerm.Cmp(epsDec) <= 0 {
			break
		}
	}

	// ln(m) = 2 * sum
	lnm := nlShared.newFloat(workPrec)
	lnm.Mul(two, sum)

	return lnm
}
