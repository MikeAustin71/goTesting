package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

// This is a backup of 01_NewtonExample
func main() {

	/*
		GolanSqrtExample()
		fmt.Println("Varification")
		GolangSqrtExampleVerify()

	*/

	Newton01()
	//TestNewtonInitialGuess()
	// TestBigFloatPower()
}

func GolanSqrtExample() {
	// We'll do computations with 200 bits of precision in the mantissa.
	const prec = 200

	// Compute the square root of 2 using Newton's Method. We start with
	// an initial estimate for sqrt(2), and then iterate:
	//     x_{n+1} = 1/2 * ( x_n + (2.0 / x_n) )

	// Since Newton's Method doubles the number of correct digits at each
	// iteration, we need at least log_2(prec) steps.
	steps := int(math.Log2(prec))

	// Initialize values we need for the computation.
	two := new(big.Float).SetPrec(prec).SetInt64(2)
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	// Use 1 as the initial estimate.
	x := new(big.Float).SetPrec(prec).SetInt64(1)

	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.
	t := new(big.Float)

	// Iterate.
	for i := 0; i <= steps; i++ {
		t.Quo(two, x)  // t = 2.0 / x_n
		t.Add(x, t)    // t = x_n + (2.0 / x_n)
		x.Mul(half, t) // x_{n+1} = 0.5 * t
	}

	// We can use the usual fmt.Printf verbs since big.Float implements fmt.Formatter
	fmt.Printf("sqrt(2) = %.50f\n", x)

	// Print the error between 2 and x*x.
	t.Mul(x, x) // t = x*x
	fmt.Printf("error = %e\n", t.Sub(two, t))

	//Output:
	//
	//	sqrt(2) = 1.41421356237309504880168872420969807856967187537695
	//	error = 0.000000e+00

}

func GolangSqrtExampleVerify() {

	testStr := "1.41421356237309504880168872420969807856967187537695"

	answer,
		ok := new(big.Float).SetString(testStr)

	if !ok {
		fmt.Printf("Error SetString(testStr) FAILED!\n"+
			"testStr = %v\n",
			testStr)
		return
	}

	t := big.NewFloat(0)
	t.Set(answer)

	result := t.Mul(t, answer)

	fmt.Printf("Result = '%v'\n",
		result.Text('f', -1))

}

func TestBigFloatPower() {

	funcName := "TestBigFloatPower"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	const prec = 4096
	power := int64(50)

	var ok bool

	base := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)
	_,
		ok = base.SetString("1.5")

	if !ok {
		panic("base.SetString(\"1.5\") FAILED!\n")
	}

	expectedResult := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	_,
		ok = expectedResult.SetString("637621500.21404958690340780691486")

	if !ok {
		panic("expectedResult.SetString(\"637621500.21404958690340780691486\") FAILED!\n")
	}

	raisedToPower := BigFloatPower(
		base,
		power,
		prec)

	fmt.Printf("BigFloatPower Results\n"+
		"      precision = %v\n"+
		"           base = %v\n"+
		"          power = %v\n"+
		"expected result = %v\n"+
		"  raisedToPower = \n%v\n\n",
		prec,
		base.Text('f', -1),
		power,
		expectedResult.Text('f', -1),
		raisedToPower.Text('f', -1))

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("  Successful Completion!\n" +
		"Function: " +
		funcName + "\n")

	fmt.Printf(breakStr + "\n")
}

func TestNewtonInitialGuess() {

	funcName := "TestNewtonInitialGuess"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	const prec = 16384

	nthRoot := int64(4)
	alphaInt64 := int64(6000000)

	alpha := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(alphaInt64)

	accuracyThreashold := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	accuracyThreashold.SetString("0.015")

	guess,
		raisedToPower,
		absAlphaDelta,
		compareResult,
		cycleCount := NewtonInitialGuess(
		nthRoot,
		alpha,
		accuracyThreashold,
		prec)

	fmt.Printf("Initial Guess Results\n"+
		"    precision = %v\n"+
		"      nthRoot = %v\n"+
		"compareResult = %v\n"+
		"   cycleCount = %v\n"+
		"        alpha = %v\n"+
		"raisedToPower = \n%v\n\n"+
		"absAlphaDelta = \n%v\n\n"+
		"        guess =\n%v\n\n",
		prec,
		nthRoot,
		compareResult,
		cycleCount,
		alpha.Text('f', -1),
		raisedToPower.Text('f', -1),
		absAlphaDelta.Text('f', -1),
		guess.Text('f', -1))

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("  Successful Completion!\n" +
		"Function: " +
		funcName + "\n")

	fmt.Printf(breakStr + "\n")
}

