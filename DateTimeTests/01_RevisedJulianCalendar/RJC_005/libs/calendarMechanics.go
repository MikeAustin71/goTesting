package libs

import (
	"fmt"
	"strings"
	"sync"
)

// Consists of low-level methods used in calendar
// date/time calculations.
//
type CalendarMechanics struct {

	lock *sync.Mutex
}

// ComputeCompleteYearInterval - Calculates the number of
// complete years between a target year and a base year.
// The returned result is always a positive value expressed
// as a type 'int64'.
//
// The target year must be a complete year. If it is a partial
// year, the returned 'Completed Years Interval' will be invalid.
//
// The base year may be a partial year or a complete year.
//
func (calMech *CalendarMechanics) ComputeCompleteYearInterval(
	baseYear,
	targetYear int64) (completedYearsInterval int64) {

	if calMech.lock == nil {
		calMech.lock = &sync.Mutex{}
	}

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	completedYearsInterval = baseYear - targetYear

	if completedYearsInterval < 0 {
		completedYearsInterval *= -1
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

	return completedYearsInterval
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
	ePrefix string) (ordinalDate int64, err error) {

	ePrefix += "CalendarMechanics.GetOrdinalDayNumber() "

	if calMech.lock == nil {
		calMech.lock = &sync.Mutex{}
	}

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

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



