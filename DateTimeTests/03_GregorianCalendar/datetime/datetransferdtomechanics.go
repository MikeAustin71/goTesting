package datetime

import (
	"errors"
	"math"
	"sync"
)

type dateTransferDtoMechanics struct {
	lock    *sync.Mutex
}

// empty - Resets the passed instance of DateTransferDto to
// an invalid and uninitialized state. All member variables
// are set to invalid values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTransDto        *DateTransferDto
//     - A pointer to an instance of DateTransferDto. This method will
//       change the values of internal member variables to achieve
//       the method's objectives.
//
//       The 'empty' operation will set all internal member variables
//       to invalid data values.
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
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (dateTransDtoMech *dateTransferDtoMechanics) empty(
	dateTransDto *DateTransferDto,
	ePrefix            string) (err error) {

	if dateTransDtoMech.lock == nil {
		dateTransDtoMech.lock = new(sync.Mutex)
	}

	dateTransDtoMech.lock.Lock()

	defer dateTransDtoMech.lock.Unlock()

	ePrefix += "dateTransferDtoMechanics.empty() "

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

	dateTransDto.astronomicalYear = math.MinInt64
	dateTransDto.yearNumberingMode = CalendarYearNumMode(0).None()
	dateTransDto.yearNumType = CalendarYearNumType(0).None()
	dateTransDto.month = -1
	dateTransDto.day = -1
	dateTransDto.hasLeapSecond = false
	dateTransDto.calendarBaseData = nil
	dateTransDto.tag = ""

	return err
}

