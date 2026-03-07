package naturalLogCalcs

import "math/big"

type naturalLogAGMMechanics struct{}

// agm computes the arithmetic-geometric mean of (a0, b0)
// at the given working precision.
func (nlAgmMech *naturalLogAGMMechanics) agm(a0, b0 *big.Float, workPrec uint) *big.Float {

	a := new(big.Float).SetPrec(workPrec)
	b := new(big.Float).SetPrec(workPrec)

	a.Set(a0)
	b.Set(b0)

	half := new(big.Float).SetPrec(workPrec)
	half.SetFloat64(0.5)

	tmpA := new(big.Float).SetPrec(workPrec)
	tmpB := new(big.Float).SetPrec(workPrec)
	diff := new(big.Float).SetPrec(workPrec)

	// Epsilon ~ 2^(-workPrec + 4)
	eps := new(big.Float).SetPrec(workPrec)
	eps.SetFloat64(1.0)
	eps.SetMantExp(eps, -int(workPrec-4))

	for {
		// a_{n+1} = (a_n + b_n) / 2
		tmpA.Add(a, b)
		tmpA.Mul(tmpA, half)

		// b_{n+1} = sqrt(a_n * b_n)
		tmpB.Mul(a, b)
		tmpB.Sqrt(tmpB)

		// |a_{n+1} - b_{n+1}|
		diff.Sub(tmpA, tmpB)
		if diff.Sign() < 0 {
			diff.Neg(diff)
		}

		if diff.Cmp(eps) <= 0 {
			a.Set(tmpA)
			break
		}

		a.Set(tmpA)
		b.Set(tmpB)
	}

	return a
}

// normalizeMantissa takes x and returns (m, k) such that:
//
//	x = m * 2^k, with m in [1, 2).
func (nlAgmMech *naturalLogAGMMechanics) normalizeMantissa(x *big.Float, workPrec uint) (*big.Float, int) {

	m := new(big.Float).SetPrec(workPrec)
	k := x.MantExp(m) // m in [0.5, 1)

	// Adjust to [1,2): m <- 2*m, k <- k-1
	two := new(big.Float).SetPrec(workPrec)
	two.SetFloat64(2.0)

	m.Mul(m, two)
	k--

	return m, k
}
