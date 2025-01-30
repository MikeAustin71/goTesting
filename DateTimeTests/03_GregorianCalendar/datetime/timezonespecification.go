package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// TimeZoneSpecification - Internal data structure used to
// store the Time Zone data elements used to describe a single
// time zone.
type TimeZoneSpecification struct {
	zoneLabel              string           // Identifying Time Zone Label. A tag or description field.
	referenceDateTime      time.Time        // The date time used in defining the Time Zone
	zoneName               string           // The Time Zone abbreviation. Examples: 'EST', 'CST', 'PST'
	zoneOffsetTotalSeconds int              // Signed number of seconds offset from UTC.
	//                                         + == East of UTC; - == West of UTC
	zoneSignValue          int              // -1 == West of UTC  +1 == East of UTC. Apply this sign
	//                                         to the offset hours, minutes and seconds value
	offsetHours            int              // Normalized Offset Hours from UTC. Always a positive number,
	//                                         refer to ZoneSign for correct sign value.
	offsetMinutes          int              // Normalized Offset Minutes offset from UTC. Always a positive number,
	//                                         refer to ZoneSign for the correct sign value
	offsetSeconds          int              // Normalized Offset Seconds offset from UTC. Always a positive number,
	//                                         refer to ZoneSign for the correct sign value
	zoneOffset             string           // A text string representing offset from UTC for this time zone.
	//                                         Example "-0600 CST" or "+0200 EET"
	zoneAbbrvLookupId      string           // A string representing the Abbreviation Id used in map lookups.
	//                                         Examples: "CST-0600", "EET+0200"
	utcOffset              string           // A text string representing the offset from UTC for this Time Zone.
	//                                         Examples: "-0600" or "+0200"
	locationPtr            *time.Location   // Pointer to a Time Zone Location
	locationName           string           // Time Zone Location Name Examples: "Local", "America/Chicago", "America/New_York"
	locationNameType       LocationNameType // Three possible values:
	//                                           None()
	//                                           NonConvertibleTimeZone()
	//                                           ConvertibleTimeZone()
	militaryTimeZoneName   string           // Full Military Time Zone text name. Examples: "Alpha", "Bravo", "Charlie", "Zulu"
	militaryTimeZoneLetter string           // Single Alphabetic Character identifying a Military Time Zone.
	tagDescription         string           // Unused - Available for classification, labeling or description by user.
	timeZoneCategory       TimeZoneCategory // Enumeration of Time Zone Category:
	//                                          TzCat.None()
	//                                          TzCat.TextName()
	//                                          TzCat.UtcOffset()
	timeZoneClass          TimeZoneClass    // Enumeration of Time Zone Class:
	//                                          TzClass.None()
	//                                          TzClass.AlternateTimeZone()
	//                                          TzClass.OriginalTimeZone()
	timeZoneType           TimeZoneType     // Enumeration of Time Zone Type:
	//                                          TzType.None()
	//                                          TzType.Iana()
	//                                          TzType.Military()
	//                                          TzType.Local()
	//                                          TzType.UtcOffset()
	timeZoneUtcOffsetStatus TimeZoneUtcOffsetStatus // Enumeration of Time Zone UTC Offset Status:
	//                                                  TzUtcStatus.None()
	//                                                  TzUtcStatus.Static()
	//                                                  TzUtcStatus.Variable()
	//
	lock *sync.Mutex       // Used for implementing thread safe operations.
}

// CopyIn - Copies the values of input parameter 'tzSpec2'
// to all of the data fields in the current instance of 
// TimeZoneSpecification (tzSpec). When completed 'tzSpec' will
// have data field values identical to those of 'tzSpec2'
//
func (tzSpec *TimeZoneSpecification) CopyIn(tzSpec2 TimeZoneSpecification) {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	tzSpecUtil := timeZoneSpecUtility{}

	tzSpecUtil.copyIn(tzSpec, &tzSpec2)

}