func Newton01() {

	funcName := "Newton01"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	const prec = 16384

	// Since Newton's Method doubles the number of correct digits at each
	// iteration, we need at least log_2(prec) steps.
	steps := int(math.Log2(prec))
	steps += 100

	// Alpha = 8
	alpha := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	alpha.SetString("8")

	// n = 3
	n_int64 := int64(23)
	// n_float := new(big.Float).SetPrec(prec).SetInt64(n_int64)

	// The first guess!
	//x_k := new(big.Float).SetPrec(prec).SetFloat64(1.0)

	accuracyThreashold := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	accuracyThreashold.SetString("0.000015")

	x_k,
		raisedToPower,
		absAlphaDelta,
		compareResult,
		cycleCount := NewtonInitialGuess(
		n_int64,
		alpha,
		accuracyThreashold,
		prec)

	fmt.Printf("First Guess Data\n"+
		"x_k           = %v\n"+
		"precision     = %v\n"+
		"raisedToPower = %v\n"+
		"absAlphaDelta = %v\n"+
		"compareResult = %v\n"+
		"  Cycle Count = %v\n\n",
		x_k.Text('f', 20),
		prec,
		raisedToPower.Text('f', 20),
		absAlphaDelta.Text('f', 20),
		compareResult,
		cycleCount)

	n_minus_1_int64 := n_int64 - 1

	n_minus_1 := new(big.Float).SetPrec(prec).SetInt64(n_minus_1_int64)

	one_over_n_rat := big.NewRat(1, n_int64)

	one_over_n_float := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).SetRat(one_over_n_rat)

	x_k1 := new(big.Float).SetPrec(prec).SetInt64(0)
	fac_1 := new(big.Float).SetPrec(prec).SetInt64(0)
	fac_2 := new(big.Float).SetPrec(prec).SetInt64(0)
	fac_3 := new(big.Float).SetPrec(prec).SetInt64(0)
	fac_4 := new(big.Float).SetPrec(prec).SetInt64(0)

	fmt.Printf("The number of steps = '%v'\n",
		steps)

	for i := 0; i <= steps; i++ {
		fac_1 = BigFloatPower(x_k, n_minus_1_int64, prec)
		fac_2.Quo(alpha, fac_1)
		fac_3.Mul(n_minus_1, x_k)
		fac_4.Add(fac_3, fac_2)
		x_k1.Mul(one_over_n_float, fac_4)
		x_k.Set(x_k1)
	}

	fmt.Printf("Calculation Factors\n"+
		"Alpha = '%v'\n"+
		"N = '%v'\n"+
		"Result: %v\n",
		alpha.Text('f', -1),
		n_int64,
		x_k.Text('f', -1))

	calcAlpha := BigFloatPower(
		x_k,
		n_int64,
		prec)

	fmt.Printf("Results Verification\n"+
		"Expected Alpha = '%v'\n"+
		"Verified Alpha = '%v'\n",
		alpha.Text('f', -1),
		calcAlpha.Text('f', -1))

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("  Successful Completion!\n" +
		"Function: " +
		funcName + "\n")

	fmt.Printf(breakStr + "\n")

	/*
			Output:

		Calculation Factors
		Alpha = '8'
		N = '3'
		Result:
		1.999999999999999999999999999999999999999999999999980000000000000000000000000000000000000003501118365183835461631804637785
	*/

}

