package datetime

import (
	"fmt"
	"math"
	"sync"
)

// CalendarGregorianBaseData - Implements the
// Calendar Base Data Interface.
//
type CalendarGregorianBaseData struct {
	lock *sync.Mutex
}

// GetCalendarSpecification - Returns the Calendar Specification ID for the
// Gregorian Calendar
func (gregCalBData *CalendarGregorianBaseData) GetCalendarSpecification() CalendarSpec {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	return CalendarSpec(0).Gregorian()
}

// GetISODayOfWeekNo - Returns the ISO Day Of The Week Number. This
// method receives a Julian Day Number and proceeds to calculate the day
// of the week number associated with that Julian Day Number Date. It
// then returns that day of the week number to the calling function
// formatted with the ISO 8601 Standard Day Of The Week Numbering System.
//
//
// ISO 8601 Standard Day Of The Week Numbering System
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
// The ISO 8601 Standard Day of the Week Numbering System is listed as follows:
//
//             ISO 8601
//             Standard
//            Day of Week       Day Of Week
//              Number             Name
//           ===============   ==============
//                 1              Monday
//                 2              Tuesday
//                 3              Wednesday
//                 4              Thursday
//                 5              Friday
//                 6              Saturday
//                 7              Sunday
//
// Use the numbering scheme listed above when accessing map data returned by
// this method for the ISO 8601 Standard Day Of The Week Numbering System.
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
//  julianDayNoDto      JulianDayNoDto
//     - This input parameter contains the data elements of a Julian Day
//       Number and Time value. Note that key Julian Day Number and Time
//       values are stored as *big.Int and *big.Float.
//
//        type JulianDayNoDto struct {
//          julianDayNo             *big.Int    // Julian Day Number expressed as integer value
//          julianDayNoFraction     *big.Float  // The Fractional Time value of the Julian
//                                              //   Day No Time
//          julianDayNoTime         *big.Float  // Julian Day Number Plus Time Fraction accurate to
//                                              //   within nanoseconds
//          julianDayNoNumericalSign int        // Sign of the Julian Day Number/Time value. This value
//                                              //   is either '+1' or '-1'
//          totalJulianNanoSeconds   int64      // Julian Day Number Time Value expressed in nanoseconds.
//                                              //   Always represents a positive value less than 36-hours
//          netGregorianNanoSeconds  int64      // Gregorian nanoseconds. Always represents a value in
//                                              //    nanoseconds which is less than 24-hours.
//          applyLeapSecond    bool // If set to 'true' it signals that the day identified
//                                  //   by this Julian Day Number has a duration go 24-hours
//                                  //   + 1-second.
//          hours               int // Gregorian Hours
//          minutes             int // Gregorian Minutes
//          seconds             int // Gregorian Seconds
//          nanoseconds         int // Gregorian Nanoseconds
//        }
//
//        The integer portion of the Julian Day Number (digits to left of
//        the decimal) represents the Julian day number and is
//        stored in 'JulianDayNoDto.julianDayNo'. The fractional
//        digits to the right of the decimal represent elapsed time
//        since noon on the Julian day number and is stored in
//        'JulianDayNoDto.julianDayNoFraction'. The combined Julian
//        Day Number Time value is stored in 'JulianDayNoDto.julianDayNoTime'.
//        All time values are expressed as Universal Coordinated Time (UTC).
//
//
//  ePrefix                       string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isoDayOfWeekNo      ISO8601DayOfWeekNo
//     - If the method completes successfully, this parameter will contain
//       an enumeration type specifying the day of the week number formatted
//       according to the ISO 8601 Standard Day Of The Week Numbering System.
//       This enumeration value will allow access to the correct day number,
//       day name and day name abbreviation. The day of the week number is
//       calculated from input parameter 'julianDayNoDto'.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note this error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'.
//
func (gregCalBData *CalendarGregorianBaseData) GetISODayOfWeekNo(
	julianDayNoDto JulianDayNoDto,
	ePrefix string) (
	isoDayOfWeekNo ISO8601DayOfWeekNo,
	err error) {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	ePrefix += "CalendarGregorianBaseData.GetISODayOfWeekNo() "

	gregCalBDataUtil := calendarGregorianBaseDataUtility{}

	isoDayOfWeekNo,
		err = gregCalBDataUtil.getISODayOfWeekNumber(
		julianDayNoDto,
		ePrefix)

	return isoDayOfWeekNo, err
}

