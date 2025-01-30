package datetime

import (
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"
)

// TimeMechanics - This type includes instance methods designed
// to perform low level time calculations.
type TimeMechanics struct {
	lock       *sync.Mutex
}

// AllocateNanoseconds - This helper function simply allocates
// total time nanoseconds to hours, minutes, seconds and nanoseconds.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  timeTotalNanoseconds  int64
//     - The total amount of time expressed in nanoseconds. This value
//       will be broken down into component hours, minutes, seconds and
//       nanoseconds. If this value is less than zero, an error will be
//       triggered. Likewise, if the 'timeTotalNanoseconds' value is
//       Greater Than 24-hours, an error will be returned. If 'timeTotalNanoseconds'
//       is EQUAL to 24-hours, the return value 'dayAdjustment' is set to
//       +1 and the returned 'hour', 'minute', 'second' and 'nanosecond'
//       are set to zero. If the time value of parameter 'timeTotalNanoseconds'
//       is GREATER THAN 24-hours, an error will be returned.
//
//
//  ePrefix               string
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
//  hour              int
//     - The number of hours represented by input parameter, 'timeTotalNanoseconds'.
//
//  minute            int
//     - The number of minutes represented by input parameter, 'timeTotalNanoseconds'.
//
//  second            int
//     - The number of seconds represented by input parameter, 'timeTotalNanoseconds'.
//
//  nanosecond        int
//     - The number of nanoseconds represented by input parameter, 'timeTotalNanoseconds'.
//
//
//  dayAdjustment       int
//     - If input parameter 'timeTotalNanoseconds' is EQUAL to 24-hours, this return
//       value is set to +1. Otherwise, this return value is set to zero.
//
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
func (timeMech *TimeMechanics) AllocateNanoseconds(
	timeTotalNanoseconds int64,
	ePrefix string) (
	hour,
	minute,
	second,
	nanosecond,
	dayAdjustment int,
	err error) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	ePrefix += "TimeMechanics.AllocateNanoseconds() "

	if timeTotalNanoseconds < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'timeTotalNanoseconds' is less than zero!\n" +
			"timeTotalNanoseconds='%v'\n", timeTotalNanoseconds)

		return hour, minute, second, nanosecond, dayAdjustment, err
	}

	hour = 0
	minute = 0
	second = 0
	nanosecond = 0
	dayAdjustment = 0
	err = nil

	if timeTotalNanoseconds == 0 {

		return hour, minute, second, nanosecond, dayAdjustment, err
	}

	temp := int64(time.Hour) * 24

	if timeTotalNanoseconds > temp {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'timeTotalNanoseconds' is greater than 24-hours!\n" +
			"timeTotalNanoseconds='%v'\n", timeTotalNanoseconds)

		return hour, minute, second, nanosecond, dayAdjustment, err
	}

	if timeTotalNanoseconds == temp {
		// timeTotalNanoseconds == 24-hours
		dayAdjustment = 1
		return hour, minute, second, nanosecond, dayAdjustment, err
	}

	temp = int64(time.Hour)

	if timeTotalNanoseconds >= temp {
		hour = int(timeTotalNanoseconds / temp)
		timeTotalNanoseconds -= int64(hour) * temp
	}

	temp = int64(time.Minute)

	if timeTotalNanoseconds >= temp {
		minute = int(timeTotalNanoseconds / temp)
		timeTotalNanoseconds -= int64(minute) * temp
	}

	temp = int64(time.Second)

	if timeTotalNanoseconds >= temp {
		second = int(timeTotalNanoseconds / temp)
		timeTotalNanoseconds -= int64(second) * temp
	}

	nanosecond = int(timeTotalNanoseconds)

	return hour, minute, second, nanosecond, dayAdjustment, err
}

