package main

import (
	"fmt"
	"math/big"
	"strings"
	"time"
)

func main() {
	// Use this as a test verification function
	startYear := int64(-4713)
	startMonth := 2
	startDay := 6
	startHour := 12
	startMinute:=0
	startSecond:=0
	startNanosecond:=0

	targetYear := int64(1)
	targetMonth := 1
	targetDay := 1
	targetHour := 12
	targetMinute:=0
	targetSecond:=0
	targetNanosecond:=0

	_,
	_,
	_,
	err :=
		testerRevisedJulianDateToJulianDayNumber(
			startYear,
			startMonth,
			startDay,
			startHour,
			startMinute,
			startSecond,
			startNanosecond,
			targetYear,
			targetMonth,
			targetDay,
			targetHour,
			targetMinute,
			targetSecond,
			targetNanosecond)

	if err != nil {
		fmt.Println("ERROR RETURN!")
	}

}


func testerRevisedJulianDateToJulianDayNumber(
	startYear int64,
	startMonth,
	startDay,
	startHours,
	startMinutes,
	startSeconds,
	startNanoseconds int,
	targetYear int64,
	targetMonth,
	targetDay,
	targetHours,
	targetMinutes,
	targetSeconds,
	targetNanoseconds int) (
	julianDayNo int64,
	julianDayNoTime,
	julianDayNoFraction *big.Float,
	err error) {

	ePrefix := "testerRevisedJulianDateToJulianDayNumber() "
	//The following constant defined midnight at the start of Revised Julian date Monday, 1 January 1 AD as the beginning of the first ordinal day. This moment was Julian Day number 1721425.5.
	//	RJepoch = 1
	//err = nil
	// https://www.liquisearch.com/revised_julian_calendar/revised_julian_calendrical_calculations/fixed_days

	julianDayNo = 0

	julianDayNoTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	julianDayNoFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	// Julian Day # 0 assumed to start at
	// 12:00-noon 2/6/-4713 (Astronomical Year Numbering)
	/*
		startYear := int64(-4713)
		startMonth := 2
		startDay := 6
		startHour := 12
		startMinute:=0
		startSecond:=0
		startNanosecond:=0

	*/
	var adjustedStartYear, beginningOrdinalDays int64
	var adjustedStartMonth, adjustedStartDay int

	beginningOrdinalDays,
		err = computeOrdinalDateFromYrMthDay(
		startYear,
		startMonth,
		startDay,
		ePrefix)

	if err != nil {
		return julianDayNo, julianDayNoTime, julianDayNoFraction, err
	}

	adjustedStartYear,
		adjustedStartMonth,
		adjustedStartDay,
		err = computeMonthDayFromOrdDate(
		beginningOrdinalDays,
		startYear,
		ePrefix)

	if err != nil {
		return julianDayNo, julianDayNoTime, julianDayNoFraction, err
	}


	firstCompleteYear := adjustedStartYear + 1 // -4712


	totalDays := beginningOrdinalDays
	yearDays := int64(0)

	for i:= firstCompleteYear; i < targetYear; i++ {

		yearDays = 365

		if rjcIsLeapYear(i) {
			yearDays = 366
		}

		totalDays += yearDays
	}

	var endingOrdinalDays int64

	endingOrdinalDays,
		err = computeOrdinalDateFromYrMthDay(
		targetYear,
		targetMonth,
		targetDay,
		ePrefix)

	//fmt.Printf("First Calc endingOrdinalDays: %v\n",
	//	endingOrdinalDays)
	//
	//fmt.Printf("First Calc targetYear: %v\n",
	//	targetYear)
	//
	//fmt.Printf("First Calc targetMonth: %v\n",
	//	targetMonth)
	//
	//
	//fmt.Printf("First Calc targetDay: %v\n\n",
	//	targetDay)


	if err != nil {
		return julianDayNo, julianDayNoTime, julianDayNoFraction, err
	}

	totalDays += endingOrdinalDays

	actualNanoSecs :=
		int64(targetHours) * int64(time.Hour) +
			int64(targetMinutes) * int64(time.Minute) +
			int64(targetSeconds) * int64(time.Second) +
			int64(targetNanoseconds)

	finalTargetYear := targetYear

	finalTargetMonth := targetMonth

	finalTargetDay := targetDay

	total12hourNanoSecs := int64(time.Hour * 12)

	if actualNanoSecs < total12hourNanoSecs {
		actualNanoSecs += total12hourNanoSecs
		totalDays--
		endingOrdinalDays--

		finalTargetYear,
			finalTargetMonth,
			finalTargetDay,
			err = computeMonthDayFromOrdDate(
			endingOrdinalDays,
			targetYear,
			ePrefix)

		if err != nil {
			return julianDayNo, julianDayNoTime, julianDayNoFraction, err
		}

	} else if actualNanoSecs == total12hourNanoSecs {

		actualNanoSecs = 0

		finalTargetYear,
			finalTargetMonth,
			finalTargetDay,
			err = computeMonthDayFromOrdDate(
			endingOrdinalDays,
			targetYear,
			ePrefix)

		if err != nil {
			return julianDayNo, julianDayNoTime, julianDayNoFraction, err
		}

	} else if actualNanoSecs > total12hourNanoSecs {
		actualNanoSecs -= total12hourNanoSecs
		totalDays++
		endingOrdinalDays ++

		finalTargetYear,
			finalTargetMonth,
			finalTargetDay,
			err = computeMonthDayFromOrdDate(
			endingOrdinalDays,
			targetYear,
			ePrefix)

		if err != nil {
			return julianDayNo, julianDayNoTime, julianDayNoFraction, err
		}

	}

	var finalTargetOrdinalDays int64

	finalTargetOrdinalDays,
		err = computeOrdinalDateFromYrMthDay(
		finalTargetYear,
		finalTargetMonth,
		finalTargetDay,
		ePrefix)

	if err != nil {
		return julianDayNo, julianDayNoTime, julianDayNoFraction, err
	}

	var finalTargetHours, finalTargetMinutes, finalTargetSeconds,
	finalTargetNanoseconds int

	remainingNanoSecs := actualNanoSecs

	if remainingNanoSecs >= int64(time.Hour) {
		finalTargetHours = int(remainingNanoSecs / int64(time.Hour))
		remainingNanoSecs -= int64(finalTargetHours) * int64(time.Hour)
	}

	if remainingNanoSecs >= int64(time.Minute) {
		finalTargetMinutes = int(remainingNanoSecs / int64(time.Minute))
		remainingNanoSecs -= int64(finalTargetMinutes) * int64(time.Minute)
	}

	if remainingNanoSecs >= int64(time.Second) {
		finalTargetSeconds = int(remainingNanoSecs / int64(time.Second))
		remainingNanoSecs -= int64(finalTargetSeconds) * int64(time.Second)
	}

	finalTargetNanoseconds = int(remainingNanoSecs)

	julianDayNo = totalDays

	total24HourNanoSecs := int64(time.Hour * 24)

	floatActNanoSecs :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(actualNanoSecs)

	float24HourNanoSecs :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(total24HourNanoSecs)

	julianDayNoFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Quo(floatActNanoSecs, float24HourNanoSecs)

	// fmt.Printf("First julianDayNoFraction: %30.20f\n",
	//	julianDayNoFraction)

	bigJulianDayNoFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(julianDayNo)

	julianDayNoTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Add(bigJulianDayNoFloat,julianDayNoFraction)

	var targetOrdinalDays int64

	targetOrdinalDays,
		err = computeOrdinalDateFromYrMthDay(
		targetYear,
		targetMonth,
		targetDay, ePrefix)

	if err != nil {
		return julianDayNo, julianDayNoTime, julianDayNoFraction, err
	}

	separator := strings.Repeat("-", 65)
	equalSeparator := strings.Repeat("=", 65)

	fmt.Println(ePrefix)
	fmt.Println(equalSeparator)
	fmt.Println("Computing Julian Day Number for")
	fmt.Printf("Revised Julian Calendar Date: %v-%v-%v %v:%v:%v.%v\n",
		targetYear,targetMonth,targetDay, targetHours, targetMinutes,
		targetSeconds, targetNanoseconds)
	fmt.Println(equalSeparator)
	fmt.Println(" The following constant defined midnight at the start of Revised")
	fmt.Println(" Julian date Monday, 1 January 1 AD as the beginning of the first")
	fmt.Println(" ordinal day. This moment was Julian Day number 1721425.5.")
	fmt.Println(" RJepoch = 1")
	fmt.Println(" err = nil")
	fmt.Println(" https://www.liquisearch.com/revised_julian_calendar/revised_julian_calendrical_calculations/fixed_days")
	fmt.Println(equalSeparator)

	fmt.Println(separator)
	fmt.Printf("Computed Julian Day No: %v\n", totalDays)
	fmt.Println(separator)
	fmt.Printf("       Original Start Year: %v\n", startYear)
	fmt.Printf("      Original Start Month: %v\n", startMonth)
	fmt.Printf("        Original Start Day: %v\n", startDay)
	fmt.Printf("       Original Start Hour: %v\n", startHours)
	fmt.Printf("     Original Start Minute: %v\n", startMinutes)
	fmt.Printf("     Original Start Second: %v\n", startSeconds)
	fmt.Printf(" Original Start Nanosecond: %v\n", startNanoseconds)
	fmt.Println(separator)
	fmt.Printf("        Adjusted Start Year: %v\n", adjustedStartYear)
	fmt.Printf("       Adjusted Start Month: %v\n", adjustedStartMonth)
	fmt.Printf("         Adjusted Start Day: %v\n", adjustedStartDay)
	fmt.Printf("        Adjusted Start Hour: %v\n", startHours)
	fmt.Printf("      Adjusted Start Minute: %v\n", startMinutes)
	fmt.Printf("      Adjusted Start Second: %v\n", startSeconds)
	fmt.Printf("  Adjusted Start Nanosecond: %v\n", startNanoseconds)
	fmt.Printf("    Start Year Ordinal Days: %v\n", beginningOrdinalDays)
	fmt.Printf("        First Complete Year: %v\n", firstCompleteYear)
	fmt.Println(separator)
	fmt.Printf("                Target Year: %v\n", targetYear)
	fmt.Printf("               Target Month: %v\n", targetMonth)
	fmt.Printf("                 Target Day: %v\n", targetDay)
	fmt.Printf("                Target Hour: %v\n", targetHours)
	fmt.Printf("              Target Minute: %v\n", targetMinutes)
	fmt.Printf("              Target Second: %v\n", targetSeconds)
	fmt.Printf("          Target Nanosecond: %v\n", targetNanoseconds)
	fmt.Printf("        Target Ordinal Days: %v\n", targetOrdinalDays)
	fmt.Println(separator)
	fmt.Printf("          Final Target Year: %v\n", finalTargetYear)
	fmt.Printf("         Final Target Month: %v\n", finalTargetMonth)
	fmt.Printf("           Final Target Day: %v\n", finalTargetDay)
	fmt.Printf("         Final Target Hours: %v\n", finalTargetHours)
	fmt.Printf("       Final Target Minutes: %v\n", finalTargetMinutes)
	fmt.Printf("       Final Target Seconds: %v\n", finalTargetSeconds)
	fmt.Printf("   Final Target Nanoseconds: %v\n", finalTargetNanoseconds)
	fmt.Printf("Final Target Ordinal Day No: %v\n", finalTargetOrdinalDays)
	fmt.Println(separator)
	fmt.Printf("        * Returned Results *\n")
	fmt.Printf("        Int64 Julian Day No:   %v\n", julianDayNo)
	fmt.Printf("         Julian Day No/Time: %30.20f\n",
		julianDayNoTime)
	fmt.Printf("     Julian Day No Fraction: %30.20f\n",
		julianDayNoFraction)
	fmt.Println(equalSeparator)
	fmt.Println(equalSeparator)

	return julianDayNo, julianDayNoTime, julianDayNoFraction, err

}