// GetUsDayOfWeekNo - Returns the US Day Of The Week Number. This method
// receives a Julian Day Number and proceeds to calculate the day of the
// week number associated with that Julian Day Number Date. It then returns
// that day of the week number to the calling function formatted with the
// US Day Of The Week Numbering System.
//
//
// US Day Of The Week Numbering System
//
// The United States, Canada, Australia and New Zealand put Sunday as the first
// day of the week on their calendars. The first day of the week, 'Sunday', is
// numbered as zero (0) with the last day of the week being numbered as six (6),
// for 'Saturday'. This system is referred to as the US Day Of The Week Numbering
// System and is listed as follows:
//
//                US
//            Day of Week        Day Of Week
//               Number             Name
//           ===============   ==============
//                 0               Sunday
//                 1               Monday
//                 2               Tuesday
//                 3               Wednesday
//                 4               Thursday
//                 5               Friday
//                 6               Saturday
//
// Use the numbering scheme listed above when accessing map data returned by
// this method for the US Day Of The Week Numbering System.
//
// For more information on the US Day Of The Week Numbering System, reference:
//   https://www.timeanddate.com/date/week-numbers.html
//   Type: UsDayOfWeekNo Source Code File: datetime/dayofweeknumberusenum.go
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  julianDayNoDto      JulianDayNoDto
//     - This input parameter contains the data elements of a Julian Day
//       Number and Time value. Note that key Julian Day Number and Time
//       values are stored as *big.Int and *big.Float.
//
//        type JulianDayNoDto struct {
//          julianDayNo             *big.Int    // Julian Day Number expressed as integer value
//          julianDayNoFraction     *big.Float  // The Fractional Time value of the Julian
//                                              //   Day No Time
//          julianDayNoTime         *big.Float  // Julian Day Number Plus Time Fraction accurate to
//                                              //   within nanoseconds
//          julianDayNoNumericalSign int        // Sign of the Julian Day Number/Time value. This value
//                                              //   is either '+1' or '-1'
//          totalJulianNanoSeconds   int64      // Julian Day Number Time Value expressed in nanoseconds.
//                                              //   Always represents a positive value less than 36-hours
//          netGregorianNanoSeconds  int64      // Gregorian nanoseconds. Always represents a value in
//                                              //    nanoseconds which is less than 24-hours.
//          applyLeapSecond    bool // If set to 'true' it signals that the day identified
//                                  //   by this Julian Day Number has a duration go 24-hours
//                                  //   + 1-second.
//          hours               int // Gregorian Hours
//          minutes             int // Gregorian Minutes
//          seconds             int // Gregorian Seconds
//          nanoseconds         int // Gregorian Nanoseconds
//        }
//
//        The integer portion of the Julian Day Number (digits to left of
//        the decimal) represents the Julian day number and is
//        stored in 'JulianDayNoDto.julianDayNo'. The fractional
//        digits to the right of the decimal represent elapsed time
//        since noon on the Julian day number and is stored in
//        'JulianDayNoDto.julianDayNoFraction'. The combined Julian
//        Day Number Time value is stored in 'JulianDayNoDto.julianDayNoTime'.
//        All time values are expressed as Universal Coordinated Time (UTC).
//
//
//  ePrefix                       string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  usDayOfWeekNo       UsDayOfWeekNo
//     - If the method completes successfully, this parameter will contain
//       an enumeration type specifying the day of the week number formatted
//       according to the US Day Of The Week Numbering System. This enumeration
//       value will allow access to the correct day number, day name and day
//       name abbreviation. The day of the week number is calculated from
//       input parameter 'julianDayNoDto'.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note this error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'.
//
func (gregCalBData *CalendarGregorianBaseData) GetUsDayOfWeekNo(
	julianDayNoDto JulianDayNoDto,
	ePrefix string) (
	usDayOfWeekNo UsDayOfWeekNo,
	err error) {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	ePrefix += "CalendarGregorianBaseData.GetUSDayOfWeekNo() "

	gregCalBDataUtil := calendarGregorianBaseDataUtility{}

	usDayOfWeekNo,
		err = gregCalBDataUtil.getUsDayOfWeekNumber(
		julianDayNoDto,
		ePrefix)

	return usDayOfWeekNo, err
}


// GetDaysInLeapYear - Returns the number of days in a
// Gregorian Calendar Leap Year.
//
func (gregCalBData *CalendarGregorianBaseData) GetDaysInLeapYear() int {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()


	return 366
}

// GetDaysInStandardYear - Returns the number of days in a
// Gregorian Calendar Standard Year or non-leap year.
//
func (gregCalBData *CalendarGregorianBaseData) GetDaysInStandardYear() int {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	return 365
}

