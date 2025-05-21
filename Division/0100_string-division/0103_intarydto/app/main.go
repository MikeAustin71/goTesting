package main

import (
	"fmt"
	"golangmikesamples/Division/0100_string-division/0103_intarydto/common"
)

/*
import (
	"bitbucket.org/AmarilloMike/GolangMikeSamples/Division/0100_string-division/0103_intarydto/common"
	"fmt"
)
D:\go\work\src\bitbucket.org\AmarilloMike\GolangMikeSamples\Division\0100_string-division\0103_intarydto\app\main.go
*/

func main() {

	t02_SetSigDigits()

}

func t02_SetSigDigits() {
	nStr := "002750.000060"
	eAryLen := 12
	eNumStr := "002750.000060"
	eIntegerLen := 6
	eSigIntegerLen := 4
	eSigFractionLen := 5
	eIsZeroValue := false
	eIsIntegerZeroValue := false
	eSignVal := 1
	ePrecision := 6

	/*
		nStr := "082.7770"
		eAryLen := 7
		eNumStr := "082.7770"
		eIntegerLen := 3
		eSigIntegerLen := 2
		eSigFractionLen := 3
		eIsZeroValue := false
		eIsIntegerZeroValue := false
		eSignVal := 1
		ePrecision := 4
	*/
	ia := common.IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		fmt.Printf("Error from ia.SetIntAryWithNumStr(nStr). Error= %v \n", err)
	}

	fmt.Println("Expected Values")
	fmt.Println("Expected NumStr: ", eNumStr)
	fmt.Println("Expected IntAryLen: ", eAryLen)
	fmt.Println("Expected Integer Len: ", eIntegerLen)
	fmt.Println("Expected SignificantIntegerLen: ", eSigIntegerLen)
	fmt.Println("Expected SignificantFractionLen: ", eSigFractionLen)
	eIsZeroValue = !eIsZeroValue
	eIsZeroValue = !eIsZeroValue
	fmt.Println("Expected IsZeroValue: ", eIsZeroValue)
	eIsIntegerZeroValue = !eIsIntegerZeroValue
	eIsIntegerZeroValue = !eIsIntegerZeroValue
	fmt.Println("Expected IsIntegerZeroValue: ", eIsIntegerZeroValue)
	fmt.Println("Expected Sign Value: ", eSignVal)
	fmt.Println("Expected Precision: ", ePrecision)
	fmt.Println("----------------------------------------------------")
	fmt.Println()
	fmt.Println("                Actual NumStr: ", ia.NumStr)
	fmt.Println("             Actual IntAryLen: ", ia.IntAryLen)
	fmt.Println("           Actual Integer Len: ", ia.IntegerLen)
	fmt.Println(" Actual SignificantIntegerLen: ", ia.SignificantIntegerLen)
	fmt.Println("Actual SignificantFractionLen: ", ia.SignificantFractionLen)
	fmt.Println("           Actual IsZeroValue: ", ia.IsZeroValue)
	fmt.Println("    Actual IsIntegerZeroValue: ", ia.IsIntegerZeroValue)
	fmt.Println("         Actual FirstDigitIdx: ", ia.FirstDigitIdx)
	fmt.Println("          Actual LastDigitIdx: ", ia.LastDigitIdx)
	fmt.Println("            Actual Sign Value: ", ia.SignVal)
	fmt.Println("             Actual Precision: ", ia.Precision)

}

func t01_SetWithInt64() {

	ia := common.IntAry{}.New()

	num := uint64(32)

	eNumStr := "0.0032"
	ePrecision := uint(4)
	eSignVal := 1

	err := ia.SetIntAryWithInt64(num, ePrecision, eSignVal)

	if err != nil {
		fmt.Printf("Error Received from ia.SetIntAryWithInt64(num, ePrecision, eSignVal) \n")
	}

	fmt.Println("SetIntAryWithInt64")
	fmt.Println("    Expected NumStr: ", eNumStr)
	fmt.Println(" Expected Precision: ", ePrecision)
	fmt.Println("Expected Sign Value: ", eSignVal)
	fmt.Println("----------------------------------")
	fmt.Println("Actual Results")
	fmt.Println("----------------------------------")
	fmt.Println("      Actual NumStr: ", ia.NumStr)
	fmt.Println("   Actual Precision: ", ia.Precision)
	fmt.Println("  Actual Sign Value: ", ia.SignVal)

}