// ComputePositiveJulianDayNumberTimeFraction - Computes the Julian Day
// Number time fraction for both positive and negative Julian Day Numbers.
// Input parameter, 'julianDayNumberSign' is used to classify the Julian
// day number as positive or negative.
//
// Note that the input parameter 'applyLeapSecond' is used to adjust for
// 24-hour days which include an extra 'leap second'. Set this parameter
// to 'false' unless you fully understand the implications of 'leap second'.
// Set this parameter to 'true' if, and ONLY if, the day for which this
// time fraction will be computed includes a leap second. For more information
// on the 'leap second', reference:
//      https://en.wikipedia.org/wiki/Leap_second
//
// The Julian Day starts a 12:00:00-hours or Noon.
//
// This method receives an int64 time value specifying the total UTC time
// in nanoseconds. It then proceeds to convert this value into a Julian Day
// Number time fraction of type *big.Float.
//
// For more information on the Julian Day Number, reference:
//   https://en.wikipedia.org/wiki/Julian_day
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  julianDayNumberSign        int
//     - This parameter MUST be set to either +1 or -1. These
//       two values signal whether the time conversion will be
//       performed on a negative Julian Day Number or a positive
//       Julian Day Number. For Negative Julian Day Numbers set
//       this value to '-1'. For Julian Day Numbers Greater than
//       or equal to zero, set this value to '+1'.
//
//
//  totalTimeNanoseconds       int64
//     - This value is converted to a fraction representing the
//       Julian Day Number Time Fraction. Ordinarily, this int64
//       value should never exceed the number of nanoseconds in a
//       24-hour day. In extremely rare circumstances involving a
//       a leap second, this value may equal 24-hours + 1-second.
//
//       For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
//  requestedResultPrecision   uint
//     - This unsigned integer is used to set the precision for the
//       *big.Float floating point time fraction returned by this
//       method. This 'precision' parameter also controls the internal
//       accuracy for interim floating point calculations performed
//       by this method. For more information on precision and type
//       *big.Float floating point numbers, reference:
//           https://golang.org/pkg/math/big/
//
//       If the value of 'requestedResultPrecision' is less than 75,
//       'requestedResultPrecision' will be automatically set to 75.
//
//
//  applyLeapSecond           bool
//     - Set this boolean parameter to 'true' if, and ONLY if, the specified
//       day for which this time fraction will be calculated includes a leap
//       second. A leap second is a one-second adjustment that is occasionally
//       applied to Coordinated Universal Time (UTC) in order to accommodate
//       the difference between precise time (as measured by atomic clocks)
//       and imprecise observed solar time (known as UT1 and which varies due
//       to irregularities and long-term slowdown in the Earth's rotation).
//       If this parameter is set to 'true', the time calculation will assume
//       a the duration of a 'day' is 24-hours and one second. Otherwise, the
//       duration of a day is assumed to consist of 24-hours. For more
//       information on the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
//  ePrefix            string
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
//  julianDayNoTimeFraction    *big.Float
//     - This value represents the Julian Day Number time fraction
//       expressed as a floating point value. Remember that the integer
//       portion of this fraction is always zero. The integer Julian Day
//       Number is NOT included in this fraction. Also, the fractional
//       returned fractional value is always positive meaning that the
//       value is always equal to or greater than zero.
//
//
//  julianDayNoAdjustment      int
//     - The Julian Day starts at 12:00:00-Noon. To aid in computing the
//       final Julian Day Number this 'julianDayNoAdjustment' is set to one
//       of three values:
//
//          1. If the computed hour is less than 12:00:00-Noon, this integer is
//             set to '-1' (minus 1).
//
//          2. If the computed hour is equal to 12:00:00-Noon this integer is
//             set to '0' (zero).
//
//          3. If the computed hour is greater than 12:00:00-Noon this integer is
//             set to '1' (plus 1).
//
//
//  err                        error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//

func (timeMech *TimeMechanics) ComputeJulianDayNumberTimeFraction(
	julianDayNumberSign int,
	totalTimeNanoseconds int64,
	requestedResultPrecision uint,
	applyLeapSecond bool,
	ePrefix string) (
	julianDayNoTimeFraction *big.Float,
	julianDayNoAdjustment int,
	err error) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	ePrefix += "TimeMechanics.ComputeJulianDayNumberTimeFraction() "

	julianDayNoTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(requestedResultPrecision).
			SetFloat64(0.0)

	julianDayNoAdjustment = 0

	err = nil

	twentyFourHourNanoseconds := int64(time.Hour) * 24

	if applyLeapSecond {

		twentyFourHourNanoseconds += int64(time.Second)

	}

	if julianDayNumberSign != -1 &&
		julianDayNumberSign != 1 {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "julianDayNumberSign",
			inputParameterValue: fmt.Sprintf("%v",
				julianDayNumberSign),
			errMsg:              "'julianDayNumberSign' must be equal to +1 or -1",
			err:                 nil,
		}
	}

	if totalTimeNanoseconds < 0 || totalTimeNanoseconds > twentyFourHourNanoseconds {

		tempStr := fmt.Sprintf("The valid range for 'totalTimeNanoseconds' is 0-%v inclusive.",
			twentyFourHourNanoseconds)

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "totalTimeNanoseconds",
			inputParameterValue: fmt.Sprintf("%v", totalTimeNanoseconds),
			errMsg:              "Error: 'totalTimeNanoseconds' is invalid!\n" +
				tempStr,
			err:                 nil,
		}

		return julianDayNoTimeFraction, julianDayNoAdjustment, err
	}

	if requestedResultPrecision < 75 {
		requestedResultPrecision = 75
	}

	noonNanoseconds := int64(time.Hour * 12)

	if julianDayNumberSign == 1 {
	// Julian Day Number is Positive

		if totalTimeNanoseconds == noonNanoseconds {

			julianDayNoTimeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					SetInt64(0)

			julianDayNoAdjustment = 0

		} else if totalTimeNanoseconds > noonNanoseconds {

			rat :=
				big.NewRat(
					totalTimeNanoseconds - noonNanoseconds,
					twentyFourHourNanoseconds)

			julianDayNoTimeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					SetRat(rat)

			julianDayNoAdjustment = 1

		} else {
			// totalTimeNanoseconds < noonNanoseconds

			rat :=
				big.NewRat(totalTimeNanoseconds + noonNanoseconds,
					twentyFourHourNanoseconds)

			julianDayNoTimeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					SetRat(rat)

			julianDayNoAdjustment = -1
		}
	} else {
		// julianDayNumberSign == -1
		// Julian Day Number is Negative

		if totalTimeNanoseconds == noonNanoseconds {

			julianDayNoTimeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					SetInt64(0)

			julianDayNoAdjustment = 0

		} else if totalTimeNanoseconds < noonNanoseconds {


			rat:=
				big.NewRat((totalTimeNanoseconds - noonNanoseconds)*-1,
					twentyFourHourNanoseconds)


			julianDayNoTimeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					SetRat(rat)

			julianDayNoAdjustment = -1

		} else {
			// totalTimeNanoseconds > noonNanoseconds

			totalTimeNanoseconds -= int64(time.Hour) * 36

			rat :=
				big.NewRat(
					totalTimeNanoseconds * -1,
					twentyFourHourNanoseconds)

			julianDayNoTimeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					SetRat(rat)

			julianDayNoAdjustment = 1
		}
	}

	return julianDayNoTimeFraction, julianDayNoAdjustment, err
}