// CopyOut - Returns a deep copy of the current Time Zone 
// Specification object as a new instance of 'TimeZoneSpecification'.
//
func (tzSpec *TimeZoneSpecification) CopyOut() TimeZoneSpecification {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	tzSpecUtil := timeZoneSpecUtility{}

	return tzSpecUtil.copyOut(tzSpec)
}

// Empty - Sets all the values of the data fields in the
// current TimeZoneSpecification to their empty or zero values.
//
func (tzSpec *TimeZoneSpecification) Empty() {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	tzSpecUtil := timeZoneSpecUtility{}

	tzSpecUtil.empty(tzSpec)

}

// Equal - Returns a boolean value of true if both the current instance
// of TimeZoneSpecification and the input parameter TimeZoneSpecification are
// equivalent in all respects.
//
// Exceptions: Note that the following private member data fields
// are NOT checked for equivalency.
//
// zone label is NOT checked for equivalency
// tagDescription is NOT checked for equivalency
//
func (tzSpec *TimeZoneSpecification) Equal( tzSpec2 TimeZoneSpecification) bool {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	tzSpecUtil := timeZoneSpecUtility{}

	return tzSpecUtil.equal(tzSpec, tzSpec2)
}

// IsEmpty() returns a boolean value of 'true' if all
// data field values are set to their empty or zero
// values.
func (tzSpec *TimeZoneSpecification) IsEmpty() bool {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	ePrefix := "TimeZoneSpecification.IsEmpty()"

	tzSpecUtil := timeZoneSpecUtility{}

	return tzSpecUtil.isEmpty(tzSpec, ePrefix)
}

// IsValid - Examines the data fields of the current
// TimeZoneSpecification instance are valid.
//
func (tzSpec *TimeZoneSpecification) IsValid(ePrefix string) error {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	ePrefix += "TimeZoneSpecification.IsValidInstanceError() "

	if strings.TrimLeft(strings.TrimRight(tzSpec.locationName, " "), " ") == "" {
		return errors.New(ePrefix +
			"\nError: locationName is an empty string!\n")
	}

	if tzSpec.locationPtr == nil {
		return errors.New(ePrefix +
			"\nError: Location Pointer is 'nil'!\n")
	}

	if tzSpec.locationPtr.String() != tzSpec.locationName {
		return fmt.Errorf(ePrefix +
			"\nError: The Location Pointer is NOT equal to the Location Name!\n" +
			"Location Pointer String='%v'\n" +
			"Location Name = '%v'\n",
			tzSpec.locationPtr.String() , tzSpec.locationName)
	}

	dtMech := DTimeNanobot{}

	locPtr, err := dtMech.LoadTzLocation(tzSpec.locationName, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError: Location Name is NOT a valid time zone!\n"+
			"tzdef.locationName='%v'\n"+
			"Returned Error='%v'\n", tzSpec.locationName, err.Error())
	}

	if locPtr.String() != tzSpec.locationName {
		return fmt.Errorf(ePrefix +
			"\nError: LoadLocation Pointer string NOT equal to tzSpec.locationName !\n" +
			"tzSpec.locationName='%v'\n" +
			"loc.String()='%v'\n", tzSpec.locationName, locPtr.String())
	}

	if tzSpec.timeZoneType == TzType.Military() &&
		(tzSpec.militaryTimeZoneLetter == "" ||
			tzSpec.militaryTimeZoneName == "") {
		return fmt.Errorf(ePrefix +
			"\nError: This time zone is classified as a 'Military' Time Zone.\n" +
			"However, one or both of the Military Time Zone name strings are empty.\n" +
			"tzSpec.militaryTimeZoneLetter='%v'\n" +
			"tzSpec.militaryTimeZoneName='%v'\n",
			tzSpec.militaryTimeZoneLetter , tzSpec.militaryTimeZoneName)
	}

	return nil
}

// GetLocationPointer - Returns the time zone location in the form of
// a pointer to 'time.Location'.
//
func (tzSpec *TimeZoneSpecification) GetLocationPointer() *time.Location {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.locationPtr
}

// GetOriginalLocationName - Returns the time zone name or time zone location.
// Examples: "Local", "America/Chicago", "America/New_York"
//
func (tzSpec *TimeZoneSpecification) GetLocationName() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.locationName
}