func Newton02() {

	funcName := "Newton02"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	const prec = 800

	// Since Newton's Method doubles the number of correct digits at each
	// iteration, we need at least log_2(prec) steps.
	steps := int(math.Log2(prec))

	// Alpha = 8
	alpha_int64 := int64(592)
	alpha := new(big.Float).SetPrec(prec).SetInt64(alpha_int64)

	// n = 5
	n_int64 := int64(5)
	//n := new(big.Float).SetPrec(prec).SetInt64(n_int64)

	n_minus_1_int64 := n_int64 - 1

	n_minus_1 := new(big.Float).SetPrec(prec).SetInt64(n_minus_1_int64)

	one_over_n_rat := big.NewRat(1, n_int64)

	numStr := one_over_n_rat.FloatString(400)

	one_over_n_float,
		ok := new(big.Float).SetPrec(prec).SetString(numStr)

	if !ok {
		fmt.Printf("Error: big.Float.SetString(%v)\n",
			numStr)
		return
	}

	x_k := new(big.Float).SetPrec(prec).SetInt64(1)

	x_k1 := new(big.Float).SetPrec(prec).SetInt64(0)
	fac_1 := new(big.Float).SetPrec(prec).SetInt64(0)
	fac_2 := new(big.Float).SetPrec(prec).SetInt64(0)
	fac_3 := new(big.Float).SetPrec(prec).SetInt64(0)
	fac_4 := new(big.Float).SetPrec(prec).SetInt64(0)

	for i := 0; i <= steps; i++ {
		fac_1 = BigFloatPower(x_k, n_minus_1_int64, prec)
		fac_2.Quo(alpha, fac_1)
		fac_3.Mul(n_minus_1, x_k)
		fac_4.Add(fac_3, fac_2)
		x_k1.Mul(one_over_n_float, fac_4)
		x_k.Set(x_k1)
	}

	fmt.Printf("Calculation Factors\n"+
		"Alpha = '%v'\n"+
		"N = '%v'\n"+
		"Result: %v\n",
		alpha_int64,
		n_int64,
		x_k.Text('f', -1))

	calcAlpha := BigFloatPower(
		x_k,
		n_int64,
		prec)

	fmt.Printf("Results Verification\n"+
		"Expected Alpha = '%v'\n"+
		"Verified Alpha = '%v'\n",
		alpha_int64,
		calcAlpha.Text('f', -1))

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("  Successful Completion!\n" +
		"Function: " +
		funcName + "\n")

	fmt.Printf(breakStr + "\n")

	/*
			Output:

		Calculation Factors
		Alpha = '8'
		N = '3'
		Result:
		1.999999999999999999999999999999999999999999999999980000000000000000000000000000000000000003501118365183835461631804637785
	*/

}