// GetDaysInYear - Returns the number of days for 'year' value supplied
// as an input parameter.
//
// The algorithm for identifying identifying the number of days in a year
// under the Gregorian Calendar, requires 'year' value formatted under the
// Astronomical Year Numbering System.
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
// In order to convert the input parameter 'year' value to a a value
// formatted in the Astronomical Year Numbering System, the algorithm
// relies on the second input parameter, 'yearNumType' of type
// CalendarYearNumType.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  year                   int64
//     - A year value. This year value is classified by type as being either
//       an Astronomical Year, Before Common Era Year or a Common Era Year
//       based on the second input parameter 'yearNumType'. This 'year'
//       parameter is then converted to an Astronomical Year Value for
//       processing purposes.
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
//  daysInYear             int
//     - If this method completes successfully, the number of days in the
//       the year value supplied by input parameter, 'year', will be
//       returned.
//
//
//  err                    error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func  (gregCalBData *CalendarGregorianBaseData) GetDaysInYear(
	year int64,
	yearNumType CalendarYearNumType,
	ePrefix string)  (
	daysInYear int,
	err error) {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	gregCalBDataMech :=
		calendarGregorianBaseDataMechanics{}

	var isLeapYear bool

	isLeapYear, err = gregCalBDataMech.isLeapYear(
		year,
		yearNumType,
		ePrefix)

	if isLeapYear {

		daysInYear = gregCalBDataMech.getLeapYearDaysInYear()

	} else {

		daysInYear = gregCalBDataMech.getStandardYearDaysInYear()

	}

	return daysInYear, err
}


// GetDaysOfWeekNames - This method returns a map consisting of the
// days of the week names.
//
// The returned map containing the week day name abbreviations will be
// indexed according to one of two Day Of The Week naming conventions:
//
//    (1) US Day Of The Week Numbering System
//                 OR
//    (2) The ISO 8601 Standard Day Of The Week Numbering System
//
// The selection of one of these two numbering systems is determined by the
// enumeration value of input parameter, 'dayOfWeekNoSysType'.
//
//
// US Day Of The Week Numbering System
//
// The United States, Canada, Australia and New Zealand put Sunday as the
// first day of the week on their calendars. The first day of the week,
// 'Sunday', is numbered as zero with the last day of the week being numbered
// as 6, for 'Saturday'. This system is referred to as the US Day Of The Week
// Numbering System.
//
// The US Day Of The Week Numbering System and associated day of the week
// names are listed as follows:
//
//                 US
//             Day of Week      Day Of The Week
//                Number             Names
//            ===============   ===============
//                  0                Sunday
//                  1                Monday
//                  2                Tuesday
//                  3                Wednesday
//                  4                Thursday
//                  5                Friday
//                  6                Saturday
//
// Use the numbering scheme listed above when accessing map data returned by
// this method for the US Day Of The Week Numbering System.
//
// For more information on the US Day Of The Week Numbering System, reference:
//   https://www.timeanddate.com/date/week-numbers.html
//   Type: UsDayOfWeekNo Source Code File: datetime/dayofweeknumberusenum.go
//
//
// ISO 8601 Standard Day Of The Week Numbering System
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
// The ISO 8601 Standard Day of the Week Numbering System and associated Day Of
// The Week names are listed as follows:
//
//             ISO 8601
//             Standard
//            Day of Week      Day Of The Week
//               Number             Names
//            ===============   ===============
//                  1               Monday
//                  2               Tuesday
//                  3               Wednesday
//                  4               Thursday
//                  5               Friday
//                  6               Saturday
//                  7               Sunday
//
// Use the numbering scheme listed above when accessing map data returned by
// this method for the ISO 8601 Standard Day Of The Week Numbering System.
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
//  dayOfWeekNoSysType            DayOfWeekNumberingSystemType
//     - This parameter is an enumeration of valid Day Of The Week Numbering
//       Systems. It MUST be set to one of two values:
//
//         (1) DayOfWeekNumberingSystemType(0).UsDayOfWeek()
//                             OR
//         (2) DayOfWeekNumberingSystemType(0).ISO8601DayOfWeek()
//
//       For more information on type DayOfWeekNumberingSystemType see the
//       source code file documentation:
//               datetime/dayofweeknumberingsystemtypeenum.go
//
//       'dayOfWeekNoSysType' is used by this method to determine the indexing
//       used to access the Day Of The Week Names contained in the map data
//       returned by this method. The indexing sequences for both the US and
//       ISO 8601 Day Of The Week Numbering Systems are discussed above.
//       If 'dayOfWeekNoSysType' is set to any value other than the two
//       listed above, an error will be triggered.
//
//
//  ePrefix                       string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  daysOfWeekNames               map[int]string
//     - If this method completes successfully, this map will contain Days
//       Of The Week Names indexed according to the input parameter,
//       'dayOfWeekNoSysType'.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note this error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'.
//
func (gregCalBData *CalendarGregorianBaseData) GetDaysOfWeekNames(
	dayOfWeekNoSysType DayOfWeekNumberingSystemType,
	ePrefix string) (
	daysOfWeekNames map[int]string,
	err error) {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	ePrefix += "CalendarGregorianBaseData.GetDaysOfWeekNames() "

	daysOfWeekNames = map[int] string {
		0 : "Error",
	}

	gregCalBDataMech := calendarGregorianBaseDataMechanics{}

	switch dayOfWeekNoSysType {

	case DayOfWeekNumberingSystemType(0).UsDayOfWeek():

		daysOfWeekNames = gregCalBDataMech.getDaysOfWeekNamesUS()

	case DayOfWeekNumberingSystemType(0).ISO8601DayOfWeek():

		daysOfWeekNames = gregCalBDataMech.getDaysOfWeekNamesISO8601()

	default:

		err = fmt.Errorf(ePrefix + "\n" +
			"ERROR: Input parameter 'dayOfWeekNoSysType' is INVALID!\n" +
			"dayOfWeekNoSysType Value ='%v'\n",
			dayOfWeekNoSysType.XValueInt())

	}

	return daysOfWeekNames, err
}