func computeMonthDayFromOrdDate(
	ordinalDate int64,
	inputYear int64,
	ePrefix string)(year int64, month, day int, err error) {

	ePrefix += "computeMonthDayFromOrdDate() "
	year = -1
	month = -1
	day = -1
	err = nil

	if ordinalDate < 0 ||
		ordinalDate > 364 {
		err = fmt.Errorf("\n" +ePrefix + "Error:\n" +
			"Input Parameter 'ordinalDate' is INVALID!\n" +
			"ordinalDate='%v'\n",
			ordinalDate)
		return year, month, day, err
	}

	ordDays := []int64 {
		0,
		31,
		59,
		90,
		120,
		151,
		181,
		212,
		243,
		273,
		304,
		334,
	}

	mthDays := []int {
		-1,
		31,
		28,
		31,
		30,
		31,
		30,
		31,
		31,
		30,
		31,
		30,
		31,
	}

	year = inputYear

	isLeapYear := rjcIsLeapYear(year)

	for i:=11; i > -1; i-- {

		if ordDays[i] <= ordinalDate {

			ordinalDate -= ordDays[i]

			month = i + 1

			if month > 2 && isLeapYear {
				ordinalDate--
			}

			if ordinalDate == 0 &&
				month==1 {
				year--
				month = 12
				day = 31

				break

			} else if ordinalDate == 0 {
				month--
				day = mthDays[month]
				break
			} else {
				day = int(ordinalDate)
				break
			}

		}
	}

	return year, month, day, err
}

