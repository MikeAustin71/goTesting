package GCal_Libs01

import (
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"
)

// CalendarGregorianUtility - This type contains methods
// used to process date arithmetic associated with the
// Gregorian Calendar
//
// Reference:
//  https://en.wikipedia.org/wiki/Gregorian_calendar
//
type CalendarGregorianUtility struct {

	lock *sync.Mutex
}


func (calGreg *CalendarGregorianUtility) GetJDN(
	targetYear int64,
	targetMonth int,
	targetDay int,
	targetHour int,
	targetMinute int,
	targetSecond int,
	targetNanosecond int,
	ePrefix string) (
	julianDayNumber int64,
	julianDayNumberTime *big.Float,
	julianDayNumberTimeFraction *big.Float,
	err error) {

	if calGreg.lock == nil {
		calGreg.lock = &sync.Mutex{}
	}

	calGreg.lock.Lock()

	defer calGreg.lock.Unlock()

	ePrefix += "CalendarGregorianUtility.GetJDN() "

	julianDayNumber = 0

	julianDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	julianDayNumberTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	err = nil

	calGregMech := calendarGregorianMechanics{}

	isLeapYear := calGregMech.isLeapYear(targetYear)

	calUtil := CalendarUtility{}

	err = calUtil.IsValidDateTimeComponents(
		isLeapYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {

		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	var baseDateTimeDto, targetDateTimeDto DateTimeTransferDto

	targetDateTimeDto, err = DateTimeTransferDto{}.New(
		isLeapYear,
		targetYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	calCycleCfg := calGregMech.getCalendarCyclesConfig()

	baseDateTimeDto = calCycleCfg.GetJDNBaseStartYearDateTime()

	// Base Date Time for Gregorian Calendar:
	// November 24, -4713 12:00:00.000000000 UTC (Noon)

	var baseTargetComparisonResult, baseRemainingDaysInYear,
	targetRemainingDaysInYear, targetYearOrdinalNumber,
	julianDayNumberSign int

	var mainCycleAdjustmentDays, remainderYears, ordinalDays *big.Int

	baseTargetComparisonResult, err = baseDateTimeDto.Compare(&targetDateTimeDto, ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	baseDateTimeDto.SetTag("baseDateTimeDto")
	targetDateTimeDto.SetTag("targetDateTimeDto")

	targetYearOrdinalNumber, err = targetDateTimeDto.GetOrdinalDayNumber(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	targetRemainingDaysInYear, err = targetDateTimeDto.GetRemainingDaysInYear(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	fmt.Printf("Target Days Remaining in Year: %v\n",
		targetRemainingDaysInYear)

	baseRemainingDaysInYear, err = baseDateTimeDto.GetRemainingDaysInYear(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	fmt.Printf("Base Days Remaining in Year: %v\n",
		baseRemainingDaysInYear)

	var mainCycleStartDateTime DateTimeTransferDto

	// Primary Decision Tree
	if baseTargetComparisonResult == 0 {
		// base and target have equivalent date/times
		// Julian Day Number is Zero and time fraction is Zero.

		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err

	} else if baseTargetComparisonResult == 1 {
		// base is greater than target date/time
		// target date/time must be negative
		// Julian Day Number must be negative

		mainCycleStartDateTime =
			calCycleCfg.GetMainCycleStartDateForNegativeJDNNo()

		mainCycleStartDateTime.SetTag("Main Cycle Start Date Time")

		fmt.Printf("%s\n",
			mainCycleStartDateTime.String())

		julianDayNumberSign = -1

		// RemainderYears =
		// BaseYear - TargetYear -1
		remainderYears =
			big.NewInt(
				mainCycleStartDateTime.GetYear() -
					targetDateTimeDto.GetYear() - 1 )

		mainCycleAdjustmentDays =
			calCycleCfg.GetMainCycleAdjustmentDaysForNegativeJDNNo()

	} else {
		// target is greater than base date/time
		// base is less than target date/time
		// target date/time could be positive or negative
		// Use Main Cycle Start Date/Time which is less Than target

		mainCycleStartDateTime =
			calCycleCfg.GetMainCycleStartDateForPositiveJDNNo()

		mainCycleStartDateTime.SetTag("Main Cycle Start Date Time")

		fmt.Printf("%s\n",
			mainCycleStartDateTime.String())

		julianDayNumberSign = 1

		// RemainderYears or Whole Years Interval =
		// TargetYear - mainCycleStartDateTime -1
		remainderYears =
			big.NewInt(
			targetDateTimeDto.GetYear() -
				mainCycleStartDateTime.GetYear() - 1)

	mainCycleAdjustmentDays =
		calCycleCfg.GetMainCycleAdjustmentDaysForPositiveJDNNo()

}

	fmt.Printf("Initial Remainder Years: %v\n",
		remainderYears.Text(10))

	ordinalDays =
		big.NewInt(
			int64(
				baseRemainingDaysInYear +
					targetYearOrdinalNumber))

	cycles := calCycleCfg.GetCalendarCycleConfigurations()
	daysDuration := big.NewInt(0)

	if remainderYears.Cmp(big.NewInt(0)) == 1 {

		var cycleCount *big.Int
		totalYears := big.NewInt(0)

		for i:=0; i < len(cycles); i++ {

			cycleCount =
				big.NewInt(0).
					Quo(remainderYears,
						cycles[i].GetYearsInCycle())

			remainderYears =
				big.NewInt(0).
					Sub(remainderYears,
						big.NewInt(0).
						Mul(cycleCount, cycles[i].GetYearsInCycle()))

			fmt.Printf("cycles[%v].cycleCount=%v\n",
				i, cycleCount.Text(10))

			fmt.Printf("cycles[%v].yearsInCycle=%v\n",
				i, cycles[i].GetYearsInCycle().Text(10))

			fmt.Printf("cycles[%v].remainderYears=%v\n",
				i, remainderYears.Text(10))


			cycles[i].SetCycleCount(cycleCount)
			cycles[i].SetRemainderYears(remainderYears)

			totalYears =
				big.NewInt(0).
					Add(totalYears, cycles[i].GetCycleCountTotalYears())

			fmt.Printf("cycles[%v].CycleCountTotalDays=%v\n",
				i, cycles[i].GetCycleCountTotalDays().Text(10))

			daysDuration =
				big.NewInt(0).
					Add(daysDuration, cycles[i].GetCycleCountTotalDays())
		}

		fmt.Printf("daysDuration: %v\n",
			daysDuration.Text(10))

		daysDuration =
			big.NewInt(0).
				Add(daysDuration, mainCycleAdjustmentDays)

		fmt.Printf("Adjusted daysDuration: %v\n",
			daysDuration.Text(10))

	}

	fmt.Printf("ordinalDays: %v\n",
		ordinalDays.Text(10))

	daysDuration =
		big.NewInt(0).
			Add(daysDuration, ordinalDays)

	fmt.Printf("Final daysDuration: %v\n",
		daysDuration.Text(10))

	if !daysDuration.IsInt64() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: 'daysDuration' is too large and cannot convert to int64!\n" +
			"daysDuration='%v'\n",
			daysDuration.Text(10))
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	julianDayNumber = daysDuration.Int64()

	// Compute Time Components

	noonNanoseconds := int64(time.Hour * 12)

	var targetDayTotalNanoseconds int64

	targetDayTotalNanoseconds, err =
		targetDateTimeDto.GetTotalTimeInNanoseconds(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}


	twentyFourHourNanosecondsFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(int64(time.Hour) * 24)

	julianDayNoTimeAdjustment := int64(0)

	if targetDayTotalNanoseconds == noonNanoseconds {

		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetFloat64(0.0)

	} else if targetDayTotalNanoseconds < noonNanoseconds {
		targetDayTotalNanoseconds += noonNanoseconds

		julianDayNoTimeAdjustment--

		targetNanosecondFloat :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(targetDayTotalNanoseconds)


		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Quo(targetNanosecondFloat, twentyFourHourNanosecondsFloat)


	} else {
		// targetDayTotalNanoseconds > noonNanoseconds


		targetDayTotalNanoseconds -= noonNanoseconds

		targetNanosecondFloat :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(targetDayTotalNanoseconds)


		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Quo(targetNanosecondFloat, twentyFourHourNanosecondsFloat)
	}

	julianDayNumber += julianDayNoTimeAdjustment

	fmt.Printf("julianDayNoTimeAdjustment: %v\n",
		julianDayNoTimeAdjustment)

	julianDayNoFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(julianDayNumber)

	julianDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Add(julianDayNoFloat, julianDayNumberTimeFraction)


	if julianDayNumberSign == -1 {

		julianDayNumber *= -1

		julianDayNumberTime =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Neg(julianDayNumberTime)
	}

	separator := strings.Repeat("-", 75)
	fmt.Println()
	fmt.Println(separator)

	fmt.Printf("Final Julian Day Number: %v\n",
		julianDayNumber)

	fmt.Println(separator)
	fmt.Println(baseDateTimeDto.String())
	fmt.Println(targetDateTimeDto.String())
	fmt.Println(separator)

	return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
}

// GetJulianDayNumber - Returns values defining the Julian Day Number
// and Time for a date/time in the Revised Julian Calendar.
//
// All time input parameters are assumed to be expressed in Coordinated
// Universal Time (UTC). For more information on Coordinated Universal
// Time, reference:
//   https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
// Julian Day Number is used to define a standard time duration, and
// perform date/time conversions, between differing calendar systems.
//
// The base date/time for Julian Day Number zero on the Gregorian
// Calendar is November 24, -4713 12:00:00.000000000 UTC (Noon) or
// November 24, 4714 BCE 12:00:00.000000000 UTC (Noon).
//
// For more information on the Julian Day Number, reference:
//  https://en.wikipedia.org/wiki/Julian_day
//
// For more information on the Gregorian Calendar, reference:
//  https://en.wikipedia.org/wiki/Gregorian_calendar
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  targetYear         int64
//     - The year number associated with this date/time specification.
//       The year value may be positive or negative. The year value must
//       conform to the astronomical year numbering system. This means
//       that year zero is valid and recognized. Example: 1/1/0000. The
//       astronomical year value -4712 is therefore equivalent to
//       -4713 BCE. All year values submitted to this method must use
//       the astronomical year numbering system. For more information
//       on the astronomical year numbering system, reference:
//              https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//  targetMonth        int
//     - The month number for this date/time specification.
//       The valid range is 1 - 12 inclusive.
//
//  targetDay          int
//     - The day number for this date/time specification. The day
//       number must fall within the limits of the month number
//       submitted above in input parameter, 'targetMonth'.
//
//  targetHour         int
//     - The hour time component for this date/time specification.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00. All time
//       parameters are assumed to be expressed in Universal
//       Coordinated Time (UTC).
//
//  targetMinute       int
//     - The minute time component for this date/time specification.
//       The valid range is 0 - 59 inclusive.  All time
//       parameters are assumed to be expressed in Universal
//       Coordinated Time (UTC).
//
//  targetSecond       int
//     - The second time component for this date/time specification.
//       The valid range is 0 - 60 inclusive. The value 60 is only
//       used in the case of leap seconds.  All time parameters are
//       assumed to be expressed in Universal Coordinated Time (UTC).
//
//  targetNanosecond   int
//     - The nanosecond time component for this date/time specification.
//       The valid range is 0 - 999,999,999 inclusive.  All time
//       parameters are assumed to be expressed in Universal
//       Coordinated Time (UTC).
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  julianDayNumber             int64
//     - The integer julian day number of the date specified by input
//       parameters 'targetYear', 'targetMonth' and 'targetDay'. This
//       value equals the number of days elapsed between the base date,
//       February 8, 4714 BCE 12:00:00 UTC (Noon), and the target date/time
//       specified by the input parameters. Both base and target date/times
//       represent moments on the Revised Julian Calendar.
//
//
//  julianDayNumberTime         *big.Float
//     - The combined Julian Day number and fractional time value represented
//       by the date/time specified by input parameters, targetYear,
//       targetMonth, targetDay, targetHour, targetMinute, targetSecond
//       and targetNanosecond. This value equals the duration in day numbers
//       calculated from the start date of February 8, 4714 BCE 12:00:00 UTC
//       (Noon) on the Revised Julian Calendar.
//
//
//  julianDayNumberTimeFraction *big.Float
//     - The fractional time value associated with the the ordinal
//       day number. This value does NOT contain the integer ordinal
//       day number. Instead, it only contains the time value
//       represented by the date/time input parameters, targetHour,
//       targetMinute, targetSecond and targetNanosecond. This value
//       equals the number of nanoseconds since midnight of the ordinal
//       day number divided by the number of nanoseconds in a
//       24-hour day.
//
//
//  err                         error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calGreg *CalendarGregorianUtility) GetJulianDayNumber(
	targetYear int64,
	targetMonth int,
	targetDay int,
	targetHour int,
	targetMinute int,
	targetSecond int,
	targetNanosecond int,
	ePrefix string) (
	julianDayNumber int64,
	julianDayNumberTime *big.Float,
	julianDayNumberTimeFraction *big.Float,
	err error) {

	if calGreg.lock == nil {
		calGreg.lock = &sync.Mutex{}
	}

	calGreg.lock.Lock()

	defer calGreg.lock.Unlock()

	ePrefix += "CalendarGregorianUtility.GetJulianDayNumber() "

	julianDayNumber = 0

	julianDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	julianDayNumberTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	err = nil

	calGreg2 := CalendarGregorianUtility{}

	isLeapYear := calGreg2.IsLeapYear(targetYear)

	calUtil := CalendarUtility{}

	err = calUtil.IsValidDateTimeComponents(
		isLeapYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {

		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	var baseDateTimeDto, targetDateTimeDto DateTimeTransferDto

	targetDateTimeDto, err = DateTimeTransferDto{}.New(
		isLeapYear,
		targetYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	// Base Date Time for Gregorian Calendar:
	// November 24, -4713 12:00:00.000000000 UTC (Noon)
	//
	// Reference: https://en.wikipedia.org/wiki/Julian_day
	baseDateTimeDto, err = DateTimeTransferDto{}.New(
		false,
		int64(-4713),
		11,
		24,
		12,
		0,
		0,
		0,
		ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	var baseTargetComparisonResult, baseTargetYearsComparison int

	baseTargetComparisonResult, err = baseDateTimeDto.Compare(&targetDateTimeDto, ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	baseTargetYearsComparison = baseDateTimeDto.CompareYears(&targetDateTimeDto)

	// baseDateTimeDto is now less than targetDateTimeDto

	baseDateTimeDto.SetTag("baseDateTimeDto")
	targetDateTimeDto.SetTag("targetDateTimeDto")

	var wholeYearsInterval, lastYearPriorToTargetYear int64

	var targetYearOrdinalNumber, baseYearOrdinalNumber int

	var lastYearPriorToTargetYearDeltaDays, targetPartialYearDeltaDays int64

	var wholeYearDays int64

	targetYearOrdinalNumber, err = targetDateTimeDto.GetOrdinalDayNumber(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	targetPartialYearDeltaDays = int64(targetYearOrdinalNumber)

	baseYearOrdinalNumber, err = baseDateTimeDto.GetOrdinalDayNumber(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	baseYearVariance :=
		baseDateTimeDto.GetYear() -
		targetDateTimeDto.GetYear()

	baseAndTargetAreAdjacentYears := false

	if baseYearVariance == -1 || baseYearVariance == 1 {
		baseAndTargetAreAdjacentYears = true
	}

	// ** Main Decision Tree **
	if baseTargetComparisonResult == 0 {
		// base and target have equivalent date/times
		// Julian Day Number is Zero and time fraction is Zero.

		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err

	} else if baseTargetYearsComparison == 0 {
		// base and target years are equal
		wholeYearsInterval = 0
		wholeYearDays = 0
		lastYearPriorToTargetYearDeltaDays = 0

		// base year is greater than target year
		targetPartialYearDeltaDays =
			int64(baseYearOrdinalNumber - targetYearOrdinalNumber)

		if targetPartialYearDeltaDays < 0 {
			targetPartialYearDeltaDays *= -1
		}
		// End of if baseTargetYearsComparison == 0
	} else if baseAndTargetAreAdjacentYears {
		// base and target years are adjacent years

		wholeYearsInterval = 0
		wholeYearDays = 0

		if baseTargetComparisonResult == 1 {
			// base year date/time is greater than target year date/time
			// base year and target year are negative

			lastYearPriorToTargetYearDeltaDays =
				baseDateTimeDto.GetYearDays() - int64(baseYearOrdinalNumber)

			targetPartialYearDeltaDays = int64(targetYearOrdinalNumber)

		} else {
			// base year is less than target year
			// base year is negative. target year
			// must be negative because base year
			// and target year are adjacent years

			lastYearPriorToTargetYearDeltaDays =
				int64(baseYearOrdinalNumber)

			targetPartialYearDeltaDays =
				targetDateTimeDto.GetYearDays() - int64(targetYearOrdinalNumber)

		}


	} else if baseTargetComparisonResult == 1 {
		// base and target years are NOT equal
		// base and target years are NOT adjacent years
		// base is greater than target date/time
		// target date/time must be negative

		wholeYearsInterval =
			targetDateTimeDto.GetYear() -
				baseDateTimeDto.GetYear() + 1

		if wholeYearsInterval < 0 {
			wholeYearsInterval *= -1
		}

		wholeYearDays =
			calGreg2.NumCalendarDaysForWholeYearsInterval(wholeYearsInterval)


		// target year is negative
		lastYearPriorToTargetYear = targetDateTimeDto.GetYear() + 1

		if calGreg2.IsLeapYear(lastYearPriorToTargetYear) {
			lastYearPriorToTargetYearDeltaDays = int64(366) - int64(baseYearOrdinalNumber)
		} else {
			lastYearPriorToTargetYearDeltaDays = int64(365) - int64(baseYearOrdinalNumber)
		}

		targetPartialYearDeltaDays = int64(targetYearOrdinalNumber)

	} else {
		// base and target year are NOT equal!
		// base and target years are NOT adjacent years
		// target is greater than base date/time
		// target date/time could be positive or
		// negative
		wholeYearsInterval =
			targetDateTimeDto.GetYear() -
				baseDateTimeDto.GetYear() - 1

		if wholeYearsInterval < 0 {
			wholeYearsInterval *= -1
		}

		wholeYearDays =
			calGreg2.NumCalendarDaysForWholeYearsInterval(wholeYearsInterval)

		if targetDateTimeDto.GetYearNumberSign() < 0 {
			// target year is negative
			lastYearPriorToTargetYear = targetDateTimeDto.GetYear() + 1

		} else {
			// target year is positive
			lastYearPriorToTargetYear = targetDateTimeDto.GetYear() - 1
		}

		if calGreg2.IsLeapYear(lastYearPriorToTargetYear) {
			lastYearPriorToTargetYearDeltaDays = int64(366) - int64(baseYearOrdinalNumber)
		} else {
			lastYearPriorToTargetYearDeltaDays = int64(365) - int64(baseYearOrdinalNumber)
		}

		targetPartialYearDeltaDays = int64(targetYearOrdinalNumber)

	}

	julianDayNumber = wholeYearDays +
		lastYearPriorToTargetYearDeltaDays +
		targetPartialYearDeltaDays

	noonNanoseconds := int64(time.Hour * 12)

	var targetDayTotalNanoseconds int64

	targetDayTotalNanoseconds, err =
		targetDateTimeDto.GetTotalTimeInNanoseconds(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	julianDayAdjustment := 0

	twentyFourHourNanosecondsFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(int64(time.Hour) * 24)


	if targetDayTotalNanoseconds == noonNanoseconds {
		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetFloat64(0.0)

	} else if targetDayTotalNanoseconds < noonNanoseconds {
		julianDayNumber--
		targetDayTotalNanoseconds += noonNanoseconds
		julianDayAdjustment--

		targetNanosecondFloat :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(targetDayTotalNanoseconds)


		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Quo(targetNanosecondFloat, twentyFourHourNanosecondsFloat)


	} else {
		// targetDayTotalNanoseconds > noonNanoseconds

		targetDayTotalNanoseconds -= noonNanoseconds

		targetNanosecondFloat :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(targetDayTotalNanoseconds)


		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Quo(targetNanosecondFloat, twentyFourHourNanosecondsFloat)

	}

	julianDayNoFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(julianDayNumber)

	julianDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Add(julianDayNoFloat, julianDayNumberTimeFraction)

	separator := strings.Repeat("-", 65)
	fmt.Println()
	fmt.Println(separator)

	fmt.Printf("julianDayAdjustment: %v\n",
		julianDayAdjustment)

	fmt.Println(separator)
	fmt.Println(baseDateTimeDto.String())
	fmt.Println(targetDateTimeDto.String())
	fmt.Println(separator)

	return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
}

// GetOrdinalDayNumberTime - Computes and returns the 'Fixed Day'
// or ordinal day number time under the Revised Julian Calendar.
//
// The ordinal day number time identifies the number of days and
// and fraction times since a fixed start date. The start date
// for the Revised Julian Calendar is Monday, 1 January 1 AD 00:00:00
// (midnight). This is defined as the beginning of the first ordinal
// day. All ordinal day numbers use this start date as the base base
// reference point.
//
// This method receives a series of input parameters specifying
// a target date and time. The times are always assumed to be
// Universal Coordinated Times (UTC). The method then returns
// the ordinal day number as integer (int64) and a type *big.Float
// which defines both the ordinal day number and the time expressed
// as a fraction of a 24-hour day.
//
// Note that the input parameter 'targetYear' is a type int64 which
// must be configured under the astronomical year numbering system.
// This system recognizes the year zero as a legitimate year value.
// Under the astronomical year numbering system the year 4713 BCE is
// formatted as -4712.
//
// Reference:
//  https://en.wikipedia.org/wiki/Revised_Julian_calendar
//  https://en.wikipedia.org/wiki/Rata_Die
//  https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  targetYear         int64
//     - The year number associated with this date/time specification.
//       The year value may be positive or negative. The year value must
//       conform to the astronomical year numbering system. This means
//       that year zero is valid and recognized. Example: 1/1/0000. The
//       astronomical year value -4712 is therefore equivalent to
//       -4713 BCE. All year values submitted to this method must use
//       the astronomical year numbering system. For more information
//       on the astronomical year numbering system, reference:
//              https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//  targetMonth        int
//     - The month number for this date/time specification.
//       The valid range is 1 - 12 inclusive.
//
//  targetDay          int
//     - The day number for this date/time specification. The day
//       number must fall within the limits of the month number
//       submitted above in input parameter, 'targetMonth'.
//
//  targetHour         int
//     - The hour time component for this date/time specification.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00. 'targetHour'
//       is assumed to represent Coordinated Universal Time (UTC).
//
//  targetMinute       int
//     - The minute time component for this date/time specification.
//       The valid range is 0 - 59 inclusive. 'targetMinute' is assumed
//       to represent  Coordinated Universal Time (UTC).
//
//  targetSecond       int
//     - The second time component for this date/time specification.
//       The valid range is 0 - 60 inclusive. The value 60 is only
//       used in the case of leap seconds.
//
//  targetNanosecond   int
//     - The nanosecond time component for this date/time specification.
//       The valid range is 0 - 999,999,999 inclusive
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  ordinalDayNumber              int64
//     - The ordinal day number of the date represented by the date
//       specified by input parameters targetYear, targetMonth and
//       targetDay. This value equals the number of days elapsed,
//       plus 1-day, since January 1, 1 CE on the Revised Julian
//       Calendar.
//
//
//  ordinalDayNumberTime          *big.Float
//     - The combined ordinal day number and fractional time
//       value for the ordinal day number represented by the
//       date/time specified by input parameters, targetYear,
//       targetMonth, targetDay, targetHour, targetMinute,
//       targetSecond and targetNanosecond. This value equals
//       the ordinal day number calculated from the start date
//       of January 1, 1 CE on the Revised Julian Calendar.
//
//
//  ordinalDayNumberTimeFraction  *big.Float
//     - The fractional time value associated with the the ordinal
//       day number. This value does NOT contain the integer ordinal
//       day number. Instead, it only contains the time value
//       represented by the date/time input parameters, targetHour,
//       targetMinute, targetSecond and targetNanosecond. This value
//       equals the number of nanoseconds since midnight of the ordinal
//       day number divided by the number of nanoseconds in a
//       24-hour day.
//
//
//  err                           error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calGreg *CalendarGregorianUtility) GetOrdinalDayNumberTime(
	targetYear int64,
	targetMonth,
	targetDay,
	targetHour,
	targetMinute,
	targetSecond,
	targetNanosecond int,
	ePrefix string) (
	ordinalDayNumber int64,
	ordinalDayNumberTime *big.Float,
	ordinalDayNumberTimeFraction *big.Float,
	err error) {

	if calGreg.lock == nil {
		calGreg.lock = &sync.Mutex{}
	}

	calGreg.lock.Lock()

	defer calGreg.lock.Unlock()

	ordinalDayNumber = 0

	ordinalDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	ordinalDayNumberTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	err = nil

	calGreg2 := CalendarGregorianUtility{}

	isTargetLeapYear := calGreg2.IsLeapYear(targetYear)

	calUtil := CalendarUtility{}

	err = calUtil.IsValidDateTimeComponents(
		isTargetLeapYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {

		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	var baseDateTimeDto, targetDateTimeDto DateTimeTransferDto

	targetDateTimeDto, err = DateTimeTransferDto{}.New(
		isTargetLeapYear,
		targetYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	// Base Date Time:  January 1, 0001 00:00:00.000000000 UTC
	baseDateTimeDto, err = DateTimeTransferDto{}.New(
		false,
		int64(1),
		1,
		1,
		0,
		0,
		0,
		0,
		ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	var baseTargetComparisonResult, baseTargetYearsComparison int

	baseTargetComparisonResult, err = baseDateTimeDto.Compare(&targetDateTimeDto, ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	baseTargetYearsComparison = baseDateTimeDto.CompareYears(&targetDateTimeDto)

	// baseDateTimeDto is now less than targetDateTimeDto

	baseDateTimeDto.SetTag("baseDateTimeDto")
	targetDateTimeDto.SetTag("targetDateTimeDto")
	// Compute wholeYearInterval
	var wholeYearsInterval int64

	var targetYearOrdinalNumber, baseYearOrdinalNumber int

	targetYearOrdinalNumber, err = targetDateTimeDto.GetOrdinalDayNumber(ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	wholeYearsInterval = 0

	if baseTargetComparisonResult == 0 {
		// base and target have equivalent date/times
		// Julian Day Number is Zero and time fraction is Zero.

		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err

	} else if baseTargetYearsComparison == 0 {
		// base and target years are equal
		ordinalDayNumber = int64(targetYearOrdinalNumber)




	} else if baseTargetComparisonResult == 1 {
		// base is greater than target date/time
		// target date/time must be negative
		wholeYearsInterval =
			targetDateTimeDto.GetYear() -
				baseDateTimeDto.GetYear() + 1

	} else {
		// target is greater than base date/time
		// target date/time could be positive or
		// negative
		wholeYearsInterval =
			targetDateTimeDto.GetYear() -
				baseDateTimeDto.GetYear() - 1

	}

	if wholeYearsInterval < 0 {
		wholeYearsInterval *= -1
	}

	var wholeYearDays int64

	if wholeYearsInterval == 0 {

		wholeYearDays = 0

	} else {

		wholeYearDays =
			calGreg2.NumCalendarDaysForWholeYearsInterval(wholeYearsInterval)

	}

	baseYearOrdinalNumber, err = baseDateTimeDto.GetOrdinalDayNumber(ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	var lastWholeTargetYear int64
	var lastWholeTargetYearDeltaDays, targetPartialYearDeltaDays int64


	if baseTargetYearsComparison == 1 {
		// base year is greater than target year
		// target year is negative

		lastWholeTargetYear = targetDateTimeDto.GetYear() + 1

	} else {
		// baseTargetYearsComparison == -1
		// base year is less than target year.
		// target could be positive or negative.
		//
		if targetDateTimeDto.GetYearNumberSign() < 0 {
			// target year is negative
			lastWholeTargetYear = targetDateTimeDto.GetYear() + 1

		} else {
			// target year is positive
			lastWholeTargetYear = targetDateTimeDto.GetYear() - 1
		}
	}

	if calGreg2.IsLeapYear(lastWholeTargetYear) {
		lastWholeTargetYearDeltaDays = int64(366) - int64(baseYearOrdinalNumber)
	} else {
		lastWholeTargetYearDeltaDays = int64(365) - int64(baseYearOrdinalNumber)
	}

	targetPartialYearDeltaDays = int64(targetYearOrdinalNumber)

	ordinalDayNumber = wholeYearDays +
		lastWholeTargetYearDeltaDays +
		targetPartialYearDeltaDays

	var targetDayTotalNanoseconds int64

	targetDayTotalNanoseconds, err =
		targetDateTimeDto.GetTotalTimeInNanoseconds(ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	twentyFourHourNanosecondsFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(int64(time.Hour) * 24)

	targetNanosecondFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(targetDayTotalNanoseconds)

	ordinalDayNumberTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Quo(targetNanosecondFloat, twentyFourHourNanosecondsFloat)

	ordinalDayNoFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(ordinalDayNumber)

	ordinalDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Add(ordinalDayNoFloat, ordinalDayNumberTimeFraction)

	separator := strings.Repeat("-", 65)
	fmt.Println()
	fmt.Println(separator)

	fmt.Println(separator)
	fmt.Println(baseDateTimeDto.String())
	fmt.Println(targetDateTimeDto.String())
	fmt.Println(separator)

	return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
}

// GetYearDays - Returns the number of days in the year
// identified by input parameter 'year' under the Gregorian
// Calendar.
//
// If the year is a standard year, this method will return 365-days.
// If the year is a leap year, this method will return 365-days.
//
// For more information on the Gregorian Calendar and leap years,
// reference:
//
//   https://en.wikipedia.org/wiki/Gregorian_calendar
//   https://www.timeanddate.com/date/leapyear.html
//
//
func (calGreg *CalendarGregorianUtility) GetYearDays(
	year int64) int {

	if calGreg.lock == nil {
		calGreg.lock = &sync.Mutex{}
	}

	calGreg.lock.Lock()

	defer calGreg.lock.Unlock()

	calGregMech := calendarGregorianMechanics{}

	isLeapYear := calGregMech.isLeapYear(year)

	if isLeapYear {
		return 366
	}

	return 365
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
//  3. The year is evenly divisible by 100 and the year is also
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
func (calGreg *CalendarGregorianUtility) IsLeapYear(
	year int64) bool {

	if calGreg.lock == nil {
		calGreg.lock = &sync.Mutex{}
	}

	calGreg.lock.Lock()

	defer calGreg.lock.Unlock()
	calGregMech := calendarGregorianMechanics{}

	return calGregMech.isLeapYear(year)
}

// JulianDayNumberTimeToDateTime - Receives a Julian Day Number Time floating
// point value and returns the equivalent date/time under the Gregorian Calendar.
//
// The start date/time for Julian Day Number calculates performed under the
// Gregorian Calendar is November 24, -4713 12:00:00.000000000 UTC (Noon). All
// calculated Julian Day Numbers after this moment are positive. All Julian Day
// Numbers prior to this moment are negative.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  julianDayNumberTime  *big.Float
//     - The Julian Day Number/Time expressed as a floating point value.
//
//
//
//  ePrefix              string
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
//  gregorianDateTime    DateTimeTransferDto
//     - If successful this method will return a new, fully populated instance
//       of type DateTimeTransferDto contain the year, month, day, hour, minute,
//       second and nanosecond date/time value equivalent to the Julian Day Number
//       Time passed as an input parameter.
//
//  err                  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calGreg *CalendarGregorianUtility) JulianDayNumberTimeToDateTime(
	julianDayNumberTime *big.Float,
	ePrefix string) (gregorianDateTime DateTimeTransferDto, err error) {

	if calGreg.lock == nil {
		calGreg.lock = &sync.Mutex{}
	}

	calGreg.lock.Lock()

	defer calGreg.lock.Unlock()

	gregorianDateTime = DateTimeTransferDto{}
	err = nil

	julianDayNumTimeNumberSign := julianDayNumberTime.Sign()

	// Convert to absolute value
	if julianDayNumTimeNumberSign == -1 {
		julianDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Neg(julianDayNumberTime)

	}

	// Truncate to integer
	julianDayNoBigInt, _ :=
		julianDayNumberTime.Int(nil)

	julianDayNoFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt(julianDayNoBigInt)

	julianDayNumberTimeFraction :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Sub(julianDayNumberTime, julianDayNoFloat)

	twentyFourHoursNanoseconds := int64(time.Hour * 24)

	twentyFourHourNanosecondsFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64( twentyFourHoursNanoseconds)

	convertedNanoseconds :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Mul(julianDayNumberTimeFraction, twentyFourHourNanosecondsFloat)

	remainingNanoseconds, _ :=
		convertedNanoseconds.Int64()

	noonNanoseconds := int64(time.Hour * 12)

	remainingNanoseconds += noonNanoseconds

	julianDayNoTimeAdjustment := 0

	if remainingNanoseconds == twentyFourHoursNanoseconds {

		remainingNanoseconds = 0
		julianDayNoTimeAdjustment = 1

	} else if remainingNanoseconds > twentyFourHoursNanoseconds {
		remainingNanoseconds -= twentyFourHoursNanoseconds
		julianDayNoTimeAdjustment = 1
	}

	var hour, minute, second, nanosecond int

	timeMech := TimeMechanics{}

	hour,
	minute,
	second,
	nanosecond,
	_,
	err = timeMech.AllocateNanoseconds(remainingNanoseconds, ePrefix)

	if err != nil {
		return gregorianDateTime, err
	}

	calGregMech := calendarGregorianMechanics{}

	calendarConfig := calGregMech.getCalendarCyclesConfig()

	cycleCount := big.NewInt(0)

	var mainCycleAdjustmentDays *big.Int

	var baseJDNStartDateTime, mainCycleStartDateTime DateTimeTransferDto

	baseJDNStartDateTime =
		calendarConfig.GetJDNBaseStartYearDateTime()

	var mainCycleStartDateYear *big.Int

	if julianDayNumTimeNumberSign == 1 {
		// Julian Day Number is a positive value.

		mainCycleStartDateTime =
			calendarConfig.GetMainCycleStartDateForPositiveJDNNo()

		mainCycleStartDateYear =
			mainCycleStartDateTime.GetYearBigInt()

		fmt.Printf("Main Cycle Start Date Year for Positive JDN.\n" +
			"Initial mainCycleStartDateYear='%v\n",
			mainCycleStartDateYear.Text(10))

		mainCycleAdjustmentDays =
			big.NewInt(0).
				Neg(calendarConfig.GetMainCycleAdjustmentDaysForPositiveJDNNo())

	} else {
		// Julian Day Number is a negative number

		mainCycleStartDateTime =
			calendarConfig.GetMainCycleStartDateForNegativeJDNNo()

		mainCycleStartDateYear =
			mainCycleStartDateTime.GetAbsoluteYearBigIntValue()

		fmt.Printf("Main Cycle Start Date Year for Negative JDN.\n" +
			"Initial mainCycleStartDateYear='%v\n",
			mainCycleStartDateYear.Text(10))

		mainCycleAdjustmentDays =
			big.NewInt(0).
				Neg(calendarConfig.GetMainCycleAdjustmentDaysForNegativeJDNNo())

	}

	fmt.Printf("Initial Remainder Days: %v\n",
		julianDayNoBigInt.Text(10))

	fmt.Printf("Main Cycle Adjustment Days: %v\n",
		mainCycleAdjustmentDays.Text(10))

	remainderDays :=
		big.NewInt(0).
			Add(julianDayNoBigInt, mainCycleAdjustmentDays)

	fmt.Printf("Adjusted Remainder Days: %v\n",
		remainderDays.Text(10))

	deltaYears := big.NewInt(0)

	cycleDtos := calendarConfig.GetCalendarCycleConfigurations()

	for i:=0; i < len(cycleDtos); i++ {

		cycleCount = big.NewInt(0).
			Quo(
				remainderDays,
				cycleDtos[i].GetDaysInCycle())

		remainderDays =
			big.NewInt(0).
				Sub(remainderDays,
					big.NewInt(0).
					Mul(cycleCount,cycleDtos[i].GetDaysInCycle()))

		cycleDtos[i].SetCycleCount(cycleCount)
		cycleDtos[i].SetRemainderDays(remainderDays)

		deltaYears = big.NewInt(0).
			Add(deltaYears, cycleDtos[i].GetCycleCountTotalYears())

		fmt.Println()
		fmt.Printf("%v-Year Cycle\n",
			cycleDtos[i].GetYearsInCycle().Text(10))
		fmt.Printf("Cycle Count: %v\n",
			cycleCount.Text(10))
		fmt.Printf("Cycle Total Days: %v\n",
			cycleDtos[i].GetCycleCountTotalDays().Text(10))
		fmt.Printf("Cycle Total Years: %v\n",
			cycleDtos[i].GetCycleCountTotalYears().Text(10))
		fmt.Printf("Cycle Remainder Days: %v\n",
			remainderDays.Text(10))
	}

	if !remainderDays.IsInt64() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Month Days cannot be converted to Int64!\n" +
			"Month Days='%v'\n",
			remainderDays.Text(10))
		return gregorianDateTime, err
	}

	fmt.Println()
	fmt.Printf("Raw Total Delta Years: %s\n",
		deltaYears.Text(10))

	fmt.Printf("Remainder Days Raw: %s\n",
		remainderDays.Text(10))

	var endingYearBigInt *big.Int

	fmt.Printf("Main Cycle Start Year Value: %v\n",
		mainCycleStartDateYear.Text(10))

	// For a negative JDN Number, mainCycleStartDateYear
	// is a positive year value.
	//
	// For a positive JDN Number, mainCycleStartDateYear
	// is a negative year value.
	endingYearBigInt =
		big.NewInt(0).
			Add(mainCycleStartDateYear, deltaYears)

	if !endingYearBigInt.IsInt64() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Ending Year Number cannot be converted to Int64!\n" +
			"Ending Year Number='%v'\n",
			endingYearBigInt.Text(10))
		return gregorianDateTime, err
	}

	fmt.Printf("Initial Ending Year: %v\n",
		endingYearBigInt)

	var baseJDNDateRemainingDaysInYear, yearAdjustment int

	yearAdjustment = 0

	baseJDNDateRemainingDaysInYear, err =
		baseJDNStartDateTime.GetRemainingDaysInYear(ePrefix)

	if err != nil {
		return gregorianDateTime, err
	}

	var ordinalDayNo int64

	ordinalDayNo = remainderDays.Int64()

	fmt.Printf("Initial Ordinal Day Number: %v\n",
		ordinalDayNo)

	if ordinalDayNo < 0 {
		// Invalid Result!
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Remainder Days (a.k.a. Ordinal Day Number) is LESS THAN ZERO!\n" +
			"Ordinal Day Number: '%v\n",
			ordinalDayNo)

		return gregorianDateTime, err

	} else if ordinalDayNo == 0 {

		ordinalDayNo = int64(365 - baseJDNDateRemainingDaysInYear)

	fmt.Printf("Remaining Day Zero Calc - Ordinal Day Number: %v\n",
		ordinalDayNo)

	} else if ordinalDayNo > 0 {

		if ordinalDayNo > int64(baseJDNDateRemainingDaysInYear) {
			yearAdjustment++
			ordinalDayNo = ordinalDayNo - int64(baseJDNDateRemainingDaysInYear)

			fmt.Printf("Ordinal Day Greater than baseJDNDateRemainingDaysInYear\n" +
				"OrdinalDay: %v\n",
				ordinalDayNo)

		} else {
			// ordinalDayNo must be less than or equal to
			// baseJDNDateRemainingDaysInYear
			ordinalDayNo = 365 - (ordinalDayNo + int64(baseJDNDateRemainingDaysInYear))

			fmt.Printf("Ordinal Day Number Less Than or Eqaul to\n" +
				"base date remaining days in year.\n" +
				"Ordinal Day Number='%v'\n",
				ordinalDayNo)
		}

	} else {
		// ordinalDayNo Less Than Zero
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Remainder Days (a.k.a. Ordinal Day) is Less Than Zero!\n" +
			"Remainder Days: %v\n",
			ordinalDayNo)

		return gregorianDateTime, err
	}

	year := endingYearBigInt.Int64()

	fmt.Printf("Initial Year Value: '%v'\n",
		year)

	fmt.Printf("Year Adjustment: '%v'\n",
		yearAdjustment)

	year += int64(yearAdjustment)


	if julianDayNumTimeNumberSign == -1 {
		year *= -1
	}

	fmt.Printf("Final Year: %v\n",
		year)

	var isLeapYear bool

	isLeapYear = calGregMech.isLeapYear(year)

	fmt.Printf("Julian Day Number Time Adjustment: %v\n",
		julianDayNoTimeAdjustment)

	ordinalDayNo += int64(julianDayNoTimeAdjustment)

	if ordinalDayNo < 0 ||
		ordinalDayNo > 367 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Month Days value is INVALID!\n" +
			"Month Days='%v'\n", ordinalDayNo)
		return gregorianDateTime, err
	}

	calUtil := CalendarUtility{}
	var month, day int
	yearAdjustment = 0

	fmt.Printf("Final Ordinal Day Number: %v\n",
		ordinalDayNo)

	yearAdjustment,
	month,
	day,
	err = calUtil.GetMonthDayFromOrdinalDayNo(
		ordinalDayNo,
		isLeapYear,
		ePrefix)

	if err != nil {
		return gregorianDateTime, err
	}

	fmt.Printf("Ordinal Day Year Adjustment Value: %v\n",
		yearAdjustment)

	if yearAdjustment != 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: calUtil.GetMonthDayFromOrdinalDayNo() returned a 'yearAdjustment'\n" +
			"for an invalid Ordinal Day Number.\n" +
			"Year Adjustment: '%v'\n," +
			"Ordinal Day Number: '%v'\n",
			yearAdjustment, ordinalDayNo)

		return gregorianDateTime, err
	}

	// JDN Time to gross nanoseconds.

	gregorianDateTime = DateTimeTransferDto{
		isLeapYear:          isLeapYear,
		year:                year,
		month:               month,
		day:                 day,
		hour:                hour,
		minute:              minute,
		second:              second,
		nanosecond:          nanosecond,
		tag:                 "",
		isThisInstanceValid: true,
		lock:                new(sync.Mutex),
	}

	return gregorianDateTime, err
}

// NumCalendarDaysForWholeYearsInterval - Computes the total
// number of 24-hour days in a period of years specified by
// input parameter 'wholeYearsInterval'. The number of total
// days is calculated in accordance with the Gregorian Calendar.
//
// Methodology:
//
// In the Gregorian calendar, three criteria must be taken
// into account to correctly identify leap years:
//
// 1. The year must be evenly divisible by 4;
//
// 2. If the year can also be evenly divided by 100, it is not
//    a leap year, unless...
//
//  3. The year is evenly divisible by 100 and the year is also
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
// The input parameter 'wholeYearsInterval' is defined as a series of
// contiguous whole, or complete, years consisting of either 365-days
// or 366-days (in the case of leap years).
//
// No partial years should be included in this interval.
//
//
func (calGreg *CalendarGregorianUtility) NumCalendarDaysForWholeYearsInterval(
	wholeYearsInterval int64) (totalDays int64) {

	if calGreg.lock == nil {
		calGreg.lock = &sync.Mutex{}
	}

	calGreg.lock.Lock()

	defer calGreg.lock.Unlock()

	totalDays = 0

	if wholeYearsInterval < 0 {
		wholeYearsInterval *= -1
	}

	if wholeYearsInterval == 0 {
		return 0
	}

	separator := strings.Repeat("*", 65)

	fmt.Println()
	fmt.Println("NumCalendarDaysForWholeYearsInterval() ")
	fmt.Println(separator)
	fmt.Printf("       Whole Years Interval: %v\n", wholeYearsInterval)

	if wholeYearsInterval >= 900 {

		numOfCycles := wholeYearsInterval / 400

		totalDays = numOfCycles * 146097

		fmt.Printf("  Number of 400-Year Cycles: %v\n", numOfCycles)
		fmt.Printf("Number of Days in %v-Cycles: %v\n", numOfCycles, totalDays)

		wholeYearsInterval = wholeYearsInterval - (numOfCycles * 400)

		fmt.Printf("  Number of Remainder Years: %v\n", wholeYearsInterval)
		fmt.Println(separator)
		fmt.Println()

	}

	totalDays += wholeYearsInterval * 365

	leapDays := wholeYearsInterval / 4

	skipLeapDays := wholeYearsInterval / 100

	totalDays += leapDays - skipLeapDays

	fmt.Println(separator)
	fmt.Printf("Total Days In wholeYearsInterval: %v\n",
		totalDays)
	fmt.Println(separator)
	fmt.Println()

	return totalDays
}
