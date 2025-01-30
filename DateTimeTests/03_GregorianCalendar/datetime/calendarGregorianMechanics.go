package datetime

import (
	"math/big"
	"sync"
)

type calendarGregorianMechanics struct {
	lock *sync.Mutex
}

// getCalendarCyclesConfig - returns a CalendarCycleConfiguration instance
// contain all the information necessary for julian day number calculations
// using the Gregorian Calendar
func (calGregMech *calendarGregorianMechanics) getCalendarCyclesConfig(
	) (
	calCyclesCfg CalendarCycleConfiguration) {

	if calGregMech.lock == nil {
		calGregMech.lock = new(sync.Mutex)
	}

	calGregMech.lock.Lock()

	defer calGregMech.lock.Unlock()

	ePrefix := "calendarGregorianMechanics.getCalendarCyclesConfig() "

	var err error

	calCycles := CalendarCycleConfiguration{}

	calCycles.mainCycleStartDateForPositiveJDNNo,
	err =
		ADateTimeDto{}.New(
			CalendarSpec(0).Gregorian(),
			int64(-5200),
			CalendarYearNumType(0).Astronomical(),
			11,
			24,
			false,
			12,
			0,
			0,
			0,
			"UTC",
			"",
			"Positive JDN No Main Cycle Start",
			ePrefix)

	if err != nil {
		return CalendarCycleConfiguration{}
	}

	calCycles.mainCycleAdjustmentYearsForPositiveJDNNo =
		big.NewInt(-487)

	calCycles.mainCycleAdjustmentDaysForPositiveJDNNo =
		big.NewInt(-177873)

	calCycles.mainCycleStartDateForNegativeJDNNo,
		err =
		ADateTimeDto{}.New(
			CalendarSpec(0).Gregorian(),
			int64(-4000),
			CalendarYearNumType(0).Astronomical(),
			11,
			24,
			false,
			12,
			0,
			0,
			0,
			"UTC",
			"",
			"Negative JDN No Main Cycle Start",
			ePrefix)

	if err != nil {
		return CalendarCycleConfiguration{}
	}

	calCycles.mainCycleAdjustmentYearsForNegativeJDNNo =
		big.NewInt(-713)

	calCycles.mainCycleAdjustmentDaysForNegativeJDNNo =
		big.NewInt(-260417)

	calCycles.jdnBaseStartYearDateTime,
		err =
		ADateTimeDto{}.New(
			CalendarSpec(0).Gregorian(),
			int64(-4713),
			CalendarYearNumType(0).Astronomical(),
			11,
			24,
			false,
			12,
			0,
			0,
			0,
			"UTC",
			"",
			"Julian Day Number Base Start Date/Time",
			ePrefix)

	if err != nil {
		return CalendarCycleConfiguration{}
	}

	calCycles.ordinalFixedDateStartYearDateTime,
		err =
		ADateTimeDto{}.New(
			CalendarSpec(0).Gregorian(),
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
			"Gregorian Ordinal Day Number Start Date/Time",
			ePrefix)

	if err != nil {
		return CalendarCycleConfiguration{}
	}

	calCycles.mainCycleConfig =
		CalendarCycleDto{
			yearsInCycle:         big.NewInt(400),
			daysInCycle:          big.NewInt(146097),
			cycleCount:           big.NewInt(0),
			cycleCountTotalDays:  big.NewInt(0),
			cycleCountTotalYears: big.NewInt(0),
			remainderYears: big.NewInt(0),
			remainderDays: big.NewInt(0),
			lock:                 new(sync.Mutex),
		}

	calCycles.calendarCyclesConfig = []CalendarCycleDto{
		{ // 0
			yearsInCycle:         big.NewInt(400),
			daysInCycle:          big.NewInt(146097),
			cycleCount:           big.NewInt(0),
			cycleCountTotalDays:  big.NewInt(0),
			cycleCountTotalYears: big.NewInt(0),
			remainderYears: big.NewInt(0),
			remainderDays: big.NewInt(0),
			lock:                 new(sync.Mutex),
		},

		{ // 1
			yearsInCycle: big.NewInt(100),
			daysInCycle:  big.NewInt(36524),
			cycleCount:   big.NewInt(0),
			cycleCountTotalDays: big.NewInt(0),
			cycleCountTotalYears: big.NewInt(0),
			remainderYears: big.NewInt(0),
			remainderDays: big.NewInt(0),
			lock:         new(sync.Mutex),
		},

		{// 2
			yearsInCycle: big.NewInt(4),
			daysInCycle:  big.NewInt(1461	),
			cycleCount:   big.NewInt(0),
			cycleCountTotalDays: big.NewInt(0),
			cycleCountTotalYears: big.NewInt(0),
			remainderYears: big.NewInt(0),
			remainderDays: big.NewInt(0),
			lock:         new(sync.Mutex),
		},

		{// 3
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

	calCycles.calendarBaseData = &CalendarGregorianBaseData{}

	calCycles.lock = new(sync.Mutex)

	return calCycles
}

// isLeapYear - Returns a boolean value signaling whether the year
// value passed as an input parameter is a leap year (366-days)
// under the Gregorian Calendar.
//
// If the method returns 'true' the input parameter 'year' qualifies
// as a leap year consisting of 366-days. If the method returns 'false',
// the input parameter 'year' is a standard year consisting of 365-days.
//
// Methodology:
//
// In the Gregorian calendar, three criteria must be taken
// into account to identify leap years:
//
// 1. The year must be evenly divisible by 4;
//
// 2. If the year can also be evenly divided by 100, it is not
//    a leap year, unless...
//
// 3. The year is evenly divisible by 100 and the year is also
//     evenly divisible by 400. Then it is a leap year.
//
// According to these rules, the years 2000 and 2400 are leap years,
// while 1800, 1900, 2100, 2200, 2300, and 2500 are not leap years.
//
// For more information on the Gregorian Calendar and leap years,
// reference:
//
//   https://en.wikipedia.org/wiki/Gregorian_calendar
//   https://www.timeanddate.com/date/leapyear.html
//
//
func (calGregMech *calendarGregorianMechanics) isLeapYear(
	year int64) bool {

	if calGregMech.lock == nil {
		calGregMech.lock = new(sync.Mutex)
	}

	calGregMech.lock.Lock()

	defer calGregMech.lock.Unlock()

	calBaseData := CalendarGregorianBaseData{}

	isLeapYear, _ := calBaseData.IsLeapYear(
		year,
		CalendarYearNumType(0).Astronomical(),
		"")

	return isLeapYear
	//if year < 0 {
	//	year *= -1
	//}
	//
	//if year % int64(4) == 0 {
	//
	//	isLeapYear = true
	//
	//	if year % 100 == 0 {
	//
	//		isLeapYear = false
	//
	//		if  year % int64(400) == 0 {
	//			isLeapYear = true
	//		}
	//	}
	//
	//} else {
	//	isLeapYear = false
	//}
	//
	//return isLeapYear
}