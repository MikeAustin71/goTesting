package datetime

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"sync"
)

type dateTransferDtoUtility struct {
	lock    *sync.Mutex
}


// copyIn - Makes a deep copy of incoming DateTransferDto instance
// 'incomingDateTransDto' and stores the data in the internal member
// variables of 'oldDateTransDto', the original DateTransferDto instance.
// All member variable data values in 'oldDateTransDto' will be
// overwritten.
//
// If this method completes successfully, the internal member variable
// data values for both 'oldDateTransDto' and 'incomingDateTransDto' will
// be identical.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  oldDateTransDto        *DateTransferDto
//     - A pointer to an instance of DateTransferDto. This method will
//       perform a deep copy of member variable data values from input
//       parameter 'incomingDateTransDto' to 'oldDateTransDto'. All original
//       member variable data values contained in 'oldDateTransDto' will be
//       overwritten.
//
//
//  incomingDateTransDto   *DateTransferDto
//     - A pointer to an instance of DateTransferDto. This method will
//       perform a deep copy and transfer member variable data values
//       from 'incomingDateTransDto' to 'oldDateTransDto'.
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
//  err                    error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//
func (dateTransUtil *dateTransferDtoUtility) copyIn(
	oldDateTransDto       *DateTransferDto,
	incomingDateTransDto  *DateTransferDto,
	ePrefix               string) (
	err error) {

	if dateTransUtil.lock == nil {
		dateTransUtil.lock = new(sync.Mutex)
	}

	dateTransUtil.lock.Lock()

	defer dateTransUtil.lock.Unlock()

	ePrefix += "dateTransferDtoUtility.copyIn() "

	err = nil

	if oldDateTransDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'oldDateTransDto' is INVALID!\n" +
			"oldDateTransDto=nil!")
		return  err
	}

	if oldDateTransDto.lock == nil {
		oldDateTransDto.lock = new(sync.Mutex)
	}

	if incomingDateTransDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'incomingDateTransDto' is INVALID!\n" +
			"incomingDateTransDto=nil!")
		return  err
	}

	if incomingDateTransDto.lock == nil {
		incomingDateTransDto.lock = new(sync.Mutex)
	}

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	_, err = dateTransDtoNanobot.testDateTransferDtoValidity(
		incomingDateTransDto,
		ePrefix + "Testing 'incomingDateTransDto' Validity ")

	if err != nil {
		return err
	}

	dateTransDtoMech := dateTransferDtoMechanics{}

	err = dateTransDtoMech.empty(
		oldDateTransDto,
		ePrefix + "Emptying 'oldDateTransDto' ")

	oldDateTransDto.astronomicalYear   = incomingDateTransDto.astronomicalYear
	oldDateTransDto.yearNumberingMode  = incomingDateTransDto.yearNumberingMode
	oldDateTransDto.yearNumType        = incomingDateTransDto.yearNumType
	oldDateTransDto.month              = incomingDateTransDto.month
	oldDateTransDto.day                = incomingDateTransDto.day
	oldDateTransDto.hasLeapSecond      = incomingDateTransDto.hasLeapSecond
	oldDateTransDto.calendarBaseData   = incomingDateTransDto.calendarBaseData.New()
	oldDateTransDto.tag                = incomingDateTransDto.tag

	_, err = dateTransDtoNanobot.testDateTransferDtoValidity(
		oldDateTransDto,
		ePrefix + "Final Validity Check 'oldDateTransDto' ")

	return err

}


