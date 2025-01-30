package datetime

import (
	"math/big"
	"sync"
)

// CalendarGregorianBaseData - Implements the
// Calendar Base Data Interface.
//
type JulianCalendarBaseData struct {
	lock *sync.Mutex
}

// GetDaysInLeapYear - Returns the number of days in a
// Gregorian Calendar Leap Year
func (gregJulianBData JulianCalendarBaseData) GetDaysInLeapYear() int {

	if gregJulianBData.lock == nil {
		gregJulianBData.lock = new(sync.Mutex)
	}

	gregJulianBData.lock.Lock()

	defer gregJulianBData.lock.Unlock()


	return 366
}

// GetDaysInStandardYear - Returns the number of days in a
// Julian Calendar Standard Year or non-leap year.
//
func (gregJulianBData JulianCalendarBaseData) GetDaysInStandardYear() int {

	if gregJulianBData.lock == nil {
		gregJulianBData.lock = new(sync.Mutex)
	}

	gregJulianBData.lock.Lock()

	defer gregJulianBData.lock.Unlock()

	return 365
}

// IsLeapYear - Determines whether the input parameter
// 'year' is a leap year under the Julian Calendar.
//
// If 'year' is a Julian leap year, this method returns 'true'.
//
// Note: This method should NOT be used to determine leap years
// for the Revised Julian Calendar or the Goucher-Parker
// calendar.
//
// Under the Julian Calendar a year is is classified as a 'leap year'
// if it is evenly divisible by '4'.
//
// Reference:
//   https://en.wikipedia.org/wiki/Julian_calendar
//
func  (gregJulianBData JulianCalendarBaseData) IsLeapYear(
	year int64) bool {

	if gregJulianBData.lock == nil {
		gregJulianBData.lock = new(sync.Mutex)
	}

	gregJulianBData.lock.Lock()

	defer gregJulianBData.lock.Unlock()

	remainder := year % 4

	if remainder == 0 {
		return true
	}

	return false
}

// calendarJulianMechanics - This type contains methods
// used to process date arithmetic associated with the
// Julian Calendar
//
// References:
//  https://en.wikipedia.org/wiki/Julian_calendar
//  https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars
//  https://www.fourmilab.ch/documents/calendar/
//
type calendarJulianMechanics struct {
	lock *sync.Mutex
}


// getCalendarCyclesConfig - returns a CalendarCycleConfiguration instance
// contain all the information necessary for julian day number calculations
// using the Julian Calendar
func (calJulianMech *calendarJulianMechanics) getCalendarCyclesConfig() (
	calCyclesCfg CalendarCycleConfiguration) {

	if calJulianMech.lock == nil {
		calJulianMech.lock = new(sync.Mutex)
	}

	calJulianMech.lock.Lock()

	defer calJulianMech.lock.Unlock()

	ePrefix := "calendarJulianMechanics.getCalendarCyclesConfig() "

	calCycles := CalendarCycleConfiguration{}

	var err error

	calCycles.mainCycleStartDateForPositiveJDNNo,
	err =
		ADateTimeDto{}.New(
		CalendarSpec(0).Julian(),
		int64(-4716),
		CalendarYearNumType(0).Astronomical(),
		1,
		1,
		false,
		12,
		0,
		0,
		0,
		"UTC",
		"",
		"Positive JDN No Main Cycle Start - Julian Calendar",
		ePrefix)

	if err != nil {
		return CalendarCycleConfiguration{}
	}


	calCycles.mainCycleAdjustmentYearsForPositiveJDNNo =
		big.NewInt(-4)

	calCycles.mainCycleAdjustmentDaysForPositiveJDNNo =
		big.NewInt(-1461)

	calCycles.mainCycleStartDateForNegativeJDNNo,
	err = ADateTimeDto{}.New(
		CalendarSpec(0).Julian(),
		-4708,
		CalendarYearNumType(0).Astronomical(),
		1,
		1,
		false,
		12,
		0,
		0,
		0,
		"UTC",
		"",
		"Negative JDN No Main Cycle Start - Julian Calendar",
		ePrefix)

	if err != nil {
		return CalendarCycleConfiguration{}
	}

	calCycles.mainCycleAdjustmentYearsForNegativeJDNNo =
		big.NewInt(-4)

	calCycles.mainCycleAdjustmentDaysForNegativeJDNNo =
		big.NewInt(-1461)

	calCycles.jdnBaseStartYearDateTime,
	err = ADateTimeDto{}.New(
		CalendarSpec(0).Julian(),
		int64(-4712),
		CalendarYearNumType(0).Astronomical(),
		1,
		1,
		false,
		12,
		0,
		0,
		0,
		"UTC",
		"",
		"",
		ePrefix)

	if err != nil {
		return CalendarCycleConfiguration{}
	}

	calCycles.jdnBaseStartYearDateTime,
	err = ADateTimeDto{}.New(
		CalendarSpec(0).Julian(),
		int64(1),
		CalendarYearNumType(0).Astronomical(),
		1,
		1,
		false,
		0,
		0,
		0,
		0,
		"UTC",
		"",
		"Julian Ordinal Day Number Start Date/Time",
		ePrefix)

	if err != nil {
		return CalendarCycleConfiguration{}
	}

	calCycles.mainCycleConfig =
		CalendarCycleDto{
			yearsInCycle: big.NewInt(4),
			daysInCycle:  big.NewInt(1461	),
			cycleCount:   big.NewInt(0),
			cycleCountTotalDays: big.NewInt(0),
			cycleCountTotalYears: big.NewInt(0),
			remainderYears: big.NewInt(0),
			remainderDays: big.NewInt(0),
			lock:         new(sync.Mutex),
		}

	calCycles.calendarCyclesConfig = []CalendarCycleDto{
		{// 0
			yearsInCycle: big.NewInt(4),
			daysInCycle:  big.NewInt(1461	),
			cycleCount:   big.NewInt(0),
			cycleCountTotalDays: big.NewInt(0),
			cycleCountTotalYears: big.NewInt(0),
			remainderYears: big.NewInt(0),
			remainderDays: big.NewInt(0),
			lock:         new(sync.Mutex),
		},

		{// 1
			yearsInCycle: big.NewInt(1),
			daysInCycle:  big.NewInt(365),
			cycleCount:   big.NewInt(0),
			cycleCountTotalDays: big.NewInt(0),
			cycleCountTotalYears: big.NewInt(0),
			remainderYears: big.NewInt(0),
			remainderDays: big.NewInt(0),
			lock:         new(sync.Mutex),
		},
	}

	// calCycles.calendarBaseData = JulianCalendarBaseData{}

	calCycles.lock = new(sync.Mutex)

	return calCycles
}

// isLeapYear - Determines whether the input parameter
// 'year' is a leap year under the Julian Calendar.
//
// If 'year' is a Julian leap year, this method returns 'true'.
//
// Note: This method should NOT be used to determine leap years
// for the Revised Julian Calendar or the Goucher-Parker
// calendar.
//
// Under the Julian Calendar a year is is classified as a 'leap year'
// if it is evenly divisible by '4'.
//
// Reference:
//   https://en.wikipedia.org/wiki/Julian_calendar
//
func (calJulianMech *calendarJulianMechanics) isLeapYear(
	year int64) bool {

	if calJulianMech.lock == nil {
		calJulianMech.lock = new(sync.Mutex)
	}

	calJulianMech.lock.Lock()

	defer calJulianMech.lock.Unlock()

	remainder := year % 4

	if remainder == 0 {
		return true
	}

	return false

}