package RJC_Libs

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"
	"time"
)

// Consists of low-level methods used in calendar
// date/time calculations.
//
type CalendarMechanics struct {

	lock *sync.Mutex
}

// GetCompleteYearInterval - Calculates the number of
// complete years between a target year and a base year.
// The returned result is always a positive value expressed
// as a type 'int64'.
//
//
func (calMech *CalendarMechanics) GetCompleteYearInterval(
	baseYear int64,
	baseMonth,
	baseDay,
	baseHour,
	baseMinute,
	baseSecond,
	baseNanosecond int,
	targetYear int64,
	targetMonth,
	targetDay,
	targetHour,
	targetMinute,
	targetSecond,
	targetNanosecond int,
	ePrefix string) (completedYearsInterval int64, err error) {

	if calMech.lock == nil {
		calMech.lock = &sync.Mutex{}
	}

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "CalendarMechanics.GetCompleteYearInterval() "
	
	completedYearsInterval = math.MaxInt64
	err = nil
	

/*
	var tempYear int64
	var tempMonth, tempDay, tempHour, tempMinute, tempSecond,
	tempNanosecond int

	if targetYear <  baseYear {
		tempYear = baseYear
		tempMonth = baseMonth
		tempDay = baseDay
		tempHour = baseHour
		tempMinute = baseMinute
		tempSecond = baseSecond
		tempNanosecond = baseNanosecond

		baseYear = targetYear
		baseMonth = targetMonth
		baseDay = targetDay
		baseHour = targetHour
		baseMinute = targetMinute
		baseSecond = targetSecond
		baseNanosecond = targetNanosecond

		targetYear = tempYear
		targetMonth = tempMonth
		targetDay = tempDay
		targetHour = tempHour
		targetMinute = tempMinute
		targetSecond = tempSecond
		targetNanosecond = tempNanosecond
	}

*/

var baseDateTimeTotalNanoseconds, targetDateTimeNanoseconds int64

	calMech2 := CalendarMechanics{}
	
	baseDateTimeTotalNanoseconds, err = 
		calMech2.GetTimeTotalNanoseconds(
			baseHour,
			baseMinute,
			baseSecond,
			baseNanosecond,
			ePrefix)

	if err != nil {
		return completedYearsInterval, err
	}
	
	isBaseYearCompleteYear := false

	if baseMonth == 1 &&
			baseDay == 1 &&
			baseDateTimeTotalNanoseconds == 0 {

		isBaseYearCompleteYear = true
		
	}

	targetDateTimeNanoseconds, err =
		calMech2.GetTimeTotalNanoseconds(
			targetHour,
			targetMinute,
			targetSecond,
			targetNanosecond,
			ePrefix)

	if err != nil {
		return completedYearsInterval, err
	}

	twentyFourHourNanosecondsMinusOne := (int64(time.Hour) * 24) - 1

	isTargetYearCompleteYear := false

	if targetMonth == 12 &&
		targetDay == 31 &&
		targetDateTimeNanoseconds >= twentyFourHourNanosecondsMinusOne {

		isTargetYearCompleteYear = true

	}

	completedYearsInterval = baseYear - targetYear

	if completedYearsInterval < 0 {
		completedYearsInterval *= -1
	}


	if isBaseYearCompleteYear {
		completedYearsInterval++
	}
	
	if isTargetYearCompleteYear {
		completedYearsInterval++
	} else {
		completedYearsInterval--
	}
	
	separator := strings.Repeat("-", 65)

	fmt.Println()
	fmt.Println("GetCompleteYearInterval")
	fmt.Println(separator)
	fmt.Printf("                 baseYear: %v\n",baseYear)
	fmt.Printf("Last Complete Target Year: %v\n", targetYear)
	fmt.Println(separator)
	fmt.Printf("Complete Years Interval: %v\n", completedYearsInterval)
	fmt.Println(separator)

	return completedYearsInterval, err
}

