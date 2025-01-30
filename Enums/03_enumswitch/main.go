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

	c.checkInitializeMaps(false)

	return  mTDurCalcTypeIntToString[int(c)]
}

func (c TDurCalcType) ParseString(
	valueString string,
	caseSensitive bool) (TDurCalcType, error) {

	ePrefix := "TDurCalcType.ParseString() "

	c.checkInitializeMaps(false)

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

func (c TDurCalcType) checkInitializeMaps(reInitialize bool) {


	if !reInitialize &&
		mTDurCalcTypeIntToString != nil &&
		len(mTDurCalcTypeIntToString) > 3 &&
		mTDurCalcTypeStringToInt != nil &&
		len(mTDurCalcTypeStringToInt) > 3 &&
		mTDurCalcTypeLwrCaseStringToInt != nil &&
		len(mTDurCalcTypeLwrCaseStringToInt) > 3 {
		fmt.Println("checkInitializeMaps == OK!")
		return
	}

	var t  = TDurCalcType(0).StdYearMth()

	mTDurCalcTypeIntToString = make(map[int] string, 0)
	mTDurCalcTypeStringToInt = make(map[string] int, 0)
	mTDurCalcTypeLwrCaseStringToInt = make(map[string] int, 0)

	s := reflect.TypeOf(t)

	r := reflect.TypeOf(int(0))
	args := [1]reflect.Value{reflect.Zero(s)}

	for i:= 0; i < s.NumMethod(); i ++ {

		f := s.Method(i).Name

		if f == "String" ||
			f == "ParseString" ||
			f == "Value" ||
			f == "checkInitializeMaps" {
			continue
		}

		value:= s.Method(i).Func.Call(args[:])[0].Convert(r).Int()
		x := int(value)
		mTDurCalcTypeIntToString[x] = f
		mTDurCalcTypeStringToInt[f] = x
		mTDurCalcTypeLwrCaseStringToInt[strings.ToLower(f)] = x
	}

	fmt.Println("Initialized Maps")
}

func main() {
	t:= TDurCalcType(0).GregorianYears()

	fmt.Println("Expecting: ", t)
	test20(t)

	t = TDurCalcType(0).CumSeconds()

	fmt.Println("Expecting: ", t)
	test20(t)

}

func test20(t TDurCalcType) {

	switch t {
	case t.StdYearMth():
		fmt.Println("Found StdYearMth")
	case t.CumMonths():
		fmt.Println("Found CumMonths")
	case t.CumWeeks():
		fmt.Println("Found CumWeeks")
	case t.CumDays():
		fmt.Println("Found CumDays")
	case t.CumHours():
		fmt.Println("Found CumHours")
	case t.CumMinutes():
		fmt.Println("Found CumMinutes")
	case t.CumSeconds():
		fmt.Println("Found CumSeconds")
	case t.CumMilliseconds():
		fmt.Println("Found CumMilliseconds")
	case t.CumMicroseconds():
		fmt.Println("Found CumMicroseconds")
	case t.CumNanoseconds():
		fmt.Println("Found CumNanoseconds")
	case t.GregorianYears():
		fmt.Println("Found GregorianYears")
	default:
		fmt.Println("Error Did Not Match Enum Value!")
	}


}