/*
func (timeMech *TimeMechanics) ComputeJulianDayNumberTimeFraction(
	julianDayNumberSign int,
	totalTimeNanoseconds int64,
	requestedResultPrecision uint,
	ePrefix string) (
	julianDayNoTimeFraction *big.Float,
	julianDayNoAdjustment int,
	err error) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	ePrefix += "TimeMechanics.ComputeJulianDayNumberTimeFraction() "

	julianDayNoTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(requestedResultPrecision).
			SetFloat64(0.0)

	julianDayNoAdjustment = 0

	err = nil

	twentyFourHourNanoseconds := int64(time.Hour) * 24

	twentyFourHoursPlusOneSecond := twentyFourHourNanoseconds + int64(time.Second)

	if julianDayNumberSign != -1 &&
		julianDayNumberSign != 1 {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "julianDayNumberSign",
			inputParameterValue: fmt.Sprintf("%v",
				julianDayNumberSign),
			errMsg:              "'julianDayNumberSign' must be equal to +1 or -1",
			err:                 nil,
		}
	}

	if totalTimeNanoseconds < 0 || totalTimeNanoseconds > twentyFourHoursPlusOneSecond {

		tempStr := fmt.Sprintf("The valid range for 'totalTimeNanoseconds' is 0-%v inclusive.",
			twentyFourHoursPlusOneSecond)

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "totalTimeNanoseconds",
			inputParameterValue: fmt.Sprintf("%v", totalTimeNanoseconds),
			errMsg:              "Error: 'totalTimeNanoseconds' is invalid!\n" +
				tempStr,
			err:                 nil,
		}

		return julianDayNoTimeFraction, julianDayNoAdjustment, err
	}

	if totalTimeNanoseconds == twentyFourHoursPlusOneSecond {
		// Assume this day includes a leap second.
		twentyFourHourNanoseconds = twentyFourHoursPlusOneSecond
	}

	noonNanoseconds := int64(time.Hour * 12)

	oneFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(requestedResultPrecision).
			SetInt64(1)

	twentyFourHourNanosecondsBigInt :=
		big.NewInt(twentyFourHourNanoseconds)

	var factor1 *big.Float

	if julianDayNumberSign == 1 {

		utcMinusNoonBigInt :=
			big.NewInt(0).
				Sub(
					big.NewInt(totalTimeNanoseconds),
					big.NewInt(noonNanoseconds) )

		utcMinus12Div24RatResult :=
			big.NewRat(1,1).
				SetFrac(
					utcMinusNoonBigInt,
					twentyFourHourNanosecondsBigInt)

		utcMinus12Div24 :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(requestedResultPrecision).
				SetRat(utcMinus12Div24RatResult)

		if totalTimeNanoseconds == noonNanoseconds {

			factor1 =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					Sub(oneFloat, utcMinus12Div24)

			julianDayNoTimeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					Sub(oneFloat, factor1)

			julianDayNoAdjustment = 0

		} else if totalTimeNanoseconds < noonNanoseconds {
			julianDayNoTimeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					Add(oneFloat, utcMinus12Div24)

			julianDayNoAdjustment = -1

		} else {
			// totalTimeNanoseconds > noonNanoseconds

			factor1 =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					Sub(oneFloat, utcMinus12Div24)

			julianDayNoTimeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					Sub(oneFloat, factor1)

			julianDayNoAdjustment = 1

		}
	} else {
		// julianDayNumberSign == -1

		utcPlusNoonBigInt :=
			big.NewInt(0).
				Add(
					big.NewInt(totalTimeNanoseconds),
					big.NewInt(noonNanoseconds) )

		utcPlus12Div24RatResult :=
			big.NewRat(1,1).
				SetFrac(
					utcPlusNoonBigInt,
					twentyFourHourNanosecondsBigInt)

		utcPlus12Div24 :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(requestedResultPrecision).
				SetRat(utcPlus12Div24RatResult)

		if totalTimeNanoseconds == noonNanoseconds {
			julianDayNoTimeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					Sub(oneFloat, utcPlus12Div24)

			julianDayNoAdjustment = 0

		} else if totalTimeNanoseconds < noonNanoseconds {

			julianDayNoTimeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					Sub(oneFloat, utcPlus12Div24)

			julianDayNoAdjustment = -1

		} else {
			// totalTimeNanoseconds > noonNanoseconds

			utcMinusNoonBigInt :=
				big.NewInt(0).
					Sub(
						big.NewInt(totalTimeNanoseconds),
						big.NewInt(noonNanoseconds) )

			utcMinus12Div24RatResult :=
				big.NewRat(1,1).
					SetFrac(
						utcMinusNoonBigInt,
						twentyFourHourNanosecondsBigInt)

			utcMinus12Div24 :=
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					SetRat(utcMinus12Div24RatResult)

			julianDayNoTimeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(requestedResultPrecision).
					Sub(oneFloat, utcMinus12Div24)

			julianDayNoAdjustment = 1
		}
	}

	return julianDayNoTimeFraction, julianDayNoAdjustment, err
}
*/

