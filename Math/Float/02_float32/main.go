package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	float64Tests02()
}

func float64Tests01() {

	funcName := "float64Tests01"

	fmt.Printf("Function %v\n",
		funcName)

	breakStr := strings.Repeat("=", 50)

	fmt.Printf(breakStr + "\n\n")

	numberStr := "1234567890.12345678901234567890"

	fmt.Printf("Original Number String:         %v\n",
		numberStr)

	f64,
		err := strconv.ParseFloat(numberStr, 64)

	if err != nil {
		fmt.Printf("\n%v-Error\n"+
			"Error return from strconv.ParseFloat(numberStr, 64)\n"+
			"numberStr='%v'\n"+
			"Error = \n%v\n",
			funcName,
			numberStr,
			err.Error())

		return
	}

	fmt.Printf("Float 64 Conversion Value .20f: %.20f\n",
		f64)

	fmt.Printf(breakStr + "\n\n")

	numberStr2 := "1.12345678901234567890"

	fmt.Printf("Original Number2 Str:           %v\n",
		numberStr2)

	f64,
		err = strconv.ParseFloat(numberStr2, 64)

	if err != nil {
		fmt.Printf("\n%v-Error\n"+
			"Error return from strconv.ParseFloat(numberStr2, 64)\n"+
			"numberStr='%v'\n"+
			"Error = \n%v\n",
			funcName,
			numberStr,
			err.Error())

		return
	}

	fmt.Printf("Float 64 numberStr2 .20f:       %.20f\n",
		f64)

	fmt.Printf(breakStr + "\n\n")

	numberStr3 := "999999.12345678901234567890"

	fmt.Printf("Original Number3 Str:           %v\n",
		numberStr3)

	f64,
		err = strconv.ParseFloat(numberStr3, 64)

	if err != nil {
		fmt.Printf("\n%v-Error\n"+
			"Error return from strconv.ParseFloat(numberStr3, 64)\n"+
			"numberStr='%v'\n"+
			"Error = \n%v\n",
			funcName,
			numberStr,
			err.Error())

		return
	}

	fmt.Printf("Float 64 numberStr3 .20f:       %.20f\n",
		f64)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

	// Printed Output

	//	Function float64Tests01
	//	==================================================
	//
	//	Original Number String:         1234567890.12345678901234567890
	//	Float 64 Conversion Value .20f: 1234567890.12345671653747558594
	//	==================================================
	//
	//	Original Number2 Str:           1.12345678901234567890
	//	Float 64 numberStr2 .20f:       1.12345678901234569125
	//	==================================================
	//
	//	Original Number3 Str:           999999.12345678901234567890
	//	Float 64 numberStr3 .20f:       999999.12345678906422108412
	//
	//	==================================================

	//
	//	Float 64 Capacity is approximately 15 to 17 digits
	//		including integer and fractional digits.

	// https://en.wikipedia.org/wiki/Floating-point_arithmetic#IEEE_754:_floating_point_in_modern_computers
	// Type		Sign  Exponent	Significand	  Total   Exponent	 Bits	     Number of
	//							   field       bits     bias    precision	decimal digits
	//          ----  --------	-----------   -----   --------  ---------	--------------
	// Single	 1		8			23			32		  127		24			~7.2
	// Double	 1		11			52			64		  1023		53		   ~15.9

}

func float64Tests02() {

	funcName := "float64Tests02"

	fmt.Printf("Function %v\n",
		funcName)

	breakStr := strings.Repeat("=", 50)

	fmt.Printf(breakStr + "\n\n")

	numberStr := "12.1234567890"

	float64Num := 12.1234567890

	fmt.Printf("Original Number String:         %v\n",
		numberStr)

	float64Str := strconv.FormatFloat(
		float64Num, 'f', -1, 64)

	fmt.Printf("float64Str:             %v\n",
		float64Str)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}

func float32Tests01() {

	funcName := "float32Tests01"

	fmt.Printf("Function %v\n",
		funcName)

	breakStr := strings.Repeat("=", 50)

	fmt.Printf(breakStr + "\n\n")

	numberStr := "1234567890.12345678901234567890"

	fmt.Printf("Original Number String:         %v\n",
		numberStr)

	f64,
		err := strconv.ParseFloat(numberStr, 32)

	if err != nil {
		fmt.Printf("\n%v-Error\n"+
			"Error return from strconv.ParseFloat(numberStr, 32)\n"+
			"numberStr='%v'\n"+
			"Error = \n%v\n",
			funcName,
			numberStr,
			err.Error())

		return
	}

	fmt.Printf("Float 32 Conversion Value .20f: %.20f\n",
		float32(f64))

	fmt.Printf(breakStr + "\n\n")

	numberStr2 := "1.12345678901234567890"

	fmt.Printf("Original Number2 Str:           %v\n",
		numberStr2)

	f64,
		err = strconv.ParseFloat(numberStr2, 32)

	if err != nil {
		fmt.Printf("\n%v-Error\n"+
			"Error return from strconv.ParseFloat(numberStr2, 32)\n"+
			"numberStr='%v'\n"+
			"Error = \n%v\n",
			funcName,
			numberStr,
			err.Error())

		return
	}

	fmt.Printf("Float 32 numberStr2 .20f:       %.20f\n",
		float32(f64))

	fmt.Printf(breakStr + "\n\n")

	numberStr3 := "999999.12345678901234567890"

	fmt.Printf("Original Number3 Str:           %v\n",
		numberStr3)

	f64,
		err = strconv.ParseFloat(numberStr3, 32)

	if err != nil {
		fmt.Printf("\n%v-Error\n"+
			"Error return from strconv.ParseFloat(numberStr3, 32)\n"+
			"numberStr='%v'\n"+
			"Error = \n%v\n",
			funcName,
			numberStr,
			err.Error())

		return
	}

	fmt.Printf("Float 32 numberStr3 .20f:       %.20f\n",
		float32(f64))

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

	// Printed Output

	//	Function float32Tests01
	//	==================================================
	//
	//	Original Number String:         1234567890.12345678901234567890
	//	Float 32 Conversion Value .20f: 1234567936.00000000000000000000
	//	==================================================
	//
	//	Original Number2 Str:           1.12345678901234567890
	//	Float 32 numberStr2 .20f:       1.12345683574676513672
	//	==================================================
	//
	//	Original Number3 Str:           999999.12345678901234567890
	//	Float 32 numberStr3 .20f:       999999.12500000000000000000
	//
	//	==================================================
	//

	//
	//	Float 32 Capacity is approximately 7 to 8 digits
	//		including integer and fractional digits.

	// https://en.wikipedia.org/wiki/Floating-point_arithmetic#IEEE_754:_floating_point_in_modern_computers
	// Type		Sign  Exponent	Significand	  Total   Exponent	 Bits	     Number of
	//							   field       bits     bias    precision	decimal digits
	//          ----  --------	-----------   -----   --------  ---------	--------------
	// Single	 1		8			23			32		  127		24			~7.2
	// Double	 1		11			52			64		  1023		53		   ~15.9

}
