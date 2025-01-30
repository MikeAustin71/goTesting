package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

type BigFloatCalcStats struct {
	BaseFloatNum big.Float

	BasePrec uint

	BaseNumIntDigits int64

	BaseNumFracDigits int64

	CalcResult big.Float

	CalcResultPrec uint

	CalcResultNumIntDigits int64

	CalcResultNumFracDigits int64
}

type PureNumberStrComponents struct {
	NumberSign int
	// -1 == negative
	//  0 == zero
	//  1 == positive

	NumberType int
	//	0 == invalid
	//	1 == integer number
	//	2 == floating point number

	NumIntegerDigits int64

	NumFractionalDigits int64

	AbsoluteValueNumStr string

	AllIntegerDigitsNumStr string
}

func main() {

	/*
		GolangSqrtExample()
		fmt.Println("Verification")
		GolangSqrtExampleVerify()

	*/

	// TestLog10InitialGuess()

	// Newton01()

	// TestNewtonInitialGuess()

	// TestBigFloatPower()

	// TestModulo()

	TestNewtonNthRoot01()

	// TestNewtonNthRoot02()
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

	var baseStr string
	var expectedResultStr string

	// ****************************
	// Set Parameters
	// ****************************
	const prec = 4096
	power := int64(2)
	baseStr = "0.2"
	expectedResultStr = "0.04"
	// ****************************

	var ok bool

	base := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)
	_,
		ok = base.SetString(baseStr)

	if !ok {
		panic("base.SetString(\"1.5\") FAILED!\n")
	}

	expectedResult := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	_,
		ok = expectedResult.SetString(expectedResultStr)

	if !ok {
		panic(fmt.Sprintf("expectedResult.SetString(%v) FAILED!\n",
			expectedResultStr))
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

	alphaValStr := ""
	accuracyDeltaStr := ""
	expectedResultStr := ""
	// prec := uint(16384)
	// *************************
	// ** Set Parameters **
	localPrec := 20
	prec := uint(4096)
	nthRoot := int64(3)
	alphaValStr = "0.00000000000000002879"
	accuracyDeltaStr = "0.0000000000000000000000000000000000005"
	expectedResultStr = "0.0000030648829156192458114498573382988"
	// *************************

	accuracyThresholdDelta := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	alpha := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	accuracyThresholdDelta.SetString(accuracyDeltaStr)

	var ok bool

	_,
		ok = alpha.SetString(alphaValStr)

	if !ok {
		fmt.Printf("\n%v\n"+
			"ERROR: alpha.SetString(alphaValStr) Failed!\n",
			funcName)

		return
	}

	var calculationPrec uint

	guess,
		raisedToPower,
		absAlphaDelta,
		compareResult,
		initializeCycleCount,
		calculationCycleCount,
		calculationPrec := NewtonInitialGuess(
		nthRoot,
		alpha,
		accuracyThresholdDelta,
		prec)

	fmt.Printf("Initial Guess Results\n"+
		"                Precision = %v\n"+
		"    Calculation Precision = %v\n"+
		"            compareResult = %v\n"+
		"     initializeCycleCount = %v\n"+
		"    calculationCycleCount = %v\n"+
		"                  nthRoot = %v\n"+
		"                    alpha = %v\n"+
		"  Expected Accuracy Delta = %v\n"+
		"    Actual Accuracy Delta = %v\n"+
		"  Summary Raised To Power = %v\n"+
		"  Summary guess           = %v\n"+
		"  Actual Result           = %v\n"+
		"raisedToPower = \n%v\n\n"+
		"absAlphaDelta = \n%v\n\n"+
		"        guess =\n%v\n\n",
		prec,
		calculationPrec,
		compareResult,
		initializeCycleCount,
		calculationCycleCount,
		nthRoot,
		alpha.Text('f', -1),
		accuracyThresholdDelta.Text('f', localPrec),
		absAlphaDelta.Text('f', localPrec),
		raisedToPower.Text('f', localPrec),
		guess.Text('f', localPrec),
		expectedResultStr,
		raisedToPower.Text('f', -1),
		absAlphaDelta.Text('f', -1),
		guess.Text('f', -1))

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("  Successful Completion!\n" +
		"Function: " +
		funcName + "\n")

	fmt.Printf(breakStr + "\n")
}

func TestNewtonNthRoot01() {

	funcName := "TestNewtonNthRoot01()"
	var initialGuessAccuracyThresholdStr,
		alphaStr string

	breakStr := strings.Repeat("=", 70)

	var localPrecision int

	//prec := uint(16384)
	// prec := uint(4096)

	// Input Parameters
	// *********************************
	n := int64(22)
	alphaStr = "58324.8123"
	initialGuessAccuracyThresholdStr = "0.00005"
	localPrecision = 50
	numOfCalcSteps := 120
	// Number of Calc Steps over and above
	//	the minimum required Steps
	expectedNthRootStr := "1.6467576806739737181931788067863"
	requiredDigitsAlphaAccuracy := int64(1000)
	// *********************************

	prec,
		err := calculateRequiredPrecision(
		0,
		0,
		2000)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"Error:\n%v\n\n",
			funcName,
			err.Error())

		return
	}

	numOfCalcSteps =
		numOfCalcSteps + int(math.Log2(float64(prec)))

	idx := strings.Index(expectedNthRootStr, ".")
	var expectedNthRootFracDigits int

	if idx == -1 {
		expectedNthRootFracDigits = 0
	} else {
		expectedNthRootFracDigits =
			len(expectedNthRootStr) - (idx + 1)
	}

	var ok bool

	expectedNthRoot := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	_,
		ok = expectedNthRoot.SetString(expectedNthRootStr)

	if !ok {
		fmt.Printf("%v\n"+
			"expectedNthRoot.SetString(expectedNthRootStr) FAILED!\n"+
			"expectedNthRootStr= %v\n",
			funcName,
			expectedNthRootStr)
		return
	}

	nthRootVariance := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	alpha := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	_,
		ok = alpha.SetString(alphaStr)

	if !ok {
		fmt.Printf("%v\n"+
			"alpha.SetString(alphaStr) FAILED!\n"+
			"alphaStr= %v\n",
			funcName,
			alphaStr)
		return
	}

	// The first guess!
	accuracyThreshold := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	accuracyThreshold.SetString(initialGuessAccuracyThresholdStr)

	var calculationPrecision uint

	fmt.Printf("\n" + breakStr + "\n")

	initialGuess,
		raisedToPower,
		absAlphaDelta,
		compareResult,
		initializeCycleCount,
		calculationCycleCount,
		calculationPrecision := NewtonInitialGuess(
		n,
		alpha,
		accuracyThreshold,
		prec)

	nthRootVariance.Sub(initialGuess, expectedNthRoot)

	var compareResultStr string

	if compareResult == -1 {

		compareResultStr = "Initial Guess Lower Than Actual Nth Root"

	} else if compareResult == 1 {

		compareResultStr = "Initial Guess Greater Than Actual Nth Root"

	} else {

		compareResultStr = "Initial Guess Equal To Actual Nth Root"

	}

	fmt.Printf(breakStr+"\n"+
		"%v\n"+
		"First Guess Data\n"+
		"              Initial Guess = %v\n"+
		"          Expected Nth Root = %v\n"+
		"         Nth Root Variance  = %v\n"+
		"                  precision = %v\n"+
		"      calculation precision = %v\n"+
		"                          n = %v\n"+
		"                      alpha = %v\n"+
		"Initial Guess raisedToPower = %v\n"+
		"              absAlphaDelta = %v\n"+
		"   Requested Accuracy Delta = %v\n"+
		"              compareResult = %v\n"+
		"      Compare Result Status = %v\n"+
		" Initialization Cycle Count = %v\n"+
		"    Calculation Cycle Count = %v\n\n",
		funcName,
		initialGuess.Text('f', localPrecision),
		expectedNthRootStr,
		nthRootVariance.Text('f', localPrecision),
		prec,
		calculationPrecision,
		n,
		alpha.Text('f', localPrecision),
		raisedToPower.Text('f', localPrecision),
		absAlphaDelta.Text('f', localPrecision),
		accuracyThreshold.Text('f', localPrecision),
		compareResult,
		compareResultStr,
		initializeCycleCount,
		calculationCycleCount)

	nthRoot,
		actualNumCalcSteps,
		actualPrecisionBits := Newton(
		alpha,
		n,
		initialGuess,
		calculationPrecision,
		numOfCalcSteps,
		requiredDigitsAlphaAccuracy)

	calculationPrecision = actualPrecisionBits

	//calcAlpha := BigFloatPower(
	//	nthRoot,
	//	n,
	//	calculationPrecision)

	calcAlpha := big.NewFloat(0)
	var resultsStats BigFloatCalcStats

	resultsStats,
		err = raiseToPowerInt(
		nthRoot,
		n)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"Error:\n%v\n\n",
			funcName,
			err.Error())

		return
	}

	calcAlpha.Copy(&resultsStats.CalcResult)

	// Results Summary
	// *********************************
	fmt.Printf(breakStr+"\n"+
		"%v\n"+
		"****** Summary Results ******\n"+
		"              Precision = %v\n"+
		"   Calculated Precision = %v\n"+
		"Actual Newton Precision = %v\n"+
		"        Local Precision = %v\n"+
		"                      n = %v\n"+
		"                  alpha = %v\n"+
		"       Calculated alpha = %v\n"+
		"          Initial Guess = %v\n"+
		"   Requested Calc Steps = %v\n"+
		"      Actual Calc Steps = %v\n"+
		breakStr+"\n"+
		"Summary Actual nth Root = %v\n"+
		"Actual nth Root Rounded = %v\n"+
		"      Expected nth Root = %v\n"+
		"          Initial Guess = %v\n"+
		breakStr+"\n",
		funcName,
		prec,
		calculationPrecision,
		actualPrecisionBits,
		localPrecision,
		n,
		alpha.Text('f', localPrecision),
		calcAlpha.Text('f', localPrecision),
		initialGuess.Text('f', localPrecision),
		numOfCalcSteps,
		actualNumCalcSteps,
		nthRoot.Text('f', localPrecision),
		nthRoot.Text('f', expectedNthRootFracDigits),
		expectedNthRootStr,
		initialGuess.Text('f', localPrecision))

	alphaStr = alpha.Text('f', int(calculationPrecision))

	numAlphaStrTotalDigits := len(alphaStr)
	numAlphaStrIntDigits := 0
	numAlphaStrDecDigits := 0
	foundDecimalPt := false

	calcAlphaStr := calcAlpha.Text('f', int(calculationPrecision))

	numCalcAlphaDigits := len(calcAlphaStr)
	numOfMatchingCalcAlphaDigits := 0
	numOfMatchingCalcAlphaDecimalDigits := 0
	foundNonMatchingCalcAlphaDigit := false

	for i := 0; i < numAlphaStrTotalDigits; i++ {

		if alphaStr[i] == '.' {

			foundDecimalPt = true

			continue
		}

		if alphaStr[i] >= '0' &&
			alphaStr[i] <= '9' {

			if foundDecimalPt == false {

				numAlphaStrIntDigits++

			} else {

				numAlphaStrDecDigits++
			}

			if i < numCalcAlphaDigits {

				if alphaStr[i] == calcAlphaStr[i] {

					if foundNonMatchingCalcAlphaDigit == false {

						numOfMatchingCalcAlphaDigits++
					}

				} else {

					foundNonMatchingCalcAlphaDigit = true
				}

			}

		}
	}

	if numOfMatchingCalcAlphaDigits > numAlphaStrIntDigits {

		numOfMatchingCalcAlphaDecimalDigits =
			numOfMatchingCalcAlphaDigits - numAlphaStrIntDigits

	}

	fmt.Printf("Alpha And Calculated Alpha Comparison\n"+
		breakStr+"\n"+
		"                Total Number Of Alpha Digits = %v\n"+
		"  Total Number of Matching Calc Alpha Digits = %v\n"+
		"           Required Alpha Digits of Accuracy = %v\n"+
		"Number of Matching Calc Alpha Decimal Digits = %v\n"+
		breakStr+"\n\n",
		numAlphaStrTotalDigits,
		numOfMatchingCalcAlphaDigits,
		requiredDigitsAlphaAccuracy,
		numOfMatchingCalcAlphaDecimalDigits)

	actualNthRootStr :=
		nthRoot.Text('f', int(calculationPrecision))

	expectedNthRootStr =
		expectedNthRoot.Text('f', int(calculationPrecision))

	// Results Detail
	// *********************************
	fmt.Printf(breakStr+"\n"+
		"%v\n"+
		"****** Detail Results ******\n"+
		"Calculated Precision = %v\n"+
		"  Actual Input Alpha = \n%v\n\n"+
		"      Verified Alpha = \n%v\n\n"+
		" Calculated nth Root =  \n%v\n\n"+
		"   Expected nth Root =  \n%v\n\n",
		funcName,
		calculationPrecision,
		alphaStr,
		calcAlphaStr,
		actualNthRootStr,
		expectedNthRootStr)

	return
}