// GetElapsedWholeOrdinalDaysInYear - Returns the number of elapsed
// 24-hour days base on specified month and day.
//
func (calMech *CalendarMechanics) GetElapsedWholeOrdinalDaysInYear(
	isLeapYear bool,
	month,
	day int,
	ePrefix string) (elapsedCompleteDays int, err error) {

	if calMech.lock == nil {
		calMech.lock = &sync.Mutex{}
	}

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	elapsedCompleteDays = math.MaxInt32
	err = nil

	ePrefix += "CalendarMechanics.GetElapsedWholeOrdinalDaysInYear() "

	calMech2 := CalendarMechanics{}

	var ordinalDayNumber int


	ordinalDayNumber, err = calMech2.GetOrdinalDayNumber(
		isLeapYear,
		month,
		day,
		ePrefix)

	if err != nil {
		return elapsedCompleteDays, err
	}

	if ordinalDayNumber > 0 {
		elapsedCompleteDays = ordinalDayNumber - 1
	} else {
		elapsedCompleteDays = 0
	}


	return elapsedCompleteDays, err
}

// GetJulianDayNoFraction - Computes the Julian Day Number time
// fraction.
//
func (calMech *CalendarMechanics) GetJulianDayNoFraction(
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (
	julianDayNoTimeFraction *big.Float,
	julianDayNoAdjustment int64,
	err error) {

	if calMech.lock == nil {
		calMech.lock = &sync.Mutex{}
	}

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "CalendarMechanics.GetJulianDayNoTimeFraction() "

	julianDayNoTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	julianDayNoAdjustment = 0

	err = nil

	calMech2 := CalendarMechanics{}

	var totalDateTimeNanoSecs, noonNanoSecs int64

	noonNanoSecs = int64(time.Hour) * 12

	totalDateTimeNanoSecs, err =
		calMech2.GetTimeTotalNanoseconds(
			hour,
			minute,
			second,
			nanosecond,
			ePrefix)

if totalDateTimeNanoSecs == noonNanoSecs {
	// Fraction is zero.
	return julianDayNoTimeFraction, julianDayNoAdjustment, err
}

if totalDateTimeNanoSecs  < noonNanoSecs {
	julianDayNoAdjustment = -1

	totalDateTimeNanoSecs += noonNanoSecs

} else {
	// totalDateTimeNanoSecs  > noonNanoSecs
	julianDayNoAdjustment = 0
	totalDateTimeNanoSecs -= noonNanoSecs
}

	twentyFourHourNanoseconds :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(int64(time.Hour) * 24)

	actualNanoSec :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(totalDateTimeNanoSecs)

	julianDayNoTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Quo(actualNanoSec, twentyFourHourNanoseconds)

	return julianDayNoTimeFraction, julianDayNoAdjustment, err
}

// GetLastCompleteYear - Given a specified date and time
// this method returns the last complete year. A complete
// year will contain either 365-days or 366-days (leap year).
// A day as used here is defined as 24-hours.
//
func (calMech *CalendarMechanics) GetLastCompleteYear(
	year int64,
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (lastCompleteYear int64, err error) {

	if calMech.lock == nil {
		calMech.lock = &sync.Mutex{}
	}

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "CalendarMechanics.GetLastCompleteYear() "

	lastCompleteYear = math.MaxInt64
	
	if month < 1 || month > 12 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'month' is invalid!\n" +
			"month='%v'\n", month)
		return lastCompleteYear, err
	}
	
	if day > 31 || day < 1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'day' is invalid!\n" +
			"day='%v'\n", day)
		return lastCompleteYear, err
	}

	calMech2 := CalendarMechanics{}

	var totalTimeNanoseconds int64

	totalTimeNanoseconds, err = calMech2.GetTimeTotalNanoseconds(
		hour,
		minute,
		second,
		nanosecond,
		ePrefix)

	if err != nil {
		return lastCompleteYear, err
	}

	twentyFourHoursMinus1 := (int64(time.Hour) * 24) - 1

	if month == 12 &&
		 day== 31 &&
		totalTimeNanoseconds >= twentyFourHoursMinus1 {

		lastCompleteYear = year

	} else {

		if year < 0 {
			lastCompleteYear = year + 1
		} else {
			lastCompleteYear = year - 1
		}

	}

	return lastCompleteYear, err
}