// GetLocationNameType - Describes and classifies the Time Zone
// Location. The return value is a LocationNameType value which
// is an enumeration time zone location name classifications.
//
// Possible return values:
//
//  LocationNameType(0).None()
//                               - The Time Zone is uninitialized. This
//                                 is an error condition.
//
//  LocationNameType(0).NonConvertibleTimeZone()
//                               - The Time Zone Location Name cannot
//                                 be converted to other time zones.
//
//  LocationNameType(0).ConvertibleTimeZone()
//                               - The Time Zone Name is a complete
//                                 and valid time zone name which is
//                                 convertible across all other
//                                 time zones.
//
// For easy access to these enumeration values, use the global variable,
// 'LocNameType'. Example: LocNameType.ConvertibleTimeZone()
//
func (tzSpec *TimeZoneSpecification) GetLocationNameType() LocationNameType {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.locationNameType
}

// GetMilitaryOrStdTimeZoneName - If the Time Zone Specification object
// is configured as a 'Military' Time Zone, the Military Time Zone Name
// is returned. If the Time Zone Specification object is configured as a
// 'Local' time zone, the time zone name 'Local' is returned. Finally,
// if the Time Zone Specification object is configured as an 'IANA' time
// zone, the appropriate 'IANA' time zone name is returned.
//
func (tzSpec *TimeZoneSpecification) GetMilitaryOrStdTimeZoneName() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	if tzSpec.timeZoneType == TzType.Military() {
		return tzSpec.militaryTimeZoneName
	}

	return tzSpec.locationName
}

// GetMilitaryTimeZoneName - Returns a string containing the military
// time zone name, if applicable. If the current TimeZoneSpecification
// instance does not define a military time zone, this return value
// is an empty string.
//
func (tzSpec *TimeZoneSpecification) GetMilitaryTimeZoneName() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.militaryTimeZoneName
}

// GetMilitaryTimeZoneLetter - Returns a string containing the military
// time zone letter or abbreviation. If the current TimeZoneSpecification
// instance does not define a military time zone, this return value
// is an empty string.
//
func (tzSpec *TimeZoneSpecification) GetMilitaryTimeZoneLetter() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.militaryTimeZoneLetter
}

// GetOffsetElements - Returns a series of string and integer
// values which taken collectively identify the offset from
// UTC for this time zone.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  offsetSignChar string - Like return value offsetSignValue, this string
//                          value signals whether the offset from UTC is West
//                          or East of UTC. This string will always have one of
//                          two values: "+" or "-". The plus sign ("+") signals
//                          that the offset is East of UTC. The minus sign ("-")
//                          signals that the offset is West of UTC.
//
//  offsetSignValue int  - Similar to return value 'offsetSignChar' above except
//                         that sign values are expressed as either a '-1' or positive
//                         '1' integer value. -1 == West of UTC  +1 == East of UTC.
//                         Apply this sign value to the offset hours, minutes and
//                         seconds value returned below.
//
//  offsetHours     int  - Normalized Offset Hours from UTC. Always a positive number,
//                         refer to ZoneSign for correct sign value.
//
//  offsetMinutes   int  - Normalized Offset Minutes offset from UTC. Always a
//                         positive number, refer to ZoneSign for the correct
//                         sign value.
//
//  offsetSeconds   int  - Normalized Offset Seconds offset from UTC. Always a
//                         positive number, refer to ZoneSign for the correct
//                         sign value.
//
func (tzSpec *TimeZoneSpecification) GetOffsetElements() (
	offsetSignChar string,
	offsetSignValue,
	offsetHours,
	offsetMinutes,
	offsetSeconds int) {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	if tzSpec.zoneSignValue < 0 {
		offsetSignChar = "-"
	} else {
		offsetSignChar = "+"
	}

	offsetSignValue = tzSpec.zoneSignValue
	offsetHours = tzSpec.offsetHours
	offsetMinutes = tzSpec.offsetMinutes
	offsetSeconds = tzSpec.offsetSeconds

	return offsetSignChar,
		offsetSignValue,
		offsetHours,
		offsetMinutes,
		offsetSeconds
}