// ComputeBigIntNanoseconds - Utility method to sum days, hours, minutes,
// seconds and subMicrosecondNanoseconds and return total subMicrosecondNanoseconds as a type *big.Int.
//
func (timeMech *TimeMechanics) ComputeBigIntNanoseconds(
	days *big.Int,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	ePrefix string) (
	totalNanoseconds *big.Int,
	err error) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	ePrefix += "TimeMechanics.ComputeBigIntNanoseconds() "

	totalNanoseconds = big.NewInt(0)
	err = nil

	if days == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "days",
			inputParameterValue: "",
			errMsg:              "Error: Input parameter 'days' is nil!",
			err:                 nil,
		}

		return totalNanoseconds, err
	}

	temp := big.NewInt(0).
		Mul(days, big.NewInt(int64(time.Hour)*24))

	totalNanoseconds.Add(totalNanoseconds, temp)

	temp = big.NewInt(int64(hours) * int64(time.Hour))

	totalNanoseconds.Add(totalNanoseconds, temp)

	temp = big.NewInt(int64(minutes) * int64(time.Minute))

	totalNanoseconds.Add(totalNanoseconds, temp)

	temp = big.NewInt(int64(seconds) * int64(time.Second))

	totalNanoseconds.Add(totalNanoseconds, temp)

	temp = big.NewInt(int64(nanoseconds))

	totalNanoseconds.Add(totalNanoseconds, temp)

	return totalNanoseconds, err
}


// ComputeFloat64TimeFracToGregorianSeconds - Utility routine to
// compute time elements to nearest second from a float64
// Julian Day Number Time. Constituent hours, minutes and
// seconds are returned as type int in Gregorian Calendar
// time.
//
// Julian Days start at noon. Gregorian days start at
// midnight. This method adjusts the hours to reflect
// Gregorian days.
func (timeMech *TimeMechanics) ComputeFloat64TimeFracToGregorianSeconds(
	julianDayNoTime float64) (
	days int64,
	hours,
	minutes,
	seconds,
	numericalSign int) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	days = 0
	hours = 0
	minutes = 0
	seconds = 0
	numericalSign = 1

	if julianDayNoTime < 0.0 {
		numericalSign = -1
		julianDayNoTime = math.Abs(julianDayNoTime)
	}

	fracSeconds :=
		int64((julianDayNoTime * 86400.0) + 0.5)

	// 86400 seconds in a 24-hour day
	if fracSeconds >= 86400 {
		days = fracSeconds / int64(86400)

		fracSeconds -= days * 86400
	}

	if fracSeconds >= 3600 {
		hours = int(fracSeconds/3600)
		fracSeconds -= int64(hours) * int64(3600)
		if hours >= 12 {
			hours -= 12
		}
	}

	if fracSeconds >= 60 {
		minutes = int(fracSeconds/60)
		fracSeconds -= int64(minutes) * int64(60)
	}

	seconds = int(fracSeconds)

	return days, hours, minutes, seconds, numericalSign
}

// ComputeTimeElementsInt64 - Utility method to break gross subMicrosecondNanoseconds
// int constituent hours, minutes, seconds and remaining subMicrosecondNanoseconds. As
// the method name implies, the return values are of type Int64.
//
func (timeMech *TimeMechanics) ComputeTimeElementsInt64(
	grossNanoSeconds int64) (
	hours,
	minutes,
	seconds,
	nanoSeconds int,
	numericalSign int) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	hours = 0
	minutes = 0
	seconds = 0
	nanoSeconds = 0

	numericalSign = 1

	if grossNanoSeconds < 0 {
		numericalSign = -1
		grossNanoSeconds = grossNanoSeconds *int64(numericalSign)
	}

	if grossNanoSeconds == 0 {
		numericalSign = 0
		return hours, minutes, seconds, nanoSeconds, numericalSign
	}

	temp := int64(time.Hour)

	if grossNanoSeconds >= temp {
		hours = int(grossNanoSeconds/temp)
		grossNanoSeconds -= int64(hours) * temp
	}

	temp = int64(time.Minute)

	if grossNanoSeconds >= temp {
		minutes = int(grossNanoSeconds/temp)
		grossNanoSeconds -= int64(minutes) * temp
	}

	temp = int64(time.Second)

	if grossNanoSeconds >= temp {
		seconds = int(grossNanoSeconds/temp)
		grossNanoSeconds -= int64(seconds) * temp
	}

	nanoSeconds = int(grossNanoSeconds)

	return hours, minutes, seconds, nanoSeconds, numericalSign
}