// copyOut - Returns a deep copy of input parameter 'dateTransDto'
// which is a pointer to an instance of 'DateTransferDto'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTransDto        *DateTransferDto
//     - A pointer to an instance of DateTransferDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       This method will validate this input parameter and
//       return an error if it is found to be invalid.
//
//
//  ePrefix            string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newDateTransDto     DateTransferDto
//     - If successful this method returns a deep copy of the input
//       parameter, 'dateTransDto'.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (dateTransUtil *dateTransferDtoUtility) copyOut(
	dateTransDto *DateTransferDto,
	ePrefix            string) (
	newDateTransDto DateTransferDto,
	err error) {

	if dateTransUtil.lock == nil {
		dateTransUtil.lock = new(sync.Mutex)
	}

	dateTransUtil.lock.Lock()

	defer dateTransUtil.lock.Unlock()

	ePrefix += "dateTransferDtoUtility.copyOut() "

	err = nil

	newDateTransDto = DateTransferDto{}

	if dateTransDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'dateTransDto' is INVALID!\n" +
			"dateTransDto=nil!")
		return newDateTransDto, err
	}

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDtoNanobot :=dateTransferDtoNanobot{}

	_,
	err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDto,
		ePrefix + "Testing validity of 'dateTransDto'. ")

	if err != nil {
		return newDateTransDto, err
	}

	dateTransDtoMech := dateTransferDtoMechanics{}

	err = dateTransDtoMech.empty(
		&newDateTransDto,
		ePrefix + "Emptying 'newDateTransDto' ")

	if err != nil {
		return newDateTransDto, err
	}

	newDateTransDto.astronomicalYear   = dateTransDto.astronomicalYear
	newDateTransDto.yearNumberingMode  = dateTransDto.yearNumberingMode
	newDateTransDto.yearNumType        = dateTransDto.yearNumType
	newDateTransDto.month              = dateTransDto.month
	newDateTransDto.day                = dateTransDto.day
	newDateTransDto.hasLeapSecond      = dateTransDto.hasLeapSecond
	newDateTransDto.calendarBaseData   = dateTransDto.calendarBaseData.New()
	newDateTransDto.tag                = dateTransDto.tag

	_, err = dateTransDtoNanobot.testDateTransferDtoValidity(
		&newDateTransDto,
		ePrefix + "Final Validity Check 'newDateTransDto' ")

	return newDateTransDto, err
}

