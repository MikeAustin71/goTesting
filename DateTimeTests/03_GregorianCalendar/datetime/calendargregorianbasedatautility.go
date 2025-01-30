package datetime

import (
	"fmt"
	"math"
	"sync"
)

type calendarGregorianBaseDataUtility struct {
	lock    *sync.Mutex
}

// getISODayOfWeekNumber - Returns the ISO Day Of The Week Number. This
// method receives a Julian Day Number and proceeds to calculate the day
// of the week number associated with that Julian Day Number Date. It
// then returns that day of the week number to the calling function
// formatted with the ISO Standard Day Of The Week Numbering System.
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
func (gregCalBDataUtil *calendarGregorianBaseDataUtility) getISODayOfWeekNumber(
	julianDayNoDto JulianDayNoDto,
	ePrefix string) (
	isoDayOfWeekNo ISO8601DayOfWeekNo,
	err error) {

	if gregCalBDataUtil.lock == nil {
		gregCalBDataUtil.lock = new(sync.Mutex)
	}

	gregCalBDataUtil.lock.Lock()

	defer gregCalBDataUtil.lock.Unlock()

	ePrefix +=
		"calendarGregorianBaseDataUtility.getUsDayOfWeekNumber() "

	isoDayOfWeekNo =ISO8601DayOfWeekNo(0).None()

	err = julianDayNoDto.IsValidInstanceError(ePrefix + "Testing validity of 'julianDayNoDto' ")

	if err != nil {
		return isoDayOfWeekNo, err
	}

	calMech := calendarMechanics{}

	var usDayOfWeekNo UsDayOfWeekNo

	usDayOfWeekNo,
	err = calMech.usDayOfWeekNumber(
		julianDayNoDto,
		ePrefix)

	if err != nil {
		return isoDayOfWeekNo, err
	}

	isoDayOfWeekNo, err =
		usDayOfWeekNo.XISO8601DayOfWeekNumber(ePrefix)

	return isoDayOfWeekNo, err
}

// getUsDayOfWeekNumber - Returns the US Day Of The Week Number. This method
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
//       an enumeration parameter which can be interrogated using methods
//       on the IDayOfWeekNumber interface. This interface will support one
//       of two types: UsDayOfWeekNo or ISO8601DayOfWeekNo. These enumeration
//       values will allow the user to access the correct day number, day name
//       and day of the week number calculated from input parameter
//       'julianDayNoDto'.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note this error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'.
//
func (gregCalBDataUtil *calendarGregorianBaseDataUtility) getUsDayOfWeekNumber(
	julianDayNoDto JulianDayNoDto,
	ePrefix string) (
	usDayOfWeekNo UsDayOfWeekNo,
	err error) {

	if gregCalBDataUtil.lock == nil {
		gregCalBDataUtil.lock = new(sync.Mutex)
	}

	gregCalBDataUtil.lock.Lock()

	defer gregCalBDataUtil.lock.Unlock()

	ePrefix +=
		"calendarGregorianBaseDataUtility.getUsDayOfWeekNumber() "

	usDayOfWeekNo = UsDayOfWeekNo(0).None()

	err = julianDayNoDto.IsValidInstanceError(ePrefix + "Testing validity of 'julianDayNoDto' ")

	if err != nil {
		return usDayOfWeekNo, err
	}

	calMech := calendarMechanics{}

	usDayOfWeekNo,
	err = calMech.usDayOfWeekNumber(
		julianDayNoDto,
		ePrefix)

	return usDayOfWeekNo, err
}