func TestNewtonNthRoot02() {

	funcName := "TestNewtonNthRoot01()"
	var initialGuessAccuracyThresholdStr,
		alphaStr string

	breakStr := strings.Repeat("=", 70)

	var localPrecision int

	//prec := uint(16384)
	// prec := uint(4096)

	// Input Parameters
	// *********************************
	n := int64(22)
	alphaStr = "58324.8123"
	initialGuessAccuracyThresholdStr = "0.00005"
	localPrecision = 50
	numOfCalcSteps := 120
	// Number of Calc Steps over and above
	//	the minimum required Steps
	expectedNthRootStr := "1.6467576806739737181931788067863"
	requiredDigitsAlphaAccuracy := int64(2000)
	// *********************************

	prec := setPrecisionFromDigits02(requiredDigitsAlphaAccuracy)

	numOfCalcSteps =
		numOfCalcSteps + int(math.Log2(float64(prec)))

	idx := strings.Index(expectedNthRootStr, ".")
	var expectedNthRootFracDigits int

	if idx == -1 {
		expectedNthRootFracDigits = 0
	} else {
		expectedNthRootFracDigits =
			len(expectedNthRootStr) - (idx + 1)
	}

	var ok bool

	expectedNthRoot := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	_,
		ok = expectedNthRoot.SetString(expectedNthRootStr)

	if !ok {
		fmt.Printf("%v\n"+
			"expectedNthRoot.SetString(expectedNthRootStr) FAILED!\n"+
			"expectedNthRootStr= %v\n",
			funcName,
			expectedNthRootStr)
		return
	}

	nthRootVariance := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	alpha := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	_,
		ok = alpha.SetString(alphaStr)

	if !ok {
		fmt.Printf("%v\n"+
			"alpha.SetString(alphaStr) FAILED!\n"+
			"alphaStr= %v\n",
			funcName,
			alphaStr)
		return
	}

	// The first guess!
	accuracyThreshold := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	accuracyThreshold.SetString(initialGuessAccuracyThresholdStr)

	var calculationPrecision uint

	fmt.Printf("\n" + breakStr + "\n")

	initialGuess,
		raisedToPower,
		absAlphaDelta,
		compareResult,
		initializeCycleCount,
		calculationCycleCount,
		calculationPrecision := NewtonInitialGuess(
		n,
		alpha,
		accuracyThreshold,
		prec)

	nthRootVariance.Sub(initialGuess, expectedNthRoot)

	var compareResultStr string

	if compareResult == -1 {

		compareResultStr = "Initial Guess Lower Than Actual Nth Root"

	} else if compareResult == 1 {

		compareResultStr = "Initial Guess Greater Than Actual Nth Root"

	} else {

		compareResultStr = "Initial Guess Equal To Actual Nth Root"

	}

	fmt.Printf(breakStr+"\n"+
		"%v\n"+
		"First Guess Data\n"+
		"              Initial Guess = %v\n"+
		"          Expected Nth Root = %v\n"+
		"         Nth Root Variance  = %v\n"+
		"                  precision = %v\n"+
		"      calculation precision = %v\n"+
		"                          n = %v\n"+
		"                      alpha = %v\n"+
		"Initial Guess raisedToPower = %v\n"+
		"              absAlphaDelta = %v\n"+
		"   Requested Accuracy Delta = %v\n"+
		"              compareResult = %v\n"+
		"      Compare Result Status = %v\n"+
		" Initialization Cycle Count = %v\n"+
		"    Calculation Cycle Count = %v\n\n",
		funcName,
		initialGuess.Text('f', localPrecision),
		expectedNthRootStr,
		nthRootVariance.Text('f', localPrecision),
		prec,
		calculationPrecision,
		n,
		alpha.Text('f', localPrecision),
		raisedToPower.Text('f', localPrecision),
		absAlphaDelta.Text('f', localPrecision),
		accuracyThreshold.Text('f', localPrecision),
		compareResult,
		compareResultStr,
		initializeCycleCount,
		calculationCycleCount)

	nthRoot,
		actualNumCalcSteps,
		actualPrecisionBits := Newton(
		alpha,
		n,
		initialGuess,
		calculationPrecision,
		numOfCalcSteps,
		requiredDigitsAlphaAccuracy)

	calculationPrecision = actualPrecisionBits

	calcAlpha := BigFloatPower(
		nthRoot,
		n,
		calculationPrecision)

	// Results Summary
	// *********************************
	fmt.Printf(breakStr+"\n"+
		"%v\n"+
		"****** Summary Results ******\n"+
		"              Precision = %v\n"+
		"   Calculated Precision = %v\n"+
		"Actual Newton Precision = %v\n"+
		"        Local Precision = %v\n"+
		"                      n = %v\n"+
		"                  alpha = %v\n"+
		"       Calculated alpha = %v\n"+
		"          Initial Guess = %v\n"+
		"   Requested Calc Steps = %v\n"+
		"      Actual Calc Steps = %v\n"+
		breakStr+"\n"+
		"Summary Actual nth Root = %v\n"+
		"Actual nth Root Rounded = %v\n"+
		"      Expected nth Root = %v\n"+
		"          Initial Guess = %v\n"+
		breakStr+"\n",
		funcName,
		prec,
		calculationPrecision,
		actualPrecisionBits,
		localPrecision,
		n,
		alpha.Text('f', localPrecision),
		calcAlpha.Text('f', localPrecision),
		initialGuess.Text('f', localPrecision),
		numOfCalcSteps,
		actualNumCalcSteps,
		nthRoot.Text('f', localPrecision),
		nthRoot.Text('f', expectedNthRootFracDigits),
		expectedNthRootStr,
		initialGuess.Text('f', localPrecision))

	alphaStr = alpha.Text('f', int(calculationPrecision))

	numAlphaStrTotalDigits := len(alphaStr)
	numAlphaStrIntDigits := 0
	numAlphaStrDecDigits := 0
	foundDecimalPt := false

	calcAlphaStr := calcAlpha.Text('f', int(calculationPrecision))

	numCalcAlphaDigits := len(calcAlphaStr)
	numOfMatchingCalcAlphaDigits := 0
	numOfMatchingCalcAlphaDecimalDigits := 0
	foundNonMatchingCalcAlphaDigit := false

	for i := 0; i < numAlphaStrTotalDigits; i++ {

		if alphaStr[i] == '.' {

			foundDecimalPt = true

			continue
		}

		if alphaStr[i] >= '0' &&
			alphaStr[i] <= '9' {

			if foundDecimalPt == false {

				numAlphaStrIntDigits++

			} else {

				numAlphaStrDecDigits++
			}

			if i < numCalcAlphaDigits {

				if alphaStr[i] == calcAlphaStr[i] {

					if foundNonMatchingCalcAlphaDigit == false {

						numOfMatchingCalcAlphaDigits++
					}

				} else {

					foundNonMatchingCalcAlphaDigit = true
				}

			}

		}
	}

	if numOfMatchingCalcAlphaDigits > numAlphaStrIntDigits {

		numOfMatchingCalcAlphaDecimalDigits =
			numOfMatchingCalcAlphaDigits - numAlphaStrIntDigits

	}

	fmt.Printf("Alpha And Calculated Alpha Comparison\n"+
		breakStr+"\n"+
		"                Total Number Of Alpha Digits = %v\n"+
		"  Total Number of Matching Calc Alpha Digits = %v\n"+
		"           Required Alpha Digits of Accuracy = %v\n"+
		"Number of Matching Calc Alpha Decimal Digits = %v\n"+
		breakStr+"\n\n",
		numAlphaStrTotalDigits,
		numOfMatchingCalcAlphaDigits,
		requiredDigitsAlphaAccuracy,
		numOfMatchingCalcAlphaDecimalDigits)

	actualNthRootStr :=
		nthRoot.Text('f', int(calculationPrecision))

	expectedNthRootStr =
		expectedNthRoot.Text('f', int(calculationPrecision))

	// Results Detail
	// *********************************
	fmt.Printf(breakStr+"\n"+
		"%v\n"+
		"****** Detail Results ******\n"+
		"Calculated Precision = %v\n"+
		"  Actual Input Alpha = \n%v\n\n"+
		"      Verified Alpha = \n%v\n\n"+
		" Calculated nth Root =  \n%v\n\n"+
		"   Expected nth Root =  \n%v\n\n",
		funcName,
		calculationPrecision,
		alphaStr,
		calcAlphaStr,
		actualNthRootStr,
		expectedNthRootStr)

	return
}

// TestNewtonNthRoot03
//
// This test showed that the Newton algorithm will
// generate accurate results when the initial guess
// is high. Accurate results still depend on an
// initial guess which is close, probably something
// on the order of + or - 0.5.
func TestNewtonNthRoot03() {

	funcName := "TestNewtonNthRoot02()"
	var alphaStr string

	breakStr := strings.Repeat("-", 40)

	var localPrecision int

	// Input Parameters
	// *********************************
	n := int64(92)
	alphaStr = "47.8"
	initialGuessStr := "1.04295"
	prec := uint(4096)
	localPrecision = 50
	numOfCalcSteps := 120
	expectedNthRootStr := "1.0429287773878709622074291041279"
	// *********************************

	var ok bool

	expectedNthRoot := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	_,
		ok = expectedNthRoot.SetString(expectedNthRootStr)

	if !ok {
		fmt.Printf("%v\n"+
			"expectedNthRoot.SetString(expectedNthRootStr) FAILED!\n"+
			"expectedNthRootStr= %v\n",
			funcName,
			expectedNthRootStr)
		return
	}

	nthRootVariance := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	alpha := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	_,
		ok = alpha.SetString(alphaStr)

	if !ok {
		fmt.Printf("%v\n"+
			"alpha.SetString(alphaStr) FAILED!\n"+
			"alphaStr= %v\n",
			funcName,
			alphaStr)
		return
	}

	// The first guess!

	initialGuess := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	_,
		ok = initialGuess.SetString(initialGuessStr)

	if !ok {
		fmt.Printf("%v\n"+
			"initialGuess.SetString(initialGuessStr) FAILED!\n"+
			"initialGuessStr= %v\n",
			funcName,
			alphaStr)
		return

	}

	nthRootVariance.Sub(initialGuess, expectedNthRoot)

	compareResult := initialGuess.Cmp(expectedNthRoot)

	var compareResultStr string

	if compareResult == -1 {

		compareResultStr = "Initial Guess Lower Than Actual Nth Root"

	} else if compareResult == 1 {

		compareResultStr = "Initial Guess Greater Than Actual Nth Root"

	} else {

		compareResultStr = "Initial Guess Equal To Actual Nth Root"

	}

	nthRoot,
		actualNumCalcSteps := Newton04(
		alpha,
		n,
		initialGuess,
		prec,
		numOfCalcSteps)

	calcAlpha := BigFloatPower(
		nthRoot,
		n,
		prec)

	// Results Summary
	// *********************************
	fmt.Printf(breakStr+"\n"+
		"%v\n"+
		"****** Summary Results ******\n"+
		"            Precision = %v\n"+
		"      Local Precision = %v\n"+
		"                    n = %v\n"+
		"                alpha = %v\n"+
		"     Calculated alpha = %v\n"+
		"        Initial Guess = %v\n"+
		"Initial Guess Compare = %v\n"+
		" Requested Calc Steps = %v\n"+
		"   Actual Calc Steps = %v\n"+
		breakStr+"\n"+
		"Summary Actual nth Root = %v\n"+
		"          Initial Guess = %v\n"+
		"      Expected nth Root = %v\n"+
		breakStr+"\n",
		funcName,
		prec,
		localPrecision,
		n,
		alpha.Text('f', localPrecision),
		calcAlpha.Text('f', localPrecision),
		initialGuess.Text('f', localPrecision),
		compareResultStr,
		numOfCalcSteps,
		actualNumCalcSteps,
		nthRoot.Text('f', localPrecision),
		initialGuess.Text('f', localPrecision),
		expectedNthRootStr)

	alphaStr = alpha.Text('f', int(prec))

	numAlphaStrTotalDigits := len(alphaStr)
	numAlphaStrIntDigits := 0
	numAlphaStrDecDigits := 0
	foundDecimalPt := false

	calcAlphaStr := calcAlpha.Text('f', int(prec))

	numCalcAlphaDigits := len(calcAlphaStr)
	numOfMatchingCalcAlphaDigits := 0
	numOfMatchingCalcAlphaDecimalDigits := 0
	foundNonMatchingCalcAlphaDigit := false

	for i := 0; i < numAlphaStrTotalDigits; i++ {

		if alphaStr[i] == '.' {

			foundDecimalPt = true

			continue
		}

		if alphaStr[i] >= '0' &&
			alphaStr[i] <= '9' {

			if foundDecimalPt == false {

				numAlphaStrIntDigits++

			} else {

				numAlphaStrDecDigits++
			}

			if i < numCalcAlphaDigits {

				if alphaStr[i] == calcAlphaStr[i] {

					if foundNonMatchingCalcAlphaDigit == false {

						numOfMatchingCalcAlphaDigits++
					}

				} else {

					foundNonMatchingCalcAlphaDigit = true
				}

			}

		}
	}

	if numOfMatchingCalcAlphaDigits > numAlphaStrIntDigits {

		numOfMatchingCalcAlphaDecimalDigits =
			numOfMatchingCalcAlphaDigits - numAlphaStrIntDigits

	}

	fmt.Printf("Alpha And Calculated Alpha Comparison\n"+
		breakStr+"\n"+
		"                Total Number Of Alpha Digits = %v\n"+
		"  Total Number of Matching Calc Alpha Digits = %v\n"+
		"Number of Matching Calc Alpha Decimal Digits = %v\n"+
		breakStr+"\n\n",
		numAlphaStrTotalDigits,
		numOfMatchingCalcAlphaDigits,
		numOfMatchingCalcAlphaDecimalDigits)

	actualNthRootStr :=
		nthRoot.Text('f', int(prec))

	expectedNthRootStr =
		expectedNthRoot.Text('f', int(prec))

	// Results Detail
	// *********************************
	fmt.Printf(breakStr+"\n"+
		"%v\n"+
		"****** Detail Results ******\n"+
		"           Precision = %v\n"+
		"  Actual Input Alpha = \n%v\n\n"+
		"      Verified Alpha = \n%v\n\n"+
		" Calculated nth Root =  \n%v\n\n"+
		"   Expected nth Root =  \n%v\n\n",
		funcName,
		prec,
		alphaStr,
		calcAlphaStr,
		actualNthRootStr,
		expectedNthRootStr)

	return
}

func TestModulo() {

	var testSteps int

	for i := 0; i < 100; i++ {

		testSteps = i % 5

		fmt.Printf("i= %v   testSteps = %v\n",
			i,
			testSteps)

	}

}