// GetMonthDayFromOrdinalDayNo - Receives an Ordinal Day Number and returns
// the associated month and day number. The input parameter 'isLeapYear'
// specifies whether the Ordinal Day Number is included in a standard year
// (365-Days) or a Leap Year (366-Days).
//
// Reference
//    https://en.wikipedia.org/wiki/Ordinal_date
//
func (calMech *CalendarMechanics) GetMonthDayFromOrdinalDayNo(
	ordinalDate int64,
	isLeapYear bool,
	ePrefix string)(year int64, month, day int, err error) {

	if calMech.lock == nil {
		calMech.lock = &sync.Mutex{}
	}

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calendarMechanics.GetMonthDayFromOrdinalDayNo() "

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

// GetOrdinalDayNumber - Computes the ordinal day number
// for any given month and day. Input parameter 'isLeapYear'
// indicates whether the year encompassing the specified
// month and day is a 'leap year' containing 366-days
// instead of the standard 365-days.
//
// Reference
//    https://en.wikipedia.org/wiki/Ordinal_date
//
func (calMech *CalendarMechanics) GetOrdinalDayNumber(
	isLeapYear bool,
	month int,
	day int,
	ePrefix string) (ordinalDayNo int, err error) {

	ePrefix += "CalendarMechanics.GetOrdinalDayNumber() "

	if calMech.lock == nil {
		calMech.lock = &sync.Mutex{}
	}

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ordinalDayNo = -1
	err = nil

	if month < 1 || month > 12 {
		err = fmt.Errorf("\n" + ePrefix + "Error:\n" +
			"Input Parameter 'month' is INVALID!\n" +
			"month='%v'\n", month)
		return ordinalDayNo, err
	}

	ordDays := []int {
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

	if month == 1 &&
		day == 0 {
		ordinalDayNo = 0
		return ordinalDayNo, err
	}

	monthDays := mthDays[month]

	if day > monthDays || day < 1 {
		err = fmt.Errorf("\n" + ePrefix + "Error:\n" +
			"Input parameter 'day' is INVALID!\n" +
			"month='%v'\n day='%v'\n",
			month, day)
		return ordinalDayNo, err
	}


	if month == 1 {
		ordinalDayNo = day
		return ordinalDayNo, err
	} else {

		ordinalDayNo = ordDays[month-1] + day

		if isLeapYear && month > 2 {
			ordinalDayNo++
		}
	}

	return ordinalDayNo, err
}


// GetTimeTotalNanoseconds - Computes the total time in nanoseconds for
// a given time of day expressed in Hours, Minutes, Seconds, and Nanoseconds.
//
func (calMech *CalendarMechanics) GetTimeTotalNanoseconds(
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (totalTimeNanoseconds int64, err error) {

	if calMech.lock == nil {
		calMech.lock = &sync.Mutex{}
	}

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "CalendarMechanics.GetTimeTotalNanoseconds() "

	totalTimeNanoseconds = math.MaxInt64
	err = nil

	if hour > 24 || hour < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hour' is invalid!\n" +
			"hour='%v'\n", hour)
		return totalTimeNanoseconds, err
	}

	if minute > 59 || minute < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'minute' is invalid!\n" +
			"minute='%v'\n", minute)
		return totalTimeNanoseconds, err
	}

	// Watch out for leap seconds
	if second > 60 || second < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'second' is invalid!\n" +
			"second='%v'\n", second)
		return totalTimeNanoseconds, err
	}

	if nanosecond > 999999999 || nanosecond < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecond' is invalid!\n" +
			"nanosecond='%v'\n", nanosecond)
		return totalTimeNanoseconds, err
	}

	totalTimeNanoseconds = 0

	totalTimeNanoseconds += int64(hour) * int64(time.Hour)

	totalTimeNanoseconds += int64(minute) * int64(time.Minute)

	totalTimeNanoseconds += int64(second) * int64(time.Second)

	totalTimeNanoseconds +=  int64(nanosecond)

	if totalTimeNanoseconds > int64(time.Hour) * 24 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Total Time in nanoseconds exceeds 24-hours\n" +
			"Total Nanosecond='%v'\n", totalTimeNanoseconds)
	}

	return totalTimeNanoseconds, err
}