// getMonthDayFromOrdinalDayNo - Receives an Ordinal Day Number and returns
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
func (gregCalBDataUtil *calendarGregorianBaseDataUtility) getMonthDayFromOrdinalDayNo(
	ordinalDate int,
	isLeapYear bool,
	ePrefix string)(
	yearAdjustment int,
	month int,
	day int,
	err error) {

	if gregCalBDataUtil.lock == nil {
		gregCalBDataUtil.lock = new(sync.Mutex)
	}

	gregCalBDataUtil.lock.Lock()

	defer gregCalBDataUtil.lock.Unlock()

	ePrefix +=
		"calendarGregorianBaseDataUtility.getMonthDayFromOrdinalDayNo() "

	yearAdjustment = math.MinInt32
	month = -1
	day = -1
	err = nil
	var ordDays, mthDays map[int]int
	var daysInYear int

	gregCalBDataMech :=
		calendarGregorianBaseDataMechanics{}

	if isLeapYear {

		daysInYear = gregCalBDataMech.getLeapYearDaysInYear()

		if ordinalDate < 0 ||
			ordinalDate > daysInYear {
			err = fmt.Errorf("\n" +ePrefix + "Error:\n" +
				"Input Parameter 'ordinalDate' is INVALID!\n" +
				"ordinalDate='%v'\n" +
				"daysInLeapYear='%v'\n",
				ordinalDate,
				daysInYear)
			return yearAdjustment, month, day, err
		}

		ordDays =
			gregCalBDataMech.getLeapYearOrdinalDays()

		mthDays =
			gregCalBDataMech.getLeapYearMonthDays()

	} else {
		// Must be a Standard Year.
		// This is NOT a leap year

		daysInYear =
			gregCalBDataMech.getStandardYearDaysInYear()

		if ordinalDate < 0 ||
			ordinalDate > daysInYear {
			err = fmt.Errorf("\n" +ePrefix + "Error:\n" +
				"Input Parameter 'ordinalDate' is INVALID!\n" +
				"ordinalDate='%v'\n" +
				"Days in Non-LeapYear='%v'\n",
				ordinalDate,
				daysInYear)
			return yearAdjustment, month, day, err
		}

		ordDays =
			gregCalBDataMech.getStandardYearOrdinalDays()

		mthDays =
			gregCalBDataMech.getStandardYearMonthDays()

	}

	if ordinalDate == 0 {
		yearAdjustment = -1
		month = len(mthDays)
		day = mthDays[month]
		return yearAdjustment, month, day, err
	}

	yearAdjustment = 0

	for i:=11; i > -1; i-- {

		if ordDays[i] < ordinalDate {

			day = ordinalDate - ordDays[i]

			month = i + 1

			if month > len(mthDays) - 1 {
				err = fmt.Errorf(ePrefix + "\n" +
					"Error: Calculated Month Number exceeds maximum months!\n" +
					"   Maximum Month Number: %v\n" +
					"Calculated Month Number: %v\n",
					len(mthDays)-1,
					month)
				return yearAdjustment, month, day, err
			}


			if day < 1 ||
				day > mthDays[month] {
				err = fmt.Errorf(ePrefix + "\n" +
					"Invalid Ordinal Day Number Result!\n" +
					"month='%v'  Ordinal Day No='%v'\n" +
					"Original Ordinal Date='%v'\n",
					month, day, ordinalDate)
				return yearAdjustment, month, day, err
			}

			break
		}
	}

	return yearAdjustment, month, day, err
}


