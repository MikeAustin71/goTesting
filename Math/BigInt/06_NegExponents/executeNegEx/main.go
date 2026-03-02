package main

import (
	"fmt"
	"math/big"

	bi "github.com/mikeaustin71/Math/BigInt/06_NegExponents/bigIntEx"
)

func main() {

	testBigRatAlgorithm()

	return
}

func testBigRatAlgorithm() {

	fmt.Printf("%v\n" +
		"Executing Big Int - Big Rat tests...\n\n")

	numerator := big.NewInt(1)

	denominator := big.NewInt(50653)

	err := new(bi.BigIntMathPowerMolecule).BigRat(
		numerator, denominator)

	if err != nil {
		fmt.Printf("Error= %v\n", err)
	}

	return
}

func testDivisionAlgorithm() {

	fmt.Printf("%v\n" +
		"Executing Big Int Division tests...\n\n")

	numerator := big.NewInt(1)

	denominator := big.NewInt(50653)

	err := new(bi.BigIntMathPowerMolecule).BigIntDivide(numerator, denominator)

	if err != nil {
		fmt.Printf("Error= %v\n", err)
	}

	return
}

func testRaiseToNegativeFractionalPower() {

	fmt.Printf("%v\n" +
		"Executing negative exponent tests...\n\n")

	base := big.NewInt(37)

	exponent := big.NewInt(-3)

	err := new(bi.BigIntMathPowerMolecule).BigIntNumRaiseToNegativeFractionalPower(
		base, exponent)

	if err != nil {
		fmt.Printf("Error= %v\n", err)
	}

	return
}
