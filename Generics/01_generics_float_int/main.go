package main

import (
	"fmt"
	"math/big"
)

type IInputParamToStr interface {
	*big.Int | *big.Float | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 | ~string | ~bool
}

func main() {
	MainSub01()
}

func MainSub01() {

	fmt.Printf("\n01_generics_float_int\n")

	var str string

	fmt.Printf("Testing 'int32'\n")
	var i32 int32

	i32 = 97843

	str = ParameterToStr(i32)

	fmt.Printf("Result: %s\n",
		str)

	str = ParamToStr2(i32)

	fmt.Printf("Type Test: %s\n\n",
		str)

	fmt.Printf("Testing 'int64'\n")

	var i64 int64

	i64 = int64(1234567890)

	str = ParameterToStr(i64)

	fmt.Printf("Result: %s\n",
		str)

	str = ParamToStr2(i64)

	fmt.Printf("Type Test: %s\n\n",
		str)

	// Testing big.Int

	fmt.Printf("Testing 'big.Int'\n")

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

	str = ParameterToStr(bigI)

	fmt.Printf("Result: %s\n",
		str)

	str = ParamToStr2(bigI)

	fmt.Printf("Type Test: %s\n\n",
		str)

	fmt.Printf("Testing 'float32'\n")
	var f32 float32
	f32 = 1.2345678

	str = ParameterToStr(f32)

	fmt.Printf("Result: %s\n",
		str)

	str = ParamToStr2(f32)

	fmt.Printf("Type Test: %s\n\n",
		str)

	fmt.Printf("Testing 'float64'\n")
	var f64 float64
	f64 = 9.0123456789012

	str = ParameterToStr(f64)

	fmt.Printf("Result: %s\n",
		str)

	str = ParamToStr2(f64)

	fmt.Printf("Type Test: %s\n\n",
		str)

	// -----------------------------------------
	// big.Float

	fmt.Printf("Testing 'big.Float'\n")

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

	str = ParameterToStr(bigF)

	fmt.Printf("Result: %s\n",
		str)

	str = ParamToStr2(bigF)

	fmt.Printf("Type Test: %s\n\n",
		str)

	var boolTest bool
	boolTest = true

	fmt.Printf("Testing 'bool'\n")
	str = ParameterToStr(boolTest)

	fmt.Printf("Result: %s\n",
		str)

	str = ParamToStr2(boolTest)

	fmt.Printf("Type Test: %s\n\n",
		str)

	var testStr string

	fmt.Printf("Testing 'string'\n")

	testStr = "How now brown cow!"
	str = ParameterToStr(testStr)

	fmt.Printf("Result: %s\n",
		str)

	str = ParamToStr2(testStr)

	fmt.Printf("Type Test: %s\n\n",
		str)

	fmt.Printf("Testing Two Input Paramters - 1-Generic 1-Not\n")

	str = Param2ToStr(bigI, 5)

	fmt.Printf("Result: %s\n",
		str)

}

func ParameterToStr[T IInputParamToStr](param T) string {

	str := fmt.Sprintf("%v",
		param)

	return str
}

func ParamToStr2[T IInputParamToStr](param T) string {

	str := fmt.Sprintf("'param' is type: %T\n",
		param)

	return str
}

func Param2ToStr[T IInputParamToStr](param T, value int) string {

	str := fmt.Sprintf("paramT = %v\n"+
		"2nd Param value = %v\n",
		param,
		value)

	return str
}