// compare - Compares two instances of DateTransferDto date values.
// The method will return one of three comparison values.
//
//    Comparison                Comparison
//      Result                    Status
//    ----------                ----------
//       -1          dateTransDtoOne < dateTransDtoTwo
//        0          dateTransDtoOne = dateTransDtoTwo
//       +1          dateTransDtoOne > dateTransDtoTwo
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dateTransDtoOne          *DateTransferDto
//     - A pointer to an instance of DateTransferDto. This method will
//       compare the date value of this DateTransferDto instance to the
//       date value of input parameter 'dateTransDtoTwo'.
//
//       If 'dateTransDtoOne' is an invalid instance, an error will
//       be returned.
//
//
//  dateTransDtoTwo          *DateTransferDto
//     - A pointer to an instance of DateTransferDto. This method will
//       the date value of input parameter 'dateTransDtoOne' to that of
//       this input parameter ('dateTransDtoTwo').
//
//       If 'dateTransDtoTwo' is an invalid instance, an error will
//       be returned.
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
//  compareResult            int
//     - If this method completes successfully, a comparison result will
//       be returned as an integer value. The result is based on a
//       comparison of date values between the input parameters,
//       'dateTransDtoOne' and 'dateTransDtoTwo'.
//
//       If 'dateTransDtoOne' is LESS THAN 'dateTransDtoTwo', this method
//       will return an integer value of minus one (-1).
//
//       If 'dateTransDtoOne' is EQUAL to 'dateTransDtoTwo', this method
//       will return an integer value of zero (0).
//
//       If 'dateTransDtoOne' is GREATER THAN 'dateTransDtoTwo', this
//       method will return an integer value of plus one (+1).
//
//
//  err                      error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func(dateTransUtil *dateTransferDtoUtility) compare(
	dateTransDtoOne *DateTransferDto,
	dateTransDtoTwo *DateTransferDto,
	ePrefix string) (
	compareResult int,
	err error) {

	if dateTransUtil.lock == nil {
		dateTransUtil.lock = new(sync.Mutex)
	}

	dateTransUtil.lock.Lock()

	defer dateTransUtil.lock.Unlock()

	ePrefix += "dateTransferDtoUtility.compare() "

	err = nil
	compareResult = math.MinInt32

	if dateTransDtoOne == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'dateTransDtoOne' is INVALID!\n" +
			"dateTransDtoOne=nil!")
		return compareResult, err
	}

	if dateTransDtoOne.lock == nil {
		dateTransDtoOne.lock = new(sync.Mutex)
	}

	if dateTransDtoTwo == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'dateTransDtoTwo' is INVALID!\n" +
			"dateTransDtoTwo=nil!")
		return compareResult, err
	}

	if dateTransDtoTwo.lock == nil {
		dateTransDtoTwo.lock = new(sync.Mutex)
	}

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	_,
	err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDtoOne,
		ePrefix + "Testing validity of 'dateTransDtoOne' ")

	if err != nil {
		return compareResult, err
	}

	_,
	err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDtoTwo,
		ePrefix + "Testing validity of 'dateTransDtoTwo' ")

	if err != nil {
		return compareResult, err
	}

	if dateTransDtoOne.astronomicalYear !=
		dateTransDtoTwo.astronomicalYear {

		if dateTransDtoOne.astronomicalYear <
			dateTransDtoTwo.astronomicalYear {
			compareResult = -1
		} else {
			compareResult = 1
		}

		return compareResult, err
	}

	var ordinalDayNoDtoOne, ordinalDayNoDtoTwo int

	ordinalDayNoDtoOne,
	err = dateTransDtoOne.GetOrdinalDayNoInYear(ePrefix)

	if err != nil {
		return compareResult, err
	}

	ordinalDayNoDtoTwo,
	err = dateTransDtoTwo.GetOrdinalDayNoInYear(ePrefix)

	if err != nil {
		return compareResult, err
	}

	if ordinalDayNoDtoOne !=
		ordinalDayNoDtoTwo {

		if ordinalDayNoDtoOne <
			ordinalDayNoDtoTwo {
			compareResult = -1
		} else {
			// Ordinal Day for dateTransDtoTwo
			// MUST BE GREATER THAN ordinalDayNoDtoOne
			compareResult = 1
		}

		return compareResult, err
	}

	// The two dates are equal
	compareResult = 0
	return compareResult, err
}

