package datetime

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"
)

type CalendarEngines struct {
	lock  *sync.Mutex
	debugCode  bool
}

// DateTimeToJulianDayNumber - Method stub for a function to calculate the Julian
// Day Number from a specific date for all supported calendars.
//
// For more information on the Julian Day Number, reference:
//   https://en.wikipedia.org/wiki/Julian_day
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  targetDateTimeDto          ADateTimeDto
//     - This type encapsulates the target date/time for which an
//       equivalent Julian Day Number will be calculated.
//
//  calCycleCfg                CalendarCycleConfiguration
//     - This type encapsulates the Calendar Cycle information used
//       in calculating Julian Day Numbers for dates under multiple
//       calendars such as Gregorian, Julian, Revised Julian and
//       Revised Goucher Parker calendars.
//
//  floatResultPrecision       uint
//     - This unsigned integer is used to set the precision for the
//       *big.Float floating point time fraction returned by this
//       method. This 'precision' parameter also controls the internal
//       accuracy for interim floating point calculations performed
//       by this method. For more information on precision and type
//       *big.Float floating point numbers, reference:
//           https://golang.org/pkg/math/big/
//
//
//  ePrefix                    string
//     - A string consisting of the method chain used to call this
//       method. In case of error, this text string is included in
//       the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//
//  julianDayNumber            *big.Int
//     - The integer value of the Julian Day Number returned as a type
//       *big.Int. As an integer type it does NOT contain the time fraction.
//       This value will be negative for negative Julian Day Numbers.
//
//  julianDayNumberTime        *big.Float
//     - This returned *big.Float floating point value contains both the
//       integer Julian Day Number and the time fraction, combined. This
//       value will be negative for negative Julian Day Numbers.
//
//  julianDayNoTimeFraction    *big.Float
//     - This value represents the Julian Day Number time fraction expressed
//       as a floating point value. Remember that the integer portion of
//       this fraction is always zero. The integer Julian Day Number is NOT
//       included in this fraction. This time fraction will always be
//       returned as a positive value. Positive or negative Julian Day numbers
//       can be determined by an analysis of returned parameters,
//       'julianDayNumber' or 'julianDayNumberTime'.
//
//
//  err                        error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calEng *CalendarEngines) DateTimeToJulianDayNumber(
	targetDateTimeDto ADateTimeDto,
	calCycleCfg CalendarCycleConfiguration,
	floatResultPrecision uint,
	ePrefix string) (
	jDNDto JulianDayNoDto,
	err error) {

	if calEng.lock == nil {
		calEng.lock = &sync.Mutex{}
	}

	calEng.lock.Lock()

	defer calEng.lock.Unlock()

	ePrefix += "CalendarEngines.DateTimeToJulianDayNumber() "

	jDNDto = JulianDayNoDto{}

	julianDayNumberTimeFraction :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(floatResultPrecision).
			SetFloat64(0.0)

	err = nil

	if calCycleCfg.GetNumberOfCalendarCycleConfigDtos() == 0 {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'calCyclesCfg' is invalid!\n" +
			"The number of calendar cycles is zero!\n")
		return jDNDto, err
	}

	err = targetDateTimeDto.IsValidInstanceError(ePrefix)

	if err != nil {
		return jDNDto, err
	}

	baseDateTimeDto := calCycleCfg.GetJDNBaseStartYearDateTime()

	err = baseDateTimeDto.IsValidInstanceError(ePrefix)

	if err != nil {
		return jDNDto, err
	}

	// baseYearOrdinalNumber,
	var baseTargetComparisonResult, julianDayNumberSign int

	var mainCycleAdjustmentDays, remainderYears, ordinalDays *big.Int

	baseTargetComparisonResult, err = baseDateTimeDto.Compare(&targetDateTimeDto, ePrefix)

	if err != nil {
		return jDNDto, err
	}

	baseDateTimeDto.SetTag("baseDateTimeDto")
	targetDateTimeDto.SetTag("targetDateTimeDto")

	var mainCycleStartDateTime ADateTimeDto

	jDNMech := julianDayNoDtoMechanics{}

	var calBaseData ICalendarBaseData

	calBaseData, err = calCycleCfg.GetCalendarBaseData(ePrefix)

	if err != nil {
		return jDNDto, err
	}


	// Primary Decision Tree
	if baseTargetComparisonResult == 0 {
		// base and target have equivalent date/times
		// Julian Day Number is Zero and time fraction is Zero.

		err = jDNMech.setBigValDto(
			&jDNDto,
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(floatResultPrecision).
				SetFloat64(0.0),
				1024,
			targetDateTimeDto.date.hasLeapSecond,
			ePrefix)

		return jDNDto, err

	} else if baseTargetComparisonResult == 1 {
		// Base Date Time is GREATER THAN Target Date Time
		// Target Date Time is LESS THAN Base Date Time
		// Target Date Time must be negative
		// Julian Day Number must be negative
		// Use Main Cycle Start Date/Time which is GREATER THAN Base

		mainCycleStartDateTime =
			calCycleCfg.GetMainCycleStartDateForNegativeJDNNo()

		mainCycleStartDateTime.SetTag("Main Cycle Start Date Time")

		if calEng.debugCode {
		fmt.Printf("%s\n",
			mainCycleStartDateTime.String())
		}

			julianDayNumberSign = -1

		// RemainderYears =
		// BaseYear - TargetYear -1
		remainderYears =
			big.NewInt(
				mainCycleStartDateTime.date.astronomicalYear -
					targetDateTimeDto.date.astronomicalYear - 1 )

		mainCycleAdjustmentDays =
			calCycleCfg.GetMainCycleAdjustmentDaysForNegativeJDNNo()

			var baseYearOrdinalNumber, targetDaysRemainingInYear int

			baseYearOrdinalNumber, err = calBaseData.GetOrdinalDayNumber(
				baseDateTimeDto.date.GetIsLeapYear(),
				baseDateTimeDto.date.month,
				baseDateTimeDto.date.day,
				ePrefix)

		if err != nil {
			return jDNDto, err
		}

		targetDaysRemainingInYear,
		err = calBaseData.GetRemainingDaysInYear(
			targetDateTimeDto.date.astronomicalYear,
			CalYearType.Astronomical(),
			targetDateTimeDto.date.month,
			targetDateTimeDto.date.day,
			ePrefix)

		if err != nil {
			return jDNDto, err
		}

		ordinalDays =
			big.NewInt(
				int64(
					baseYearOrdinalNumber +
						targetDaysRemainingInYear))

		if calEng.debugCode {

			fmt.Println()
			fmt.Println("Target Date is Less Than Base Date")
			fmt.Println("Julian Day Number is Negative")
			fmt.Println("baseYearOrdinalNumber+targetDaysRemainingInYear")
			fmt.Printf("    baseYearOrdinalNumber: %v\n",
				baseYearOrdinalNumber)
			fmt.Printf("targetDaysRemainingInYear: %v\n",
				targetDaysRemainingInYear)
			fmt.Printf("             Ordinal Days: %v\n",
				ordinalDays)
			fmt.Println()

		}

	} else {
		// Base Date Time is LESS THAN Target Date Time
		// Target Date Time is GREATER THAN Base Date Time
		// Target Date Time may be positive or negative
		// Julian Day Number must be positive
		// Use Main Cycle Start Date/Time which is LESS THAN Base

		mainCycleStartDateTime =
			calCycleCfg.GetMainCycleStartDateForPositiveJDNNo()

		mainCycleStartDateTime.SetTag("Main Cycle Start Date Time")

		if calEng.debugCode {
			fmt.Printf("%s\n",
				mainCycleStartDateTime.String())
		}

		julianDayNumberSign = 1

		// RemainderYears or Whole Years Interval =
		// TargetYear - mainCycleStartDateTime -1
		remainderYears =
			big.NewInt(
				targetDateTimeDto.date.astronomicalYear -
					mainCycleStartDateTime.date.astronomicalYear - 1)

		mainCycleAdjustmentDays =
			calCycleCfg.GetMainCycleAdjustmentDaysForPositiveJDNNo()

		var baseRemainingDaysInYear, targetYearOrdinalNumber int

		baseRemainingDaysInYear, err = calBaseData.GetRemainingDaysInYear(
			baseDateTimeDto.date.astronomicalYear,
			CalYearType.Astronomical(),
			baseDateTimeDto.date.month,
			baseDateTimeDto.date.day,
			ePrefix)

		if err != nil {
			return jDNDto, err
		}

		targetYearOrdinalNumber, err = calBaseData.GetOrdinalDayNumber(
			targetDateTimeDto.date.GetIsLeapYear(),
			targetDateTimeDto.date.month,
			targetDateTimeDto.date.day,
			ePrefix)

		if err != nil {
			return jDNDto, err
		}

		ordinalDays =
			big.NewInt(
				int64(
					baseRemainingDaysInYear +
						targetYearOrdinalNumber))

		if calEng.debugCode {

			fmt.Println()
			fmt.Println("Target Date is Greater Than Base Date")
			fmt.Println("Julian Day Number is Positive")
			fmt.Println("baseRemainingDaysInYear+targetYearOrdinalNumber")
			fmt.Printf("    baseRemainingDaysInYear: %v\n",
				baseRemainingDaysInYear)
			fmt.Printf("targetYearOrdinalNumber: %v\n",
				targetYearOrdinalNumber)
			fmt.Printf("             Ordinal Days: %v\n",
				ordinalDays)
			fmt.Println()

		}

	}

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


			if calEng.debugCode {
				fmt.Printf("cycles[%v].cycleCount=%v\n",
					i, cycleCount.Text(10))

				fmt.Printf("cycles[%v].yearsInCycle=%v\n",
					i, cycles[i].GetYearsInCycle().Text(10))

				fmt.Printf("cycles[%v].remainderYears=%v\n",
					i, remainderYears.Text(10))
			}

			cycles[i].SetCycleCount(cycleCount)
			cycles[i].SetRemainderYears(remainderYears)

			totalYears =
				big.NewInt(0).
					Add(totalYears, cycles[i].GetCycleCountTotalYears())

			if calEng.debugCode {

				fmt.Printf("cycles[%v].CycleCountTotalDays=%v\n",
					i, cycles[i].GetCycleCountTotalDays().Text(10))

			}

			daysDuration =
				big.NewInt(0).
					Add(daysDuration, cycles[i].GetCycleCountTotalDays())
		}

		if calEng.debugCode {
			fmt.Println()
			fmt.Printf("  totalYears: %v\n",
				totalYears.Text(10))

			fmt.Printf("daysDuration: %v\n",
				daysDuration.Text(10))
			fmt.Printf("mainCycleAdjustmentDays: %v\n",
				mainCycleAdjustmentDays.Text(10))
		}

		daysDuration =
			big.NewInt(0).
				Add(daysDuration, mainCycleAdjustmentDays)

		if calEng.debugCode {
			fmt.Printf("Adjusted daysDuration: %v\n",
				daysDuration.Text(10))
			fmt.Println()
		}

	}

	if calEng.debugCode {
		fmt.Printf(" ordinalDays: %v\n",
			ordinalDays.Text(10))
		fmt.Printf("daysDuration: %v\n",
			daysDuration.Text(10))
	}

	daysDuration =
		big.NewInt(0).
			Add(daysDuration, ordinalDays)

	if calEng.debugCode {
		fmt.Printf("Final daysDuration: %v\n",
			daysDuration.Text(10))
		fmt.Println()
	}

	julianDayNumber :=
		big.NewInt(0).
		Set(daysDuration)

	if calEng.debugCode {
		fmt.Printf("Initial Julian Day Number: %v\n",
			julianDayNumber.Text(10))
	}

	// Compute Time Components

	var targetDayTotalNanoseconds int64

	var julianDayNoTimeAdjustment int

	targetDayTotalNanoseconds, err =
		targetDateTimeDto.GetTotalTimeInNanoseconds(ePrefix)

	if err != nil {
		return jDNDto, err
	}

	timeMech := TimeMechanics{}

	julianDayNumberTimeFraction,
		julianDayNoTimeAdjustment,
		err = timeMech.ComputeJulianDayNumberTimeFraction(
		julianDayNumberSign,
		targetDayTotalNanoseconds,
		floatResultPrecision,
		targetDateTimeDto.date.hasLeapSecond,
		ePrefix)


	if julianDayNumberSign == 1  {
		// Base Date Time is LESS THAN Target Date Time
		// Target Date Time is GREATER THAN Base Date Time
		// Target Date Time may be positive or negative
		// Julian Day Number must be positive

		if calEng.debugCode {

			fmt.Println("julianDayNumberSign == 1")

			fmt.Printf("julianDayNoTimeAdjustment: %v\n",
				julianDayNoTimeAdjustment)

			fmt.Printf("Pre-Time Adjustment Julian Day Number: %v\n",
				julianDayNumber)
		}

		if julianDayNoTimeAdjustment == -1 {
			julianDayNumber =
				big.NewInt(0).
					Add(julianDayNumber,
						big.NewInt(-1))
		}


	} else {
		//  julianDayNumberSign == -1
		// Base Date Time is GREATER THAN Target Date Time
		// Target Date Time must be negative
		// Julian Day Number must be negative


		if calEng.debugCode {

			fmt.Println("julianDayNumberSign == -1")

			fmt.Printf("julianDayNoTimeAdjustment: %v\n",
				julianDayNoTimeAdjustment)

			fmt.Printf("Pre-Time Adjustment Julian Day Number: %v\n",
				julianDayNumber)
		}

		if julianDayNoTimeAdjustment == 1 {
			julianDayNumber =
				big.NewInt(0).
					Add(julianDayNumber,
						big.NewInt(-1))
		}

	}


	if calEng.debugCode {

		fmt.Printf("Post-Adjustment Julian Day Number: %v\n",
			julianDayNumber)

	}

	separator := strings.Repeat("-", 75)

	if calEng.debugCode {
		fmt.Println()
		fmt.Println(separator)

		fmt.Printf("Final Julian Day Number: %v\n",
			julianDayNumber.Text(10))

		fmt.Println(separator)
		fmt.Println(baseDateTimeDto.String())
		fmt.Println(targetDateTimeDto.String())
		fmt.Println(separator)
	}

	mathBFMech := MathBigFloatHelper{}

	julianDayNoTime :=
		mathBFMech.CombineIntFracValues(
			julianDayNumber,
			julianDayNumberTimeFraction,
			julianDayNumberSign,
			1024)

	err = jDNMech.setBigValDto(
		&jDNDto,
		julianDayNoTime,
		1024,
		targetDateTimeDto.date.hasLeapSecond,
		ePrefix)

	return jDNDto, err
}