// Newton
//
// alpha = radicand
//
// In mathematics, an nth root of a number 'alpha'
// is a number r which, when raised to the power n,
// yields alpha:
//
//	r^n = alpha
//
// https://en.wikipedia.org/wiki/Nth_root
//
// This approach taken from :
//
//	https://pkg.go.dev/math/big
//		Example (Sqrt2)
func Newton(
	alpha *big.Float,
	n_int64 int64,
	initialGuess *big.Float,
	prec uint,
	numOfCalcSteps int,
	requiredDigitsAlphaAccuracy int64) (
	nthRoot *big.Float,
	actualNumCalcSteps int,
	actualPrecisionBits uint) {

	funcName := "Newton()"

	// Since Newton's Method doubles the number of correct digits at each
	// iteration, we need at least log_2(prec) steps.
	//steps := int(math.Log2(float64(prec)))
	//steps += 100

	//estimatedMinPrecisionUint :=
	//	setPrecisionFromDigits02(requiredDigitsAlphaAccuracy)

	estimatedMinPrecisionUint,
		err := calculateRequiredPrecision(
		0,
		0,
		requiredDigitsAlphaAccuracy)

	if err != nil {
		fmt.Printf("\n\n%v\n"+
			"Error:\n%v\n\n",
			funcName,
			err.Error())

		return
	}

	if prec < estimatedMinPrecisionUint {
		prec = estimatedMinPrecisionUint
	}

	outputDigitsOfAccuracy :=
		int(requiredDigitsAlphaAccuracy + 100)

	alphaStr := alpha.Text('f', outputDigitsOfAccuracy)

	lenAlphaStr := len(alphaStr)

	n_minus_1_int64 := n_int64 - 1

	n_minus_1 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(n_minus_1_int64)

	one_over_n_rat := big.NewRat(1, n_int64)

	one_over_n_float := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetRat(one_over_n_rat)

	x_k1 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	fac_1 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	fac_2 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	fac_3 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	fac_4 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	x_k := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	x_k.Set(initialGuess)

	actualNumCalcSteps = numOfCalcSteps

	numberOfAccurateDigits := int64(0)

	calcAlpha := big.NewFloat(0)

	var calcAlphaStr string

	var resultsStats BigFloatCalcStats

	for i := 0; i <= numOfCalcSteps; i++ {

		//fac_1 = BigFloatPower(x_k, n_minus_1_int64, prec)

		resultsStats,
			err = raiseToPowerInt(
			x_k,
			n_minus_1_int64)

		if err != nil {

			fmt.Printf("\n\n%v\n"+
				"Error:\n%v\n\n",
				funcName,
				err.Error())

			return
		}

		fac_1.Copy(
			&resultsStats.CalcResult)

		fac_2.Quo(alpha, fac_1)

		fac_3.Mul(n_minus_1, x_k)

		fac_4.Add(fac_3, fac_2)

		x_k1.Mul(one_over_n_float, fac_4)

		x_k.Set(x_k1)

		testSteps := i % 15

		if i > 0 && testSteps == 0 {

			numberOfAccurateDigits = 0

			//calcAlpha = BigFloatPower(
			//	x_k,
			//	n_int64,
			//	prec)

			resultsStats,
				err = raiseToPowerInt(
				x_k,
				n_int64)

			if err != nil {

				fmt.Printf("\n\n%v\n"+
					"Error:\n%v\n\n",
					funcName,
					err.Error())

				return
			}

			calcAlpha.Copy(
				&resultsStats.CalcResult)

			calcAlphaStr = calcAlpha.Text('f', outputDigitsOfAccuracy)

			for i := 0; i < lenAlphaStr; i++ {

				if alphaStr[i] == calcAlphaStr[i] {
					numberOfAccurateDigits++
				} else {
					break
				}
			}

			if numberOfAccurateDigits >= requiredDigitsAlphaAccuracy {

				actualNumCalcSteps = i + 1

				break
			}

		}

	}

	return x_k, actualNumCalcSteps, prec
}

// Newton05
//
// alpha = radicand
//
// In mathematics, an nth root of a number 'alpha'
// is a number r which, when raised to the power n,
// yields alpha:
//
//	r^n = alpha
//
// https://en.wikipedia.org/wiki/Nth_root
//
// This approach taken from :
//
//	https://pkg.go.dev/math/big
//		Example (Sqrt2)
func Newton05(
	alpha *big.Float,
	n_int64 int64,
	initialGuess *big.Float,
	prec uint,
	numOfCalcSteps int,
	requiredDigitsAlphaAccuracy int64) (
	nthRoot *big.Float,
	actualNumCalcSteps int,
	actualPrecisionBits uint) {

	// Since Newton's Method doubles the number of correct digits at each
	// iteration, we need at least log_2(prec) steps.
	//steps := int(math.Log2(float64(prec)))
	//steps += 100

	estimatedMinPrecisionUint :=
		setPrecisionFromDigits02(requiredDigitsAlphaAccuracy)

	if prec < estimatedMinPrecisionUint {
		prec = estimatedMinPrecisionUint
	}

	outputDigitsOfAccuracy :=
		int(requiredDigitsAlphaAccuracy + 100)

	alphaStr := alpha.Text('f', outputDigitsOfAccuracy)

	lenAlphaStr := len(alphaStr)

	n_minus_1_int64 := n_int64 - 1

	n_minus_1 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(n_minus_1_int64)

	one_over_n_rat := big.NewRat(1, n_int64)

	one_over_n_float := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetRat(one_over_n_rat)

	x_k1 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	fac_1 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	fac_2 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	fac_3 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	fac_4 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	x_k := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	x_k.Set(initialGuess)

	actualNumCalcSteps = numOfCalcSteps

	numberOfAccurateDigits := int64(0)

	var calcAlpha *big.Float

	var calcAlphaStr string

	for i := 0; i <= numOfCalcSteps; i++ {

		fac_1 = BigFloatPower(x_k, n_minus_1_int64, prec)

		fac_2.Quo(alpha, fac_1)

		fac_3.Mul(n_minus_1, x_k)

		fac_4.Add(fac_3, fac_2)

		x_k1.Mul(one_over_n_float, fac_4)

		x_k.Set(x_k1)

		testSteps := i % 15

		if i > 0 && testSteps == 0 {

			numberOfAccurateDigits = 0

			calcAlpha = BigFloatPower(
				x_k,
				n_int64,
				prec)

			calcAlphaStr = calcAlpha.Text('f', outputDigitsOfAccuracy)

			for i := 0; i < lenAlphaStr; i++ {

				if alphaStr[i] == calcAlphaStr[i] {
					numberOfAccurateDigits++
				} else {
					break
				}
			}

			if numberOfAccurateDigits >= requiredDigitsAlphaAccuracy {

				actualNumCalcSteps = i + 1

				break
			}

		}

	}

	return x_k, actualNumCalcSteps, prec
}

// Newton04
//
// alpha = radicand
//
// In mathematics, an nth root of a number 'alpha'
// is a number r which, when raised to the power n,
// yields alpha:
//
//	r^n = alpha
//
// https://en.wikipedia.org/wiki/Nth_root
//
// This approach taken from :
//
//	https://pkg.go.dev/math/big
//		Example (Sqrt2)
func Newton04(
	alpha *big.Float,
	n_int64 int64,
	initialGuess *big.Float,
	prec uint,
	numOfCalcSteps int) (
	nthRoot *big.Float,
	actualNumCalcSteps int) {

	// Since Newton's Method doubles the number of correct digits at each
	// iteration, we need at least log_2(prec) steps.
	//steps := int(math.Log2(float64(prec)))
	//steps += 100

	n_minus_1_int64 := n_int64 - 1

	n_minus_1 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(n_minus_1_int64)

	one_over_n_rat := big.NewRat(1, n_int64)

	one_over_n_float := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetRat(one_over_n_rat)

	x_k1 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	fac_1 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	fac_2 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	fac_3 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	fac_4 := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	x_k := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	x_k.Set(initialGuess)

	actualNumCalcSteps = numOfCalcSteps

	factor3_3 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	factor3_3.SetFloat64(3.3219789132197891321978913219789)

	estimatedNumOfDigits := new(big.Float).
		SetPrec(prec).
		SetMode(big.AwayFromZero).
		SetInt64(int64(prec))

	estimatedNumOfDigits.Quo(estimatedNumOfDigits, factor3_3)

	tI64,
		_ := estimatedNumOfDigits.Int64()

	estimatedNumOfDigitsInt := int(tI64 + 8)

	var last_x_k_str, new_x_k_str string

	last_x_k_str = x_k.Text('f', estimatedNumOfDigitsInt)

	for i := 0; i <= numOfCalcSteps; i++ {

		fac_1 = BigFloatPower(x_k, n_minus_1_int64, prec)

		fac_2.Quo(alpha, fac_1)

		fac_3.Mul(n_minus_1, x_k)

		fac_4.Add(fac_3, fac_2)

		x_k1.Mul(one_over_n_float, fac_4)

		x_k.Set(x_k1)

		new_x_k_str =
			x_k.Text('f', estimatedNumOfDigitsInt)

		if new_x_k_str == last_x_k_str {
			actualNumCalcSteps = i + 1
			break
		}

		last_x_k_str = new_x_k_str
	}

	return x_k, actualNumCalcSteps
}