// GetDaysOfWeekNameAbbreviations - Returns a map containing the abbreviated week
// day name abbreviations. These week day name abbreviations will contain either
// two (2) or three (3) character week day name abbreviations depending on the
// value of input parameter 'numberOfCharsInAbbreviation'.
//
// The returned map containing the week day name abbreviations will be indexed
// according to one of two Day Of The Week naming conventions:
//
//    (1) US Day Of The Week Numbering System
//                 OR
//    (2) The ISO 8601 Standard Day Of The Week Numbering System
//
// The selection of one of these two numbering systems is determined by the
// enumeration value of input parameter, 'dayOfWeekNoSysType'.
//
//
// US Day Of The Week Numbering System
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
//                        Day of           Day of
//        US             Week Name        Week Name
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
// Use the numbering scheme listed above when accessing map data returned by
// this method for the US Day Of The Week Numbering System.
//
// For more information on the US Day Of The Week Numbering System, reference:
//   https://www.timeanddate.com/date/week-numbers.html
//   Type: UsDayOfWeekNo Source Code File: datetime/dayofweeknumberusenum.go
//
//
// ISO 8601 Standard Day Of The Week Numbering System
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
// Use the numbering scheme listed above when accessing map data returned by
// this method for the ISO 8601 Standard Day Of The Week Numbering System.
//
// For more information on the ISO 8601 Standard Day Of The Week Numbering System,
// reference:
//   https://en.wikipedia.org/wiki/ISO_8601#Week_dates
//   https://www.timeanddate.com/date/week-numbers.html
//   Type: ISO8601DayOfWeekNo Source Code File: datetime/dayofweeknumberiso8601enum.go
//
//
// IMPORTANT
//
// Again, remember that only two and three character week day name abbreviations are
// currently supported.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dayOfWeekNoSysType            DayOfWeekNumberingSystemType
//     - This parameter is an enumeration of valid Day Of The Week Numbering
//       Systems. It MUST be set to one of two values:
//
//         (1) DayOfWeekNumberingSystemType(0).UsDayOfWeek()
//                             OR
//         (2) DayOfWeekNumberingSystemType(0).ISO8601DayOfWeek()
//
//       For more information on type DayOfWeekNumberingSystemType see the
//       source code file documentation:
//               datetime/dayofweeknumberingsystemtypeenum.go
//
//
//  numberOfCharsInAbbreviation   int
//     - This integer value determines the number of characters in the day of
//       week name abbreviations returned by this method. Currently, only
//       two or three character day of the week name abbreviations are supported.
//       If the input parameter, 'numberOfCharsInAbbreviation', contains any value
//       other than '2' or '3', an error will be returned.
//
//
//  ePrefix                       string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  weekDayNameAbbrvs             map[int] string
//     - If the method completes successfully, a map containing the day of week
//       name abbreviations indexed according to the value of input parameter,
//       'dayOfWeekNoSysType', will be returned. The abbreviated week day names
//       will contain either two or three character abbreviations depending on
//       the value of input parameter, 'numberOfCharsInAbbreviation'.
//
//
//  err                           error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (gregCalBData *CalendarGregorianBaseData) GetDaysOfWeekNameAbbreviations(
	dayOfWeekNoSysType DayOfWeekNumberingSystemType,
	numberOfCharsInAbbreviation int,
	ePrefix string) (
	weekDayNameAbbrvs map[int] string,
	err error) {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	ePrefix += "CalendarGregorianBaseData.GetDaysOfWeekNameAbbreviations() "

	weekDayNameAbbrvs = map[int] string {
		0 : "Error",
	}

	gregCalBDataMech := calendarGregorianBaseDataMechanics{}

	switch dayOfWeekNoSysType {

	case DayOfWeekNumberingSystemType(0).UsDayOfWeek():

		weekDayNameAbbrvs,
		err = gregCalBDataMech.getDaysOfWeekNameAbbrvsUS(
			numberOfCharsInAbbreviation,
			ePrefix)

	case DayOfWeekNumberingSystemType(0).ISO8601DayOfWeek():

		weekDayNameAbbrvs,
			err = gregCalBDataMech.getDaysOfWeekNameAbbrvsISO8601(
			numberOfCharsInAbbreviation,
			ePrefix)

	default:
		err = fmt.Errorf(ePrefix + "\n" +
			"ERROR: Input parameter 'dayOfWeekNoSysType' is INVALID!\n" +
			"dayOfWeekNoSysType='%v'\n",
			dayOfWeekNoSysType.XValueInt())

	}

	return weekDayNameAbbrvs, err
}


