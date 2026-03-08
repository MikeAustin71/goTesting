package naturalLogCalcs

import "math/big"

type naturalLogTaylorMechanics struct{}

// lnTaylorCore computes ln(m) for m in (0, 2] using the
// atanh-style series:
//
//	t = (m - 1) / (m + 1)
//	ln(m) = 2 * (t + t^3/3 + t^5/5 + ...)
//
// It assumes m > 0 and uses workPrec bits.
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

  // --------------------------------------------------------------------
  // 006 Fix
  // Corrected stopping criterion:
  //
  // Stop when |next term| < 2^(-(prec + 4))
  //
  // This ensures the Taylor series runs one decimal digit beyond the
  // target precision (≈ 4 bits of safety), preventing premature exit
  // for values near 1.
  // --------------------------------------------------------------------
  eps := nlShared.newFloat(workPrec)
  eps.SetFloat64(1.0)
  eps.SetMantExp(eps, -int(workPrec-64+4))
  // Explanation:
  //   workPrec = prec + 64
  //   We want threshold = 2^(-(prec + 4))
  //   So exponent = -(workPrec - 64 + 4)

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
    if absTerm.Cmp(eps) <= 0 {
      break
    }
  }

  // ln(m) = 2 * sum
  lnm := nlShared.newFloat(workPrec)
  lnm.Mul(two, sum)

  return lnm
}