// JulianDayNoToDateTime - Generates an equivalent date/time value for a
// Julian Day Number/Time value.
func (calEng *CalendarEngines) JulianDayNoToDateTime(
	julianDayNoDto JulianDayNoDto,
	calCycleCfg CalendarCycleConfiguration,
	ePrefix string) (
	targetDateTimeDto ADateTimeDto,
	err error) {

	if calEng.lock == nil {
		calEng.lock = &sync.Mutex{}
	}

	calEng.lock.Lock()

	defer calEng.lock.Unlock()

	ePrefix += "CalendarEngines.JulianDayNoToDateTime() "

	targetDateTimeDto = ADateTimeDto{}

	var julianDayNoTime *big.Float

	julianDayNoTime, err = julianDayNoDto.GetDayNoTimeBigFloat(ePrefix)

	if err != nil {
		return targetDateTimeDto, err
	}

	if len(calCycleCfg.calendarCyclesConfig) == 0 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "calCycleCfg",
			inputParameterValue: "",
			errMsg:              "'calCycleCfg' contains zero CalendarCycleConfigurations!",
			err:                 nil,
		}

		return targetDateTimeDto, err
	}

	precision := uint(1024)

	mathBFloatMech := MathBigFloatHelper{}

	// Number sign is +1, 0 or -1
	julianDayNumTimeNumberSign := julianDayNoTime.Sign()

	if julianDayNumTimeNumberSign == 0 {
		// If the Julian day number is 0.000000, simply return
		// the base calendar date/time
		targetDateTimeDto, err =
			calCycleCfg.jdnBaseStartYearDateTime.CopyOut(ePrefix + "- julianDayNumTimeNumberSign == 0 ")

		return targetDateTimeDto, err
	}

	// Number sign is now +1 or -1

	julianDayNumber, _ :=
		mathBFloatMech.Floor(julianDayNoTime, precision).Int(nil)

	if calEng.debugCode{
		fmt.Println()
		fmt.Println(ePrefix)
		fmt.Printf("Starting Integer julianDayNumber: %v\n",
			julianDayNumber.Text(10))
		fmt.Println()
	}

	julianDayNumber = julianDayNumber.Abs(julianDayNumber)

	// Convert julianDayNoTime to an absolute value
	julianDayNoTime, _ =
		mathBFloatMech.Abs(julianDayNoTime, precision)

	var hour, minute, second, nanosecond,
	julianDayNoTimeAdjustment int

	timeMech := TimeMechanics{}

	if calEng.debugCode {
		fmt.Printf("Absolute Value julianDayNumber: %v\n",
			julianDayNumber.Text(10))

		fmt.Println("Preparing To Calculate Time Hours, Minutes etc.")
		fmt.Printf("julianDayNumTimeNumberSign: %v\n",
			julianDayNumTimeNumberSign)
		fmt.Printf("julianDayNoTime: %46.30f\n",
			julianDayNoTime)
		fmt.Println()
	}

	hour,
		minute,
		second,
		nanosecond,
		julianDayNoTimeAdjustment,
		err = timeMech.ComputeUTCTimeFromJulianDayNoFrac(
		julianDayNumTimeNumberSign,
		julianDayNoTime,
		precision,
		julianDayNoDto.hasLeapSecond,
		ePrefix)

	if err != nil {
		return targetDateTimeDto, err
	}

	if calEng.debugCode {
		fmt.Println("Post Time Calculation")
		fmt.Printf("Hour: %v Minute: %v Second: %v Nanosecond: %v\n",
			hour, minute, second, nanosecond)
		fmt.Printf("julianDayNoTimeAdjustment: %v\n",
			julianDayNoTimeAdjustment)
	}

	if calEng.debugCode {
		fmt.Println("Time Adjustment Post JDNSign Conversion")
		fmt.Printf("julianDayNoTimeAdjustment: %v\n",
			julianDayNoTimeAdjustment)
		fmt.Println()
	}

	cycleCount := big.NewInt(0)

	var mainCycleAdjustmentDays *big.Int

	var  mainCycleStartDateTime ADateTimeDto

	var mainCycleStartDateYear *big.Int

	var mainCycleRemainingDaysInYear,
	mainCycleOrdinalDayNo int

	var calBaseData ICalendarBaseData

	calBaseData, err = calCycleCfg.GetCalendarBaseData(ePrefix)

	if err != nil {
		return targetDateTimeDto, err
	}

 if julianDayNumTimeNumberSign == 1 {
	// julianDayNumTimeNumberSign == 1
	// Julian Day Number is a positive value.

	if julianDayNoTimeAdjustment == -1 {
		 // If Time is Less than noon add 1-day
		 julianDayNoTimeAdjustment = 1
	} else {
		 julianDayNoTimeAdjustment = 0
	}

		mainCycleStartDateTime =
			calCycleCfg.GetMainCycleStartDateForPositiveJDNNo()

		mainCycleStartDateYear =
			mainCycleStartDateTime.GetYearBigInt()

	mainCycleRemainingDaysInYear,
	err = calBaseData.GetRemainingDaysInYear(
		mainCycleStartDateTime.date.astronomicalYear,
		CalYearType.Astronomical(),
		mainCycleStartDateTime.date.month,
		mainCycleStartDateTime.date.day,
		ePrefix)


	 if err != nil {
		 return targetDateTimeDto, err
	 }

		mainCycleAdjustmentDays =
			big.NewInt(0).
				Neg(calCycleCfg.GetMainCycleAdjustmentDaysForPositiveJDNNo())

	 if calEng.debugCode {
		 fmt.Printf("Main Cycle Start Date Year for Positive JDN.\n" +
			" Initial mainCycleStartDateYear: %v\n" +
			"Initial mainCycleAdjustmentDays: %v\n" +
		 	" mainCycleRemainingDaysInYear: %v\n",
			mainCycleStartDateYear.Text(10),
			 mainCycleAdjustmentDays,
			 mainCycleRemainingDaysInYear)

		 fmt.Println()
	 }

	} else {
		// Julian Day Number is a negative number
		// julianDayNumTimeNumberSign == -1

	 if julianDayNoTimeAdjustment >= 0 {
		 // Time is greater than or equal to noon.
		 // Add 1-day
		 julianDayNoTimeAdjustment = 1
	 } else {
		 julianDayNoTimeAdjustment = 0
	 }

		mainCycleStartDateTime =
			calCycleCfg.GetMainCycleStartDateForNegativeJDNNo()

		mainCycleStartDateYear =
			mainCycleStartDateTime.GetAbsoluteYearBigIntValue()

	 mainCycleOrdinalDayNo, err = calBaseData.GetOrdinalDayNumber(
		mainCycleStartDateTime.date.GetIsLeapYear(),
		mainCycleStartDateTime.date.month,
		mainCycleStartDateTime.date.day,
		ePrefix)


	 if err != nil {
		 return targetDateTimeDto, err
	 }

		mainCycleAdjustmentDays =
			big.NewInt(0).
				Neg(calCycleCfg.GetMainCycleAdjustmentDaysForNegativeJDNNo())

		if calEng.debugCode {

			fmt.Printf("Main Cycle Start Date Year for Negative JDN.\n" +
				" Initial mainCycleStartDateYear: %v\n" +
				"Initial mainCycleAdjustmentDays: %v\n" +
				" mainCycleOrdinalDayNo: %v\n",
				mainCycleStartDateYear.Text(10),
				mainCycleAdjustmentDays,
				mainCycleOrdinalDayNo)

			fmt.Println()
		}
	}

	if calEng.debugCode {
		fmt.Printf("Initial Remainder Days: %v\n",
			julianDayNumber.Text(10))

		fmt.Printf("Main Cycle Adjustment Days: %v\n",
			mainCycleAdjustmentDays.Text(10))
	}

	cycleRemainderDays :=
		big.NewInt(0).
			Add(julianDayNumber, mainCycleAdjustmentDays)

	if calEng.debugCode {
		fmt.Printf("Adjusted Remainder Days: %v\n",
			cycleRemainderDays.Text(10))
		fmt.Println()
	}

	cycleDeltaYears := big.NewInt(0)

	cycleDtos := calCycleCfg.GetCalendarCycleConfigurations()

	for i:=0; i < len(cycleDtos); i++ {

		cycleCount = big.NewInt(0).
			Quo(
				cycleRemainderDays,
				cycleDtos[i].GetDaysInCycle())

		cycleRemainderDays =
			big.NewInt(0).
				Sub(cycleRemainderDays,
					big.NewInt(0).
						Mul(cycleCount,cycleDtos[i].GetDaysInCycle()))

		cycleDtos[i].SetCycleCount(cycleCount)
		cycleDtos[i].SetRemainderDays(cycleRemainderDays)

		cycleDeltaYears = big.NewInt(0).
			Add(cycleDeltaYears, cycleDtos[i].GetCycleCountTotalYears())

		if calEng.debugCode {
			fmt.Printf("%v-Year Cycle\n",
				cycleDtos[i].GetYearsInCycle().Text(10))
			fmt.Printf("Cycle Count: %v\n",
				cycleCount.Text(10))
			fmt.Printf("Cycle Total Days: %v\n",
				cycleDtos[i].GetCycleCountTotalDays().Text(10))
			fmt.Printf("Cycle Total Years: %v\n",
				cycleDtos[i].GetCycleCountTotalYears().Text(10))
			fmt.Printf("Cycle Remainder Days: %v\n",
				cycleRemainderDays.Text(10))
			fmt.Println()
		}

	}

	if !cycleRemainderDays.IsInt64() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Month Days cannot be converted to Int64!\n" +
			"Month Days='%v'\n",
			cycleRemainderDays.Text(10))
		return targetDateTimeDto, err
	}

	if calEng.debugCode {
		fmt.Println("---------------------------------------")
		fmt.Printf("Raw Total Delta Years: %s\n",
			cycleDeltaYears.Text(10))

		fmt.Printf("Remainder Days Raw: %s\n",
			cycleRemainderDays.Text(10))

		fmt.Printf("Main Cycle Start Year Value: %v\n",
			mainCycleStartDateYear.Text(10))
		fmt.Println("---------------------------------------")
		fmt.Println()
	}

	var endingYearBigInt *big.Int

	// For a negative JDN Number, mainCycleStartDateYear
	// precedes the actual base start date.
	//
	// For a positive JDN Number, mainCycleStartDateYear
	// is a negative year value.
	endingYearBigInt =
		big.NewInt(0).
			Add(mainCycleStartDateYear, cycleDeltaYears)

	if !endingYearBigInt.IsInt64() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Ending Year Number cannot be converted to Int64!\n" +
			"Ending Year Number='%v'\n",
			endingYearBigInt.Text(10))
		return targetDateTimeDto, err
	}

	if calEng.debugCode {
		fmt.Printf("Initial Ending Year: %v\n",
			endingYearBigInt)
	}

	endingYearSign := 1

	if endingYearBigInt.Sign() == -1 {
	// Generate absolute value for ending year
		endingYearSign = -1

		endingYearBigInt =
			big.NewInt(0).
				Neg(endingYearBigInt)
	}

	targetYear := endingYearBigInt.Int64()

	maxInt := big.NewInt(math.MaxInt32)

	minInt := big.NewInt(math.MinInt32)

	if cycleRemainderDays.Cmp(maxInt) == 1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Big Int 'cycleRemainderDays' is larger than the maximum int32 value!\n" +
			"Maximum int32='%v'" +
			"cycleRemainderDays='%v'\n ",
			maxInt.Text(10),
			cycleRemainderDays.Text(10))
		return targetDateTimeDto, err
	}

	if cycleRemainderDays.Cmp(minInt) == -1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Big Int 'cycleRemainderDays' is less than the minimum int32 value!\n" +
			"Minimum int32='%v'" +
			"cycleRemainderDays='%v'\n ",
			minInt.Text(10),
			cycleRemainderDays.Text(10))
		return targetDateTimeDto, err
	}

	cycleRemainderDaysInt := int(cycleRemainderDays.Int64())

	if calEng.debugCode {
		fmt.Printf("cycleRemainderDaysInt: %v\n",
			cycleRemainderDaysInt)
	}

	var isLeapYear bool

	var targetYearDaysInYear int

	if calEng.debugCode {
		fmt.Printf("Julian Day Number Time Adjustment: %v\n",
			julianDayNoTimeAdjustment)
	}

	cycleRemainderDaysInt += julianDayNoTimeAdjustment

	if calEng.debugCode {
		fmt.Printf("cycleRemainderDaysInt No Adjusted For Time: %v\n",
			cycleRemainderDaysInt)
	}

	var targetOrdinalDayNo int

	// Analyze and Adjust Ordinal Day Number
	if julianDayNumTimeNumberSign == 1 {

		if cycleRemainderDaysInt > mainCycleRemainingDaysInYear {
			// julianDayNumTimeNumberSign == 1
			// Julian Day Number is Positive
			// cycleRemainderDaysInt > mainCycleRemainingDaysInYear
			// targetYear is adjusted


			if endingYearSign == 1 {
				// If the Ending Year Sign is '+',
				// add 1-year to the Absolute Value of
				// EndingYear
				targetYear += 1
			} else {
				// If the Ending Year Sign is '-',
				// subtract 1-year from the Absolute Value
				// of Ending Year
				targetYear += -1
			}


			if calEng.debugCode{
				fmt.Println("julianDayNumTimeNumberSign == 1")
				fmt.Println("Julian Day Number is Positive")
				fmt.Println("cycleRemainderDaysInt > mainCycleRemainingDaysInYear")
				fmt.Println("Ending Year Adjustment = +1")
				fmt.Printf("New Target Ending Year: %v\n",
					targetYear)
				fmt.Printf("Initial cycleRemainderDaysInt: %v\n",
					cycleRemainderDaysInt)
				fmt.Printf(" mainCycleRemainingDaysInYear: %v\n",
					mainCycleRemainingDaysInYear)
			}

			targetOrdinalDayNo = cycleRemainderDaysInt -  mainCycleRemainingDaysInYear

			if endingYearSign == -1 ||
				julianDayNumTimeNumberSign == -1 {
				targetYear *= -1
			}

			isLeapYear, err =
				calBaseData.IsLeapYear(
					targetYear,
					CalYearType.Astronomical(),
					ePrefix)

			if err != nil {
				return targetDateTimeDto, err
			}


			if isLeapYear {
				targetYearDaysInYear = calBaseData.GetDaysInLeapYear()
			} else {
				targetYearDaysInYear = calBaseData.GetDaysInStandardYear()
			}

			if calEng.debugCode {
				fmt.Printf("                  targetYear: %v\n",
					targetYear)
				fmt.Printf("                        isLeapYear: %v\n",
					isLeapYear)
				fmt.Printf("          Ending Year Days In Year: %v\n",
					targetYearDaysInYear)
				fmt.Printf("          Final Ordinal Day Number: %v\n",
					targetOrdinalDayNo)
				fmt.Printf("      mainCycleRemainingDaysInYear: %v\n",
					mainCycleRemainingDaysInYear)
				fmt.Println()
			}

			// Correct Ordinal Day is set

		} else {
			// julianDayNumTimeNumberSign == 1
			// Julian Day Number is Positive
			// cycleRemainderDaysInt <= mainCycleRemainingDaysInYear
			// targetYear is NOT adjusted

			if calEng.debugCode{
				fmt.Println("julianDayNumTimeNumberSign == 1")
				fmt.Println("Julian Day Number is Positive")
				fmt.Println("cycleRemainderDaysInt <= mainCycleRemainingDaysInYear")
				fmt.Println("targetYear is NOT adjusted.")
				fmt.Println("Ending Year Adjustment = 0")
				fmt.Printf("Initial computed cycleRemainderDaysInt: %v\n",
					cycleRemainderDaysInt)
				fmt.Printf(" mainCycleRemainingDaysInYear: %v\n",
					mainCycleRemainingDaysInYear)
			}

			if endingYearSign == -1 ||
				julianDayNumTimeNumberSign == -1 {
				targetYear *= -1
			}

			isLeapYear, err = calBaseData.IsLeapYear(
				targetYear,
				CalYearType.Astronomical(),
				ePrefix)

			if err != nil {
				return targetDateTimeDto, err
			}

			if isLeapYear {
				targetYearDaysInYear = calBaseData.GetDaysInLeapYear()
			} else {
				targetYearDaysInYear = calBaseData.GetDaysInStandardYear()
			}

			targetOrdinalDayNo =
				targetYearDaysInYear - ( mainCycleRemainingDaysInYear - cycleRemainderDaysInt)

			if calEng.debugCode {

				fmt.Printf("                  targetYear: %v\n",
					targetYear)
				fmt.Printf("                        isLeapYear: %v\n",
					isLeapYear)
				fmt.Printf("          Ending Year Days In Year: %v\n",
					targetYearDaysInYear)
				fmt.Printf("Ordinal Day Adjusted For Base Days: %v\n",
					targetOrdinalDayNo)
				fmt.Printf("      mainCycleRemainingDaysInYear: %v\n",
					mainCycleRemainingDaysInYear)
				fmt.Println()
			}
		}

	} else {
		// julianDayNumTimeNumberSign == -1
		// Julian Day Number is Negative

		if cycleRemainderDaysInt >= mainCycleOrdinalDayNo {
			// julianDayNumTimeNumberSign == -1
			// Julian Day Number is Negative
			// cycleRemainderDaysInt >= mainCycleOrdinalDayNo
			// targetYear IS Adjusted

			targetYear++

			if calEng.debugCode{
				fmt.Println("julianDayNumTimeNumberSign == -1")
				fmt.Println("Julian Day Number is Negative")
				fmt.Println("cycleRemainderDaysInt >= mainCycleOrdinalDayNo")
				fmt.Println("Ending Year Adjustment = +1")
				fmt.Printf("Adjusted Target Ending Year: %v\n",
					targetYear)
				fmt.Printf("Initial    cycleRemainderDaysInt: %v\n",
					cycleRemainderDaysInt)
				fmt.Printf("      mainCycleOrdinalDayNo: %v\n",
					mainCycleOrdinalDayNo)
			}

			if endingYearSign == -1 ||
				julianDayNumTimeNumberSign == -1 {
				targetYear *= -1
			}

			isLeapYear,
			err = calBaseData.IsLeapYear(
				targetYear,
				CalYearType.Astronomical(),
				ePrefix)

			if err != nil {
				return targetDateTimeDto, err
			}

			if isLeapYear {
				targetYearDaysInYear = calBaseData.GetDaysInLeapYear()
			} else {
				targetYearDaysInYear = calBaseData.GetDaysInStandardYear()
			}

			targetOrdinalDayNo =
				targetYearDaysInYear - (cycleRemainderDaysInt - mainCycleOrdinalDayNo)

			if calEng.debugCode {
				fmt.Printf("                  targetYear: %v\n",
					targetYear)
				fmt.Printf("                        isLeapYear: %v\n",
					isLeapYear)
				fmt.Printf("     Target Ending Year DaysInYear: %v\n",
					targetYearDaysInYear)
				fmt.Printf("                 Final Ordinal Day: %v\n",
					targetOrdinalDayNo)
				fmt.Printf("             mainCycleOrdinalDayNo: %v\n",
					mainCycleOrdinalDayNo)
				fmt.Println()
			}

			// Correct Ordinal Day is set

		} else {
			// julianDayNumTimeNumberSign == -1
			// Julian Day Number is Negative
			// cycleRemainderDaysInt < mainCycleOrdinalDayNo
			// There IS NO Target Ending Year Adjustment

			if calEng.debugCode{
				fmt.Println("julianDayNumTimeNumberSign == -1")
				fmt.Println("Julian Day Number is Negative")
				fmt.Println("cycleRemainderDaysInt < mainCycleOrdinalDayNo")
				fmt.Println("Ending Year Adjustment = 0")
				fmt.Printf("Initial cycleRemainderDaysInt: %v\n",
					cycleRemainderDaysInt)
				fmt.Printf("   mainCycleOrdinalDayNo: %v\n",
					mainCycleOrdinalDayNo)
			}

			targetOrdinalDayNo = mainCycleOrdinalDayNo - cycleRemainderDaysInt

			if endingYearSign == -1 ||
				julianDayNumTimeNumberSign == -1 {
				targetYear *= -1
			}

			isLeapYear,
			err = calBaseData.IsLeapYear(
				targetYear,
				CalYearType.Astronomical(),
				ePrefix)

			if err != nil {
				return targetDateTimeDto, err
			}

			if isLeapYear {
				targetYearDaysInYear = calBaseData.GetDaysInLeapYear()
			} else {
				targetYearDaysInYear = calBaseData.GetDaysInStandardYear()
			}

			if calEng.debugCode {
				fmt.Printf("          targetYear: %v\n",
					targetYear)
				fmt.Printf("Days In Target Ending Year: %v\n",
					targetYearDaysInYear)
				fmt.Printf("                isLeapYear: %v\n",
					isLeapYear)
				fmt.Printf("  Final Ordinal Day Number: %v\n",
					targetOrdinalDayNo)
				fmt.Println()
			}
		}
	}

	if targetOrdinalDayNo < 0 {
		// Invalid Result!
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Remainder Days (a.k.a. Ordinal Day Number) is LESS THAN ZERO!\n" +
			"Ordinal Day Number: '%v\n",
			targetOrdinalDayNo)

		return targetDateTimeDto, err
	}

	if targetOrdinalDayNo > targetYearDaysInYear {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Month Days value is INVALID!\n" +
			"Month Days='%v'\n" +
			"Max Days In Year='%v'\n",
			targetOrdinalDayNo,
			targetYearDaysInYear)

		return targetDateTimeDto, err
	}

	var month, day int

	yearAdjustment := 0

	yearAdjustment,
		month,
		day,
		err = calBaseData.GetMonthDayFromOrdinalDayNo(
		targetOrdinalDayNo,
		isLeapYear,
		ePrefix)

	if err != nil {
		return targetDateTimeDto, err
	}

	if calEng.debugCode {
		fmt.Printf("Ordinal Day Year Adjustment Value: %v\n",
			yearAdjustment)
	}

	if yearAdjustment != 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: calBaseData.GetMonthDayFromOrdinalDayNo() returned a 'yearAdjustment'\n" +
			"for an invalid Ordinal Day Number.\n" +
			"Year Adjustment: '%v'\n," +
			"Ordinal Day Number: '%v'\n",
			yearAdjustment, targetOrdinalDayNo)

		return targetDateTimeDto, err
	}

	// JDN Time to gross nanoseconds.

	if err != nil {
	return targetDateTimeDto, err
}
	targetDateTimeDto,
	err = ADateTimeDto{}.New(
		calBaseData.GetCalendarSpecification(),
		targetYear,
		CalYearType.Astronomical(),
		month,
		day,
		false,
		hour,
		minute,
		second,
		nanosecond,
		"UTC",
		"",
		"",
		ePrefix)

	return targetDateTimeDto, err
}

// SetCodeDebug - If 'debugOn == true, the print statements
// will be activated to assist in code debugging and monitoring.
//
func (calEng *CalendarEngines)SetCodeDebug(debugOn bool) {
	calEng.debugCode = debugOn
}