// GetReferenceDateTime - Returns the reference Date Time
//
func (tzSpec *TimeZoneSpecification) GetReferenceDateTime() time.Time {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.referenceDateTime
}

// GetOriginalTagDescription - Returns the private member variable
// "tagDescription". This field is available for users to
// tag, classify or otherwise attach descriptive information
// to this TimeZoneSpecification instance.
//
func (tzSpec *TimeZoneSpecification) GetTagDescription() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.tagDescription
}

// GetTimeZoneAbbreviation - Returns the Time Zone abbreviation. The
// term 'Time Zone Abbreviation' is a synonym for 'Zone Name'.
//
// Examples of Time Zone Abbreviations are:
//  'EST', 'CST', 'PST', 'EDT', 'CDT', 'PDT'
//
func (tzSpec *TimeZoneSpecification) GetTimeZoneAbbreviation() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneName

}

// GetTimeZoneCategory - Returns the Time Zone Category description .
// Time Zone Category is an enumeration identifying the time zone by
// category of time zone name.
//
// Possible Return Values:
//
//  TimeZoneCategory(0).None()       -  Signals that Time Zone Category is uninitialized.
//                                      This represents an error condition.
//
//  TimeZoneCategory(0).TextName()   -  Signals that the Time Zone is identified
//                                      by a standard IANA Text Name. Examples:
//                                        "America/Chicago"
//                                        "Asia/Amman"
//                                        "Atlantic/Bermuda"
//                                        "Australia/Sydney"
//                                        "Europe/Rome"
//
//  TimeZoneCategory(0).UtcOffset()   -  Signals that the Time Zone is identified
//                                       by a valid UTC Offset and has no associated
//                                       text name. Examples:
//                                         "+07"
//                                         "+10"
//
// For easy access to these enumeration values, use the global variable
// 'TzCat'. Example: TzCat.None()
//
func (tzSpec *TimeZoneSpecification) GetTimeZoneCategory() TimeZoneCategory {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.timeZoneCategory
}

// GetTimeZoneClass - Returns the Time Zone Class description.
// Time Zone Class is an enumeration identifying the time zone
// status.
//
// Possible Return Values:
//
// TimeZoneClass(0).None()              - An Error Condition
//
// TimeZoneClass(0).AlternateTimeZone() - Generated Time Zone from
//                                        Time Zone Abbreviation
//
// TimeZoneClass(0).OriginalTimeZone()  - Original Valid Time Zone
//
// For easy access to these enumeration values, use the global variable
// 'TzClass'. Example: TzClass.AlternateTimeZone()
//
func (tzSpec *TimeZoneSpecification) GetTimeZoneClass() TimeZoneClass {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.timeZoneClass
}

// GetTimeZoneName - Returns the time zone name, also known
// as the Time Zone 'Location' Name.
//
func (tzSpec *TimeZoneSpecification) GetTimeZoneName() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.locationName
}


// GetTimeZoneSpecFlags - Returns all internal flags for the
// current TimeZoneSpecification instance.
//
func (tzSpec *TimeZoneSpecification) GetTimeZoneSpecFlags() (
	LocationNameType,
	TimeZoneCategory,
	TimeZoneClass,
	TimeZoneType,
	TimeZoneUtcOffsetStatus) {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.locationNameType,
					tzSpec.timeZoneCategory,
					tzSpec.timeZoneClass,
					tzSpec.timeZoneType,
					tzSpec.timeZoneUtcOffsetStatus
}

// GetTimeZoneType - Returns the Time Zone Type description.
// Time Zone Type is an enumeration identifying the time zone
// source.
//
// Possible return types.
//
//  TimeZoneType(0).None()      - Time Zone type is uninitialized
//                                and has no significant value.
//
//  TimeZoneType(0).Iana()      - Identifies an IANA Time Zone
//
//  TimeZoneType(0).Local()     - Identifies this as a 'Local' Time Zone
//
//  TimeZoneType(0).Military()  - Identifies a Military Time Zone
//
// For easy access to these enumeration values, use the global variable
// 'TzType'. Example: TzType.Military()
//
func (tzSpec *TimeZoneSpecification) GetTimeZoneType() TimeZoneType {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.timeZoneType
}

