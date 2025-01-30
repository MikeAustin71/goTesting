package datetime

import (
	"fmt"
	"sync"
)

type calendarGregorianBaseDataMechanics struct {
	lock    *sync.Mutex
}

// getDaysOfWeekNameAbbrvsISO8601 - Returns a map containing the abbreviated week
// day names indexed according to the ISO 8601 Standard Day Of The Week Numbering
// System. This method will return two or three character day of the week name
// abbreviations depending on input parameter 'numberOfCharsInAbbreviation'.
//
// Again, only two and three character week day name abbreviations are currently
// supported.
//
// ISO stands for the International Organization for Standardization (ISO).
//     https://www.iso.org/home.html
//
// The ISO 8601 Standard is the most common day of the week numbering system used
// internationally. This standard is used in Western Europe, Scandinavia, and most
// of Eastern Europe as well as many other nations across the globe.
//
// The ISO 8601 standard specifies that the week begins on Monday. Days of the week
// are numbered beginning with one (1) for Monday and ending with seven (7) for
// Sunday. Western European Calendars therefore show the first day of the week as
// Monday.
//
// The ISO 8601 Standard Day of the Week Numbering System and associated 2 and
// 3-character day of the week name abbreviations are listed as follows:
//
//     ISO 8601           Day of           Day of
//     Standard          Week Name        Week Name
//    Day of Week       Abbreviation     Abbreviation
//       Number           3-Chars          2-Chars
//   ===============   ==============   ==============
//         1                Mon              Mo
//         2                Tue              Tu
//         3                Wed              We
//         4                Thu              Th
//         5                Fri              Fr
//         6                Sat              Sa
//         7                Sun              Su
//
// Use the numbering scheme listed above when accessing the map data returned
// by this method.
//
// For more information on the ISO 8601 Standard Day Of The Week Numbering System,
// reference:
//   https://en.wikipedia.org/wiki/ISO_8601#Week_dates
//   https://www.timeanddate.com/date/week-numbers.html
//   Type: ISO8601DayOfWeekNo Source Code File: datetime/dayofweeknumberiso8601enum.go
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numberOfCharsInAbbreviation   int
//     - This integer value determines the number of characters in the day of
//       week name abbreviations returned by this method. Currently, only
//       two or three character day of the week name abbreviations are supported.
//       If the input parameter, 'numberOfCharsInAbbreviation', contains any value
//       other than '2' or '3', an error will be returned.
//
//
//  ePrefix                string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  ISO8601DaysOfWeekNameAbbrvs   map[int] string
//     - If the method completes successfully, a map containing the day of week
//       name abbreviations indexed according to the ISO 8601 Standard will be
//       returned. The abbreviated week day names will contain either two or
//       three character abbreviations depending on the value of input parameter,
//       'numberOfCharsInAbbreviation'.
//
//
//  err                           error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (gregCalBDataMech *calendarGregorianBaseDataMechanics) getDaysOfWeekNameAbbrvsISO8601(
	numberOfCharsInAbbreviation int,
	ePrefix string) (
	ISO8601DaysOfWeekNameAbbrvs map[int] string,
	err error) {

	if gregCalBDataMech.lock == nil {
		gregCalBDataMech.lock = new(sync.Mutex)
	}

	gregCalBDataMech.lock.Lock()

	defer gregCalBDataMech.lock.Unlock()

	if numberOfCharsInAbbreviation == 3 {

		ISO8601DaysOfWeekNameAbbrvs = map[int] string {
			1: "Mon",
			2: "Tue",
			3: "Wed",
			4: "Thu",
			5: "Fri",
			6: "Sat",
			7: "Sun",
		}

	} else if numberOfCharsInAbbreviation == 2 {

		ISO8601DaysOfWeekNameAbbrvs = map[int] string {
			1: "Mo",
			2: "Tu",
			3: "We",
			4: "Th",
			5: "Fr",
			6: "Sa",
			7: "Su",
		}

	} else {

		ISO8601DaysOfWeekNameAbbrvs = map[int] string {
			1: "ERROR",
		}

		err = fmt.Errorf(ePrefix + "\n" +
			"ERROR: Input parameter 'numberOfCharsInAbbreviation' is INVALID!\n" +
			"Valid number of characters in Day of Week Name Abbreviation is either '2' or '3'.\n" +
			"numberOfCharsInAbbreviation='%v'\n",
			numberOfCharsInAbbreviation)
	}

	return ISO8601DaysOfWeekNameAbbrvs, err
}