// compareAreYearsAdjacent - Compares the astronomical year values
// encapsulated by input parameters 'dateTransDtoOne', and
// 'dateTransDtoTwo' to determine if the two year values are
// adjacent. Adjacent years are defined here as differing by a value
// of plus or minus one year. If the year values are 'Adjacent', the
// difference as calculated by subtracting the astronomical year
// value of 'dateTransDtoTwo.year' from 'dateTransDtoOne.year'
// (dateTransDtoOne.year - dateTransDtoTwo.year) will always equal
// to plus or minus 1 (+1 or -1).
//
// The 'Adjacent Year' comparison implemented by calculating the
// difference in astronomical year values will always have one of
// three possible outcomes (dateTransDtoOne.year - dateTransDtoTwo.year):
//
//  1. If 'dateTransDtoOne.year' IS ADJACENT to 'dateTransDtoTwo.year',
//     and the difference as calculated by 'dateTransDtoOne.year -
//     dateTransDtoTwo.year' is plus one (+1), it signals that
//     'dateTransDtoOne.year' is GREATER than 'dateTransDtoTwo.year'
//     by a value of '+1'. For this case, the returned boolean value
//     of 'areYearsAdjacent' is set to 'true' and the returned integer
//     value of 'compareResult' is set to +1.
//
//  2. If 'dateTransDtoOne.year' IS NOT ADJACENT to 'dateTransDtoTwo.year',
//     the calculation result of 'dateTransDtoOne.year - dateTransDtoTwo.year'
//     will be GREATER THAN plus one (+1) or minus one (-1). In this
//     case, the returned boolean value of 'areYearsAdjacent' will be set
//     to 'false' and the returned integer value of 'compareResult' will
//     be set to zero ('0').
//
//  3. If 'dateTransDtoOne.year' IS ADJACENT to 'dateTransDtoTwo.year',
//     and the difference as calculated by 'dateTransDtoOne.year -
//     dateTransDtoTwo.year' is minus one (-1), it signals that
//     'dateTransDtoOne.year' is LESS THAN 'dateTransDtoTwo.year' by a
//     value of minus one (-1). For this case, the returned boolean
//     value of 'areYearsAdjacent' is set to 'true' and the returned
//     integer value of 'compareResult' is set to -1.
//
// If either of input parameters 'dateTransDtoOne' or 'dateTransDtoTwo'
// is judged to be invalid, an error condition will be triggered.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTransDtoOne          *DateTransferDto
//     - A pointer to an instance of DateTransferDto. The year value
//       encapsulated by this instance will be compared to the year
//       value contained in input parameter, 'dateTransDtoTwo' in
//       order to determine whether the two years are adjacent years
//       as defined in the narrative above.
//
//       If 'dateTransDtoOne' is an invalid instance, an error will
//       be returned.
//
//
//  dateTransDtoTwo          *DateTransferDto
//     - A pointer to an instance of DateTransferDto. The year value
//       encapsulated by this instance will be compared to the year
//       value contained in input parameter, 'dateTransDtoOne' in
//       order to determine whether the two years are adjacent years
//       as defined in the narrative above.
//
//       If 'dateTransDtoTwo' is an invalid instance, an error will
//       be returned.
//
//
//  ePrefix                  string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  areYearsAdjacent         bool
//     - This boolean return value will signal whether the year values
//       contained in input parameters 'dateTransDtoOne' and
//       'dateTransDtoTwo' are adjacent years. Adjacent years differ by
//       a value of plus or minus 1-year.
//
//       If the 'dateTransDtoOne' and 'dateTransDtoTwo' are adjacent
//       years this boolean value is set to 'true'. Otherwise, it is set
//       to 'false'.
//
//
//  compareResult            int
//     - This integer return value will specify the comparison result
//       obtained from determining whether input parameters 'dateTransDtoOne'
//       and 'dateTransDtoTwo' are adjacent years. As such, this return
//       parameter will be populated with one of three possible values.
//
//       1. If 'dateTransDtoOne' IS ADJACENT to 'dateTransDtoTwo', and the
//          difference as calculated by 'dateTransDtoOne.year -
//          dateTransDtoTwo.year' is plus one (+1), it signals that
//          'dateTransDtoOne.year' is GREATER than 'dateTransDtoTwo.year'
//          by a value of '+1'. For this case, the returned integer value
//          of 'compareResult' is set to +1.
//
//       2. If 'dateTransDtoOne.year' IS NOT ADJACENT to 'dateTransDtoTwo.year',
//          the calculation result of 'dateTransDtoOne.year - dateTransDtoTwo.year'
//          will be GREATER THAN plus one (+1) or minus one (-1). In this
//          case the returned integer value of 'compareResult' will be set
//          to zero ('0').
//
//       3. If 'dateTransDtoOne.year' IS ADJACENT to 'dateTransDtoTwo.year',
//          and the difference as calculated by 'dateTransDtoOne.year -
//          'dateTransDtoTwo.year' is minus one (-1), it signals that
//          'dateTransDtoOne.year' is LESS THAN 'dateTransDtoTwo.year' by a
//          value of minus one (-1). For this case, the returned integer
//          value of 'compareResult' is set to -1.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func(dateTransUtil *dateTransferDtoUtility) compareAreYearsAdjacent(
	dateTransDtoOne *DateTransferDto,
	dateTransDtoTwo *DateTransferDto,
	ePrefix string) (
	areYearsAdjacent bool,
	compareResult int,
	err error) {

	if dateTransUtil.lock == nil {
		dateTransUtil.lock = new(sync.Mutex)
	}

	dateTransUtil.lock.Lock()

	defer dateTransUtil.lock.Unlock()

	ePrefix += "dateTransferDtoUtility.compareAreYearsAdjacent() "

	areYearsAdjacent = false
	compareResult = math.MinInt32
	err = nil

	if dateTransDtoOne == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'dateTransDtoOne' is INVALID!\n" +
			"dateTransDtoOne=nil!")
		return areYearsAdjacent,compareResult, err
	}

	if dateTransDtoOne.lock == nil {
		dateTransDtoOne.lock = new(sync.Mutex)
	}

	if dateTransDtoTwo == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'dateTransDtoTwo' is INVALID!\n" +
			"dateTransDtoTwo=nil!")
		return areYearsAdjacent,compareResult, err
	}

	if dateTransDtoTwo.lock == nil {
		dateTransDtoTwo.lock = new(sync.Mutex)
	}

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	_,
	err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDtoOne,
		ePrefix + "Testing validity of dateTransDtoOne. ")

	if err != nil {
		return areYearsAdjacent,compareResult, err
	}

	_,
	err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDtoTwo,
		ePrefix + "Testing validity of dateTransDtoTwo. ")

	if err != nil {
		return areYearsAdjacent,compareResult, err
	}

	calMech := calendarMechanics{}

	areYearsAdjacent,
	compareResult = calMech.areAstronomicalYearsAdjacent(
		dateTransDtoOne.astronomicalYear,
		dateTransDtoTwo.astronomicalYear)

	return areYearsAdjacent,compareResult, err
}


