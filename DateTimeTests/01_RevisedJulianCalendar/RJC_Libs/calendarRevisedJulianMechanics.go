package RJC_Libs

import (
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"
)

// This type contains methods used to process date
// arithmetic associated with the Revised Julian
// Calendar.
//
// Reference:
//  https://en.wikipedia.org/wiki/Revised_Julian_calendar
//
type CalendarRevisedJulianMechanics struct {

	lock *sync.Mutex
}

// GetTimeOfDayFraction - Computes the time fraction for a day
// in the Revised Julian Calendar. The Revised Julian Calendar
// day starts at midnight.
//
// Reference:
//  https://en.wikipedia.org/wiki/Revised_Julian_calendar
//
func (calRJM *CalendarRevisedJulianMechanics) GetTimeOfDayFraction(
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (
	timeFraction *big.Float,
	err error) {

	if calRJM.lock == nil {
		calRJM.lock = &sync.Mutex{}
	}

	calRJM.lock.Lock()

	defer calRJM.lock.Unlock()

	ePrefix += "CalendarMechanics.GetTimeOfDayFraction() "

	timeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	err = nil

	calMech := CalendarMechanics{}

	var totalDateTimeNanoSecs int64

	totalDateTimeNanoSecs, err =
		calMech.GetTimeTotalNanoseconds(
			hour,
			minute,
			second,
			nanosecond,
			ePrefix)

	if err != nil {
		return timeFraction, err
	}

	twentyFourHourNanoseconds :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(int64(time.Hour) * 24)

	actualNanoseconds :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(totalDateTimeNanoSecs)

	timeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Quo(actualNanoseconds, twentyFourHourNanoseconds)

	return timeFraction, err
}

// IsLeapYear - Returns a boolean value signaling whether the year
// value passed as an input parameter is a leap year (366-days) under
// the Revised Julian Calendar.
//
// Reference:
//  https://en.wikipedia.org/wiki/Revised_Julian_calendar
//
func (calRJM *CalendarRevisedJulianMechanics) IsLeapYear(year int64) bool {

	if calRJM.lock == nil {
		calRJM.lock = &sync.Mutex{}
	}

	calRJM.lock.Lock()

	defer calRJM.lock.Unlock()

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

// NumCalendarDaysForWholeYearsInterval - Computes the total
// number of 24-hour days in a period of years specified
// by input parameter 'wholeYearsInterval'. The number of total
// days is calculated in accordance with the Revised Julian
// Calendar.
//
// 'wholeYearsInterval' is defined as a series of contiguous
// whole, or complete, years consisting of either 365-days
// or 366-days (in the case of leap years).
//
// No partial years should be included in this interval.
//
func (calRJM *CalendarRevisedJulianMechanics) NumCalendarDaysForWholeYearsInterval(
	wholeYearsInterval int64) (totalDays int64) {

	if calRJM.lock == nil {
		calRJM.lock = &sync.Mutex{}
	}

	calRJM.lock.Lock()

	defer calRJM.lock.Unlock()

	separator := strings.Repeat("*", 65)

	totalDays = 0

	if wholeYearsInterval < 0 {
		wholeYearsInterval *= -1
	}

	if wholeYearsInterval == 0 {
		return 0
	}

	fmt.Println()
	fmt.Println("NumCalendarDaysForWholeYearsInterval() ")
	fmt.Println(separator)
	fmt.Printf("       Whole Years Interval: %v\n", wholeYearsInterval)

	if wholeYearsInterval >= 900 {

		numOfCycles := wholeYearsInterval / 900

		totalDays = numOfCycles * 328718

		fmt.Printf("  Number of 900-Year Cycles: %v\n", numOfCycles)
		fmt.Printf("Number of Days in %v-Cycles: %v\n", numOfCycles, totalDays)

		wholeYearsInterval = wholeYearsInterval - (numOfCycles * 900)

		fmt.Printf("  Number of Remainder Years: %v\n", wholeYearsInterval)
		fmt.Println(separator)
		fmt.Println()

	}

	totalDays += wholeYearsInterval * 365

	leapDays := wholeYearsInterval / 4

	skipLeapDays := wholeYearsInterval / 100

	addLeapDays := int64(0)

	if wholeYearsInterval >= 200 {
		addLeapDays++
	}

	if wholeYearsInterval >= 600 {
		addLeapDays++
	}

	totalDays += leapDays + addLeapDays - skipLeapDays

	fmt.Println(separator)
	fmt.Printf("Total Days In wholeYearsInterval: %v\n",
		totalDays)
	fmt.Println(separator)
	fmt.Println()

	return totalDays
}
