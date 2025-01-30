package main

import (
	"fmt"
	"math/big"
)

func main() {

	/*	Value: 123.456
		Format	Verb	Description
		1.234560e+02	%e	Scientific notation
		123.456000	%f	Decimal point, no exponent
		123.46	%.2f	Default width, precision 2
		␣␣123.46	%8.2f	Width 8, precision 2
		123.456	%g	Exponent as needed, necessary digits only
	*/

	//Float32ToString()

	//Float32NegToString()

	Float64ToString()
}

func Float32ToString() {
	f := float32(1.349)

	fmt.Printf("\n\n")
	fmt.Printf("Float32ToString\n")
	fmt.Printf("'f' Format:  %f\n",
		f)
	fmt.Printf("'v' Format:  %v\n",
		f)

}

func Float32NegToString() {
	f := float32(-1.349)

	fmt.Printf("\n\n")
	fmt.Printf("Float32ToString\n")
	fmt.Printf("'f' Format:  %f\n",
		f)
	fmt.Printf("'v' Format:  %v\n",
		f)

}

func Float64ToString() {

	var f float64

	f = 9.123456789012345

	fmt.Printf("\n\n")
	fmt.Printf("Float64ToString\n")
	fmt.Printf("Postive Float64 To String\n")
	fmt.Printf("'f' Format:  %f\n",
		f)
	fmt.Printf("'v' Format:  %v\n",
		f)

	f = -9.123456789012345

	fmt.Printf("\n\n")
	fmt.Printf("Float64ToString\n")
	fmt.Printf("Negative Float64 To String\n")
	fmt.Printf("'f' Format:  %f\n",
		f)
	fmt.Printf("'v' Format:  %v\n",
		f)

}

func BigFloatToString() {

	var flag bool

	flag = true

	fmt.Printf("\nBoolean Value to String\n")
	fmt.Printf("'v' Format: %v\n",
		flag)

	flag = false
	fmt.Printf("'v' Format: %v\n",
		flag)

	bigF := big.NewFloat(0)

	bigF.SetPrec(100)
	bigF.SetMode(big.ToNearestAway)

	var actualBase int

	var err error

	bigF,
		actualBase,
		err =
		bigF.Parse("1.2345679112345678921234567893", 10)

	if err != nil {
		fmt.Printf("Error: \n"+
			"bigF.Parse(\"1.234567901234567890\", 10)\n"+
			"%v\n",
			err.Error())

		return
	}

	if actualBase != 10 {
		fmt.Printf("Error: \n"+
			"bigF.Parse(\"1.234567901234567890\", 10)\n"+
			"'actualBase' is NOT EQUAL to '10'\n"+
			"actualBase = %v\n",
			actualBase)

		return
	}

	fmt.Printf("\nBig Float Value to String\n")
	fmt.Printf("'v' Format: %v\n\n",
		bigF)

	fmt.Printf("\nTesting Big Int\n")
	bigI := big.NewInt(0)
	var ok bool

	bigI,
		ok =
		bigI.SetString("123456789112345678921234567893", 10)

	if !ok {
		fmt.Printf("Error: \n" +
			"bigI.SetString(\"123456789112345678921234567893\", 10)\n" +
			"This String Parsing Function FAILED!\n")

		return
	}

	fmt.Printf("\nBig Float Value to String\n")
	fmt.Printf("'v' Format: %v\n\n",
		bigI)

}