// ComputeTimeElementsBigInt - Utility routine to break gross subMicrosecondNanoseconds
// int constituent hours, minutes, seconds and remaining subMicrosecondNanoseconds. As
// the method name implies, the return values are of type 'int'.
//
func (timeMech *TimeMechanics) ComputeTimeElementsBigInt(
	grossNanoSeconds *big.Int) (
	days *big.Int,
	hours,
	minutes,
	seconds,
	nanoseconds,
	numericalSign int) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	days = big.NewInt(0)
	hours = 0
	minutes = 0
	seconds = 0
	nanoseconds = 0
	hoursBig := big.NewInt(0)
	minutesBig := big.NewInt(0)
	secondsBig := big.NewInt(0)
	numericalSign = 1

	bigDayNanoSecs := big.NewInt(int64(time.Hour) * 24)
	bigHourNanoSecs := big.NewInt(int64(time.Hour))
	bigMinuteNanoSecs := big.NewInt(int64(time.Minute))
	bigSecondNanoSecs := big.NewInt(int64(time.Second))

	compareResult := big.NewInt(0).Cmp(grossNanoSeconds)

	if compareResult < 0 {
		numericalSign = -1
		grossNanoSeconds = big.NewInt(0).Abs(grossNanoSeconds)
	}

	if compareResult == 0 {
		numericalSign = 0
		return days, hours, minutes, seconds, nanoseconds, numericalSign
	}

	var temp *big.Int

	compareResult =  grossNanoSeconds.Cmp(bigDayNanoSecs)

	if compareResult > -1 {
		days = big.NewInt(0).Div(grossNanoSeconds, bigDayNanoSecs)
		temp = big.NewInt(0).Mul(days, bigDayNanoSecs)
		grossNanoSeconds = big.NewInt(0).Sub(grossNanoSeconds, temp)
	}

	compareResult = grossNanoSeconds.Cmp(bigHourNanoSecs)

	if compareResult > -1 {
		hoursBig = big.NewInt(0).Div(grossNanoSeconds,bigHourNanoSecs)
		temp = big.NewInt(0).Mul(hoursBig, bigHourNanoSecs)
		hours = int(hoursBig.Int64())
		grossNanoSeconds = big.NewInt(0).Sub(grossNanoSeconds, temp)
	}

	compareResult = grossNanoSeconds.Cmp(bigMinuteNanoSecs)

	if compareResult > -1 {
		minutesBig = big.NewInt(0).Div(grossNanoSeconds,bigMinuteNanoSecs)
		temp = big.NewInt(0).Mul(minutesBig, bigMinuteNanoSecs)
		minutes = int(minutesBig.Int64())
		grossNanoSeconds = big.NewInt(0).Sub(grossNanoSeconds, temp)
	}

	compareResult = grossNanoSeconds.Cmp(bigSecondNanoSecs)

	if compareResult > -1 {
		secondsBig = big.NewInt(0).Div(grossNanoSeconds,bigSecondNanoSecs)
		temp = big.NewInt(0).Mul(secondsBig, bigSecondNanoSecs)
		seconds = int(secondsBig.Int64())
		grossNanoSeconds = big.NewInt(0).Sub(grossNanoSeconds, temp)
	}

	nanoseconds = int(grossNanoSeconds.Int64())

	return days, hours, minutes, seconds, nanoseconds, numericalSign
}

// ComputeTimeElementsInt - Utility routine to break gross subMicrosecondNanoseconds
// int constituent hours, minutes, seconds and remaining subMicrosecondNanoseconds. As
// the method name implies, the return values are of type 'int'.
//
func (timeMech *TimeMechanics) ComputeTimeElementsInt(
	grossNanoSeconds int64) (
	hours,
	minutes,
	seconds,
	nanoSeconds int,
	numericalSign int) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	hours = 0
	minutes = 0
	seconds = 0
	nanoSeconds = 0
	numericalSign = 1

	if grossNanoSeconds < 0 {
		numericalSign = -1
		grossNanoSeconds *= -1
	}

	if grossNanoSeconds == 0 {
		numericalSign = 0
		return hours, minutes, seconds, nanoSeconds, numericalSign
	}

	temp := int64(time.Hour)

	if grossNanoSeconds >= temp {
		hours = int(grossNanoSeconds/temp)
		grossNanoSeconds -= int64(hours) * temp
	}

	temp = int64(time.Minute)

	if grossNanoSeconds >= temp {
		minutes = int(grossNanoSeconds/temp)
		grossNanoSeconds -= int64(minutes) * temp
	}

	temp = int64(time.Second)

	if grossNanoSeconds >= temp {
		seconds = int(grossNanoSeconds/temp)
		grossNanoSeconds -= int64(seconds) * temp
	}

	nanoSeconds = int(grossNanoSeconds)

	return hours, minutes, seconds, nanoSeconds, numericalSign
}