func Newton03() {

	funcName := "Newton03"

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

	alpha.SetString("64")

	// n = 3
	n_int64 := int64(3)
	// n_float := new(big.Float).SetPrec(prec).SetInt64(n_int64)

	// 2^3 = 8

	// The first guess!
	//x_k := new(big.Float).SetPrec(prec).SetFloat64(1.0)

	accuracyThreshold := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	accuracyThreshold.SetString("0.000015")

	var calculationPrecision uint

	x_k,
		raisedToPower,
		absAlphaDelta,
		compareResult,
		initializeCycleCount,
		calculationCycleCount,
		calculationPrecision := NewtonInitialGuess(
		n_int64,
		alpha,
		accuracyThreshold,
		prec)

	fmt.Printf("First Guess Data\n"+
		"                  x_k = %v\n"+
		"                 precision = %v\n"+
		"     calculation precision = %v\n"+
		"             raisedToPower = %v\n"+
		"             absAlphaDelta = %v\n"+
		"             compareResult = %v\n"+
		"Initialization Cycle Count = %v\n"+
		"   Calculation Cycle Count = %v\n\n",
		x_k.Text('f', 20),
		prec,
		calculationPrecision,
		raisedToPower.Text('f', 20),
		absAlphaDelta.Text('f', 20),
		compareResult,
		initializeCycleCount,
		calculationCycleCount)

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

	accuracyThreshold := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	accuracyThreshold.SetString("0.000015")

	var calculationPrecision uint

	x_k,
		raisedToPower,
		absAlphaDelta,
		compareResult,
		initializeCycleCount,
		calculationCycleCount,
		calculationPrecision := NewtonInitialGuess(
		n_int64,
		alpha,
		accuracyThreshold,
		prec)

	fmt.Printf("First Guess Data\n"+
		"                  x_k = %v\n"+
		"                 precision = %v\n"+
		"     calculation precision = %v\n"+
		"             raisedToPower = %v\n"+
		"             absAlphaDelta = %v\n"+
		"             compareResult = %v\n"+
		"Initialization Cycle Count = %v\n"+
		"   Calculation Cycle Count = %v\n\n",
		x_k.Text('f', 20),
		prec,
		calculationPrecision,
		raisedToPower.Text('f', 20),
		absAlphaDelta.Text('f', 20),
		compareResult,
		initializeCycleCount,
		calculationCycleCount)

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

func Newton01() {

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

func calculateRequiredPrecision(
	integerDigits int64,
	fractionalDigits int64,
	spareBufferDigits int64) (
	uint,
	error) {

	funcName := "calculateRequiredPrecision"

	totalDigits :=
		integerDigits +
			fractionalDigits +
			spareBufferDigits

	totalDigitsFloat := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(totalDigits)

	factorEightFloat := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(8)

	precToDigitsFactor := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	// 3.3219789132197891321978913219789

	var err error
	_,
		ok := precToDigitsFactor.SetString("3.3219789132197891321978913219789")

	if !ok {
		err = fmt.Errorf("\n%v\n"+
			"expectedNthRoot.SetString(\"3.3219789132197891321978913219789\") FAILED!\n",
			funcName)

		return 0, err
	}

	totalDigitsFloat.Mul(totalDigitsFloat, precToDigitsFactor)

	totalDigitsFloat.Quo(totalDigitsFloat, factorEightFloat)

	baseDigits,
		accuracy := totalDigitsFloat.Int64()

	if accuracy == -1 {
		baseDigits++
	}

	baseDigits = baseDigits * 8

	return uint(baseDigits), err
}

// setPrecisionFromDigits
//
// Sets the number precision bits required to store
// X Numerical Digits in the mantissa of a big.Float
// number.
func setPrecisionFromDigits02(
	totalRequiredNumericalDigits int64) (
	precision uint) {

	funcName := "setPrecisionFromDigits"

	totalRequiredNumericalDigits += 100

	precToDigitsFactor := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	// 3.3219789132197891321978913219789

	_,
		ok := precToDigitsFactor.SetString("3.3219789132197891321978913219789")

	if !ok {
		fmt.Printf("%v\n"+
			"expectedNthRoot.SetString(\"3.3219789132197891321978913219789\") FAILED!\n",
			funcName)

		return
	}

	estimatedMinPrecision := new(big.Float).
		SetInt64(totalRequiredNumericalDigits).
		SetMode(big.AwayFromZero)

	estimatedMinPrecision.Mul(
		estimatedMinPrecision, precToDigitsFactor)

	tFloat := new(big.Float).
		SetInt64(0).
		SetMode(big.AwayFromZero)

	tFloat.SetInt64(8)

	estimatedMinPrecision.Quo(estimatedMinPrecision, tFloat)

	estimatedMinPrecision.Mul(estimatedMinPrecision, tFloat)

	tFloat.SetInt64(40)

	estimatedMinPrecision.Add(estimatedMinPrecision, tFloat)

	precUint64,
		_ := estimatedMinPrecision.Uint64()

	precision = uint(precUint64)

	return precision
}

func TestLog10InitialGuess() {

	prec := uint(200)

	alpha := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	alpha.SetString("64")

	nFloat := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	nFloat.SetString("3")

	// 2^3 = 8
	Log10InitialGuess(
		alpha,
		nFloat,
		prec)

}

func Log10InitialGuess(
	alpha *big.Float,
	nFloat *big.Float,
	prec uint) (
	guess *big.Float) {

	funcName := "Log10InitialGuess"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	localPrec := 20

	fmt.Printf("Input Values\n"+
		"Fractional Digits = %v\n"+
		" alpha = %v\n"+
		"nFloat = %v\n",
		localPrec,
		alpha.Text('f', localPrec),
		nFloat.Text('f', localPrec))

	guess = new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	alpha_float64,
		accuracy := alpha.Float64()

	if accuracy != big.Exact {

		fmt.Printf("alpha.Float64() accuracy NOT EXACT!\n"+
			"accuracy = %v\n",
			accuracy)

	}

	var n_Float64 float64

	n_Float64,
		accuracy = nFloat.Float64()

	if accuracy != big.Exact {

		fmt.Printf("nFloat.Float64() accuracy NOT EXACT!\n"+
			"accuracy = %v\n",
			accuracy)

	}

	log10_alpha := math.Log10(alpha_float64)

	log10_n := math.Log10(n_Float64)

	log_n_of_alpha := log10_alpha / log10_n

	log_n_of_alpha_Float :=
		new(big.Float).
			SetPrec(prec).
			SetMode(big.AwayFromZero).
			SetFloat64(log_n_of_alpha)

	fmt.Printf("Results: log_n_of_alpha\n"+
		"Fractional Digits = %v\n"+
		"log_n_of_alpha_Float = %v\n",
		localPrec,
		log_n_of_alpha_Float.Text('f', localPrec))

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("  Successful Completion!\n" +
		"Function: " +
		funcName + "\n")

	fmt.Printf(breakStr + "\n")

	return
}

func NewtonInitialGuess(
	n_int64 int64, // 18
	alpha *big.Float, // 592
	accuracyThresholdDelta *big.Float,
	prec uint) (
	guess *big.Float,
	raisedToPower *big.Float,
	absAlphaDelta *big.Float,
	raisedToPowerAlphaCompareResult int,
	initializeCycleCount uint64,
	calculationCycleCount uint64,
	calculationPrec uint) {

	funcName := "NewtonInitialGuess()"

	var ok bool

	calculationPrec = prec + 32

	tempStr := ""

	tempStr02 := ""

	percent50 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(0.5)

	guess = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPower = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	absAlphaDelta = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	internalIncrementFactor := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	highLowGuessAbsDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPowerAlphaCompareResult = 0

	initializeCycleCount = 0
	calculationCycleCount = 0

	lastHighGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	percentVariance := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	tempStr = fmt.Sprintf("%v",
		alpha.Text('f', -1))

	if tempStr[0] == '0' {

		tempStr02 = "0."

	} else {

		tempStr02 = "1."
	}

	lenStr := len(tempStr)

	idx := strings.Index(tempStr, ".")

	numFracDigits := 0

	if idx == -1 {

		numFracDigits = 1

	} else {

		numFracDigits = lenStr - idx
	}

	tempStr02 += strings.Repeat("0", numFracDigits)
	tempStr = tempStr02
	tempStr += "1"   // Low Guess
	tempStr02 += "5" // High Guess

	_,
		ok = lastLowGuess.SetString(tempStr)

	if !ok {
		fmt.Printf("%v\n"+
			"lastLowGuess.SetString(tempStr)!\n"+
			"tempStr= %v\n",
			funcName,
			tempStr)
		return
	}

	guess.Set(lastLowGuess)

	_,
		ok = lastHighGuess.SetString(tempStr02)

	if !ok {
		fmt.Printf("%v\n"+
			"lastHighGuess.SetString(tempStr02) FAILED!\n"+
			"tempStr02= %v\n",
			funcName,
			tempStr02)
		return
	}

	internalIncrementFactor.SetInt64(10)

	guess.Mul(guess, internalIncrementFactor)

	// Set the High And Low Guesses
	//raisedToPower = BigFloatPower(
	//	guess,
	//	n_int64,
	//	calculationPrec)

	var resultsStats BigFloatCalcStats
	var err error

	resultsStats,
		err = raiseToPowerInt(
		guess,
		n_int64)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"Error:\n%v\n\n",
			funcName,
			err.Error())

		return
	}

	raisedToPower.Copy(
		&resultsStats.CalcResult)

	raisedToPowerAlphaCompareResult =
		raisedToPower.Cmp(alpha)

	internalIncrementFactor.SetInt64(2)

	for raisedToPowerAlphaCompareResult < 1 {

		absAlphaDelta.Sub(alpha, raisedToPower)

		if raisedToPowerAlphaCompareResult == 0 {

			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				initializeCycleCount,
				calculationCycleCount,
				calculationPrec

		}

		lastLowGuess.Set(guess)
		lastLowAbsAlphaDelta.Set(absAlphaDelta)
		lastLowRaisedToPower.Set(raisedToPower)

		percentVariance.Quo(alpha, raisedToPower)

		guess.Mul(guess, internalIncrementFactor)

		//raisedToPower = BigFloatPower(
		//	guess,
		//	n_int64,
		//	calculationPrec)

		resultsStats,
			err = raiseToPowerInt(
			guess,
			n_int64)

		if err != nil {

			fmt.Printf("\n\n%v\n"+
				"X2\n"+
				"Error:\n%v\n\n",
				funcName,
				err.Error())

			return
		}

		raisedToPower.Copy(
			&resultsStats.CalcResult)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		initializeCycleCount++
	}

	lastHighGuess.Set(guess)
	lastHighAbsAlphaDelta.Sub(raisedToPower, alpha)
	lastHighRaisedToPower.Set(raisedToPower)

	guess.Set(lastLowGuess)

	fmt.Printf("\n\n%v\n"+
		"Initial Guesses\n"+
		"lastHighGuess = %v\n"+
		"lastLowGuess  = %v\n"+
		"       guess  = %v\n\n",
		funcName,
		lastHighGuess.Text('f', 50),
		lastLowGuess.Text('f', 50),
		guess.Text('f', 50))

	closeEnough := false

	for closeEnough == false {

		//raisedToPower = BigFloatPower(
		//	guess,
		//	n_int64,
		//	prec)

		resultsStats,
			err = raiseToPowerInt(
			guess,
			n_int64)

		if err != nil {

			fmt.Printf("\n\n%v\n"+
				"Error:\n%v\n\n",
				funcName,
				err.Error())

			return
		}

		raisedToPower.Copy(
			&resultsStats.CalcResult)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		if raisedToPowerAlphaCompareResult == 0 {
			// Target Located!

			absAlphaDelta.Sub(alpha, raisedToPower)

			// On target
			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				initializeCycleCount,
				calculationCycleCount,
				calculationPrec

		} else if raisedToPowerAlphaCompareResult == -1 {
			// Guess is low. Alpha is higher than
			// raisedToPower. Increase Guess.
			// Last High Guess is higher than
			//	Last Low Guess

			absAlphaDelta.Sub(alpha, raisedToPower)

			highLowGuessAbsDelta.Sub(lastHighGuess, guess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			lastLowGuess.Set(guess)
			lastLowAbsAlphaDelta.Set(absAlphaDelta)
			lastLowRaisedToPower.Set(raisedToPower)

			guess.Sub(lastHighGuess, highLowGuessAbsDelta)

		} else {
			// MUST BE -
			//	Guess is high. raisedToPower is higher
			//	than Alpha. Reduce Guess.

			absAlphaDelta.Sub(raisedToPower, alpha)

			highLowGuessAbsDelta.Sub(guess, lastLowGuess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			lastHighGuess.Set(guess)
			lastHighAbsAlphaDelta.Set(absAlphaDelta)
			lastHighRaisedToPower.Set(raisedToPower)
			guess.Add(lastLowGuess, highLowGuessAbsDelta)

		}

		if lastLowAbsAlphaDelta.Cmp(accuracyThresholdDelta) < 1 {

			guess.Set(lastLowGuess)
			raisedToPower.Set(lastLowRaisedToPower)
			absAlphaDelta.Set(lastLowAbsAlphaDelta)
			raisedToPowerAlphaCompareResult = -1

			closeEnough = true
		}

		if calculationCycleCount > 10000 {

			if raisedToPowerAlphaCompareResult == 1 {
				guess.Set(lastLowGuess)
				raisedToPower.Set(lastLowRaisedToPower)
				absAlphaDelta.Sub(alpha, lastLowRaisedToPower)
				raisedToPowerAlphaCompareResult = -1
			}

			closeEnough = true

		}

		calculationCycleCount++

	}

	return guess,
		raisedToPower,
		absAlphaDelta,
		raisedToPowerAlphaCompareResult,
		initializeCycleCount,
		calculationCycleCount,
		calculationPrec
}

func NewtonInitialGuess08(
	n_int64 int64, // 18
	alpha *big.Float, // 592
	accuracyThresholdDelta *big.Float,
	prec uint) (
	guess *big.Float,
	raisedToPower *big.Float,
	absAlphaDelta *big.Float,
	raisedToPowerAlphaCompareResult int,
	initializeCycleCount uint64,
	calculationCycleCount uint64,
	calculationPrec uint) {

	funcName := "NewtonInitialGuess()"

	var ok bool

	calculationPrec = prec + 32

	tempStr := ""

	tempStr02 := ""

	percent50 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(0.5)

	percent100 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(1.0)

	percent200 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(2.00)

	percent500 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(5.00)

	percent1000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(10.00)

	percent2000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(20.000)

	percent3000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(30.000)

	percent4000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(40.000)

	percent6000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(60.000)

	percent8000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(60.000)

	percent10000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(100.000)

	guess = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPower = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	absAlphaDelta = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	internalIncrementFactor := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	highLowGuessAbsDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPowerAlphaCompareResult = 0

	initializeCycleCount = 0
	calculationCycleCount = 0

	lastHighGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	percentVariance := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	tempStr = fmt.Sprintf("%v",
		alpha.Text('f', -1))

	if tempStr[0] == '0' {

		tempStr02 = "0."

	} else {

		tempStr02 = "1."
	}

	lenStr := len(tempStr)

	idx := strings.Index(tempStr, ".")

	numFracDigits := 0

	if idx == -1 {

		numFracDigits = 1

	} else {

		numFracDigits = lenStr - idx
	}

	tempStr02 += strings.Repeat("0", numFracDigits)
	tempStr = tempStr02
	tempStr += "1"   // Low Guess
	tempStr02 += "5" // High Guess

	_,
		ok = lastLowGuess.SetString(tempStr)

	if !ok {
		fmt.Printf("%v\n"+
			"lastLowGuess.SetString(tempStr)!\n"+
			"tempStr= %v\n",
			funcName,
			tempStr)
		return
	}

	guess.Set(lastLowGuess)

	_,
		ok = lastHighGuess.SetString(tempStr02)

	if !ok {
		fmt.Printf("%v\n"+
			"lastHighGuess.SetString(tempStr02) FAILED!\n"+
			"tempStr02= %v\n",
			funcName,
			tempStr02)
		return
	}

	internalIncrementFactor.SetInt64(10)

	guess.Mul(guess, internalIncrementFactor)

	// Set the High And Low Guesses
	raisedToPower = BigFloatPower(
		guess,
		n_int64,
		calculationPrec)

	raisedToPowerAlphaCompareResult =
		raisedToPower.Cmp(alpha)

	for raisedToPowerAlphaCompareResult < 1 {

		absAlphaDelta.Sub(alpha, raisedToPower)

		if raisedToPowerAlphaCompareResult == 0 {

			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				initializeCycleCount,
				calculationCycleCount,
				calculationPrec

		}

		lastLowGuess.Set(guess)
		lastLowAbsAlphaDelta.Set(absAlphaDelta)
		lastLowRaisedToPower.Set(raisedToPower)

		percentVariance.Quo(alpha, raisedToPower)

		if percentVariance.Cmp(percent10000) > -1 {

			// 1-billion
			internalIncrementFactor.SetInt64(1000000000)

		} else if percentVariance.Cmp(percent8000) > -1 {

			// 100-million
			internalIncrementFactor.SetInt64(100000000)

		} else if percentVariance.Cmp(percent6000) > -1 {

			// 100,000 one-hundred thousand
			internalIncrementFactor.SetInt64(100000)

		} else if percentVariance.Cmp(percent4000) > -1 {

			// 10,000 ten thousand
			internalIncrementFactor.SetInt64(10000)

		} else if percentVariance.Cmp(percent3000) > -1 {

			// 1,000 one thousand
			internalIncrementFactor.SetInt64(1000)

		} else if percentVariance.Cmp(percent2000) > -1 {

			// 500 times
			internalIncrementFactor.SetInt64(500)

		} else if percentVariance.Cmp(percent1000) > -1 {

			// 10 times
			internalIncrementFactor.SetInt64(10)

		} else if percentVariance.Cmp(percent500) > -1 {

			// 2 - double
			internalIncrementFactor.SetInt64(2)

		} else if percentVariance.Cmp(percent200) > -1 {

			// Factor 1.75
			internalIncrementFactor.SetFloat64(1.75)

		} else if percentVariance.Cmp(percent100) > -1 {

			// Factor 1.5
			internalIncrementFactor.SetFloat64(1.5)

		} else if percentVariance.Cmp(percent50) > -1 {

			// Factor 1.25
			internalIncrementFactor.SetFloat64(1.25)

		} else {
			// 1.10
			internalIncrementFactor.SetFloat64(1.10)
		}

		guess.Mul(guess, internalIncrementFactor)

		raisedToPower = BigFloatPower(
			guess,
			n_int64,
			calculationPrec)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		initializeCycleCount++
	}

	lastHighGuess.Set(guess)
	lastHighAbsAlphaDelta.Sub(raisedToPower, alpha)
	lastHighRaisedToPower.Set(raisedToPower)

	guess.Set(lastLowGuess)

	fmt.Printf("\n\n%v\n"+
		"Initial Guesses\n"+
		"lastHighGuess = %v\n"+
		"lastLowGuess  = %v\n"+
		"       guess  = %v\n\n",
		funcName,
		lastHighGuess.Text('f', 50),
		lastLowGuess.Text('f', 50),
		guess.Text('f', 50))

	closeEnough := false

	for closeEnough == false {

		raisedToPower = BigFloatPower(
			guess,
			n_int64,
			prec)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		if raisedToPowerAlphaCompareResult == 0 {
			// Target Located!

			absAlphaDelta.Sub(alpha, raisedToPower)

			// On target
			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				initializeCycleCount,
				calculationCycleCount,
				calculationPrec

		} else if raisedToPowerAlphaCompareResult == -1 {
			// Guess is low. Alpha is higher than
			// raisedToPower. Increase Guess.
			// Last High Guess is higher than
			//	Last Low Guess

			absAlphaDelta.Sub(alpha, raisedToPower)

			highLowGuessAbsDelta.Sub(lastHighGuess, guess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			lastLowGuess.Set(guess)
			lastLowAbsAlphaDelta.Set(absAlphaDelta)
			lastLowRaisedToPower.Set(raisedToPower)

			guess.Sub(lastHighGuess, highLowGuessAbsDelta)

		} else {
			// MUST BE -
			//	Guess is high. raisedToPower is higher
			//	than Alpha. Reduce Guess.

			absAlphaDelta.Sub(raisedToPower, alpha)

			highLowGuessAbsDelta.Sub(guess, lastLowGuess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			lastHighGuess.Set(guess)
			lastHighAbsAlphaDelta.Set(absAlphaDelta)
			lastHighRaisedToPower.Set(raisedToPower)
			guess.Add(lastLowGuess, highLowGuessAbsDelta)

		}

		if lastLowAbsAlphaDelta.Cmp(accuracyThresholdDelta) < 1 {

			guess.Set(lastLowGuess)
			raisedToPower.Set(lastLowRaisedToPower)
			absAlphaDelta.Set(lastLowAbsAlphaDelta)
			raisedToPowerAlphaCompareResult = -1

			closeEnough = true
		}

		if calculationCycleCount > 10000 {

			if raisedToPowerAlphaCompareResult == 1 {
				guess.Set(lastLowGuess)
				raisedToPower.Set(lastLowRaisedToPower)
				absAlphaDelta.Sub(alpha, lastLowRaisedToPower)
				raisedToPowerAlphaCompareResult = -1
			}

			closeEnough = true

		}

		calculationCycleCount++

	}

	return guess,
		raisedToPower,
		absAlphaDelta,
		raisedToPowerAlphaCompareResult,
		initializeCycleCount,
		calculationCycleCount,
		calculationPrec
}

func NewtonInitialGuess07(
	n_int64 int64, // 18
	alpha *big.Float, // 592
	accuracyThresholdDelta *big.Float,
	prec uint) (
	guess *big.Float,
	raisedToPower *big.Float,
	absAlphaDelta *big.Float,
	raisedToPowerAlphaCompareResult int,
	initializeCycleCount uint64,
	calculationCycleCount uint64,
	calculationPrec uint) {

	funcName := "NewtonInitialGuess07()"

	var ok bool

	calculationPrec = prec + 16

	tempStr := ""

	percent50 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(0.5)

	percent100 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(1.0)

	percent200 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(2.00)

	percent500 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(5.00)

	percent1000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(10.00)

	percent2000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(20.000)

	percent3000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(30.000)

	percent4000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(40.000)

	percent6000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(60.000)

	percent8000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(60.000)

	percent10000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(100.000)

	guess = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPower = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	absAlphaDelta = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	internalIncrementFactor := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	highLowGuessAbsDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPowerAlphaCompareResult = 0

	initializeCycleCount = 0
	calculationCycleCount = 0

	lastHighGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	percentVariance := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	tempStr = fmt.Sprintf("%v",
		alpha.Text('f', -1))

	lenStr := len(tempStr)

	idx := strings.Index(tempStr, ".")

	numFracDigits := 0

	if idx == -1 {

		tempStr += ".0"
		numFracDigits = 1

	} else {

		numFracDigits = lenStr - idx
	}

	tempStr = "0."
	tempStr += strings.Repeat("0", numFracDigits)
	tempStr += "9"

	_,
		ok = guess.SetString(tempStr)

	if !ok {
		fmt.Printf("%v\n"+
			"alpha greater than or equal to '1'"+
			"guess.SetString(tempStr) FAILED!\n"+
			"tempStr= %v\n",
			funcName,
			tempStr)
		return
	}

	lastLowGuess.Set(guess)

	// Set the High And Low Guesses
	raisedToPower = BigFloatPower(
		guess,
		n_int64,
		calculationPrec)

	raisedToPowerAlphaCompareResult =
		raisedToPower.Cmp(alpha)

	internalIncrementFactor.SetInt64(10)

	for raisedToPowerAlphaCompareResult < 1 {

		absAlphaDelta.Sub(alpha, raisedToPower)

		if raisedToPowerAlphaCompareResult == 0 {

			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				initializeCycleCount,
				calculationCycleCount,
				calculationPrec

		}

		lastLowGuess.Set(guess)
		lastLowAbsAlphaDelta.Set(absAlphaDelta)
		lastLowRaisedToPower.Set(raisedToPower)

		percentVariance.Quo(alpha, raisedToPower)

		tempStr = percentVariance.Text('f', 3)

		if percentVariance.Cmp(percent10000) > -1 {

			// 1-billion
			internalIncrementFactor.SetInt64(1000000000)

		} else if percentVariance.Cmp(percent8000) > -1 {

			// 100-million
			internalIncrementFactor.SetInt64(100000000)

		} else if percentVariance.Cmp(percent6000) > -1 {

			// 100,000 one-hundred thousand
			internalIncrementFactor.SetInt64(100000)

		} else if percentVariance.Cmp(percent4000) > -1 {

			// 10,000 ten thousand
			internalIncrementFactor.SetInt64(10000)

		} else if percentVariance.Cmp(percent3000) > -1 {

			// 1,000 one thousand
			internalIncrementFactor.SetInt64(1000)

		} else if percentVariance.Cmp(percent2000) > -1 {

			// 500 times
			internalIncrementFactor.SetInt64(500)

		} else if percentVariance.Cmp(percent1000) > -1 {

			// 20 times
			internalIncrementFactor.SetInt64(20)

		} else if percentVariance.Cmp(percent500) > -1 {

			// 2 - double
			internalIncrementFactor.SetInt64(2)

		} else if percentVariance.Cmp(percent200) > -1 {

			// Factor 1.75
			internalIncrementFactor.SetFloat64(1.75)

		} else if percentVariance.Cmp(percent100) > -1 {

			// Factor 1.5
			internalIncrementFactor.SetFloat64(1.5)

		} else if percentVariance.Cmp(percent50) > -1 {

			// Factor 1.25
			internalIncrementFactor.SetFloat64(1.25)

		} else {
			// 1.10
			internalIncrementFactor.SetFloat64(1.10)
		}

		guess.Mul(guess, internalIncrementFactor)

		raisedToPower = BigFloatPower(
			guess,
			n_int64,
			calculationPrec)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		initializeCycleCount++
	}

	lastHighGuess.Set(guess)
	lastHighAbsAlphaDelta.Sub(raisedToPower, alpha)
	lastHighRaisedToPower.Set(raisedToPower)

	guess.Set(lastLowGuess)

	fmt.Printf("\n\n%v\n"+
		"Initial Guesses\n"+
		"lastHighGuess = %v\n"+
		"lastLowGuess  = %v\n"+
		"       guess  = %v\n\n",
		funcName,
		lastHighGuess.Text('f', 50),
		lastLowGuess.Text('f', 50),
		guess.Text('f', 50))

	closeEnough := false

	for closeEnough == false {

		raisedToPower = BigFloatPower(
			guess,
			n_int64,
			prec)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		if raisedToPowerAlphaCompareResult == 0 {
			// Target Located!

			absAlphaDelta.Sub(alpha, raisedToPower)

			// On target
			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				initializeCycleCount,
				calculationCycleCount,
				calculationPrec

		} else if raisedToPowerAlphaCompareResult == -1 {
			// Guess is low. Alpha is higher than
			// raisedToPower. Increase Guess.
			// Last High Guess is higher than
			//	Last Low Guess

			absAlphaDelta.Sub(alpha, raisedToPower)

			highLowGuessAbsDelta.Sub(lastHighGuess, guess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			lastLowGuess.Set(guess)
			lastLowAbsAlphaDelta.Set(absAlphaDelta)
			lastLowRaisedToPower.Set(raisedToPower)

			guess.Sub(lastHighGuess, highLowGuessAbsDelta)

		} else {
			// MUST BE -
			//	Guess is high. raisedToPower is higher
			//	than Alpha. Reduce Guess.

			absAlphaDelta.Sub(raisedToPower, alpha)

			highLowGuessAbsDelta.Sub(guess, lastLowGuess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			lastHighGuess.Set(guess)
			lastHighAbsAlphaDelta.Set(absAlphaDelta)
			lastHighRaisedToPower.Set(raisedToPower)
			guess.Add(lastLowGuess, highLowGuessAbsDelta)

		}

		if lastLowAbsAlphaDelta.Cmp(accuracyThresholdDelta) < 1 {

			guess.Set(lastLowGuess)
			raisedToPower.Set(lastLowRaisedToPower)
			absAlphaDelta.Set(lastLowAbsAlphaDelta)
			raisedToPowerAlphaCompareResult = -1

			closeEnough = true
		}

		if calculationCycleCount > 10000 {

			if raisedToPowerAlphaCompareResult == 1 {
				guess.Set(lastLowGuess)
				raisedToPower.Set(lastLowRaisedToPower)
				absAlphaDelta.Sub(alpha, lastLowRaisedToPower)
				raisedToPowerAlphaCompareResult = -1
			}

			closeEnough = true

		}

		calculationCycleCount++

	}

	return guess,
		raisedToPower,
		absAlphaDelta,
		raisedToPowerAlphaCompareResult,
		initializeCycleCount,
		calculationCycleCount,
		calculationPrec
}

func NewtonInitialGuess06(
	n_int64 int64, // 18
	alpha *big.Float, // 592
	accuracyThresholdDelta *big.Float,
	prec uint) (
	guess *big.Float,
	raisedToPower *big.Float,
	absAlphaDelta *big.Float,
	raisedToPowerAlphaCompareResult int,
	initializeCycleCount uint64,
	calculationCycleCount uint64,
	calculationPrec uint) {

	funcName := "NewtonInitialGuess06()"

	var ok bool

	calculationPrec = prec + 16

	tempStr := ""

	percent50 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(0.5)

	percent100 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(1.0)

	percent200 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(2.00)

	percent500 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(5.00)

	percent1000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(10.00)

	percent2000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(20.000)

	percent3000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(30.000)

	percent4000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(40.000)

	percent6000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(60.000)

	guess = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPower = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	absAlphaDelta = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	internalIncrementFactor := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	highLowGuessAbsDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPowerAlphaCompareResult = 0

	initializeCycleCount = 0
	calculationCycleCount = 0

	lastHighGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	percentVariance := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	tempStr = fmt.Sprintf("%v",
		alpha.Text('f', -1))

	lenStr := len(tempStr)

	idx := strings.Index(tempStr, ".")

	numFracDigits := 0

	if idx == -1 {

		tempStr += ".0"
		numFracDigits = 1

	} else {

		numFracDigits = lenStr - idx
	}

	tempStr = "0."
	tempStr += strings.Repeat("0", numFracDigits)
	tempStr += "9"

	_,
		ok = guess.SetString(tempStr)

	if !ok {
		fmt.Printf("%v\n"+
			"alpha greater than or equal to '1'"+
			"guess.SetString(tempStr) FAILED!\n"+
			"tempStr= %v\n",
			funcName,
			tempStr)
		return
	}

	lastLowGuess.Set(guess)

	// Set the High And Low Guesses
	raisedToPower = BigFloatPower(
		guess,
		n_int64,
		calculationPrec)

	raisedToPowerAlphaCompareResult =
		raisedToPower.Cmp(alpha)

	internalIncrementFactor.SetInt64(10)

	for raisedToPowerAlphaCompareResult < 1 {

		absAlphaDelta.Sub(alpha, raisedToPower)

		if raisedToPowerAlphaCompareResult == 0 {

			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				initializeCycleCount,
				calculationCycleCount,
				calculationPrec

		}

		lastLowGuess.Set(guess)
		lastLowAbsAlphaDelta.Set(absAlphaDelta)
		lastLowRaisedToPower.Set(raisedToPower)

		percentVariance.Quo(alpha, raisedToPower)

		tempStr = percentVariance.Text('f', 3)

		if percentVariance.Cmp(percent6000) > -1 {

			// 1-billion
			internalIncrementFactor.SetInt64(1000000000)

		} else if percentVariance.Cmp(percent4000) > -1 {

			// 1-million
			internalIncrementFactor.SetInt64(1000000)

		} else if percentVariance.Cmp(percent3000) > -1 {

			// 100,000 one-hundred thousand
			internalIncrementFactor.SetInt64(100000)

		} else if percentVariance.Cmp(percent2000) > -1 {

			// 10,000 ten thousand
			internalIncrementFactor.SetInt64(10000)

		} else if percentVariance.Cmp(percent1000) > -1 {

			// 1,000 one thousand
			internalIncrementFactor.SetInt64(1000)

		} else if percentVariance.Cmp(percent500) > -1 {

			// 2 - double
			internalIncrementFactor.SetInt64(2)

		} else if percentVariance.Cmp(percent200) > -1 {

			// Factor 1.75
			internalIncrementFactor.SetFloat64(1.75)

		} else if percentVariance.Cmp(percent100) > -1 {

			// Factor 1.5
			internalIncrementFactor.SetFloat64(1.5)

		} else if percentVariance.Cmp(percent50) > -1 {

			// Factor 1.25
			internalIncrementFactor.SetFloat64(1.25)

		} else {
			// 1.10
			internalIncrementFactor.SetFloat64(1.10)
		}

		guess.Mul(guess, internalIncrementFactor)

		raisedToPower = BigFloatPower(
			guess,
			n_int64,
			calculationPrec)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		initializeCycleCount++
	}

	lastHighGuess.Set(guess)
	lastHighAbsAlphaDelta.Sub(raisedToPower, alpha)
	lastHighRaisedToPower.Set(raisedToPower)

	guess.Set(lastLowGuess)

	fmt.Printf("\n\n%v\n"+
		"Initial Guesses\n"+
		"lastHighGuess = %v\n"+
		"lastLowGuess  = %v\n"+
		"       guess  = %v\n\n",
		funcName,
		lastHighGuess.Text('f', 50),
		lastLowGuess.Text('f', 50),
		guess.Text('f', 50))

	closeEnough := false

	for closeEnough == false {

		raisedToPower = BigFloatPower(
			guess,
			n_int64,
			prec)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		if raisedToPowerAlphaCompareResult == 0 {
			// Target Located!

			absAlphaDelta.Sub(alpha, raisedToPower)

			// On target
			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				initializeCycleCount,
				calculationCycleCount,
				calculationPrec

		} else if raisedToPowerAlphaCompareResult == -1 {
			// Guess is low. Alpha is higher than
			// raisedToPower. Increase Guess.
			// Last High Guess is higher than
			//	Last Low Guess

			absAlphaDelta.Sub(alpha, raisedToPower)

			highLowGuessAbsDelta.Sub(lastHighGuess, guess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			lastLowGuess.Set(guess)
			lastLowAbsAlphaDelta.Set(absAlphaDelta)
			lastLowRaisedToPower.Set(raisedToPower)

			guess.Sub(lastHighGuess, highLowGuessAbsDelta)

		} else {
			// MUST BE -
			//	Guess is high. raisedToPower is higher
			//	than Alpha. Reduce Guess.

			absAlphaDelta.Sub(raisedToPower, alpha)

			highLowGuessAbsDelta.Sub(guess, lastLowGuess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			lastHighGuess.Set(guess)
			lastHighAbsAlphaDelta.Set(absAlphaDelta)
			lastHighRaisedToPower.Set(raisedToPower)
			guess.Add(lastLowGuess, highLowGuessAbsDelta)

		}

		if lastLowAbsAlphaDelta.Cmp(accuracyThresholdDelta) < 1 {

			guess.Set(lastLowGuess)
			raisedToPower.Set(lastLowRaisedToPower)
			absAlphaDelta.Set(lastLowAbsAlphaDelta)
			raisedToPowerAlphaCompareResult = -1

			closeEnough = true
		}

		if calculationCycleCount > 10000 {

			if raisedToPowerAlphaCompareResult == 1 {
				guess.Set(lastLowGuess)
				raisedToPower.Set(lastLowRaisedToPower)
				absAlphaDelta.Sub(alpha, lastLowRaisedToPower)
				raisedToPowerAlphaCompareResult = -1
			}

			closeEnough = true

		}

		calculationCycleCount++

	}

	return guess,
		raisedToPower,
		absAlphaDelta,
		raisedToPowerAlphaCompareResult,
		initializeCycleCount,
		calculationCycleCount,
		calculationPrec
}

func NewtonInitialGuess05(
	n_int64 int64, // 18
	alpha *big.Float, // 592
	accuracyThresholdDelta *big.Float,
	prec uint) (
	guess *big.Float,
	raisedToPower *big.Float,
	absAlphaDelta *big.Float,
	raisedToPowerAlphaCompareResult int,
	initializeCycleCount uint64,
	calculationCycleCount uint64,
	calculationPrec uint) {

	funcName := "NewtonInitialGuess05()"

	var ok bool

	calculationPrec = prec + 16

	tempStr := ""

	percent50 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(0.5)

	percent75 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(0.5)

	//percent150 := new(big.Float).
	//	SetInt64(0).
	//	SetPrec(calculationPrec).
	//	SetMode(big.AwayFromZero).
	//	SetFloat64(1.50)

	percent200 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(2.00)

	percent500 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(5.00)

	//percent1000 := new(big.Float).
	//	SetInt64(0).
	//	SetPrec(calculationPrec).
	//	SetMode(big.AwayFromZero).
	//	SetFloat64(10.00)

	percent2000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(20.000)

	percent4000 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero).
		SetFloat64(40.000)

	guess = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPower = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	absAlphaDelta = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	internalIncrementFactor := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	highLowGuessAbsDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPowerAlphaCompareResult = 0

	initializeCycleCount = 0
	calculationCycleCount = 0

	lastHighGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	percentVariance := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	tempStr = fmt.Sprintf("%v",
		alpha.Text('f', -1))

	lenStr := len(tempStr)

	idx := strings.Index(tempStr, ".")

	numFracDigits := 0

	if idx == -1 {

		tempStr += ".0"
		numFracDigits = 1

	} else {

		numFracDigits = lenStr - idx
	}

	tempStr = "0."
	tempStr += strings.Repeat("0", numFracDigits)
	tempStr += "9"

	_,
		ok = guess.SetString(tempStr)

	if !ok {
		fmt.Printf("%v\n"+
			"alpha greater than or equal to '1'"+
			"guess.SetString(tempStr) FAILED!\n"+
			"tempStr= %v\n",
			funcName,
			tempStr)
		return
	}

	lastLowGuess.Set(guess)

	// Set the High And Low Guesses
	raisedToPower = BigFloatPower(
		guess,
		n_int64,
		calculationPrec)

	raisedToPowerAlphaCompareResult =
		raisedToPower.Cmp(alpha)

	internalIncrementFactor.SetInt64(10)

	for raisedToPowerAlphaCompareResult < 1 {

		absAlphaDelta.Sub(alpha, raisedToPower)

		if raisedToPowerAlphaCompareResult == 0 {

			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				initializeCycleCount,
				calculationCycleCount,
				calculationPrec

		}

		lastLowGuess.Set(guess)
		lastLowAbsAlphaDelta.Set(absAlphaDelta)
		lastLowRaisedToPower.Set(raisedToPower)

		percentVariance.Quo(alpha, raisedToPower)

		tempStr = percentVariance.Text('f', 3)

		if percentVariance.Cmp(percent4000) > -1 {

			// 1,000,000 1-million
			internalIncrementFactor.SetInt64(1000000)

		} else if percentVariance.Cmp(percent2000) > -1 {

			// 100,000 One Hundred Thousand
			internalIncrementFactor.SetInt64(100000)

		} else if percentVariance.Cmp(percent500) > -1 {

			// 10,000 ten thousand
			internalIncrementFactor.SetInt64(10000)

		} else if percentVariance.Cmp(percent200) > -1 {

			// 1,000 one-thousand
			internalIncrementFactor.SetInt64(1000)

		} else if percentVariance.Cmp(percent75) > -1 {

			// 100 one-hundred
			internalIncrementFactor.SetInt64(100)

		} else {
			// 10 1-ten
			internalIncrementFactor.SetInt64(10)
		}

		guess.Mul(guess, internalIncrementFactor)

		raisedToPower = BigFloatPower(
			guess,
			n_int64,
			calculationPrec)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		initializeCycleCount++
	}

	lastHighGuess.Set(guess)
	lastHighAbsAlphaDelta.Sub(raisedToPower, alpha)
	lastHighRaisedToPower.Set(raisedToPower)

	guess.Set(lastLowGuess)

	fmt.Printf("\n\n%v\n"+
		"Initial Guesses\n"+
		"lastHighGuess = %v\n"+
		"lastLowGuess  = %v\n"+
		"       guess  = %v\n\n",
		funcName,
		lastHighGuess.Text('f', 50),
		lastLowGuess.Text('f', 50),
		guess.Text('f', 50))

	closeEnough := false

	for closeEnough == false {

		raisedToPower = BigFloatPower(
			guess,
			n_int64,
			prec)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		if raisedToPowerAlphaCompareResult == 0 {
			// Target Located!

			absAlphaDelta.Sub(alpha, raisedToPower)

			// On target
			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				initializeCycleCount,
				calculationCycleCount,
				calculationPrec

		} else if raisedToPowerAlphaCompareResult == -1 {
			// Guess is low. Alpha is higher than
			// raisedToPower. Increase Guess.
			// Last High Guess is higher than
			//	Last Low Guess

			absAlphaDelta.Sub(alpha, raisedToPower)

			highLowGuessAbsDelta.Sub(lastHighGuess, guess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			lastLowGuess.Set(guess)
			lastLowAbsAlphaDelta.Set(absAlphaDelta)
			lastLowRaisedToPower.Set(raisedToPower)

			guess.Sub(lastHighGuess, highLowGuessAbsDelta)

		} else {
			// MUST BE -
			//	Guess is high. raisedToPower is higher
			//	than Alpha. Reduce Guess.

			absAlphaDelta.Sub(raisedToPower, alpha)

			highLowGuessAbsDelta.Sub(guess, lastLowGuess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			lastHighGuess.Set(guess)
			lastHighAbsAlphaDelta.Set(absAlphaDelta)
			lastHighRaisedToPower.Set(raisedToPower)
			guess.Add(lastLowGuess, highLowGuessAbsDelta)

		}

		if lastLowAbsAlphaDelta.Cmp(accuracyThresholdDelta) < 1 {

			guess.Set(lastLowGuess)
			raisedToPower.Set(lastLowRaisedToPower)
			absAlphaDelta.Set(lastLowAbsAlphaDelta)
			raisedToPowerAlphaCompareResult = -1

			closeEnough = true
		}

		if calculationCycleCount > 10000 {

			if raisedToPowerAlphaCompareResult == 1 {
				guess.Set(lastLowGuess)
				raisedToPower.Set(lastLowRaisedToPower)
				absAlphaDelta.Sub(alpha, lastLowRaisedToPower)
				raisedToPowerAlphaCompareResult = -1
			}

			closeEnough = true

		}

		calculationCycleCount++

	}

	return guess,
		raisedToPower,
		absAlphaDelta,
		raisedToPowerAlphaCompareResult,
		initializeCycleCount,
		calculationCycleCount,
		calculationPrec
}

func NewtonInitialGuess04(
	nthRoot int64, // 18
	alpha *big.Float, // 592
	accuracyThresholdDelta *big.Float,
	prec uint) (
	guess *big.Float,
	raisedToPower *big.Float,
	absAlphaDelta *big.Float,
	raisedToPowerAlphaCompareResult int,
	initializeCycleCount uint64,
	calculationCycleCount uint64,
	calculationPrec uint) {

	funcName := "NewtonInitialGuess04()"

	var ok bool

	calculationPrec = prec + 16

	tempStr := ""

	percent50 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)
	_,
		ok = percent50.SetString("0.5")

	if !ok {
		fmt.Printf("%v\n"+
			"percent50.SetString(\"0.5\") FAILED!\n"+
			"tempStr= %v\n",
			funcName,
			tempStr)
		return
	}

	tempFloat := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	guess = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPower = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	absAlphaDelta = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	highLowGuessAbsDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPowerAlphaCompareResult = 0

	initializeCycleCount = 0
	calculationCycleCount = 0

	lastHighGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	tempStr = fmt.Sprintf("%v",
		alpha.Text('f', -1))

	lenStr := len(tempStr)

	idx := strings.Index(tempStr, ".")

	numFracDigits := 0

	if idx == -1 {

		tempStr += ".0"
		numFracDigits = 1

	} else {

		numFracDigits = lenStr - idx
	}

	tempStr = "0."
	tempStr += strings.Repeat("0", numFracDigits)
	tempStr += "9"

	_,
		ok = guess.SetString(tempStr)

	if !ok {
		fmt.Printf("%v\n"+
			"alpha greater than or equal to '1'"+
			"guess.SetString(tempStr) FAILED!\n"+
			"tempStr= %v\n",
			funcName,
			tempStr)
		return
	}

	// Set the High And Low Guesses
	raisedToPower = BigFloatPower(
		guess,
		nthRoot,
		calculationPrec)

	raisedToPowerAlphaCompareResult =
		raisedToPower.Cmp(alpha)

	tempFloat.SetInt64(1000)

	for raisedToPowerAlphaCompareResult < 1 {

		if raisedToPowerAlphaCompareResult == 0 {

			absAlphaDelta.Sub(raisedToPower, alpha)

			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				initializeCycleCount,
				calculationCycleCount,
				calculationPrec

		}

		lastLowGuess.Set(guess)
		lastLowAbsAlphaDelta.Sub(alpha, raisedToPower)
		lastLowRaisedToPower.Set(raisedToPower)

		guess.Mul(guess, tempFloat)

		raisedToPower = BigFloatPower(
			guess,
			nthRoot,
			calculationPrec)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		initializeCycleCount++

	}

	lastHighGuess.Set(guess)
	lastHighAbsAlphaDelta.Sub(raisedToPower, alpha)
	lastHighRaisedToPower.Set(raisedToPower)

	guess.Set(lastLowGuess)

	fmt.Printf("%v\n"+
		"Initial Guesses\n"+
		"lastHighGuess = %v\n"+
		"lastLowGuess  = %v\n"+
		"       guess  = %v\n\n",
		funcName,
		lastHighGuess.Text('f', 50),
		lastLowGuess.Text('f', 50),
		guess.Text('f', 50))

	closeEnough := false

	var lastHighGuessStr, lastLowGuessStr,
		guessStr, newGuessStr, lastGuessStr,
		absAlphaDeltaStr, highLowGuessAbsDeltaStr string

	for closeEnough == false {

		if len(guessStr) > 0 {
			guessStr = ""
			newGuessStr = ""
		}

		guessStr = guess.Text('f', 50)

		raisedToPower = BigFloatPower(
			guess,
			nthRoot,
			prec)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		if raisedToPowerAlphaCompareResult == 0 {
			// Target Located!

			absAlphaDelta.Sub(alpha, raisedToPower)

			// On target
			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				initializeCycleCount,
				calculationCycleCount,
				calculationPrec

		} else if raisedToPowerAlphaCompareResult == -1 {
			// Guess is low. Alpha is higher than
			// raisedToPower. Increase Guess.
			// Last High Guess is higher than
			//	Last Low Guess

			absAlphaDelta.Sub(alpha, raisedToPower)

			absAlphaDeltaStr =
				absAlphaDelta.Text('f', 50)

			highLowGuessAbsDelta.Sub(lastHighGuess, guess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			highLowGuessAbsDeltaStr =
				highLowGuessAbsDelta.Text('f', 50)

			lastLowGuess.Set(guess)
			lastLowAbsAlphaDelta.Set(absAlphaDelta)
			lastLowRaisedToPower.Set(raisedToPower)

			guess.Sub(lastHighGuess, highLowGuessAbsDelta)

			lastGuessStr = guessStr
			newGuessStr = guess.Text('f', 50)

			lastLowGuessStr =
				lastLowGuess.Text('f', 50)

			tempStr = ""

			//fmt.Printf("Last Low Guess: %v\n",
			//	lastLowGuessStr)

		} else {
			// MUST BE -
			//	Guess is high. raisedToPower is higher
			//	than Alpha. Reduce Guess.

			absAlphaDelta.Sub(raisedToPower, alpha)

			absAlphaDeltaStr =
				absAlphaDelta.Text('f', 50)

			highLowGuessAbsDelta.Sub(guess, lastLowGuess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			highLowGuessAbsDeltaStr =
				highLowGuessAbsDelta.Text('f', 50)

			lastLowGuessStr =
				lastLowGuess.Text('f', 50)

			lastHighGuess.Set(guess)
			lastHighAbsAlphaDelta.Set(absAlphaDelta)
			lastHighRaisedToPower.Set(raisedToPower)
			guess.Add(lastLowGuess, highLowGuessAbsDelta)

			lastGuessStr = guessStr
			newGuessStr = guess.Text('f', 50)

			lastHighGuessStr =
				lastHighGuess.Text('f', 50)

			tempStr = ""
			//	fmt.Printf("Last High Guess: %v\n",
			//		lastHighGuessStr)

		}

		if len(lastHighGuessStr) > 0 ||
			len(lastLowGuessStr) > 0 ||
			len(newGuessStr) > 0 ||
			len(lastGuessStr) > 0 ||
			len(absAlphaDeltaStr) > 0 ||
			len(highLowGuessAbsDeltaStr) > 0 {

			lastHighGuessStr = ""

			lastLowGuessStr = ""
			newGuessStr = ""
			lastGuessStr = ""

		}

		/*
			if raisedToPowerAlphaCompareResult == -1 &&
				lastLowAbsAlphaDelta.Cmp(accuracyThresholdDelta) < 1 {

		*/

		if lastLowAbsAlphaDelta.Cmp(accuracyThresholdDelta) < 1 {

			//fmt.Printf("First Guess Exit\n"+
			//	"accuracyThresholdDelta = %v\n"+
			//	"lastLowAbsAlphaDelta   = %v\n",
			//	accuracyThresholdDelta.Text('f', 50),
			//	lastLowAbsAlphaDelta.Text('f', 50))

			guess.Set(lastLowGuess)
			raisedToPower.Set(lastLowRaisedToPower)
			absAlphaDelta.Set(lastLowAbsAlphaDelta)
			raisedToPowerAlphaCompareResult = -1

			closeEnough = true
		}

		if calculationCycleCount > 10000 {

			if raisedToPowerAlphaCompareResult == 1 {
				guess.Set(lastLowGuess)
				raisedToPower.Set(lastLowRaisedToPower)
				absAlphaDelta.Sub(alpha, lastLowRaisedToPower)
				raisedToPowerAlphaCompareResult = -1
			}

			closeEnough = true

		}

		calculationCycleCount++

	}

	return guess,
		raisedToPower,
		absAlphaDelta,
		raisedToPowerAlphaCompareResult,
		initializeCycleCount,
		calculationCycleCount,
		calculationPrec
}

func NewtonInitialGuess03(
	nthRoot int64, // 18
	alpha *big.Float, // 592
	accuracyThresholdDelta *big.Float,
	prec uint) (
	guess *big.Float,
	raisedToPower *big.Float,
	absAlphaDelta *big.Float,
	raisedToPowerAlphaCompareResult int,
	cycleCount uint64,
	calculationPrec uint) {

	funcName := "NewtonInitialGuess03()"

	var ok bool

	calculationPrec = prec + 16

	tempStr := ""

	percent50 := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)
	_,
		ok = percent50.SetString("0.5")

	if !ok {
		fmt.Printf("%v\n"+
			"percent50.SetString(\"0.5\") FAILED!\n"+
			"tempStr= %v\n",
			funcName,
			tempStr)
		return
	}

	tempFloat := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	guess = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPower = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowRaisedToPower := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	absAlphaDelta = new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastHighAbsAlphaDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	highLowGuessAbsDelta := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	raisedToPowerAlphaCompareResult = 0

	cycleCount = 0

	lastHighGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	lastLowGuess := new(big.Float).
		SetInt64(0).
		SetPrec(calculationPrec).
		SetMode(big.AwayFromZero)

	tempStr = fmt.Sprintf("%v",
		alpha.Text('f', -1))

	lenStr := len(tempStr)

	idx := strings.Index(tempStr, ".")

	numFracDigits := 0

	if idx == -1 {

		tempStr += ".0"
		numFracDigits = 1

	} else {

		numFracDigits = lenStr - idx
	}

	tempStr = "0."
	tempStr += strings.Repeat("0", numFracDigits)
	tempStr += "9"

	_,
		ok = guess.SetString(tempStr)

	if !ok {
		fmt.Printf("%v\n"+
			"alpha greater than or equal to '1'"+
			"guess.SetString(tempStr) FAILED!\n"+
			"tempStr= %v\n",
			funcName,
			tempStr)
		return
	}

	// Set the High And Low Guesses
	raisedToPower = BigFloatPower(
		guess,
		nthRoot,
		calculationPrec)

	raisedToPowerAlphaCompareResult =
		raisedToPower.Cmp(alpha)

	tempFloat.SetInt64(1000)

	for raisedToPowerAlphaCompareResult < 1 {

		if raisedToPowerAlphaCompareResult == 0 {

			absAlphaDelta.Sub(raisedToPower, alpha)

			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				cycleCount,
				calculationPrec

		}

		lastLowGuess.Set(guess)
		lastLowAbsAlphaDelta.Sub(alpha, raisedToPower)
		lastLowRaisedToPower.Set(raisedToPower)

		guess.Mul(guess, tempFloat)

		raisedToPower = BigFloatPower(
			guess,
			nthRoot,
			calculationPrec)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		cycleCount++

	}

	lastHighGuess.Set(guess)
	lastHighAbsAlphaDelta.Sub(raisedToPower, alpha)
	lastHighRaisedToPower.Set(raisedToPower)

	guess.Set(lastLowGuess)

	fmt.Printf("%v\n"+
		"Initial Guesses\n"+
		"lastHighGuess = %v\n"+
		"lastLowGuess  = %v\n"+
		"       guess  = %v\n\n",
		funcName,
		lastHighGuess.Text('f', 50),
		lastLowGuess.Text('f', 50),
		guess.Text('f', 50))

	closeEnough := false

	var lastHighGuessStr, lastLowGuessStr,
		guessStr, newGuessStr, lastGuessStr,
		absAlphaDeltaStr, highLowGuessAbsDeltaStr string

	for closeEnough == false {

		if len(guessStr) > 0 {
			guessStr = ""
			newGuessStr = ""
		}

		guessStr = guess.Text('f', 50)

		raisedToPower = BigFloatPower(
			guess,
			nthRoot,
			prec)

		raisedToPowerAlphaCompareResult =
			raisedToPower.Cmp(alpha)

		if raisedToPowerAlphaCompareResult == 0 {
			// Target Located!

			absAlphaDelta.Sub(alpha, raisedToPower)

			// On target
			return guess,
				raisedToPower,
				absAlphaDelta,
				raisedToPowerAlphaCompareResult,
				cycleCount,
				calculationPrec

		} else if raisedToPowerAlphaCompareResult == -1 {
			// Guess is low. Alpha is higher than
			// raisedToPower. Increase Guess.
			// Last High Guess is higher than
			//	Last Low Guess

			absAlphaDelta.Sub(alpha, raisedToPower)

			absAlphaDeltaStr =
				absAlphaDelta.Text('f', 50)

			highLowGuessAbsDelta.Sub(lastHighGuess, guess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			highLowGuessAbsDeltaStr =
				highLowGuessAbsDelta.Text('f', 50)

			lastLowGuess.Set(guess)
			lastLowAbsAlphaDelta.Set(absAlphaDelta)
			lastLowRaisedToPower.Set(raisedToPower)

			guess.Sub(lastHighGuess, highLowGuessAbsDelta)

			lastGuessStr = guessStr
			newGuessStr = guess.Text('f', 50)

			lastLowGuessStr =
				lastLowGuess.Text('f', 50)

			tempStr = ""

			//fmt.Printf("Last Low Guess: %v\n",
			//	lastLowGuessStr)

		} else {
			// MUST BE -
			//	Guess is high. raisedToPower is higher
			//	than Alpha. Reduce Guess.

			absAlphaDelta.Sub(raisedToPower, alpha)

			absAlphaDeltaStr =
				absAlphaDelta.Text('f', 50)

			highLowGuessAbsDelta.Sub(guess, lastLowGuess)

			highLowGuessAbsDelta.Mul(highLowGuessAbsDelta, percent50)

			highLowGuessAbsDeltaStr =
				highLowGuessAbsDelta.Text('f', 50)

			lastLowGuessStr =
				lastLowGuess.Text('f', 50)

			lastHighGuess.Set(guess)
			lastHighAbsAlphaDelta.Set(absAlphaDelta)
			lastHighRaisedToPower.Set(raisedToPower)
			guess.Add(lastLowGuess, highLowGuessAbsDelta)

			lastGuessStr = guessStr
			newGuessStr = guess.Text('f', 50)

			lastHighGuessStr =
				lastHighGuess.Text('f', 50)

			tempStr = ""
			//	fmt.Printf("Last High Guess: %v\n",
			//		lastHighGuessStr)

		}

		if len(lastHighGuessStr) > 0 ||
			len(lastLowGuessStr) > 0 ||
			len(newGuessStr) > 0 ||
			len(lastGuessStr) > 0 ||
			len(absAlphaDeltaStr) > 0 ||
			len(highLowGuessAbsDeltaStr) > 0 {

			lastHighGuessStr = ""

			lastLowGuessStr = ""
			newGuessStr = ""
			lastGuessStr = ""

		}

		if raisedToPowerAlphaCompareResult == -1 &&
			absAlphaDelta.Cmp(accuracyThresholdDelta) < 1 {

			guess.Set(lastLowGuess)
			raisedToPower.Set(lastLowRaisedToPower)
			absAlphaDelta.Sub(alpha, lastLowRaisedToPower)
			raisedToPowerAlphaCompareResult = -1

			closeEnough = true
		}

		if cycleCount > 10000 {

			if raisedToPowerAlphaCompareResult == 1 {
				guess.Set(lastLowGuess)
				raisedToPower.Set(lastLowRaisedToPower)
				absAlphaDelta.Sub(alpha, lastLowRaisedToPower)
				raisedToPowerAlphaCompareResult = -1
			}

			closeEnough = true

		}

		cycleCount++

	}

	return guess,
		raisedToPower,
		absAlphaDelta,
		raisedToPowerAlphaCompareResult,
		cycleCount,
		calculationPrec
}

func NewtonInitialGuess02(
	nthRoot int64, // 18
	alpha *big.Float, // 592
	accuracyThresholdDelta *big.Float,
	prec uint) (
	guess *big.Float,
	raisedToPower *big.Float,
	absAlphaDelta *big.Float,
	compareResult int,
	cycleCount uint64) {

	funcName := "NewtonInitialGuess02()"

	var ok bool

	tempCompare := 0

	tempFloat := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	tStrAbsAlphaDelta := ""
	tStrAbsPercentAlphaDelta := ""
	tStrRaisedToPower := ""
	tStrGuess := ""
	localOutputPrec := 30

	numberZero := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	numberZero.SetString("0.0")

	numberOne := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	numberOne.SetString("1")

	guess = new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	alphaStr := fmt.Sprintf("%v",
		alpha.Text('f', -1))

	decIndex := strings.Index(alphaStr, ".")

	if decIndex == -1 {
		alphaStr += ".0"
	}

	_,
		fractionalDigits,
		ok := strings.Cut(alphaStr, ".")

	if !ok {
		fmt.Printf("\n%v\n"+
			"ERROR: strings.Cut(alphaStr, \".\") FAILED!\n"+
			"alphaStr = '%v'\n",
			funcName,
			alphaStr)

		return
	}

	numOfFractionalDigits := len(fractionalDigits)

	guessFractionalDigits := strings.Repeat("0", numOfFractionalDigits+50)
	guessFractionalDigits += "5"

	var guessStr string

	if alpha.Cmp(numberOne) > -1 {

		guessStr = "1."

	} else {

		guessStr = "0."

	}

	guessStr += guessFractionalDigits

	guess.SetString(guessStr)

	fmt.Printf("Starting Guess: %v\n\n",
		guess.Text('f', -1))

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

	lastGuess := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

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

	percent1000 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent1000.SetString("100.00")

	percent600 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent600.SetString("6.00")

	percent400 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent400.SetString("4.00")

	percent200 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent200.SetString("2.00")

	percent100 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent100.SetString("1.00")

	percent90 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent90.SetString("0.9")

	percent75 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent75.SetString("0.75")

	percent65 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent65.SetString("0.65")

	percent50 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent50.SetString("0.50")

	percent35 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent35.SetString("0.35")

	percent25 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent25.SetString("0.25")

	percent15 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent15.SetString("0.15")

	percent10 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent10.SetString("0.10")

	percent05 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent05.SetString("0.05")

	percent03 := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	percent03.SetString("0.03")

	deltaIncrement := new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	deltaIncrement.SetString("0.0")

	cycleCount = uint64(0)

	closeEnough := false

	deltaIncrementValue := "0.5"

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
			// raisedToPower. Increase Guess.

			lastLowGuess.Set(guess)

			lastGuess.Set(guess)

			lastLowRaisedToPower.Set(raisedToPower)

			absAlphaDelta.Sub(alpha, lastLowRaisedToPower)

			absPercentAlphaDelta.Quo(absAlphaDelta, alpha)

			tStrGuess =
				guess.Text('f', localOutputPrec)

			tStrRaisedToPower =
				raisedToPower.Text('f', localOutputPrec)

			tStrAbsAlphaDelta = absAlphaDelta.Text('f', localOutputPrec)

			tStrAbsPercentAlphaDelta =
				absPercentAlphaDelta.Text('f', localOutputPrec)

			if len(tStrAbsAlphaDelta) > 0 ||
				len(tStrAbsPercentAlphaDelta) > 0 ||
				len(tStrRaisedToPower) > 0 ||
				len(tStrGuess) > 0 {
				ok = false
			}

			//fmt.Printf("lastLowGuess  = %v\n"+
			//	"lastLowRaisedToPower = %v\n"+
			//	"       absAlphaDelta = %v\n"+
			//	"absPercentAlphaDelta = %v\n",
			//	lastLowGuess.Text('f', 20),
			//	lastLowRaisedToPower.Text('f', 20),
			//	absAlphaDelta.Text('f', 20),
			//	absPercentAlphaDelta.Text('f', 20))

			deltaIncrement.Set(half)

			deltaIncrementValue = "0.5"

			if numOfLowGuesses > 0 {

				if absPercentAlphaDelta.Cmp(percent1000) == 1 {

					deltaIncrementValue = "100.0"

				} else if absPercentAlphaDelta.Cmp(percent600) == 1 {

					deltaIncrementValue = "10.0"

				} else if absPercentAlphaDelta.Cmp(percent400) == 1 {

					deltaIncrementValue = "5.0"

				} else if absPercentAlphaDelta.Cmp(percent200) == 1 {

					deltaIncrementValue = "3.0"

				} else if absPercentAlphaDelta.Cmp(percent100) == 1 {

					deltaIncrementValue = "2.0"

				} else if absPercentAlphaDelta.Cmp(percent75) == 1 {

					deltaIncrementValue = "1.0"

				} else if absPercentAlphaDelta.Cmp(percent65) == 1 {

					deltaIncrementValue = "0.75"

				} else if absPercentAlphaDelta.Cmp(percent50) == 1 {

					deltaIncrementValue = "0.50"

				} else if absPercentAlphaDelta.Cmp(percent35) == 1 {

					deltaIncrementValue = "0.50"

				} else if absPercentAlphaDelta.Cmp(percent25) == 1 {

					deltaIncrementValue = "0.35"

				} else if absPercentAlphaDelta.Cmp(percent15) == 1 {

					deltaIncrementValue = "0.30"

				} else if absPercentAlphaDelta.Cmp(percent10) == 1 {

					deltaIncrementValue = "0.25"

				} else if absPercentAlphaDelta.Cmp(percent05) == 1 {

					deltaIncrementValue = "0.15"

				} else if absPercentAlphaDelta.Cmp(percent03) == 1 {

					deltaIncrementValue = "0.10"

				} else {

					deltaIncrementValue = "0.05"

				}

			}

			lastLowAbsAlphaDelta.Set(absAlphaDelta)

			_,
				ok = deltaIncrement.SetString(deltaIncrementValue)

			if !ok {
				panic("#1 " +
					"deltaIncrement.SetString(deltaIncrementValue) " +
					"FAILED!\n")
			}

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

			if guess.Cmp(numberZero) < 1 {

				tStr := "0."

				tStr += strings.Repeat("0", 100)

				tStr += "5"

				_,
					ok = guess.SetString(tStr)

				if !ok {
					panic("High #5 guess.SetString(tStr) Failed!\n")
				}

				lastLowGuess.Mul(guess, percent90)
				lastGuess.Set(guess)
				tempFloat.SetFloat64(1.15)
				lastHighGuess.Mul(guess, tempFloat)

			}

			//fmt.Printf("Old Guess = %v\n",
			//	guess.Text('f', -1))
			//
			//fmt.Printf("deltaIncrement = %v\n",
			//	deltaIncrement.Text('f', -1))
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
			//	than Alpha. Reduce Guess.

			lastHighGuess.Set(guess)

			lastHighRaisedToPower.Set(raisedToPower)

			absAlphaDelta.Sub(lastHighRaisedToPower, alpha)

			absPercentAlphaDelta.Quo(absAlphaDelta, alpha)

			deltaIncrement.Set(half)

			deltaIncrementValue = "0.5"

			if numOfHighGuesses > 0 {

				deltaIncrementValue = "0.5"

				if absPercentAlphaDelta.Cmp(percent600) == 1 {

					deltaIncrementValue = "10.0"

				} else if absPercentAlphaDelta.Cmp(percent400) == 1 {

					deltaIncrementValue = "5.0"

				} else if absPercentAlphaDelta.Cmp(percent200) == 1 {

					deltaIncrementValue = "3.0"

				} else if absPercentAlphaDelta.Cmp(percent100) == 1 {

					deltaIncrementValue = "2.0"

				} else if absPercentAlphaDelta.Cmp(percent75) == 1 {

					deltaIncrementValue = "1.0"

				} else if absPercentAlphaDelta.Cmp(percent65) == 1 {

					deltaIncrementValue = "0.75"

				} else if absPercentAlphaDelta.Cmp(percent50) == 1 {

					deltaIncrementValue = "0.50"

				} else if absPercentAlphaDelta.Cmp(percent35) == 1 {

					deltaIncrementValue = "0.50"

				} else if absPercentAlphaDelta.Cmp(percent25) == 1 {

					deltaIncrementValue = "0.35"

				} else if absPercentAlphaDelta.Cmp(percent15) == 1 {

					deltaIncrementValue = "0.30"

				} else if absPercentAlphaDelta.Cmp(percent10) == 1 {

					deltaIncrementValue = "0.25"

				} else if absPercentAlphaDelta.Cmp(percent05) == 1 {

					deltaIncrementValue = "0.15"

				} else if absPercentAlphaDelta.Cmp(percent03) == 1 {

					deltaIncrementValue = "0.10"

				} else {

					deltaIncrementValue = "0.05"

				}

			}

			//fmt.Printf("lastHighGuess  = %v\n"+
			//	"lastHighRaisedToPower = %v\n",
			//	lastHighGuess.Text('f', 10),
			//	lastHighRaisedToPower.Text('f', 10))

			lastHighAbsAlphaDelta.Set(absAlphaDelta)

			_,
				ok = deltaIncrement.SetString(deltaIncrementValue)

			if !ok {
				panic("#2 " +
					"deltaIncrement.SetString(deltaIncrementValue) " +
					"FAILED!\n")
			}

			tempCompare = lastLowGuess.Cmp(lastHighGuess)

			if tempCompare == 1 {
				// Low Guess is higher than High Guess
				absGuessDelta.Mul(lastHighGuess, deltaIncrement)
				guess.Sub(guess, absGuessDelta)

			} else if tempCompare == 0 {
				// Low Guess is equal to High Guess
				absGuessDelta.Mul(lastHighGuess, deltaIncrement)
				guess.Sub(guess, absGuessDelta)

			} else {
				// High Guess is higher than Low Guess
				absGuessDelta.Sub(lastHighGuess, lastLowGuess)

				absGuessDelta.Mul(absGuessDelta, deltaIncrement)
			}

			guess.Sub(guess, absGuessDelta)

			if guess.Cmp(numberZero) < 1 {

				tStr := "0."

				tStr += strings.Repeat("0", 100)

				tStr += "5"

				_,
					ok = guess.SetString(tStr)

				if !ok {
					panic("High #5 guess.SetString(tStr) Failed!\n")
				}

				lastLowGuess.Set(guess)

				lastHighGuess.Set(guess)
			}

			//fmt.Printf("New Guess = %v\n",
			//	guess.Text('f', -1))
			//
			//fmt.Printf("deltaIncrement = %v\n",
			//	deltaIncrement.Text('f', -1))
			//
			//fmt.Printf("absGuessDelta = %v\n",
			//	absGuessDelta.Text('f', -1))

			numOfHighGuesses++
		}

		cycleCount++

		if compareResult == -1 &&
			absAlphaDelta.Cmp(accuracyThresholdDelta) < 1 {

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

func NewtonInitialGuess01(
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
	raisedToPower *big.Float) {

	raisedToPower =
		new(big.Float).
			SetPrec(prec).
			SetMode(big.AwayFromZero)

	var ok bool
	_,
		ok = raisedToPower.SetString("1.0")

	if !ok {
		panic("raisedToPower.SetString(\"1.0\") Failed!")
	}

	for i := int64(0); i < power; i++ {

		raisedToPower.Mul(raisedToPower, base)
	}

	return raisedToPower
}

func raiseToPowerInt(
	base *big.Float,
	exponent int64) (
	resultsStats BigFloatCalcStats,
	err error) {

	funcName := "raiseToPowerInt()"

	//breakStr := strings.Repeat("=", 70)
	//
	//fmt.Printf("\n\nFunction: %v\n",
	//	funcName)
	//
	//fmt.Printf(breakStr + "\n\n")

	// var pBase *big.Float

	base.SetPrec(base.MinPrec())

	resultsStats.BaseFloatNum = *base
	resultsStats.BasePrec = base.Prec()

	numStr := base.Text('f', -1)

	var pureNumStrStats PureNumberStrComponents

	pureNumStrStats,
		err = breakNumStrToComponents(
		numStr)

	if err != nil {
		return resultsStats, err
	}

	if exponent < 0 {

		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'exponent' is INVALID!\n"+
			"'exponent' has a value less than zero.\n"+
			"exponent = %v\n",
			funcName,
			exponent)

		return resultsStats, err
	}

	resultsStats.BaseNumIntDigits =
		pureNumStrStats.NumIntegerDigits

	resultsStats.BaseNumFracDigits =
		pureNumStrStats.NumFractionalDigits

	var ok bool

	bigIntBase,
		ok := big.NewInt(0).
		SetString(
			pureNumStrStats.AllIntegerDigitsNumStr,
			10)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error: bigIntBase=SetString(AllIntegerDigitsNumStr)\n"+
			"SetString Failed!\n"+
			"AllIntegerDigitsNumStr = %v\n",
			funcName,
			pureNumStrStats.AllIntegerDigitsNumStr)

		return resultsStats, err
	}

	bigIntExponent := big.NewInt(
		exponent)

	raisedToPowerFracDigits :=
		exponent * pureNumStrStats.NumFractionalDigits

	bigIntBase.Exp(bigIntBase, bigIntExponent, nil)

	numStr =
		bigIntBase.Text(10)

	lenNumStrInt64 := int64(len(numStr))

	if raisedToPowerFracDigits == 0 {

		resultsStats.CalcResultNumIntDigits =
			lenNumStrInt64

		resultsStats.CalcResultNumFracDigits = 0

	} else {

		resultsStats.CalcResultNumIntDigits =
			lenNumStrInt64 - raisedToPowerFracDigits

		resultsStats.CalcResultNumFracDigits =
			raisedToPowerFracDigits

		numStr =
			numStr[0:resultsStats.CalcResultNumIntDigits] +
				"." +
				numStr[resultsStats.CalcResultNumIntDigits:]

	}

	var calcPrecision uint

	calcPrecision,
		err = calculateRequiredPrecision(
		resultsStats.CalcResultNumIntDigits,
		resultsStats.CalcResultNumFracDigits,
		5)

	if err != nil {
		return resultsStats, err
	}

	_,
		ok = resultsStats.CalcResult.
		SetPrec(calcPrecision).
		SetMode(big.AwayFromZero).
		SetString(numStr)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error: CalcResult=SetString(numStr)\n"+
			"SetString Failed!\n"+
			"numStr = %v\n",
			funcName,
			numStr)

	}

	resultsStats.CalcResult.SetPrec(
		resultsStats.CalcResult.MinPrec())

	resultsStats.CalcResultPrec =
		resultsStats.CalcResult.Prec()

	//fmt.Printf("\n\n\t\t%v\n"+
	//	"\tRaise To Exponent Results\n"+
	//	"                   Base = %v\n"+
	//	"        Raw Base NumStr = %v\n"+
	//	"Original Base Precision = %v\n"+
	//	"         Base Precision = %v\n"+
	//	"    Base Integer Digits = %v\n"+
	//	" Base Fractional Digits = %v\n"+
	//	"               Exponent = %v\n"+
	//	"     Raw Results NumStr = %v\n"+
	//	"     Calculation Result = %v\n"+
	//	"        Expected Result = %v\n"+
	//	"  Calc Result Precision = %v\n"+
	//	" Calc Result Int Digits = %v\n"+
	//	"Calc Result Frac Digits = %v\n\n",
	//	funcName,
	//	base.Text('f', -1),
	//	rawBaseNumStr,
	//	originalBasePrec,
	//	base.Prec(),
	//	resultsStats.BaseNumIntDigits,
	//	resultsStats.BaseNumFracDigits,
	//	exponent,
	//	rawResultsStr,
	//	resultsStats.CalcResult.Text('f', -1),
	//	expectedResult,
	//	resultsStats.CalcResult.Prec(),
	//	resultsStats.CalcResultNumIntDigits,
	//	resultsStats.CalcResultNumFracDigits)
	//
	//fmt.Printf("\n\n%v\n"+
	//	"   Successful Completion!\n"+
	//	"Function: %v\n%v\n\n",
	//	breakStr,
	//	funcName,
	//	breakStr)

	return resultsStats, err
}

func breakNumStrToComponents(
	numberString string) (
	pureNumStrStats PureNumberStrComponents,
	err error) {

	funcName := ""

	lenNumberStr := len(numberString)

	if lenNumberStr == 0 {
		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'numberString'\n"+
			"is a zero length string and INVALID!\n",
			funcName)

		return pureNumStrStats, err
	}

	pureNumStrStats.NumberSign = 0
	pureNumStrStats.NumberType = 2

	if numberString[0] == '-' {

		pureNumStrStats.NumberSign = -1

		numberString = numberString[1:]

		lenNumberStr--

		if lenNumberStr == 0 {
			err = fmt.Errorf("\n\n%v\n"+
				"Error: Input parameter 'numberString'\n"+
				"is a zero length string and INVALID!\n",
				funcName)

			return pureNumStrStats, err
		}

	} else {

		pureNumStrStats.NumberSign = 1

	}

	idx := strings.Index(numberString, ".")

	if idx == -1 {

		pureNumStrStats.NumberType = 1

		pureNumStrStats.NumIntegerDigits =
			int64(lenNumberStr)

		pureNumStrStats.NumFractionalDigits =
			0

		pureNumStrStats.AbsoluteValueNumStr =
			numberString

		pureNumStrStats.AllIntegerDigitsNumStr =
			numberString

	} else {

		pureNumStrStats.NumberType = 2

		pureNumStrStats.NumIntegerDigits =
			int64(idx)

		pureNumStrStats.NumFractionalDigits =
			int64(lenNumberStr - (idx + 1))

		pureNumStrStats.AbsoluteValueNumStr =
			numberString

		pureNumStrStats.AllIntegerDigitsNumStr =
			numberString[0:idx] +
				numberString[idx+1:]
	}
	//
	//fmt.Printf("\n%v\n"+
	//	"                          numberString = %v\n"+
	//	"   pureNumStrStats.AbsoluteValueNumStr = %v\n"+
	//	"pureNumStrStats.AllIntegerDigitsNumStr = %v\n",
	//	funcName,
	//	numberString,
	//	pureNumStrStats.AbsoluteValueNumStr,
	//	pureNumStrStats.AllIntegerDigitsNumStr)

	isZeroValue := true
	lenNumberStr = len(pureNumStrStats.AbsoluteValueNumStr)

	for i := 0; i < lenNumberStr; i++ {

		if pureNumStrStats.AbsoluteValueNumStr[i] >= '0' &&
			pureNumStrStats.AbsoluteValueNumStr[i] <= '9' {

			isZeroValue = false

			break
		}
	}

	if isZeroValue {
		pureNumStrStats.NumberSign = 0
	}

	return pureNumStrStats, err
}