func computeOrdinalDateFromYrMthDay(
	year int64,
	month int,
	day int,
	ePrefix string) (ordinalDate int64, err error) {

	ePrefix += "computeOrdinalDateFromYrMthDay() "
	ordinalDate = -1
	err = nil

	if month < 1 || month > 12 {
		err = fmt.Errorf("\n" + ePrefix + "Error:\n" +
			"Input Parameter 'month' is INVALID!\n" +
			"month='%v'\n", month)
		return ordinalDate, err
	}


	ordDays := []int64 {
		0,
		31,
		59,
		90,
		120,
		151,
		181,
		212,
		243,
		273,
		304,
		334,
	}

	mthDays := []int {
		-1,
		31,
		28,
		31,
		30,
		31,
		30,
		31,
		31,
		30,
		31,
		30,
		31,
	}

	monthDays := mthDays[month]

	if day > monthDays || day < 1 {
		err = fmt.Errorf("\n" + ePrefix + "Error:\n" +
			"Input parameter 'day' is INVALID!\n" +
			"month='%v'\n day='%v'\n",
			month, day)
		return ordinalDate, err
	}

	isLeapYear := rjcIsLeapYear(year)

	if month == 1 {
		ordinalDate = int64(day)
		return ordinalDate, err
	} else {

		ordinalDate = ordDays[month-1] + int64(day)

		if isLeapYear && month > 2 {
			ordinalDate++
		}
	}

	return ordinalDate, err
}

func rjcIsLeapYear(year int64) bool {

	var isLeapYear bool

	if year < 0 {
		year *= -1
	}

	if year % int64(4) == 0 {

		isLeapYear = true

		if year % 100 == 0 {

			isLeapYear = false

			mod900 := year % int64(900)

			if  mod900 == 200 ||
				mod900 == 600 {
				isLeapYear = true
			}

		}

	} else {
		isLeapYear = false
	}

	return isLeapYear
}
