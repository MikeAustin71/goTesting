package main

import (
	"fmt"
	"reflect"
	"strings"
)

var mTDurCalcTypeIntToString = map[int]string{}

var mTDurCalcTypeStringToInt = map[string]int{}

var mTDurCalcTypeLwrCaseStringToInt = map[string]int{}

type TDurCalcType int


// StdYearMth - Value 0 Was TDurCalcTypeSTDYEARMTH
// Standard Year Month calculation
func (TDurCalcType) StdYearMth() TDurCalcType {return TDurCalcType(0)}

// CumMonths - Value 1 Was TDurCalcTypeCUMMONTHS
// Cumulative Months calculation
func (TDurCalcType) CumMonths() TDurCalcType {return TDurCalcType(1)}

// CumWeeks - Value 2 Was TDurCalcTypeCUMWEEKS
func (TDurCalcType) CumWeeks() TDurCalcType {return TDurCalcType(2)}

// CumDays - Value 3 Was TDurCalcTypeCUMDAYS
func (TDurCalcType) CumDays() TDurCalcType {return TDurCalcType(3)}

// CumHours - Value 4 Was TDurCalcTypeCUMHOURS
func (TDurCalcType) CumHours() TDurCalcType {return TDurCalcType(4)}

// CumMinutes - Value 5  Was TDurCalcTypeCUMMINUTES
func (TDurCalcType) CumMinutes() TDurCalcType {return TDurCalcType(5)}

// CumSeconds - Value 6 TDurCalcTypeCUMSECONDS
func (TDurCalcType) CumSeconds() TDurCalcType {return TDurCalcType(6)}

// CumMilliseconds - Value 7 Cumulative Milliseconds
func (TDurCalcType) CumMilliseconds() TDurCalcType {return TDurCalcType(7)}

// CumMicroseconds - Value 8 Cumulative Microseconds
func (TDurCalcType) CumMicroseconds() TDurCalcType {return TDurCalcType(8)}

// CumNanoSeconds - Value 9 Cumulative Nanoseconds
func (TDurCalcType) CumNanoseconds() TDurCalcType {return TDurCalcType(9)}

// GregorianYears - Value 10  Was TDurCalcTypeGregorianYrs
func (TDurCalcType) GregorianYears() TDurCalcType {return TDurCalcType(10)}

func (c TDurCalcType) String() string {
	 return  mTDurCalcTypeIntToString[int(c)]
}

func (c TDurCalcType) ParseString(
	valueString string,
	caseSensitive bool) (TDurCalcType, error) {

	ePrefix := "TDurCalcType.ParseString() "
	result := TDurCalcType(0)
	if len(valueString) < 3 {
		return result,
		fmt.Errorf(ePrefix +
			"Input parameter 'valueString' is INVALID! valueString='%v' ", valueString)
	}

	var ok bool
	var idx int

	if caseSensitive {

		idx, ok = mTDurCalcTypeStringToInt[valueString]

		if !ok {
			return TDurCalcType(0),
				fmt.Errorf(ePrefix +
					"'valueString' did NOT MATCH a TDurCalcType. valueString='%v' ", valueString)
		}

		if !ok {
			return TDurCalcType(0),
				fmt.Errorf(ePrefix +
					"'valueString' did NOT MATCH a TDurCalcType. valueString='%v' ", valueString)
		}

		result = TDurCalcType(idx)

	} else {


		idx, ok = mTDurCalcTypeLwrCaseStringToInt[strings.ToLower(valueString)]

		if !ok {
			return TDurCalcType(0),
				fmt.Errorf(ePrefix +
					"'valueString' did NOT MATCH a TDurCalcType. valueString='%v' ", valueString)
		}

		result =
			TDurCalcType(idx)
	}

	return result, nil
}

func (c TDurCalcType) Value() int {
	return int(c)
}

func main() {
	initializeCalcTypeMaps()
	printMaps()

}

func parseTypes3() error {
	ePrefix := "parseTypes3() "
	initializeCalcTypeMaps()

	fmt.Println("Expecting Cumulative Seconds Value 6")

	t, err := TDurCalcType(0).ParseString("cumseconds", false)

	if err != nil {
		return fmt.Errorf(ePrefix + "%v", err.Error())
	}

	fmt.Println("Expecting t.String() = ", "CumSeconds")
	fmt.Println("   Actual t.String() = ", t.String())
	fmt.Println("    Expected t value = ", "6")
	fmt.Println("      Actual t value = ", t.Value())

	return nil
}