// GetTimeZoneUtcOffsetStatus - Returns a description of the UTC Offset
// Status for this time zone. Some Time Zones have a constant UTC Offset
// throughout the year. Others, primarily those which observe Daylight
// Savings Time have a UTC Offset which varies at different times of the
// year.
//
// Possible return values:
//
// TimeZoneUtcOffsetStatus(0).None()
//               - Signals that Time Zone UTC Offset
//                 Status is uninitialized and contains
//                 no significant or valid value. This
//                 is an error condition.
//
// TimeZoneUtcOffsetStatus(0).Static()
//               - Signals that the UTC Offset associated
//                 with a given Time Zone is constant
//                 throughout the year and never changes.
//                 Typically, this means that Daylight
//                 Savings Time is NOT observed in the
//                 specified Time Zone.
//
// TimeZoneUtcOffsetStatus(0).Variable()
//               - Signals that the UTC Offset associated
//                 with a given Time Zone is not constant,
//                 and varies at least once during the year.
//                 This usually means that Daylight Savings
//                 Time is observed within the designated
//                 Time Zone.
//
// For easy access to these enumeration values, use the global variable
// 'TzUtcStatus'. Example: TzUtcStatus.Variable()
//
func (tzSpec *TimeZoneSpecification) GetTimeZoneUtcOffsetStatus() TimeZoneUtcOffsetStatus {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.timeZoneUtcOffsetStatus
}

// GetUtcOffset - returns a text string representing the
// offset from UTC for this time zone.
//
//  Examples: "-0600", "+0200"
//
func (tzSpec *TimeZoneSpecification) GetUtcOffset() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.utcOffset
}

// GetZoneAbbrvLookupId - Returns a text string containing the
// Time Zone abbreviation plus the UTC offset. This text value
// is used to look up time zone data in various internal data
// maps. Examples: "CST-0600", "EET+0200"
//
// Note: To access the time zone abbreviation, see method
// TimeZoneSpecification.GetOriginalZoneName()
//
func (tzSpec *TimeZoneSpecification) GetZoneAbbrvLookupId() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneAbbrvLookupId
}

// GetZoneLabel - Returns the Zone Label, a tag or text
// description field available for use by the user.
//
func (tzSpec *TimeZoneSpecification) GetZoneLabel() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneLabel
}

// GetZoneName - Returns the 'Zone Name'. 'Zone Name' is the
// the Time Zone Abbreviation.
//
// Zone Name Examples:
//   'EST', 'CST', 'PST', 'EDT', 'CDT', 'PDT'
//
func (tzSpec *TimeZoneSpecification) GetZoneName() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneName
}

// GetZoneOffset - Returns data field 'zoneOffset'. This is
// a text string representing the offset from UTC for this
// time zone. The returned offset string consists of two
// components, the hours and minutes of offset and the time
// zone abbreviation.
//
// Example: "-0600 CST" or "+0200 EET"
//
func (tzSpec *TimeZoneSpecification) GetZoneOffset() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneOffset
}

// GetZoneOffsetTotalSeconds - Returns the total offset seconds
// from 'UTC' for this time zone. The returned value is a signed
// value. Positive ('+') values identify seconds East of UTC.
// Negative ('-') values identify seconds West of UTC.
//
func (tzSpec *TimeZoneSpecification) GetZoneOffsetTotalSeconds() int {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneOffsetTotalSeconds
}

// GetZoneSignChar - Returns the sign character as a string.
// This is a string value identifying whether the time zone
// offset from UTC is east or west of UTC. The returned string
// will hold one of only two values: a positive "+" or a negative
// "-".  A negative minus ('-') indicates an offset West of UTC
// while a positive plus ('+') identifies and offset East of UTC.
// This string is designed to be used with the unsigned or positive
// values for offset hours, minutes and seconds.
//
func (tzSpec *TimeZoneSpecification) GetZoneSignChar() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	if tzSpec.zoneSignValue < 0 {
		return "-"
	}

	return "+"
}