// ComputeTotalTimeNanoseconds - Receives input parameters for
// hours, minutes, seconds and nanoseconds. The method then proceeds
// to calculate and return the total time value in nanoseconds.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  hour               int
//     - The hour time component for this time value.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//  minute             int
//     - The minute time component for this time value.
//       The valid range is 0 - 59 inclusive
//
//  second             int
//     - The second time component for this time value.
//       The valid range is 0 - 60 inclusive. The value
//       60 should only used in the case of leap seconds.
///       For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//  nanosecond         int
//     - The nanosecond time component for time value.
//       The valid range is 0 - 999,999,999 inclusive
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
//  totalTimeNanoseconds     int64
//     - If successful, this method will return the total time value in
//       nanoseconds for the combined input parameters 'hour', 'minute',
//       'second' and 'nanosecond'. If no errors occurred during method
//       execution, this returned value is always greater than or equal
//       to zero.
//
//
//  totalTimeExceeds24Hours  bool
//     - If successful, this method will return a boolean value indicating
//       whether the return 'totalTimeNanoseconds' exceeds 24-hours. If this
//       value is 'true', the time value of 'totalTimeNanoseconds' is greater
//       than 24-hours. Otherwise, it is not greater than 24-hours.
//
//  err                      error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (timeMech *TimeMechanics) ComputeTotalTimeNanoseconds(
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (
	totalTimeNanoseconds int64,
	totalTimeExceeds24Hours bool,
	err error) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	ePrefix += "TimeMechanics.ComputeTotalTimeNanoseconds() "

	totalTimeNanoseconds = 0

	totalTimeExceeds24Hours = false

	err = nil

	if hour < 0 || hour > 23 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hour' is invalid!\n" +
			"Valid range for 'hour' is 0 through 23, inclusive.\n" +
			"hour='%v'\n", hour)

		return totalTimeNanoseconds, totalTimeExceeds24Hours, err
	}

	if minute < 0 || minute > 59 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'minute' is invalid!\n" +
			"Valid range for 'minute' is 0 through 59, inclusive.\n" +
			"minute='%v'\n", minute)

		return totalTimeNanoseconds, totalTimeExceeds24Hours, err
	}

	if second < 0 || second > 60 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'second' is invalid!\n" +
			"Valid range for 'second' is 0 through 60, inclusive.\n" +
			"second='%v'\n", second)

		return totalTimeNanoseconds, totalTimeExceeds24Hours, err
	}

	if nanosecond < 0 || nanosecond > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecond' is invalid!\n" +
			"Valid range for 'nanosecond' is 0 through 999999999, inclusive.\n" +
			"nanosecond='%v'\n", second)

		return totalTimeNanoseconds, totalTimeExceeds24Hours, err
	}

	if hour == 0 &&
		minute == 0 &&
		second == 0 &&
		nanosecond == 0 {
		return totalTimeNanoseconds, totalTimeExceeds24Hours, err
	}

	totalTimeNanoseconds += int64(hour) * int64(time.Hour)

	totalTimeNanoseconds += int64(minute) * int64(time.Minute)

	totalTimeNanoseconds += int64(second) * int64(time.Second)

	totalTimeNanoseconds +=  int64(nanosecond)

	if totalTimeNanoseconds > int64(time.Hour) * 24 {
		totalTimeExceeds24Hours = true
	}

	return totalTimeNanoseconds, totalTimeExceeds24Hours, err
}