// CompareYears - Compares the astronomical year values for the current
// DateTransferDto instance and another DateTransferDto instance passed
// as an input parameter ('dateTransDto2').
//
// The comparison result is returned as an integer value signaling whether
// DateTransferDto instance astronomical year value is GREATER THAN, EQUAL TO,
// or LESS THAN that of input parameter 'dateTransDto2'. That means that one
// of three possible comparison results will be returned by this method:
//
//      Return                Comparison
//      Value                   Result
//
//       -1 - The current DateTransferDto instance astronomical year value is
//            LESS THAN the 'timeTransDto2' astronomical year value.
//
//        0 - The current DateTransferDto instance astronomical year value is
//            EQUAL to the 'timeTransDto2' astronomical year value.
//
//       +1 - The current DateTransferDto instance astronomical year value is
//            GREATER THAN the 'timeTransDto2' astronomical year value.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTransDtoOne          *DateTransferDto
//     - A pointer to an instance of DateTransferDto. The astronomical
//       year value encapsulated by this instance will be compared to
//       the astronomical year value contained in input parameter,
//       'dateTransDtoTwo'.
//
//       If 'dateTransDtoOne' is an invalid instance, an error will
//       be returned.
//
//
//  dateTransDtoTwo          *DateTransferDto
//     - A pointer to an instance of DateTransferDto. The astronomical
//       year value encapsulated by this instance will be compared to
//       the astronomical year value contained in input parameter,
//       'dateTransDtoOne'.
//
//       If 'dateTransDtoTwo' is an invalid instance, an error will
//       be returned.
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
//  comparisonResult         int
//     - This method will return will return an integer value describing
//       the result of comparing the astronomical year values for the
//       input parameters 'dateTransDtoOne' and 'dateTransDto2'.
//
//       Comparison Result
//       One of three possible comparison results will be returned.
//
//      Return                Comparison
//      Value                   Result
//
//       -1 - Input parameter 'dateTransDtoOne' astronomical year
//            value is LESS THAN the 'timeTransDto2' astronomical year
//            value.
//
//        0 - Input parameter 'dateTransDtoOne' astronomical year value is
//            EQUAL to the 'timeTransDto2' astronomical year value.
//
//       +1 - Input parameter 'dateTransDtoOne' astronomical year value is
//            GREATER THAN the 'timeTransDto2' astronomical year value.
//
//
//  err                      error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func(dateTransUtil *dateTransferDtoUtility) compareYears(
	dateTransDtoOne *DateTransferDto,
	dateTransDtoTwo *DateTransferDto,
	ePrefix string) (
	comparisonResult int,
	err error) {

	if dateTransUtil.lock == nil {
		dateTransUtil.lock = new(sync.Mutex)
	}

	dateTransUtil.lock.Lock()

	defer dateTransUtil.lock.Unlock()

	ePrefix += "dateTransferDtoUtility.compareAreYearsAdjacent() "
	comparisonResult = math.MinInt32
	err = nil

	if dateTransDtoOne == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'dateTransDtoOne' is INVALID!\n" +
			"dateTransDtoOne=nil!")
		return comparisonResult, err
	}

	if dateTransDtoOne.lock == nil {
		dateTransDtoOne.lock = new(sync.Mutex)
	}

	if dateTransDtoTwo == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'dateTransDtoTwo' is INVALID!\n" +
			"dateTransDtoTwo=nil!")
		return comparisonResult, err
	}

	if dateTransDtoTwo.lock == nil {
		dateTransDtoTwo.lock = new(sync.Mutex)
	}

	dateTransDtoNanobot := dateTransferDtoNanobot{}
	_,
		err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDtoOne,
		ePrefix + "Testing validity of dateTransDtoOne. ")

	if err != nil {
		return comparisonResult, err
	}

	_,
		err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDtoTwo,
		ePrefix + "Testing validity of dateTransDtoTwo. ")

	if err != nil {
		return comparisonResult, err
	}

	if dateTransDtoOne.astronomicalYear ==
			dateTransDtoTwo.astronomicalYear {

		comparisonResult = 0

	} else if dateTransDtoOne.astronomicalYear >
		dateTransDtoTwo.astronomicalYear {

		comparisonResult = 1

	} else {

		// dateTransDtoOne.astronomicalYear <
		//		dateTransDtoTwo.astronomicalYear

		comparisonResult = -1

	}

	return comparisonResult, err
}