// getDaysOfWeekNameAbbrvsUS - Returns a map containing the abbreviated day of
// the week names indexed according to the US Day Of The Week Numbering System.
// This method will return two or three character day of the week day name
// abbreviations depending on input parameter 'numberOfCharsInAbbreviation'.
//
// Again, only two and three character week day name abbreviations are currently
// supported.
//
// The United States, Canada, Australia and New Zealand put Sunday as the
// first day of the week on their calendars. The first day of the week,
// 'Sunday', is numbered as zero with the last day of the week being numbered
// as 6, for 'Saturday'. This system is referred to as the US Day Of The Week
// Numbering System.
//
// The US Day Of The Week Numbering System and associated 2 and  3-character
// day of the week name abbreviations are listed as follows:
//
//        US              Day of           Day of
//     Standard          Week Name        Week Name
//    Day of Week       Abbreviation     Abbreviation
//       Number           3-Chars          2-Chars
//   ===============   ==============   ==============
//         0                Sun              Su
//         1                Mon              Mo
//         2                Tue              Tu
//         3                Wed              We
//         4                Thu              Th
//         5                Fri              Fr
//         6                Sat              Sa
//
// Use the numbering scheme listed above when accessing the map data returned
// by this method.
//
// For more information on the US Day Of The Week Numbering System, reference:
//   https://www.timeanddate.com/date/week-numbers.html
//   Type: UsDayOfWeekNo Source Code File: datetime/dayofweeknumberusenum.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numberOfCharsInAbbreviation   int
//     - This integer value determines the number of characters in the day of
//       week name abbreviations returned by this method. Currently, only
//       two or three character day of the week name abbreviations are supported.
//       If the input parameter, 'numberOfCharsInAbbreviation', contains any value
//       other than '2' or '3', an error will be returned.
//
//
//  ePrefix                string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  UsDaysOfWeekNameAbbrvs        map[int] string
//     - If the method completes successfully, a map containing the day of week
//       name abbreviations indexed according to the US Day Of The Week Numbering
//       system will be returned. The abbreviated week day names will contain either
//       two or three character abbreviations depending on the value of input parameter,
//       'numberOfCharsInAbbreviation'.
//
//
//  err                           error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (gregCalBDataMech *calendarGregorianBaseDataMechanics) getDaysOfWeekNameAbbrvsUS(
	numberOfCharsInAbbreviation int,
	ePrefix string) (
	UsDaysOfWeekNameAbbrvs map[int] string,
	err error) {

	if gregCalBDataMech.lock == nil {
		gregCalBDataMech.lock = new(sync.Mutex)
	}

	gregCalBDataMech.lock.Lock()

	defer gregCalBDataMech.lock.Unlock()

	if numberOfCharsInAbbreviation == 3 {

		UsDaysOfWeekNameAbbrvs = map[int] string {
			0: "Sun",
			1: "Mon",
			2: "Tue",
			3: "Wed",
			4: "Thu",
			5: "Fri",
			6: "Sat",
		}

	} else if numberOfCharsInAbbreviation == 2 {

		UsDaysOfWeekNameAbbrvs = map[int] string {
			0: "Su",
			1: "Mo",
			2: "Tu",
			3: "We",
			4: "Th",
			5: "Fr",
			6: "Sa",
		}

	} else {

		UsDaysOfWeekNameAbbrvs = map[int] string {
			1: "ERROR",
		}

		err = fmt.Errorf(ePrefix + "\n" +
			"ERROR: Input parameter 'numberOfCharsInAbbreviation' is INVALID!\n" +
			"Valid number of characters in Day of Week Name Abbreviation is either '2' or '3'.\n" +
			"numberOfCharsInAbbreviation='%v'\n",
			numberOfCharsInAbbreviation)
	}

	return UsDaysOfWeekNameAbbrvs, err
}


