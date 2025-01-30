package datetime

import "sync"

type calendarRevisedJulianMechanics struct {
	lock *sync.Mutex
}

// isLeapYear - Determines whether the input parameter 'year'
// is a leap year under the Revised Julian Calendar.
//
// The Revised Julian calendar has the same months and month
// lengths as the Julian calendar, but, in the Revised Julian
// calendar, years evenly divisible by 100 are not leap years,
// except that years with remainders of 200 or 600 when divided
// by 900 remain leap years, e.g. 2000 and 2400 as in the Gregorian
// Calendar.
//
// For additional information, reference:
//    https://en.wikipedia.org/wiki/Revised_Julian_calendar
//
// Summary
//
// 1. Years evenly divisible by 4 are leap years unless they are
//    century years.
//
// 2. Years evenly divisible by 100 are not leap years unless when
//    divided by 900 those years have remainders of 200 or 600 in
//    which case they are leap years.
//
func (calRevJulianMech calendarRevisedJulianMechanics) isLeapYear(
	year int64) bool {

	if calRevJulianMech.lock == nil {
		calRevJulianMech.lock = new(sync.Mutex)
	}

	calRevJulianMech.lock.Lock()

	defer calRevJulianMech.lock.Unlock()


	var by4Remainder, by100Remainder, by900Remainder int64

	by100Remainder = year % 100

	if by100Remainder == 0 {

		by900Remainder = year % 900

		if by900Remainder == 200 ||
			by900Remainder == 900 {
			return true
		}

		return false
	}

	by4Remainder = year % 4

	if by4Remainder == 0 {
		return true
	}

	return false
}