// GetLeapYearOrdinalDays - This method returns a map containing the
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
func (gregCalBData *CalendarGregorianBaseData) GetLeapYearOrdinalDays(
	) map[int] int {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	gregCalBDataMech := calendarGregorianBaseDataMechanics{}

	return gregCalBDataMech.getLeapYearOrdinalDays()
}

// GetLeapYearMonthDays - This method returns a map containing the
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
func (gregCalBData *CalendarGregorianBaseData) GetLeapYearMonthDays(
	) map[int] int {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	gregCalBDataMech := calendarGregorianBaseDataMechanics{}

	return gregCalBDataMech.getLeapYearOrdinalDays()
}

// GetMonthDayFromOrdinalDayNo - Receives an Ordinal Day Number and returns
// the associated month and day number. The input parameter 'isLeapYear'
// specifies whether the Ordinal Day Number is included in a standard year
// (365-Days) or a Leap Year (366-Days).
//
// The return value 'yearAdjustment' this value will be populated with one
// of three values: Plus One (+1), Zero (0) or Minus One (-1).
//
//        1. Plus One (+1)       -
//             A value of plus one (+1) signals that the ordinal day
//             day number represents the first day of the next year
//             (January 1st of the next year). Effectively the year
//             value is equal to year + 1, month is equal to '1' and
//             the day value is equal to '1'.
//
//       2.  Zero (0)       -
//             A value of Zero signals that no year adjustment is
//             required. The ordinal day number was converted to
//             month and day number in the current year.
//
//       3.  Minus One (-1) -
//             A value of Minus One indicates that the ordinal date
//             represents December 31st of the prior year. 'year'
//             is therefore equal to year - 1, month = 12 and day = 31.
//
// For more information on Ordinal Date, reference:
//    https://en.wikipedia.org/wiki/Ordinal_date
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  ordinalDate        int64
//     - A value with a valid range of 1-366 inclusive which specifies
//       the day of the year expressed as an ordinal day number or ordinal
//       date. For more information on 'ordinal dates', reference:
//         https://en.wikipedia.org/wiki/Ordinal_date
//
//
//  isLeapYear         bool
//     - If 'true' it signals that the input parameter 'ordinalDate' represents
//       an ordinal date within a leap year.
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  yearAdjustment     int
//     - If the method completes successfully, this value will
//       be populated with one of three values:
//
//           One (+1)       -
//             A value of plus one (+1) signals that the ordinal day
//             day number represents the first day of the next year
//             (January 1st of the next year). Effectively the year
//             value is equal to year + 1, month is equal to '1' and
//             the day value is equal to '1'.
//
//           Zero (0)       -
//             A value of Zero signals that no year adjustment is
//             required. The ordinal day number was converted to
//             month and day number in the current year.
//
//           Minus One (-1) -
//             A value of Minus One indicates that the ordinal date
//             represents December 31st of the prior year. 'year'
//             is therefore equal to year - 1, month = 12 and day = 31.
//
//  month             int
//     - If the method completes successfully, this value
//       will contain the month number associated with the
//       input parameter 'ordinalDate'.
//
//
//  day                int
//     - If successful this value will contain the day number
//       associated with the input parameter, 'ordinalDate'.
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
func (gregCalBData *CalendarGregorianBaseData) GetMonthDayFromOrdinalDayNo(
	ordinalDate int,
	isLeapYear bool,
	ePrefix string)(
	yearAdjustment int,
	month int,
	day int,
	err error) {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	ePrefix += "CalendarGregorianBaseData.GetMonthDayFromOrdinalDayNo() "

	calGregBDataUtil :=
		calendarGregorianBaseDataUtility{}

	yearAdjustment,
	month,
	day,
	err = calGregBDataUtil.getMonthDayFromOrdinalDayNo(
		ordinalDate,
		isLeapYear,
		ePrefix)

	return yearAdjustment, month, day, err
}

