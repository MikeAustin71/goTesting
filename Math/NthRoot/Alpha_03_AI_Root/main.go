package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Square Root of 25

	fmt.Println("Calculate 4th Root of 625")
	fmt.Println()

	bigNum := new(big.Rat).SetFrac64(25, 1)

	var bigRatRoot *big.Rat

	bigRatRoot = NthRoot(2, bigNum)

	var bFloatAnswer string

	bFloatAnswer = bigRatRoot.FloatString(5)

	fmt.Printf("Answer\n")
	fmt.Printf("%v\n\n", bFloatAnswer)

}

func NthRoot(n int64, x *big.Rat) *big.Rat {
	if n == 0 {
		return nil
	}
	if x.Sign() == 0 {
		return x
	}
	if n == 1 {
		return x
	}
	if n == 2 {
		return Sqrt(x)
	}
	var (
		neg = false
		r   = new(big.Rat).Set(x)
		n1  = new(big.Int).SetInt64(n - 1)
		n2  = new(big.Int).SetInt64(n - 2)

		/*
			n3  = new(big.Int).SetInt64(n - 3)
			n4  = new(big.Int).SetInt64(n - 4)
			n5  = new(big.Int).SetInt64(n - 5)
			n6  = new(big.Int).SetInt64(n - 6)
		*/
	)
	if r.Sign() < 0 {
		neg = true
		r.Neg(r)
	}
	var (
		t1 = new(big.Rat).Set(r)
		t2 = new(big.Rat).Quo(t1, big.NewRat(int64(n), 1))
		t3 = new(big.Rat).Add(t2, big.NewRat(1, int64(n)))
		t4 = new(big.Rat).Quo(t3, big.NewRat(2, int64(n)))
	)
	for i := 0; i < 100; i++ {
		t1.Quo(r, Exp(t4, n1))
		t2.Mul(t4, big.NewRat(int64(n-1), int64(n)))
		t3.Quo(t1, Exp(t4, n2))
		t4.Add(t4, t3)
	}
	if neg && (n%2) == 0 {
		return nil
	}
	return t4
}

func Exp(x *big.Rat, y *big.Int) *big.Rat {
	if y.Sign() == -1 {
		panic("negative exponent")
	}
	if y.Sign() == 0 {
		return big.NewRat(1, 1)
	}
	if y.BitLen() == 0 {
		return big.NewRat(1, 1)
	}
	z := new(big.Rat).Set(x)
	r := big.NewRat(1, 1)
	for i := y.BitLen() - 2; i >= 0; i-- {
		r.Mul(r, r)
		if y.Bit(i) != 0 {
			r.Mul(r, z)
		}
	}
	return r
}

func Sqrt(x *big.Rat) *big.Rat {
	return NthRoot(2, x)
}