// GetZoneSignValue - Returns the data field 'zoneSignValue'. This
// is a signed integer value identifying whether the time zone
// offset from UTC is east or west of UTC. The returned integer
// will hold one of only two values: a positive '1' or a negative
// '-1'.  '-1' indicates an offset West of UTC while a positive
// '1' identifies and offset East of UTC. Apply this sign to the
// unsigned values for offset hours, minutes and seconds.
//
func (tzSpec *TimeZoneSpecification) GetZoneSignValue() int {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneSignValue
}

// New - Returns a new TimeZoneSpecification instance with member
// variables initialized to zero values.
//
func (tzSpec TimeZoneSpecification) New() (TimeZoneSpecification, error) {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	ePrefix := "TimeZoneSpecification.New() "

	tzSpec2 := TimeZoneSpecification{}

	tzSpecUtil := timeZoneSpecUtility{}

	err := tzSpecUtil.setZeroTimeZoneSpec(
		&tzSpec2,
		ePrefix)

	return tzSpec2, err
}

// NewRefDate - Returns a new instance of TimeZoneSpecification
// based on a reference date time input parameter.
//
// Input parameter 'referenceDateTime' is used to extract the
// time zone name. If parameters 'militaryTimeZoneName' or
// 'militaryTimeZoneLetter' are populated, they will control
// and an associated military time zone definition will be
// created.
//
// Note: Input parameters zoneLabel and 'tagDescription' are
// available to the user for adding descriptive narrative text
// to this TimeZoneSpecification instance.
//
func (tzSpec TimeZoneSpecification) NewRefDate(
	referenceDateTime      time.Time,
	militaryTimeZoneName   string,
	militaryTimeZoneLetter string,
	zoneLabel              string,
	tagDescription         string,
	timeZoneClass          TimeZoneClass,
	ePrefix string) (TimeZoneSpecification, error) {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	ePrefix += "TimeZoneSpecification.NewRefDate() "

	tzSpecUtil := timeZoneSpecUtility{}

	tzSpecOut := TimeZoneSpecification{}

	err := tzSpecUtil.setTimeZone(
		&tzSpecOut,
		referenceDateTime,
		militaryTimeZoneName,
		militaryTimeZoneLetter,
		zoneLabel,
		tagDescription,
		timeZoneClass,
		ePrefix)

	return tzSpecOut, err
}

// SetOriginalTagDescription - Sets the value of member variable
// and data field, TimeZoneSpecification.tagDescription. This
// field is available for users to tag, classify or
// otherwise attach descriptive information to this
// TimeZoneSpecification instance.
//
func (tzSpec *TimeZoneSpecification) SetTagDescription(tagDescription string) {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	tzSpec.tagDescription = tagDescription
}

// SetZoneLabel - Sets the value of data field "Zone Label". 'Zone
// Label' a tag or text description field available for use
// by the user.
//
func (tzSpec *TimeZoneSpecification) SetZoneLabel() string {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneLabel
}

// SetTimeZone - Sets the data values of the current Time Zone
// Specification Structure (TimeZoneSpecification).
//
func (tzSpec *TimeZoneSpecification) SetTimeZone(
	referenceDateTime      time.Time,
	militaryTimeZoneLetter string,
	militaryTimeZoneName   string,
	zoneLabel              string,
	tagDescription         string,
	timeZoneClass          TimeZoneClass,
	ePrefix string) error {

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	ePrefix += "TimeZoneSpecification.SetTimeZone() "

	tzSpecUtil := timeZoneSpecUtility{}

	return tzSpecUtil.setTimeZone(
		tzSpec,
		referenceDateTime,
		militaryTimeZoneLetter,
		militaryTimeZoneName,
		zoneLabel,
		tagDescription,
		timeZoneClass,
		ePrefix)
}