// GetOrdinalDayNumber - Computes the ordinal day number within a
// year for any given month and day under the Gregorian Calendar.
// Input parameter 'isLeapYear' indicates whether the year encompassing
// the specified month and day is a 'leap year' containing 366-days
// instead of the standard year containing 365-days.
//
// Reference
//    https://en.wikipedia.org/wiki/Ordinal_date
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  isLeapYear          bool
//     - If set to 'true' this boolean value signals that the year
//       encompassing the 'month' and 'day' input parameters is a
//       Gregorian Calendar leap year containing 366-days. If this
//       parameter is set to 'false', it signals that the year
//       encompassing the 'month' and 'day' input parameters is a
//       Gregorian Calendar standard year containing 365-days.
//
//
//  month               int
//     - The month number of the month encompassing the Ordinal Day
//       Number.
//
//
//  day                 int
//     - The day number within parameter 'month' which designates
//       Ordinal Day to be calculated.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  ordinalDayNo        int
//     - The Ordinal Day Number of the month and day specified as
//       input parameters.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func  (gregCalBData *CalendarGregorianBaseData) GetOrdinalDayNumber(
	isLeapYear bool,
	month int,
	day int,
	ePrefix string) (
	ordinalDayNo int,
	err error) {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	ePrefix += "CalendarGregorianBaseData.GetOrdinalDayNumber() "

	calGregorianBaseDataUtil :=
		calendarGregorianBaseDataUtility{}

	ordinalDayNo,
	err = calGregorianBaseDataUtil.getOrdinalDayNumber(
		isLeapYear,
		month,
		day,
		ePrefix)

	return ordinalDayNo, err
}

// GetOrdinalDayNoFromDate - Computes the ordinal day number within a
// year for any given month and day under the Gregorian Calendar. This
// method is identical to method 'CalendarGregorianBaseData.GetOrdinalDayNumber()'
// with the sole exception that this method takes input parameters of
// 'year' and 'yearType'.
//
// Referred to variously as ordinal day number or ordinal date, the
// returned ordinal day number value represents the sequential day number
// within the designated year. For more information on 'Ordinal Day Number',
// reference:
//    https://en.wikipedia.org/wiki/Ordinal_date
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  year                     int64
//     - An int64 value containing the year number. Note that this year
//       value should be interpreted in the context the year type return
//       parameter discussed below. This year value could be formatted
//       according to the Astronomical Year Numbering System or the
//       Common Era Year Numbering System.
//
//
//  yearType                 CalendarYearNumType
//     - The year number type associated with the return value 'yearValue'
//       described above. 'yearType' classifies return parameter 'yearValue'
//       as one of three year types:
//
//         1. Astronomical Year
//         2. BCE - Before Common Era
//         3. CE  - Common Era
//
//       For more information on Astronomical and Common Era Year
//       Numbering, reference:
//           Source File: datetime\calendaryearnumbertypeenum.go
//           https://en.wikipedia.org/wiki/Astronomical_year_numbering
//           https://en.wikipedia.org/wiki/Common_Era
//
//
//  month               int
//     - The month number of the month encompassing the Ordinal Day
//       Number.
//
//
//  day                 int
//     - The day number within parameter 'month' which designates
//       Ordinal Day to be calculated.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  ordinalDayNo        int
//     - The Ordinal Day Number of the month and day specified as
//       input parameters.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//