// getDaysInYear - Returns the number of days contained in
// in the year component of the date value encapsulated in
// the input parameter 'dateTransDto'.
//
// Be advised that this method will validate the input
// parameter 'dateTransDto'. If it is judged to be invalid
// this method will return an error and set the return
// parameter 'daysInYear' to minus one (-1).
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTransDto        *DateTransferDto
//     - A pointer to an instance of DateTransferDto.
//
//       This method will validate this input parameter and
//       return an error if it is found to be invalid.
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
// daysInYear           int
//     - If this method completes successfully, this return parameter will
//       be populated with the number of days in the year for the year
//       value encapsulated in input parameter 'dateTransDto'. If an error
//       is encountered, this value will be set to minus one (-1).
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func(dateTransUtil *dateTransferDtoUtility) getDaysInYear(
	dateTransDto *DateTransferDto,
	ePrefix            string) (
	daysInYear int,
	err error) {

	if dateTransUtil.lock == nil {
		dateTransUtil.lock = new(sync.Mutex)
	}

	dateTransUtil.lock.Lock()

	defer dateTransUtil.lock.Unlock()

	ePrefix += "dateTransferDtoUtility.getDaysInYear() "

	err = nil

	daysInYear = -1

	if dateTransDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'dateTransDto' is INVALID!\n" +
			"dateTransDto=nil!")
		return daysInYear, err
	}

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	_,
		err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDto,
		ePrefix + "Testing validity of 'dateTransDto'. ")

	if err != nil {
		return daysInYear, err
	}

	daysInYear,
	err =  dateTransDto.calendarBaseData.GetDaysInYear(
		dateTransDto.astronomicalYear,
		CalendarYearNumType(0).Astronomical(),
		ePrefix)

	return daysInYear, err
}