func parseTypes2() error {
	ePrefix := "parseTypes2() "
	initializeCalcTypeMaps()

	fmt.Println("Expecting Cumulative Seconds Value 6")

	t, err := TDurCalcType(0).ParseString("CumSeconds", false)

	if err != nil {
		return fmt.Errorf(ePrefix + "%v", err.Error())
	}

	fmt.Println("Expecting t.String() = ", "CumSeconds")
	fmt.Println("   Actual t.String() = ", t.String())
	fmt.Println("    Expected t value = ", "6")
	fmt.Println("      Actual t value = ", t.Value())

	return nil
}

func parseTypes() {
	initializeCalcTypeMaps()
	var t  = TDurCalcType(0).CumSeconds()
	fmt.Println("Expecting Cumulative Seconds Value 6")
	str := t.String()
	fmt.Println("t.String() ", str)

	idx, ok := mTDurCalcTypeStringToInt[str]

	if !ok {
		fmt.Println("parseTypes() Error string not found in mTDurCalcTypeStringToInt[str]")
		return
	}

	v := TDurCalcType(idx)

	fmt.Println("v=", v.String())
	fmt.Println("v value= ", v.Value())
}

func initializeCalcTypeMaps() {

	var t  = TDurCalcType(0).CumSeconds()

	mTDurCalcTypeIntToString = make(map[int] string, 0)
	mTDurCalcTypeStringToInt = make(map[string] int, 0)
	mTDurCalcTypeLwrCaseStringToInt = make(map[string] int, 0)

	s := reflect.TypeOf(t)

	r := reflect.TypeOf(int(0))
	args := [1]reflect.Value{reflect.Zero(s)}

	for i:= 0; i < s.NumMethod(); i ++ {

		f := s.Method(i).Name

		if f == "String" ||
				f== "ParseString" ||
					f=="Value"{
			continue
		}

		value:= s.Method(i).Func.Call(args[:])[0].Convert(r).Int()
		x := int(value)
		mTDurCalcTypeIntToString[x] = f
		mTDurCalcTypeStringToInt[f] = x
		mTDurCalcTypeLwrCaseStringToInt[strings.ToLower(f)] = x
	}



}

func printMaps() {

	fmt.Println("mTDurCalcTypeIntToString")

	for key, value := range mTDurCalcTypeIntToString {
		fmt.Println("Key=", key, " Value=", value)
	}

	fmt.Println()
	fmt.Println("mTDurCalcTypeStringToInt")

	for key, value := range mTDurCalcTypeStringToInt {
		fmt.Println("Key=", key, " Value=", value)
	}

	fmt.Println()
	fmt.Println("mTDurCalcTypeLwrCaseStringToInt")

	for key, value := range mTDurCalcTypeLwrCaseStringToInt{
		fmt.Println("Key=", key, " Value=", value)
	}


}


func testStringReturn08() {

	initializeCalcTypeMaps()
	var t  = TDurCalcType(0).CumSeconds()
	fmt.Println("Expecting Cumulative Seconds")
	fmt.Println("t.String() ", t.String())

}

func test09() {

	var t  = TDurCalcType(0).CumSeconds()

	fmt.Println("CumSeconds() string: ", t.String())

	s := reflect.TypeOf(t)

	r := reflect.TypeOf(int(0))

	fmt.Println("NumOfMethods: ", s.NumMethod())
	fmt.Println()
	args := [1]reflect.Value{reflect.Zero(s)}

	for i:= 0; i < s.NumMethod(); i ++ {
		f := s.Method(i).Name
		if f == "String" {
			continue
		}

		fmt.Println("Method Name: ", f)
		fmt.Println("-----------------")
		value:= s.Method(i).Func.Call(args[:])[0].Convert(r)

		fmt.Printf("This Value: %v \n", value)
		fmt.Println()
		fmt.Println()
	}

}

func test10() {

	var t  = TDurCalcType(0).CumSeconds()

	fmt.Println("CumSeconds() string: ", t.String())

	s := reflect.TypeOf(&t)

	fmt.Println("NumOfMethods: ", s.NumMethod())
	fmt.Println()
	for i:= 0; i < s.NumMethod(); i ++ {
		f := s.Method(i).Name
		fmt.Println("Method Name: ", f)
	}

}