func (gregCalBData *CalendarGregorianBaseData) GetOrdinalDayNoFromDate(
	year int64,
	yearType CalendarYearNumType,
	month int,
	day int,
	ePrefix string) (
	ordinalDayNo int,
	err error) {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	ePrefix += "CalendarGregorianBaseData.GetOrdinalDayNoFromDate() "

	ordinalDayNo = math.MinInt32

	gregCalBDataMech :=
		calendarGregorianBaseDataMechanics{}

	var isLeapYear bool

	isLeapYear,
		err = gregCalBDataMech.isLeapYear(
		year,
		yearType,
		ePrefix)

	if err != nil {
		return ordinalDayNo, err
	}

	calGregorianBaseDataUtil :=
		calendarGregorianBaseDataUtility{}

	ordinalDayNo,
		err = calGregorianBaseDataUtil.getOrdinalDayNumber(
		isLeapYear,
		month,
		day,
		ePrefix)

	return ordinalDayNo, err
}

// GetRemainingDaysInYear - Returns a positive integer value identifying
// the number of days remaining in the year based on the date/time encapsulated
// by the current ADateTimeDto instance.
//
// This calculation is performed by subtracting the date/time's ordinal day
// day number from the number of days in the current year. The accuracy of
// the result depends largely on the accuracy of the 'isLeap' year setting
// for the current ADateTimeDto instance. The 'isLeap' setting can
// be controlled through instance method, 'SetIsLeapYear()'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  year                int64
//     - The year value defined by astronomical year numbering. Astronomical
//       year numbering incorporates a zero year or year '0'. For more
//       information on astronomical year numbering reference:
//           https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//
//  month               int
//     - The month number.
//
//
//  day                 int
//     - The day number.
//
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note this error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'.
//
//
func (gregCalBData *CalendarGregorianBaseData) GetRemainingDaysInYear(
	year int64,
	yearNumType CalendarYearNumType,
	month int,
	day int,
	ePrefix string) (
	remainingDaysOfYear int,
	err error) {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	ePrefix += "CalendarGregorianBaseData.GetRemainingDaysInYear() "

	remainingDaysOfYear = -1

	err = nil

	gregCalBDataMech :=
		calendarGregorianBaseDataMechanics{}

	var isLeapYear bool

	isLeapYear,
	err = gregCalBDataMech.isLeapYear(
		year,
		yearNumType,
		ePrefix)

	if err != nil {
		return remainingDaysOfYear, err
	}

	var daysInYear int

	if isLeapYear {
		daysInYear =
			gregCalBDataMech.getLeapYearDaysInYear()
	} else {
		daysInYear =
			gregCalBDataMech.getStandardYearDaysInYear()
	}

	var ordinalDayNo int

	gregCalBDataUtil := calendarGregorianBaseDataUtility{}

	ordinalDayNo, err = gregCalBDataUtil.getOrdinalDayNumber(
		isLeapYear,
		month,
		day,
		ePrefix)

	if err != nil {
		return remainingDaysOfYear, err
	}

	remainingDaysOfYear = daysInYear - ordinalDayNo

	return remainingDaysOfYear, err
}

// GetStandardYearMonthDays  - Returns a map containing the number of days
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
func (gregCalBData *CalendarGregorianBaseData) GetStandardYearMonthDays(
	) map[int] int {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	gregCalBDataMech := calendarGregorianBaseDataMechanics{}

	return gregCalBDataMech.getStandardYearMonthDays()
}


// GetStandardYearOrdinalDays  - Returns a map containing the number of days
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
func (gregCalBData *CalendarGregorianBaseData) GetStandardYearOrdinalDays(
	) map[int] int {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	gregCalBDataMech := calendarGregorianBaseDataMechanics{}

	return gregCalBDataMech.getStandardYearOrdinalDays()
}

