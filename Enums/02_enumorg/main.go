package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*
 Adapted from Jeffery Richter Video

Jeffrey Richter Using Reflection to implement enumerated types
https://www.youtube.com/watch?v=DyXJy_0v0_U
https://github.com/JeffreyRichter/enum

 */

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

type TCalc struct {
	 TimeCalc TDurCalcType
}

var TimeDurationCalc = TDurCalcType(0)


func main() {

	test19()

}

func test18() {

 tCalc :=	TCalc{}.TimeCalc.CumMicroseconds()

 fmt.Println("tCalc: ", tCalc.String())

}

func test19() {
	a :=	TimeDurationCalc.CumDays()

	fmt.Println("a=CumDays: ", a.String())

	TimeDurationCalc = TDurCalcType(10)

	fmt.Println("a=CumDays: ", a.String())

	b := TimeDurationCalc.CumDays()

	fmt.Println("b=CumDays: ", b.String())

	t := TDurCalcType(0).CumHours()

	fmt.Println("t=", t)

}

func test20() {

	t := TDurCalcType(0).CumDays()

	fmt.Println("Testing CumDays")
	fmt.Println("Testing t.String() ", t)
	fmt.Println("Testing t.Value() ", t.Value())
	fmt.Println()

	v := TDurCalcType(0).StdYearMth()
	fmt.Println("Testing StdYearMth")
	fmt.Println("Testing v.String() ", v)
	fmt.Println("Testing v.Value() ", v.Value())
	fmt.Println()

	x := TDurCalcType(0).CumSeconds()
	fmt.Println("Testing CumSeconds")
	fmt.Println("Testing x.String() ", x)
	fmt.Println("Testing x.Value() ", x.Value())
	fmt.Println()

}