func NewtonInitialGuess(
	nthRoot int64, // 18
	alpha *big.Float, // 592
	accuracyThreshold *big.Float,
	prec uint) (
	guess *big.Float,
	raisedToPower *big.Float,
	absAlphaDelta *big.Float,
	compareResult int,
	cycleCount uint64) {

	// SetMode sets z's rounding mode to mode and returns an exact z.
	// z remains unchanged otherwise.
	// z.SetMode(z.Mode()) is a cheap way to set z's accuracy to Exact.

	numOfLowGuesses := uint64(0)
	numOfHighGuesses := uint64(0)

	absAlphaDelta = new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	absPercentAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	absGuessDelta := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	half := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	half.SetString("0.5")

	guess = new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	guess.SetString("1.015")

	lastLowGuess := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	lastLowGuess.SetString("0.0")

	lastLowRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	lastLowRaisedToPower.SetString("0.0")

	lastLowAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	lastLowAbsAlphaDelta.SetString("0.0")

	lastHighGuess := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	lastHighGuess.SetString("1.5")

	lastHighRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	lastHighRaisedToPower.SetString("1.5")

	lastHighAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	lastHighAbsAlphaDelta.SetString("1.5")

	numberOne := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	numberOne.SetString("1.0")

	numberTwo := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	numberTwo.SetString("2.0")

	deltaIncrement := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	deltaIncrement.SetString("0.0")

	tempCompare := 0

	cycleCount = uint64(0)

	closeEnough := false

	for closeEnough == false {

		raisedToPower = BigFloatPower(
			guess,
			nthRoot,
			prec)

		compareResult = raisedToPower.Cmp(alpha)

		if compareResult == 0 {

			absAlphaDelta.Sub(alpha, raisedToPower)

			// On target
			return guess,
				raisedToPower,
				absAlphaDelta,
				compareResult,
				cycleCount

		} else if compareResult == -1 {
			// Guess is low. Alpha is higher than
			// raisedToPower

			lastLowGuess.Set(guess)

			lastLowRaisedToPower.Set(raisedToPower)

			absAlphaDelta.Sub(alpha, lastLowRaisedToPower)

			absPercentAlphaDelta.Quo(absAlphaDelta, alpha)

			//fmt.Printf("lastLowGuess  = %v\n"+
			//	"lastLowRaisedToPower = %v\n"+
			//	"       absAlphaDelta = %v\n"+
			//	"absPercentAlphaDelta = %v\n",
			//	lastLowGuess.Text('f', 20),
			//	lastLowRaisedToPower.Text('f', 20),
			//	absAlphaDelta.Text('f', 20),
			//	absPercentAlphaDelta.Text('f', 20))

			deltaIncrement.Set(half)

			if numOfLowGuesses > 2 {

				if absPercentAlphaDelta.Cmp(half) == 1 {

					deltaIncrement.Set(numberTwo)

				}

			}

			lastLowAbsAlphaDelta.Set(absAlphaDelta)

			tempCompare = lastLowGuess.Cmp(lastHighGuess)

			if tempCompare == 1 {
				// Last Low Guess is higher than
				// Last High Guess
				absGuessDelta.Mul(guess, deltaIncrement)

			} else if tempCompare == 0 {
				// Last Low Guess is Equal To
				// Last High Guess
				absGuessDelta.Mul(guess, deltaIncrement)

			} else {
				// Last High Guess is higher than
				//	Last Low Guess
				absGuessDelta.Sub(lastHighGuess, lastLowGuess)

				absGuessDelta.Mul(absGuessDelta, deltaIncrement)

			}

			//fmt.Printf("Old Guess = %v\n",
			//	guess.Text('f', -1))
			//
			//fmt.Printf("absGuessDelta = %v\n",
			//	absGuessDelta.Text('f', -1))

			guess.Add(guess, absGuessDelta)

			//fmt.Printf("New Guess = %v\n",
			//	guess.Text('f', -1))

			numOfLowGuesses++

		} else {
			// MUST BE -
			//	Guess is high. raisedToPower is higher
			//	than Alpha

			lastHighGuess.Set(guess)

			lastHighRaisedToPower.Set(raisedToPower)

			absAlphaDelta.Sub(lastHighRaisedToPower, alpha)

			absPercentAlphaDelta.Quo(absAlphaDelta, alpha)

			deltaIncrement.Set(half)

			if numOfHighGuesses > 2 {

				if absPercentAlphaDelta.Cmp(half) == 1 {

					deltaIncrement.Set(numberTwo)

				}

			}

			//fmt.Printf("lastHighGuess  = %v\n"+
			//	"lastHighRaisedToPower = %v\n",
			//	lastHighGuess.Text('f', 10),
			//	lastHighRaisedToPower.Text('f', 10))

			lastHighAbsAlphaDelta.Set(absAlphaDelta)

			tempCompare = lastLowGuess.Cmp(lastHighGuess)

			if tempCompare == 1 {
				// Low Guess is higher than High Guess
				absGuessDelta.Mul(lastHighGuess, half)
				guess.Sub(guess, absGuessDelta)

			} else if tempCompare == 0 {
				absGuessDelta.Mul(lastHighGuess, half)
				guess.Sub(guess, absGuessDelta)

			} else {
				// High Guess is higher than Low Guess
				absGuessDelta.Sub(lastHighGuess, lastLowGuess)
				absGuessDelta.Mul(absGuessDelta, half)
			}

			guess.Sub(guess, absGuessDelta)

			//fmt.Printf("New Guess = %v\n",
			//	guess.Text('f', -1))

			numOfHighGuesses++
		}

		cycleCount++

		if compareResult == -1 &&
			absAlphaDelta.Cmp(accuracyThreshold) < 1 {

			guess.Set(lastLowGuess)
			raisedToPower.Set(lastLowRaisedToPower)

			closeEnough = true
		}

		if cycleCount > 10000 {

			if compareResult == 1 {
				guess.Set(lastLowGuess)
				raisedToPower.Set(lastLowRaisedToPower)
				absAlphaDelta.Sub(alpha, raisedToPower)

			}

			closeEnough = true

		}

	}

	compareResult = raisedToPower.Cmp(alpha)

	absAlphaDelta.Sub(raisedToPower, alpha)

	if absAlphaDelta.Sign() == -1 {
		absAlphaDelta.Neg(absAlphaDelta)
	}

	return guess,
		raisedToPower,
		absAlphaDelta,
		compareResult,
		cycleCount
}

func BigFloatPower(
	base *big.Float,
	power int64,
	prec uint) (
	rasedToPower *big.Float) {

	rasedToPower =
		new(big.Float).
			SetPrec(prec).
			SetMode(big.AwayFromZero)

	var ok bool
	_,
		ok = rasedToPower.SetString("1.0")

	if !ok {
		panic("rasedToPower.SetString(\"1.0\") Failed!")
	}

	for i := int64(0); i < power; i++ {

		rasedToPower.Mul(rasedToPower, base)
	}

	return rasedToPower
}