// ComputeUTCTimeFromJulianDayNoFrac - Computes the Universal Coordinated
// Time value from a Julian Day Number/Time Fraction. This method receives
// a floating point number representing the fractional portion of a 24-Hour
// Julian Day Number.  The method then computes and returns the equivalent
// UTC 'hour', 'minute' 'second' and 'nanosecond'. 'nanosecond' is rounded
// to the nearest nanosecond.
//
// Note that the input parameter 'applyLeapSecond' is used to adjust for
// 24-hour days which include an extra 'leap second'. Set this parameter
// to 'false' unless you fully understand the implications of 'leap second'.
// Set this parameter to 'true' if, and ONLY if, the day for which this
// UTC time value will be computed includes a leap second. For more information
// on the 'leap second', reference:
//      https://en.wikipedia.org/wiki/Leap_second
//
// The Julian Day starts at 12:00:00-hours or noon. The 24-hour day used by
// Universal Coordinated Time begins at 00:00:00-hours or midnight.
//
// For more information on the Julian Day Number and its associated time fraction
// see:
//          https://en.wikipedia.org/wiki/Julian_day
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  julianDayNumberSign       int
//     - This parameter MUST be set to either +1 or -1. These
//       two values signal whether the time conversion will be
//       performed on a negative Julian Day Number or a positive
//       Julian Day Number. For Negative Julian Day Numbers set
//       this value to '-1'. For Julian Day Numbers Greater than
//       or equal to zero, set this value to '+1'.
//
//       The accuracy of this parameter is of paramount importance.
//       If this value is incorrect it will product an incorrect
//       result.
//
//
//  julianDayNoTimeFraction   *big.Float
//     - This floating point value may submitted either as a Julian Day
//       Number/Time fraction including integer day number or as a time
//       fraction with a zero integer value to the left of the decimal
//       point. This method will only use the absolute value of the fractional
//       portion of 'julianDayNoTimeFraction'. Julian 'Days' identified by
//       the Julian Day Number begin at 12:00:00-hours or noon.
//       For more information on the Julian Day Number/Time see:
//          https://en.wikipedia.org/wiki/Julian_day
//
//
//  requestedPrecision        uint
//     - This unsigned integer value will determine the precision of
//       internal calculations performed on the input parameters. The
//       'requestedPrecision' value applies to the internal accuracy
//        maintained by type *big.Float floating point values. For those
//        seeking to maximize accuracy. Try values of '512' or '1024'.
//        For more information on precision and type *big.Float floating
//        point numbers, reference:
//               https://golang.org/pkg/math/big/
//
//
//  applyLeapSecond           bool
//     - Set this boolean parameter to 'true' if, and ONLY if, the specified
//       day for which this time value will be calculated includes a leap
//       second. A leap second is a one-second adjustment that is occasionally
//       applied to Coordinated Universal Time (UTC) in order to accommodate
//       the difference between precise time (as measured by atomic clocks)
//       and imprecise observed solar time (known as UT1 and which varies due
//       to irregularities and long-term slowdown in the Earth's rotation).
//       If this parameter is set to 'true', the time calculation will assume
//       the duration of a 'day' is 24-hours plus one second. Otherwise, the
//       duration of a day is assumed to consist of exactly 24-hours. For more
//       information on the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
//  ePrefix               string
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
//  hour              int
//     - The number of hours represented by input parameter, 'julianDayNoTimeFraction'.
//
//  minute            int
//     - The number of minutes represented by input parameter, 'julianDayNoTimeFraction'.
//
//  second            int
//     - The number of seconds represented by input parameter, 'julianDayNoTimeFraction'.
//
//  nanosecond        int
//     - The number of nanoseconds represented by input parameter, 'julianDayNoTimeFraction'.
//
//
//  dayAdjustment       int
//     - If return parameter is set to one of three values depending on the
//       final computed UTC Time:
//
//       Less Than 12:00:00.000000000    - dayAdjustment = -1
//       12:00:00.000000000              - dayAdjustment = 0
//       Greater Than 12:00:00.000000000 - dayAdjustment = +1
//
//       The Julian Day begin at 12:00:00-hours or noon.  For more information
//       on the Julian Day Number/Time reference:
//         https://en.wikipedia.org/wiki/Julian_day
//
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
func (timeMech *TimeMechanics) ComputeUTCTimeFromJulianDayNoFrac(
	julianDayNumberSign int,
	julianDayNoTimeFraction *big.Float,
	requestedPrecision uint,
	applyLeapSecond bool,
	ePrefix string) (
	hour,
	minute,
	second,
	nanosecond,
	julianDayNoTimeAdjustment int,
	err error) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	ePrefix += "TimeMechanics.ComputeUTCTimeFromJulianDayNoFrac() "

	hour = 0
	minute = 0
	second = 0
	nanosecond = 0
	julianDayNoTimeAdjustment = 0
	err = nil

	if julianDayNoTimeFraction == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "julianDayNoTimeFraction",
			inputParameterValue: "nil",
			errMsg:              "'julianDayNoTimeFraction' is INVALID!\n" +
				"'nil' pointer!",
			err:                 nil,
		}

		return hour,
			minute,
			second,
			nanosecond,
			julianDayNoTimeAdjustment,
			err
	}

	if julianDayNumberSign != -1 &&
		julianDayNumberSign != 1 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "julianDayNumberSign",
			inputParameterValue: fmt.Sprintf("%d",julianDayNumberSign),
			errMsg:              "'julianDayNumberSign' MUST be equal to one of two values " +
				"+1 or -1.",
			err:                 nil,
		}

		return hour,
			minute,
			second,
			nanosecond,
			julianDayNoTimeAdjustment,
			err
	}

	if julianDayNoTimeFraction.Prec() > requestedPrecision {
		requestedPrecision = julianDayNoTimeFraction.Prec() + 100
	}

	mathBFloatMech := MathBigFloatHelper{}

	julianDayNoTimeFracFloat, _ :=
		mathBFloatMech.Abs(julianDayNoTimeFraction, requestedPrecision)

	julianDayNoTimeFracFloat,
	_ =
		mathBFloatMech.FloatFractionalValue(
			julianDayNoTimeFracFloat,
			requestedPrecision)

	if julianDayNoTimeFracFloat.Sign() == 0 {

		hour = 12

		return hour,
			minute,
			second,
			nanosecond,
			julianDayNoTimeAdjustment,
			err
	}

	var twentyFourHoursNanoseconds int64

	twentyFourHoursNanoseconds = int64(time.Hour * 24)

	if applyLeapSecond == true {

		twentyFourHoursNanoseconds += int64(time.Second)

	}

	twentyFourHourNanosecondsFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64( twentyFourHoursNanoseconds)

	twelveHourNanoseconds := int64(time.Hour * 12)

	var remainingNanosecondsFloat *big.Float
	var remainingNanosecondsInt64 int64

	remainingNanosecondsFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Mul(twentyFourHourNanosecondsFloat, julianDayNoTimeFracFloat)

	remainingNanosecondsFloat =
		mathBFloatMech.RoundHalfAwayFromZero(
			remainingNanosecondsFloat,
			requestedPrecision,
			0)

	remainingNanosecondsInt64, _ =
		remainingNanosecondsFloat.Int64()

	pointFiveValueRat :=
		big.NewRat(1, 2)

	julianDayNoTimeRat, accuracy :=
		julianDayNoTimeFracFloat.Rat(nil)

	if accuracy != 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error occurred converting 'julianDayNoTimeFraction' to a " +
			"Rational Number (Type Rat).\n" +
			"Accuracy='%v'\n", accuracy)

		return hour,
			minute,
			second,
			nanosecond,
			julianDayNoTimeAdjustment,
			err

	}

	cmpFracToPointFive := julianDayNoTimeRat.Cmp(pointFiveValueRat)

	if julianDayNumberSign == -1 {
		// julianDayNumberSign == -1
		// Julian Day Number is Negative

		if cmpFracToPointFive == 1 {
			// julianDayNumberSign == -1
			// julianDayNoTimeFracFloat is greater than 0.5
			// 0.5 is a.k.a Midnight

			remainingNanosecondsInt64 *= -1
			remainingNanosecondsInt64 += int64(time.Hour) * 36

		} else {
			// julianDayNumberSign == -1
			// julianDayNoTimeFracFloat is less than or equal to 0.5
			// 0.5 is a.k.a Midnight
			remainingNanosecondsInt64 *= -1
			remainingNanosecondsInt64 += twelveHourNanoseconds

		}

	// End of if julianDayNumberSign == -1 {
	} else {
		// julianDayNumberSign == +1
		// Julian Day Number is Positive

		if cmpFracToPointFive == -1 {
			// julianDayNumberSign == +1
			// Time Faction is Less Than 0.5
			// 0.5 is a.k.a Midnight
			remainingNanosecondsInt64 += twelveHourNanoseconds

		} else {
			// julianDayNumberSign == +1
			// Time Fraction is Greater Than or Equal to 0.5
			// 0.5 is a.k.a Midnight
			remainingNanosecondsInt64 -= twelveHourNanoseconds
		}
	}

	if remainingNanosecondsInt64 < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Calculation Error: Computed value 'remainingNanosecondsInt64' " +
			"is less than zero!\n")

		return hour,
			minute,
			second,
			nanosecond,
			julianDayNoTimeAdjustment,
			err

	}

	if remainingNanosecondsInt64 == twelveHourNanoseconds {
		julianDayNoTimeAdjustment = 0
	} else if remainingNanosecondsInt64 > twelveHourNanoseconds {
		julianDayNoTimeAdjustment = 1
	} else {
		// remainingNanosecondsInt64 < twelveHourNanoseconds
		// time is less than noon
		julianDayNoTimeAdjustment = -1
	}

	temp := int64(time.Hour)

	if remainingNanosecondsInt64 >= temp {
		hour = int(remainingNanosecondsInt64 / temp)
		remainingNanosecondsInt64 -= int64(hour) * temp
	}

	temp = int64(time.Minute)

	if remainingNanosecondsInt64 >= temp {
		minute = int(remainingNanosecondsInt64 / temp)
		remainingNanosecondsInt64 -= int64(minute) * temp
	}

	temp = int64(time.Second)

	if remainingNanosecondsInt64 >= temp {
		second = int(remainingNanosecondsInt64 / temp)
		remainingNanosecondsInt64 -= int64(second) * temp
	}

	nanosecond = int(remainingNanosecondsInt64)

	return hour,
		minute,
		second,
		nanosecond,
		julianDayNoTimeAdjustment,
		err
}