// exchangeValues - Receives pointers to two DateTransferDto objects and
// proceeds to exchange the data values of all internal member variables.
// If the method completes successfully, 'dateTransDtoOne' will be populated
// with the original data values of 'dateTransDtoTwo' and 'dateTransDtoTwo'
// will be populated with the original values of 'dateTransDtoOne'.
//
//
// IMPORTANT
//
// This method does NOT validate the current DateTransferDto instance
// before returning the value. To run a validity check on the
// DateTransferDto instance first call one of the two following
// methods:
//
//  DateTransferDto.IsValidInstance() bool
//                OR
//  DateTransferDto.IsValidInstanceError(ePrefix string) error
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTransDtoOne          *DateTransferDto
//     - A pointer to an instance of DateTransferDto. This is one of the 
//       DateTransferDto objects used in the data exchange. Data values from
//       this instance will be copied to input parameter 'dateTransDtoTwo'
//       while the original values from 'dateTransDtoTwo' will be copied
//       to this instance, 'dateTransDtoOne'.
//
//
//  dateTransDtoTwo          *DateTransferDto
//     - A pointer to an instance of DateTransferDto. This is one of the 
//       DateTransferDto objects used in the data exchange. Data values from
//       this instance will be copied to input parameter 'dateTransDtoOne'
//       while the original values from 'dateTransDtoOne' will be copied
//       to this instance, 'dateTransDtoTwo'.
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
//  err                      error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (dateTransDtoMech *dateTransferDtoMechanics) exchangeValues(
	dateTransDtoOne *DateTransferDto,
	dateTransDtoTwo *DateTransferDto,
	ePrefix            string) (err error) {

	if dateTransDtoMech.lock == nil {
		dateTransDtoMech.lock = new(sync.Mutex)
	}

	dateTransDtoMech.lock.Lock()

	defer dateTransDtoMech.lock.Unlock()

	ePrefix += "dateTransferDtoMechanics.exchangeValues() "

	if dateTransDtoOne == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'dateTransDtoOne' is INVALID!\n" +
			"dateTransDtoOne=nil!")
		return err
	}

	if dateTransDtoOne.lock == nil {
		dateTransDtoOne.lock = new(sync.Mutex)
	}

	if dateTransDtoTwo == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'dateTransDtoTwo' is INVALID!\n" +
			"dateTransDtoTwo=nil!")
		return err
	}

	if dateTransDtoTwo.lock == nil {
		dateTransDtoTwo.lock = new(sync.Mutex)
	}

	dateTransDtoNanobot := dateTransferDtoNanobot{}

	_,
	err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDtoOne,
		ePrefix +
			"Testing validity of 'dateTransDtoOne'. ")

	if err != nil {
		return err
	}

	_,
	err = dateTransDtoNanobot.testDateTransferDtoValidity(
		dateTransDtoTwo,
		ePrefix +
			"Testing validity of 'dateTransDtoTwo'. ")

	if err != nil {
		return err
	}

	tempDateTransDto := DateTransferDto{}.NewZeroInstance()
	
	tempDateTransDto.astronomicalYear   = dateTransDtoOne.astronomicalYear
	tempDateTransDto.yearNumberingMode  = dateTransDtoOne.yearNumberingMode
	tempDateTransDto.yearNumType        = dateTransDtoOne.yearNumType
	tempDateTransDto.month              = dateTransDtoOne.month
	tempDateTransDto.day                = dateTransDtoOne.day
	tempDateTransDto.hasLeapSecond      = dateTransDtoOne.hasLeapSecond
	tempDateTransDto.calendarBaseData   = dateTransDtoOne.calendarBaseData.New()
	tempDateTransDto.tag                = dateTransDtoOne.tag
	
	dateTransDtoOne.astronomicalYear   = dateTransDtoTwo.astronomicalYear
	dateTransDtoOne.yearNumberingMode  = dateTransDtoTwo.yearNumberingMode
	dateTransDtoOne.yearNumType        = dateTransDtoTwo.yearNumType
	dateTransDtoOne.month              = dateTransDtoTwo.month
	dateTransDtoOne.day                = dateTransDtoTwo.day
	dateTransDtoOne.hasLeapSecond      = dateTransDtoTwo.hasLeapSecond
	dateTransDtoOne.calendarBaseData   = dateTransDtoTwo.calendarBaseData.New()
	dateTransDtoOne.tag                = dateTransDtoTwo.tag
	
	dateTransDtoTwo.astronomicalYear   = tempDateTransDto.astronomicalYear
	dateTransDtoTwo.yearNumberingMode  = tempDateTransDto.yearNumberingMode
	dateTransDtoTwo.yearNumType        = tempDateTransDto.yearNumType
	dateTransDtoTwo.month              = tempDateTransDto.month
	dateTransDtoTwo.day                = tempDateTransDto.day
	dateTransDtoTwo.hasLeapSecond      = tempDateTransDto.hasLeapSecond
	dateTransDtoTwo.calendarBaseData   = tempDateTransDto.calendarBaseData.New()
	dateTransDtoTwo.tag                = tempDateTransDto.tag

	return err
}