// setDateTransferDto - Populates the internal member variables of a
// DateTransferDto instance initialized from input parameters.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dateTransDto        *DateTransferDto
//     - A pointer to an instance of DateTransferDto. This method will
//       set the data values of 'dateTransDto' member variables based
//       on the following input parameters.
//
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
//     - The month number
//
//
//  day                 int
//    - The day number. This is the day number within the
//      the 'month' identified in the 'month' input parameter,
//      above.
//
//
//  hasLeapSecond       bool
//     - If this parameter is set to 'true', it signals that the day identified
//       by year, month and day input parameters contains a leap second.
//       This parameter is rarely used and is almost always set to 'false'.
//
//       A leap second is a one-second adjustment that is occasionally applied
//       to Coordinated Universal Time (UTC) in order to accommodate the
//       difference between precise time (as measured by atomic clocks) and
//       imprecise observed solar time (known as UT1 and which varies due to
//       irregularities and long-term slowdown in the Earth's rotation). If
//       this return parameter is set to 'true',      the time calculation
//       will assume the duration of the relevant 'day' is 24-hours plus one
//       second. Otherwise, the duration of a day is assumed to consist of
//       exactly 24-hours. For more information on the 'leap second',
//       reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
//  dayOfWeekNoSysType  DayOfWeekNumberingSystemType
//     -Set this parameter to one of two possible values:
//            DayOfWeekNumberingSystemType(0).UsDayOfWeek
//                         OR
//            DayOfWeekNumberingSystemType(0).ISO8601DayOfWeek
//
//       If this parameter is NOT set to one of the two values shown
//       above, it will be defaulted to:
//             DayOfWeekNumberingSystemType(0).UsDayOfWeek
//
//       Currently, the datetime software library supports two types of
//       Day of the Week Numbering Systems:
//
//         1. US Day Of The Week Numbering System.
//             DayOfWeekNumberingSystemType(0).UsDayOfWeek
//
//         2. ISO 8601 Standard Day of the Week Numbering System.
//             DayOfWeekNumberingSystemType(0).ISO8601DayOfWeek
//
//       US Day Of The Week Numbering System
//
//       The United States, Canada, Australia and New Zealand put Sunday as the
//       first day of the week on their calendars. The first day of the week,
//       'Sunday', is numbered as zero with the last day of the week being numbered
//       as 6 for 'Saturday'. US Day of the Week Numbers are listed below:
//
//                    Week         Week Day
//                    Day           Number
//                  ========     ===========
//                   Sunday          = 0
//                   Monday          = 1
//                   Tuesday         = 2
//                   Wednesday       = 3
//                   Thursday        = 4
//                   Friday          = 5
//                   Saturday        = 6
//
//       For more information on the US Day of the Week Numbering System,
//       reference type UsDayOfWeekNo, Source Code File:
//          datetime/dayofweeknumberusenum.go
//
//       ISO 8601 Standard Day Of The Week Numbering System
//
//       The most common day of the week numbering system used internationally,
//       is the ISO 8601 standard which specifies that the week begins on Monday.
//       ISO stands for the International Organization for Standardization (ISO).
//          https://www.iso.org/home.html
//
//       ISO 8601 standard is used in Western Europe, Scandinavia, and most of
//       Eastern Europe as well as many other nations across the globe. It is
//       considered an international standard.
//
//       Under the ISO 8601 Standard, days of the week are numbered beginning with
//       one (1) for Monday and ending with seven (7) for Sunday. ISO 8601 week day
//       numbers are listed as follows:
//
//                    Week         Week Day
//                    Day           Number
//                  ========     ===========
//                   Monday          = 1
//                   Tuesday         = 2
//                   Wednesday       = 3
//                   Thursday        = 4
//                   Friday          = 5
//                   Saturday        = 6
//                   Sunday          = 7
//
//       For more information on the ISO 8601 Standard, reference type,
//       ISO8601DayOfWeekNo defined in source file:
//          datetime/dayofweeknumberiso8601enum.go
//
//
//  calendarSystem  CalendarSpec
//     - An enumeration type designating the Calendar System
//       associated with the generated date. Reference:
//       Source File: datetime\calendarspecenum.go
//
//       Possible Calendar System values include:
//       Gregorian, Julian, Revised Julian, or Revised Goucher-Parker.
//
//       Possible Enumeration Values:
//         CalendarSpec(0).Gregorian()
//         CalendarSpec(0).Julian()
//         CalendarSpec(0).RevisedJulian()
//         CalendarSpec(0).RevisedGoucherParker()
//
//
//  tag                 string
//     - A string description to be associated with the newly created DateTransferDto instance
//       generated by this method.
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
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func(dateTransUtil *dateTransferDtoUtility) setDateTransferDto(
	dateTransDto *DateTransferDto,
	year int64,
	yearType CalendarYearNumType,
	month int,
	day int,
	hasLeapSecond bool,
	calendarSystem CalendarSpec,
	tag string,
	ePrefix string) (err error) {

	if dateTransUtil.lock == nil {
		dateTransUtil.lock = new(sync.Mutex)
	}

	dateTransUtil.lock.Lock()

	defer dateTransUtil.lock.Unlock()

	ePrefix += "dateTransferDtoUtility.setDateTransferDto() "

	err = nil

	if dateTransDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'dateTransDto' is INVALID!\n" +
			"dateTransDto=nil!")
		return err
	}

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}


	if !yearType.XIsValid() {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "yearType",
			inputParameterValue: strconv.Itoa(yearType.XValueInt()),
			errMsg:              "Input parameter 'yearType' is INVALID!",
			err:                 nil,
		}

		return err
	}

	if !calendarSystem.XIsValid() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter, 'calendarSystem' is INVALID!\n" +
			"calendarSystem='%v'\n",
			calendarSystem.XValueInt())

		return err
	}

	var calYearNumMode CalendarYearNumMode

	calYearNumMode = yearType.XGetCalendarYearNumberMode()

	calMech := calendarMechanics{}

	var astronomicalYear int64

	astronomicalYear,
	err = calMech.convertAnyYearToAstronomicalYear(
		year,
		yearType,
		ePrefix)

	if err != nil {
		return err
	}

	var calendarBaseData ICalendarBaseData

	switch calendarSystem {

	case CalendarSpec(0).Gregorian():
		calendarBaseData = &CalendarGregorianBaseData{}

	default:
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'calendarSystem' is INVALID!\n" +
			"Currently, the Gregorian Calendar System is the only supported calendar.\n" +
			"calendarSystem='%v'\n",
			calendarSystem.XValueInt())
		return err
	}

	var isLeapYear bool

	isLeapYear,
	err =
		calendarBaseData.IsLeapYear(
			astronomicalYear,
			CalendarYearNumType(0).Astronomical(),
			ePrefix)

	if err != nil {
		return err
	}

	var yearMonthDays map[int] int

	if isLeapYear {
		yearMonthDays = calendarBaseData.GetLeapYearMonthDays()
	} else {
		yearMonthDays = calendarBaseData.GetStandardYearMonthDays()
	}

	if month < 1 || month > len(yearMonthDays) {

		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'month' is INVALID!\n" +
			"The minimum month number is '1' and the\n" +
			"maximum number of months in this year is '%v'\n" +
			"month='%v'\n",
			len(yearMonthDays),
			month)

		return err
	}

	maxDaysInMonth := yearMonthDays[month]

	if day < 0 || day > maxDaysInMonth {

		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'day' is INVALID!\n" +
			"The minimum day number is '1' and the\n" +
			"maximum number of days in this month is '%v'\n" +
			"month='%v'\n" +
			"day='%v'",
			len(yearMonthDays),
			month,
			day)

		return err
	}

	dateTransDto.astronomicalYear = astronomicalYear
	dateTransDto.yearNumberingMode = calYearNumMode
	dateTransDto.yearNumType = yearType
	dateTransDto.month = month
	dateTransDto.day = day
	dateTransDto.hasLeapSecond = hasLeapSecond
	dateTransDto.calendarBaseData = calendarBaseData.New()
	dateTransDto.tag = tag

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	_, err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDto,
		ePrefix + "Test validity of updated 'dateTransDto'. ")

	return err
}