/*
func (timeMech *TimeMechanics) ComputeUTCTimeFromJulianDayNoFrac(
	julianDayNoTimeFraction *big.Float,
	requestedPrecision uint,
	ePrefix string) (
	hour,
	minute,
	second,
	nanosecond,
	julianDayNoTimeAdjustment int,
	err error) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	ePrefix += "TimeMechanics.ComputeUTCTimeFromJulianDayNoFrac() "

	hour = 0
	minute = 0
	second = 0
	nanosecond = 0
	julianDayNoTimeAdjustment = 0
	err = nil

	if julianDayNoTimeFraction == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "julianDayNoTimeFraction",
			inputParameterValue: "nil",
			errMsg:              "'julianDayNoTimeFraction' is INVALID!\n" +
				"'nil' pointer!",
			err:                 nil,
		}

		return hour,
			minute,
			second,
			nanosecond,
			julianDayNoTimeAdjustment,
			err
	}

	mathBFloatMech := MathBigFloatHelper{}

	if julianDayNoTimeFraction.Prec() > requestedPrecision {
		requestedPrecision = julianDayNoTimeFraction.Prec() + 100
	}

	julianDayNoTimeFraction, _ =
		mathBFloatMech.Abs(julianDayNoTimeFraction, requestedPrecision)

	if julianDayNoTimeFraction.Sign() == 0 {

		hour = 12

		return hour,
			minute,
			second,
			nanosecond,
			julianDayNoTimeAdjustment,
			err
	}

	julianDayNoTimeFracFloat,
	_ :=
		mathBFloatMech.FloatFractionalValue(
			julianDayNoTimeFraction,
			requestedPrecision)

	twentyFourHoursNanoseconds := int64(time.Hour * 24)

	twentyFourHourNanosecondsFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64( twentyFourHoursNanoseconds)

	convertedNanosecondsFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Mul(julianDayNoTimeFracFloat, twentyFourHourNanosecondsFloat)

	convertedNanosecondsFloat =
		mathBFloatMech.RoundHalfAwayFromZero(
			convertedNanosecondsFloat,
			requestedPrecision,
			0)

	remainingNanoseconds, _ :=
		convertedNanosecondsFloat.Int64()

	noonNanoseconds := int64(time.Hour * 12)

	remainingNanoseconds += noonNanoseconds

	if remainingNanoseconds == twentyFourHoursNanoseconds {

		remainingNanoseconds = 0
		julianDayNoTimeAdjustment = 1

	} else if remainingNanoseconds > twentyFourHoursNanoseconds {
		remainingNanoseconds -= twentyFourHoursNanoseconds
		julianDayNoTimeAdjustment = 1
	}

	temp := int64(time.Hour)

	if remainingNanoseconds >= temp {
		hour = int(remainingNanoseconds / temp)
		remainingNanoseconds -= int64(hour) * temp
	}

	temp = int64(time.Minute)

	if remainingNanoseconds >= temp {
		minute = int(remainingNanoseconds / temp)
		remainingNanoseconds -= int64(minute) * temp
	}

	temp = int64(time.Second)

	if remainingNanoseconds >= temp {
		second = int(remainingNanoseconds / temp)
		remainingNanoseconds -= int64(second) * temp
	}

	nanosecond = int(remainingNanoseconds)


	return hour,
		minute,
		second,
		nanosecond,
		julianDayNoTimeAdjustment,
		err
}
*/