// getDaysOfWeekNamesISO8601 - Returns a map containing the day of week names
// indexed according to the ISO 8601 Standard Day Of The Week Numbering System.
//
// ISO stands for the International Organization for Standardization (ISO).
//     https://www.iso.org/home.html
//
// The ISO 8601 Standard is the most common day of the week numbering system used
// internationally. This standard is used in Western Europe, Scandinavia, and most
// of Eastern Europe as well as many other nations across the globe.
//
// The ISO 8601 standard specifies that the week begins on Monday. Days of the week
// are numbered beginning with one (1) for Monday and ending with seven (7) for
// Sunday. Western European Calendars therefore show the first day of the week as
// Monday.
//
//              Week         Week Day
//              Day           Number
//            ========     ===========
//             Monday          = 1
//             Tuesday         = 2
//             Wednesday       = 3
//             Thursday        = 4
//             Friday          = 5
//             Saturday        = 6
//             Sunday          = 7
//
// For more information on the ISO 8601 Standard Day Of The Week Numbering System,
// reference:
//   https://en.wikipedia.org/wiki/ISO_8601#Week_dates
//   https://www.timeanddate.com/date/week-numbers.html
//   Type: ISO8601DayOfWeekNo Source Code File: datetime/dayofweeknumberiso8601enum.go
//
func (gregCalBDataMech *calendarGregorianBaseDataMechanics) getDaysOfWeekNamesISO8601(
	) map[int] string {

	if gregCalBDataMech.lock == nil {
		gregCalBDataMech.lock = new(sync.Mutex)
	}

	gregCalBDataMech.lock.Lock()

	defer gregCalBDataMech.lock.Unlock()

	ISO8601DaysOfWeekNames := map[int] string {
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	return ISO8601DaysOfWeekNames
}

// getDaysOfWeekNamesUS - Returns a map containing the day of week names indexed
// according to the US Day Of The Week Numbering System.
//
// The United States, Canada, Australia and New Zealand put Sunday as the
// first day of the week on their calendars. The first day of the week,
// 'Sunday', is numbered as zero with the last day of the week being numbered
// as 6, for 'Saturday'.
//
// United States day of the week numbers as listed below:
//
//              Week         Week Day
//              Day           Number
//            ========     ===========
//             Sunday          = 0
//             Monday          = 1
//             Tuesday         = 2
//             Wednesday       = 3
//             Thursday        = 4
//             Friday          = 5
//             Saturday        = 6
//
// For more information on the US Day Of The Week Numbering System, reference:
//   https://www.timeanddate.com/date/week-numbers.html
//   Type: UsDayOfWeekNo Source Code File: datetime/dayofweeknumberusenum.go
//
func (gregCalBDataMech *calendarGregorianBaseDataMechanics) getDaysOfWeekNamesUS(
 ) map[int] string {

	if gregCalBDataMech.lock == nil {
		gregCalBDataMech.lock = new(sync.Mutex)
	}

	gregCalBDataMech.lock.Lock()

	defer gregCalBDataMech.lock.Unlock()

	USDaysOfWeekNames := map[int] string {
		0: "Sunday",
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
	}

	return USDaysOfWeekNames
}

// getLeapYearDaysInYear - Returns the number of days in a Leap Year under
// the Gregorian Calendar.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//       -- NONE --
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method returns the number of days in a Gregorian Calendar Year
//       Leap Year (366).
//
func (gregCalBDataMech *calendarGregorianBaseDataMechanics) getLeapYearDaysInYear() int {

	if gregCalBDataMech.lock == nil {
		gregCalBDataMech.lock = new(sync.Mutex)
	}

	gregCalBDataMech.lock.Lock()

	defer gregCalBDataMech.lock.Unlock()

	return 366
}

// getLeapYearMonthDays - Returns a map containing the number of days in
// each month month for a Gregorian Calendar Leap Year. The key is the
// integer month number and the value is the number of days in that month
// number.
//
// Remember, these month days apply only to Leap Years on the Gregorian
// Calendar.
//
// For more information on the Gregorian Calendar, reference:
//    https://en.wikipedia.org/wiki/Gregorian_calendar
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//      --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  map[int]int
//     - If this method completes successfully, it will return a map implementing
//       an 'int' key and an 'int' value. The key represents the month numbers
//       in a Gregorian Calendar Year, 1 through 12. The returned 'int' value
//       represents the corresponding number of days in that month number within
//       a Gregorian Calendar Leap Year.
//
//
func (gregCalBDataMech *calendarGregorianBaseDataMechanics) getLeapYearMonthDays(
) map[int]int {

	if gregCalBDataMech.lock == nil {
		gregCalBDataMech.lock = new(sync.Mutex)
	}

	gregCalBDataMech.lock.Lock()

	defer gregCalBDataMech.lock.Unlock()

	leapYearMthDays := map[int]int {
		1: 31,
		2: 29,
		3: 31,
		4: 30,
		5: 31,
		6: 30,
		7: 31,
		8: 31,
		9: 30,
		10: 31,
		11: 30,
		12: 31,
	}

	return leapYearMthDays
}

// getLeapYearOrdinalDays - This method returns a map containing the
// number of ordinal days at the beginning of each month within a
// Gregorian Calendar Leap Year. The key is the integer month number
// and the value is the number of ordinal days elapsed during the year
// at the beginning of that month.
//
// Remember, these ordinal days apply only to Leap Years on the
// Gregorian Calendar.
//
// For more information on the Gregorian Calendar, reference:
//    https://en.wikipedia.org/wiki/Gregorian_calendar
//
// For more information on Ordinal Day Numbers, reference:
//    https://en.wikipedia.org/wiki/Ordinal_date
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   -- NONE --
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  map[int]int
//     - This method will return a map implementing an 'int' key and an
//       'int' value. The key represents the month numbers in a
//       Gregorian Calendar Leap Year (months 1 through 12). The returned
//       'int' value represents the corresponding number of ordinal days,
//       which have elapsed during a Gregorian Calendar Leap Year, at the
//       beginning of that month.
//
func (gregCalBDataMech *calendarGregorianBaseDataMechanics) getLeapYearOrdinalDays(
) map[int] int {

	if gregCalBDataMech.lock == nil {
		gregCalBDataMech.lock = new(sync.Mutex)
	}

	gregCalBDataMech.lock.Lock()

	defer gregCalBDataMech.lock.Unlock()

	leapYearOrdDays := map[int] int {
		1: 0,
		2: 31,
		3: 60,
		4: 91,
		5: 121,
		6: 152,
		7: 182,
		8: 213,
		9: 244,
		10: 274,
		11: 305,
		12: 335,
	}

	return leapYearOrdDays
}

// getStandardYearDaysInYear - Returns the number of days in a standard
// Gregorian Calendar Year.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//       -- NONE --
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method returns the number of days in a standard Gregorian
//       Calendar Year (365).
//
func (gregCalBDataMech *calendarGregorianBaseDataMechanics) getStandardYearDaysInYear() int {

	if gregCalBDataMech.lock == nil {
		gregCalBDataMech.lock = new(sync.Mutex)
	}

	gregCalBDataMech.lock.Lock()

	defer gregCalBDataMech.lock.Unlock()

	return 365
}

// getStandardYearMonthDays  - Returns a map containing the number of days
// in each month month for a Gregorian Calendar Standard Year. The key is
// the month number and the value is the number of days in that month
// number.
//
// Remember, these month days apply only to Standard Years on the Gregorian
// Calendar.
//
// For more information on the Gregorian Calendar, reference:
//    https://en.wikipedia.org/wiki/Gregorian_calendar
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//      --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  map[int]int
//     - If this method completes successfully, it will return a map implementing
//       an 'int' key and an 'int' value. The key represents the month numbers
//       in a Gregorian Calendar Year, 1 through 12. The returned 'int' value
//       represents the corresponding number of days in that month number within
//       a Gregorian Calendar Standard Year.
//
//
func (gregCalBDataMech *calendarGregorianBaseDataMechanics) getStandardYearMonthDays(
) map[int] int {

	if gregCalBDataMech.lock == nil {
		gregCalBDataMech.lock = new(sync.Mutex)
	}

	gregCalBDataMech.lock.Lock()

	defer gregCalBDataMech.lock.Unlock()

	stdYrMthDays := map[int] int {
		1: 31,
		2: 28,
		3: 31,
		4: 30,
		5: 31,
		6: 30,
		7: 31,
		8: 31,
		9: 30,
		10: 31,
		11: 30,
		12: 31,
	}

	return stdYrMthDays
}

// getStandardYearOrdinalDays - This method returns a map containing
// the number of ordinal days at the beginning of each month within
// a Gregorian Calendar Standard Year. The key is the integer month
// number and the value is the number of ordinal days elapsed during
// the year at the beginning of that month.
//
// Remember, these ordinal days apply only to Standard Years on the
// Gregorian Calendar.
//
// For more information on the Gregorian Calendar, reference:
//    https://en.wikipedia.org/wiki/Gregorian_calendar
//
// For more information on Ordinal Day Numbers, reference:
//    https://en.wikipedia.org/wiki/Ordinal_date
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   -- NONE --
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  map[int]int
//     - This method will return a map implementing an 'int' key and an
//       'int' value. The key represents the month numbers in a
//       Gregorian Calendar Standard Year (months 1 through 12). The
//       returned 'int' value represents the corresponding number of
//       ordinal days, which have elapsed during a Gregorian Calendar
//       Standard Year, at the beginning of that month.
//
func (gregCalBDataMech *calendarGregorianBaseDataMechanics) getStandardYearOrdinalDays(
) map[int] int {

	if gregCalBDataMech.lock == nil {
		gregCalBDataMech.lock = new(sync.Mutex)
	}

	gregCalBDataMech.lock.Lock()

	defer gregCalBDataMech.lock.Unlock()

	stdYearOrdDays := map[int] int {
		1:0,
		2:31,
		3:59,
		4:90,
		5:120,
		6:151,
		7:181,
		8:212,
		9:243,
		10:273,
		11:304,
		12: 334,
	}

	return stdYearOrdDays
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
// The algorithm for identifying Gregorian Calendar Leap Years requires
// an input value in 'years' formatted under the Astronomical Year Numbering
// System.
//
// Astronomical Year Numbering System
//
// The Astronomical Year Numbering System includes a year zero (0). In
// other words, the date January 1, year 1 is immediately preceded by
// the date December 31, year zero (0). Likewise, the date January 1,
// year 0 is immediately preceded by the date December 31, year -1.
//
// Years occurring before year zero are preceded with a minus sign
// (January 1, -1). To illustrate, years before year one (1) are
// numbered like this: 0, -1, -2, -3, -4, -5 etc. For more information
// on the the 'Astronomical Year Numbering System', reference:
//      https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
// In order to convert a 'year' value to the Astronomical Year Numbering
// System, the algorithm relies on the second input parameter, 'yearNumType'
// of type CalendarYearNumType.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  year                   int64
//     - A year value. This year value is classified by type as being either
//       an Astronomical Year, Before Common Era Year or a Common Era Year.
//       This year parameter is then converted to an Astronomical Year Value
//       and returned to the calling function.
//
//  yearNumType            CalendarYearNumType
//     - Year number type is an enumeration which determines the type of
//       conversion algorithm which will be applied to convert input parameter
//       'year' to an Astronomical Year Value. 'yearNumType' supports three
//       types of year values.
//
//       1. Astronomical Year Numbering
//            All years less than zero, equal to zero or greater than zero prefixed
//            with the appropriate numerical sign (minus sign for negative years).
//            For example, years prior to year 0 are numbered -1, -2, -3, -4 etc.
//
//       2. Before Common Era Numbering
//            All years preceding the date January 1, year 1. These years typically
//            have the suffix BCE, for Before Common Era. For example, the date
//            immediately preceding January 1, 0001 CE is December 31, 0001 BCE.
//            year preceding year 0001 BCE are numbered as 2 BCE, 3 BCE, 4 BCE etc.
//
//       3. Common Era Numbering
//            All years following, or greater than, the date December 31, 1 BCE. These
//            years often have the suffix CE signaling they are years in the Common Era.
//            For example years following the date December 31, 1 BCE are numbered 1 CE,
//            2 CE, 3 CE, 4 CE etc.
//
//       For more information on the 'Common Era Year Numbering System', reference:
//            https://en.wikipedia.org/wiki/Common_Era
//
//       For more information on CalendarYearNumType reference the source code file:
//            datetime/calendaryearnumbertypeenum.go
//
//
//  ePrefix                string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isLeapYear             bool
//     - This return value is set to 'false' if the input parameter, 'year',
//       does NOT qualify as a leap year. Otherwise this value is set to
//       'true' signaling the 'year' is a genuine leap year under the
//       Gregorian Calendar.
//
//
//  err                    error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (gregCalBDataMech *calendarGregorianBaseDataMechanics) isLeapYear(
	year int64,
	yearNumType CalendarYearNumType,
	ePrefix string) (
	isLeapYear bool,
	err error) {

	if gregCalBDataMech.lock == nil {
		gregCalBDataMech.lock = new(sync.Mutex)
	}

	gregCalBDataMech.lock.Lock()

	defer gregCalBDataMech.lock.Unlock()

	isLeapYear = false
	err = nil

	ePrefix += "calendarGregorianBaseDataMechanics.isLeapYear() "

	if ! yearNumType.XIsValid() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Input parameter 'yearNumType' is INVALID!\n" +
			"yearNumType='%v'\n",
			yearNumType.XValueInt())

		return isLeapYear, err
	}

	calMech := calendarMechanics{}

	var astronomicalYear int64

	astronomicalYear,
		err = calMech.convertAnyYearToAstronomicalYear(
		year,
		yearNumType,
		ePrefix)

	if err != nil {
		return isLeapYear, err
	}

	if astronomicalYear < 0 {
		astronomicalYear *= -1
	}

	if year % int64(4) == 0 {

		isLeapYear = true

		if year % 100 == 0 {

			isLeapYear = false

			if  year % int64(400) == 0 {
				isLeapYear = true
			}
		}

	} else {
		isLeapYear = false
	}

	return isLeapYear, err
}