// getOrdinalDayNumber - Computes the ordinal day number within a
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
func (gregCalBDataUtil *calendarGregorianBaseDataUtility) getOrdinalDayNumber(
	isLeapYear bool,
	month int,
	day int,
	ePrefix string) (
	ordinalDayNo int,
	err error) {

	if gregCalBDataUtil.lock == nil {
		gregCalBDataUtil.lock = new(sync.Mutex)
	}

	gregCalBDataUtil.lock.Lock()

	defer gregCalBDataUtil.lock.Unlock()

	ePrefix += "calendarGregorianBaseDataUtility.getOrdinalDayNumber() "

	ordinalDayNo = -1
	err = nil

	var monthDays, ordinalDays map[int] int

	gregCalBDataMech := calendarGregorianBaseDataMechanics{}


	if isLeapYear {
		monthDays = gregCalBDataMech.getLeapYearMonthDays()
		ordinalDays = gregCalBDataMech.getLeapYearOrdinalDays()
	} else {
		monthDays = gregCalBDataMech.getStandardYearMonthDays()
		ordinalDays = gregCalBDataMech.getStandardYearOrdinalDays()
	}

	maxMonthNo := len(monthDays)

	if month < 1 || month > maxMonthNo {
		err = fmt.Errorf("\n" + ePrefix + "Error:\n" +
			"Input Parameter 'month' is INVALID!\n" +
			"The valid range for 'month' number is 1 through %v inclusive.\n" +
			"month='%v'\n",
			maxMonthNo,
			month)

		return ordinalDayNo, err
	}


	if month == 1 &&
		day == 0 {
		ordinalDayNo = 0
		return ordinalDayNo, err
	}

	daysInMonth := monthDays[month]

	if day < 1 || day > daysInMonth {
		err = fmt.Errorf("\n" + ePrefix + "Error:\n" +
			"Input parameter 'day' is INVALID!\n" +
			"The month number is '%v'\n" +
			"The valid range for 'day' is 1 through %v inclusive.\n" +
			"day='%v'\n",
			month,
			daysInMonth,
			day)

		return ordinalDayNo, err
	}


	ordinalDayNo = ordinalDays[month] + day

	return ordinalDayNo, err
}