// GetYearMonthDayFromOrdinalDayNo - Receives an Ordinal Day Number and returns
// the associated year, month and day numbers. Note that tye returned year value
// is expressed using Astronomical Year Numbering.
//
// For more information on Ordinal Date, reference:
//    https://en.wikipedia.org/wiki/Ordinal_date
//
// For more information on Astronomical Year Numbering, reference:
//    https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ordinalDate         int64
//     - A value with a valid range of 1-366 inclusive which specifies
//       the day of the year expressed as an ordinal day number or ordinal
//       date. For more information on 'ordinal dates', reference:
//         https://en.wikipedia.org/wiki/Ordinal_date
//
//
//  year                int64
//     - The year associated with the input parameter, 'ordinalDate'.
//       This year value is classified by type as being either an
//       Astronomical Year, Before Common Era Year or a Common Era Year.
//       This method will convert this year value to an Astronomical Year
//       value for use in the Ordinal Day Number conversion algorithm.
//       The year number type for this year value is determined by input
//       parameter yearNumType CalendarYearNumType.
//
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
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  astronomicalYear    int64
//     - This is the year value associated with return values
//       'month' and 'day' returned below. This year value is
//       always formatted as an Astronomical Year value meaning
//       that conforms to the Astronomical Year Numbering System.
//
//       The Astronomical Year Numbering System includes a year
//       zero (0). In other words, the date January 1, year 1 is
//       immediately preceded by the date December 31, year zero (0).
//       Likewise, the date January 1, year 0 is immediately preceded
//       by the date December 31, year -1. Years occurring before year
//       zero (0) are preceded with a minus sign (January 1, -1). To
//       illustrate, years before year one (1) are numbered as follows:
//                  0, -1, -2, -3, -4, -5 etc.
//       For more information on the the Astronomical Year Numbering
//       System, reference:
//            https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//
//  month             int
//     - If the method completes successfully, this value
//       will contain the month number associated with the
//       input parameter 'ordinalDate'.
//
//
//  day                int
//     - If successful this value will contain the day number
//       associated with the input parameter, 'ordinalDate'.
//
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
func (gregCalBData *CalendarGregorianBaseData) GetYearMonthDayFromOrdinalDayNo(
	ordinalDate int,
	year int64,
	yearNumType CalendarYearNumType,
	ePrefix string)(
	astronomicalYear int64,
	month int,
	day int,
	err error) {

	astronomicalYear = math.MinInt64
	month = -1
	day = -1
	err = nil

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	ePrefix += "CalendarGregorianBaseData.GetMonthDayFromOrdinalDayNoYear() "

	gregCalBDataUtil := calendarGregorianBaseDataUtility{}

	astronomicalYear,
	month,
	day,
	err = gregCalBDataUtil.getYearMonthDayFromOrdinalDayNo(
		ordinalDate,
		year,
		yearNumType,
		ePrefix)

	return astronomicalYear, month, day, err
}

// IsLeapYear - Returns a boolean value signaling whether the year
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
func (gregCalBData *CalendarGregorianBaseData) IsLeapYear(
	year int64,
	yearNumType CalendarYearNumType,
	ePrefix string) (
	isLeapYear bool,
	err error ) {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	ePrefix += "CalendarGregorianBaseData.IsLeapYear() "

	gregCalBDataMech :=
		calendarGregorianBaseDataMechanics{}

	isLeapYear,
	err = gregCalBDataMech.isLeapYear(
		year,
		yearNumType,
		ePrefix)

	return isLeapYear, err
}

// IsValidDate - This method will evaluate date input parameters to
// determine if the date is valid under the Gregorian Calendar. If
// the date is 'invalid', this method will return a boolean value
// of 'false' plus an error type containing a detailed error message
// identifying the invalid member variable.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  year                   int64
//     - The date's year value. This year value is classified by type as being either
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
//  month                  int
//     - The month number element of the date.
//
//  day                    int
//     - The day number within the 'month' input parameter.
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
//  isValid                bool
//     - If the date comprised of input parameters, 'year', 'month' and
//       'day' is valid under the Gregorian Calendar, this return value
//       is set to 'true'. If 'false' is returned, it signals that the
//       composite date is invalid under the Gregorian Calendar.
//
//
//  err                    error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//       In addition, an error will be returned if the composite date
//       represented by input parameters, 'year', 'month' and 'day' are
//       invalid. In this case the returned error type will contain a
//       detailed message identifying the invalid date element.
//
func (gregCalBData *CalendarGregorianBaseData) IsValidDate(
	year int64,
	yearNumType CalendarYearNumType,
	month int,
	day int,
	ePrefix string) (
	isValid bool,
	err error) {

	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	ePrefix += "CalendarGregorianBaseData.IsLeapYear() "

	gregCalBDataUtil :=
		calendarGregorianBaseDataUtility{}

	isValid,
	err = gregCalBDataUtil.isValidDate(
		year,
		yearNumType,
		month,
		day,
		ePrefix)

	return isValid, err
}

// New - Returns a new instance of CalendarGregorianBaseData
//
func (gregCalBData *CalendarGregorianBaseData) New() ICalendarBaseData {
	if gregCalBData.lock == nil {
		gregCalBData.lock = new(sync.Mutex)
	}

	gregCalBData.lock.Lock()

	defer gregCalBData.lock.Unlock()

	newBaseData :=  &CalendarGregorianBaseData{}

	return newBaseData
}