// getConvertedYearValueByType - This this method receives a 'yearValue'
// and a 'yearType' describing that 'yearValue' and proceeds to covert
// the 'yearValue' to another year numbering mode. Effectively, this method
// is a means of converting year values between year numbering systems.
//
// For a usage example, consider a year value of '47' and a year type of
// of BCE (Before Common Era or 47 BCE). If the target year number mode
// is 'Astronomical Year', then 47 BCE will be converted and returned as
// converted Year Value = -46 and converted Year Number Type = 'Astronomical'
// ('Astronomical' = Astronomical Year Numbering System).
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  yearValue                int64
//     - The year value to be converted.
//
//
//  yearType                 CalendarYearNumType
//     - The Calendar Year Numbering Type describing input parameter
//       'yearValue'.
//
//     - 'CalendarYearNumType' is an enumeration which will classify
//        input parameter 'yearValue' as one of three year numbering
//        types:
//         1. Astronomical Year
//         2. BCE - Before Common Era
//         3. CE  - Common Era
//       For more information on Astronomical and Common Era Year
//       Numbering, reference:
//           Source File: datetime\calendaryearnumbertypeenum.go
//           https://en.wikipedia.org/wiki/Astronomical_year_numbering
//           https://en.wikipedia.org/wiki/Common_Era
//
//
//  targetYearNumberingMode  CalendarYearNumMode
//     - The year numbering mode for the converted year value which will
//       be returned by this method.
//
//       'CalendarYearNumMode' is an enumeration specifying a year numbering
//       system as either Astronomical Year Numbering System or Common Era
//       Year Number System. Therefore, this parameter must be set to one
//       of two possible enumeration values:
//        1. CalendarYearNumMode(0).Astronomical()
//        2. CalendarYearNumMode(0).CommonEra()
//
//       For more information on Astronomical and Common Era Year Numbering
//       Systems, reference:
//          https://en.wikipedia.org/wiki/Astronomical_year_numbering
//          https://en.wikipedia.org/wiki/Common_Era
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
//  convertedYearValue       int64
//     - The converted year value which conforms to the year numbering mode
//       specification contained in input parameter, 'targetYearNumberingMode'.
//
//
//  convertedYearNumberType  CalendarYearNumType
//     - The year number type associated with the 'convertedYearValue'
//       identified above. 'convertedYearNumberType' classifies return parameter
//       'convertedYearValue' as one of three year types:
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
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//
func (dateTransDtoMech *dateTransferDtoMechanics) getConvertedYearValueByType(
	yearValue int64,
	yearType CalendarYearNumType,
	targetYearNumberingMode CalendarYearNumMode,
	ePrefix string) (
	convertedYearValue int64,
	convertedYearNumberType CalendarYearNumType,
	err error) {

	if dateTransDtoMech.lock == nil {
		dateTransDtoMech.lock = new(sync.Mutex)
	}

	dateTransDtoMech.lock.Lock()

	defer dateTransDtoMech.lock.Unlock()

	ePrefix += "dateTransferDtoUtility.getConvertedYearValueByType() "

	calMech := calendarMechanics{}

	var astronomicalYearValue int64

	convertedYearValue = math.MinInt64
	convertedYearNumberType = CalendarYearNumType(0).None()
	err = nil

	astronomicalYearValue,
		err = calMech.convertAnyYearToAstronomicalYear(
		yearValue,
		yearType,
		ePrefix)

	if err != nil {
		return convertedYearValue,
			convertedYearNumberType,
			err
	}

	convertedYearValue,
		convertedYearNumberType,
		err = calMech.getCalendarYearByType(
		astronomicalYearValue,
		targetYearNumberingMode,
		ePrefix)

	return convertedYearValue,
		convertedYearNumberType,
		err
}

// setInstanceToZero - This method receives a pointer to a DateTransferDto
// instance and proceeds to set all of the member variable data values to
// their 'zero' state. When this operation is completed some, but not all,
// of the member variable data values will be 'invalid'. If after calling
// this method, the DateTransferDto instance is submitted for validity
// testing, it will fail those tests.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTransDto        *DateTransferDto
//     - A pointer to an instance of DateTransferDto. This method will
//       set all of the internal member variables to their native 'zero'
//       states. Thereafter, this instance will fail any subsequent
//       validity tests.
//
func (dateTransDtoMech *dateTransferDtoMechanics) setInstanceToZero(
	dateTransDto *DateTransferDto) {

	if dateTransDtoMech.lock == nil {
		dateTransDtoMech.lock = new(sync.Mutex)
	}

	dateTransDtoMech.lock.Lock()

	defer dateTransDtoMech.lock.Unlock()

	if dateTransDto == nil {
		panic("dateTransferDtoMechanics.setInstanceToZero()\n" +
			"Input parameter 'dateTransDto' is a nil pointer.")
	}

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	dateTransDto.astronomicalYear = 0
	dateTransDto.yearNumberingMode = CalendarYearNumMode(0).None()
	dateTransDto.yearNumType = CalendarYearNumType(0).None()
	dateTransDto.month = 0
	dateTransDto.day = 0
	dateTransDto.hasLeapSecond = false
	dateTransDto.calendarBaseData = nil
	dateTransDto.tag = ""

}