// GetYearMonthDayFromOrdinalDayNo - Receives an Ordinal Day Number and returns
// the associated year, month and day number. Note that tye returned year value
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
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
func (gregCalBDataUtil *calendarGregorianBaseDataUtility) getYearMonthDayFromOrdinalDayNo(
	ordinalDate int,
	year int64,
	yearNumType CalendarYearNumType,
	ePrefix string)(
	astronomicalYear int64,
	month int,
	day int,
	err error) {

	if gregCalBDataUtil.lock == nil {
		gregCalBDataUtil.lock = new(sync.Mutex)
	}

	gregCalBDataUtil.lock.Lock()

	defer gregCalBDataUtil.lock.Unlock()

	ePrefix +=
		"calendarGregorianBaseDataUtility.getMonthDayFromOrdinalDayNo() "

	astronomicalYear = math.MinInt64
	month = -1
	day = -1
	err = nil

	calMech := calendarMechanics{}

	astronomicalYear,
	err = calMech.convertAnyYearToAstronomicalYear(
		year,
		yearNumType,
		ePrefix)

	if err != nil {
		return astronomicalYear, month, day, err
	}

	gregCalBDataMech :=
		calendarGregorianBaseDataMechanics{}

	var isLeapYear bool

	isLeapYear,
	err = gregCalBDataMech.isLeapYear(
		astronomicalYear,
		CalendarYearNumType(0).Astronomical(),
		ePrefix)

	if err != nil {
		return astronomicalYear, month, day, err
	}

	var ordinalDays, monthDays map[int]int
	var daysInYear int

	if isLeapYear {

		daysInYear = gregCalBDataMech.getLeapYearDaysInYear()

		if ordinalDate < 0 ||
			ordinalDate > daysInYear {
			err = fmt.Errorf("\n" +ePrefix + "Error:\n" +
				"Input Parameter 'ordinalDate' is INVALID!\n" +
				"ordinalDate='%v'\n" +
				"Days In Leap Year='%v'\n",
				ordinalDate,
				daysInYear)

			return astronomicalYear, month, day, err
		}

		ordinalDays =
			gregCalBDataMech.getLeapYearOrdinalDays()

		monthDays =
			gregCalBDataMech.getLeapYearMonthDays()

	} else {
		// Must be a Standard Year.
		// This is NOT a Leap Year

		daysInYear =
			gregCalBDataMech.getStandardYearDaysInYear()

		if ordinalDate < 0 ||
			ordinalDate > daysInYear {
			err = fmt.Errorf(ePrefix + "\n" +
				"Input Parameter 'ordinalDate' is INVALID!\n" +
				"ordinalDate='%v'\n" +
				"Days in Standard Year or Non-LeapYear='%v'\n",
				ordinalDate,
				daysInYear)
			return astronomicalYear, month, day, err
		}

		ordinalDays =
			gregCalBDataMech.getStandardYearOrdinalDays()

		monthDays =
			gregCalBDataMech.getStandardYearMonthDays()

	}

	// monthsInYear should be '12'
	monthsInYear := len(monthDays)
	var testOrdDays int
	var ok bool

	for i:= monthsInYear; i > 0 ; i-- {

		testOrdDays, ok = ordinalDays[i]

		if !ok {
			err = fmt.Errorf(ePrefix + "Error:\n" +
				"Map Ordinal Days Look-Up Failed!\n" +
				"ordinalDays[i] Failed to return a value!\n" +
				"i (a.k.a month number) ='%v'\n",
				i)

			return astronomicalYear, month, day, err
		}

		if ordinalDate >= testOrdDays {
			day = ordinalDate - testOrdDays
			month = i
			return astronomicalYear, month, day, err
		}
	}

	err = fmt.Errorf(ePrefix + "Error: \n" +
		"After searching the 'ordinalDays' map,\n" +
		"no 'month' or 'day' value was returned!\n")

	return astronomicalYear, month, day, err
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
func (gregCalBDataUtil *calendarGregorianBaseDataUtility) isValidDate(
	year int64,
	yearNumType CalendarYearNumType,
	month int,
	day int,
	ePrefix string) (
	isValid bool,
	err error) {

	if gregCalBDataUtil.lock == nil {
		gregCalBDataUtil.lock = new(sync.Mutex)
	}

	gregCalBDataUtil.lock.Lock()

	defer gregCalBDataUtil.lock.Unlock()

	isValid = false
	err = nil

	ePrefix += "calendarGregorianBaseDataUtility.isValidDate() "

	if ! yearNumType.XIsValid() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Input parameter 'yearNumType' is INVALID!\n" +
			"yearNumType='%v'\n",
			yearNumType.XValueInt())

		return isValid, err
	}

	var astronomicalYear int64

	calMech := calendarMechanics{}

	astronomicalYear,
	err = calMech.convertAnyYearToAstronomicalYear(
		year,
		yearNumType,
		ePrefix)

	if err != nil {
		return isValid, err
	}

	gregCalBDataMech :=
		calendarGregorianBaseDataMechanics{}

	var isLeapYear bool

	isLeapYear,
	err = gregCalBDataMech.isLeapYear(
		astronomicalYear,
		CalendarYearNumType(0).Astronomical(),
		ePrefix)

	if err != nil {
		return isValid, err
	}

	var mthDays map [int] int

	if isLeapYear {
		mthDays =
			gregCalBDataMech.getLeapYearMonthDays()
	} else {
		mthDays =
			gregCalBDataMech.getStandardYearMonthDays()
	}

	maxMonthValue := len(mthDays)

	if month < 1 ||
			month > maxMonthValue {
		err = fmt.Errorf(ePrefix + "\n" +
			"The Date INVALID!\n" +
			"'month' parameter is outside the valid range.\n" +
			"The minimum value for 'month' is '1'.\n" +
			"The maximum value for 'month' is '%v'.\n" +
			"The actual value of month is '%v'.\n",
			maxMonthValue,
			month)
		return isValid, err
	}

	maxDaysInMonth, ok := mthDays[month]

	if !ok {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input prameter 'month' is invalid.\n" +
			"The month number could NOT be located in the month-days map.\n" +
			"month='%v'\n", month)
		return isValid, err
	}

	if day < 1 ||
		day > maxDaysInMonth {
		err = fmt.Errorf(ePrefix + "\n" +
			"The Date INVALID!\n" +
			"'day' parameter is outside the valid range.\n" +
			"The minimum value for 'day' is '1'.\n" +
			"The maximum value for 'day' is '%v'.\n" +
			"The actual value of 'day' is '%v'.\n",
			maxDaysInMonth,
			day)
		return isValid, err
	}

	isValid = true
	err = nil

	return isValid, err
}