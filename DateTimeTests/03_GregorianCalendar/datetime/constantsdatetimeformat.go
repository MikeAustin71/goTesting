package datetime

import "sync"

/*
 Date Time Format Constants


 This source file is located in source code repository:
 		'https://github.com/MikeAustin71/datetimeopsgo.git'

 This source code file is located at:
		'MikeAustin71\datetimeopsgo\datetime\constantsdatetime.go'



Overview and General Usage


This source file contains a series of constants useful in formatting
date time.

Types of constants defined here include:

    Date Time string formats

*/
const (
	// Date Time Format Constants
	// ================================================================================
	// FmtDateTimeSecondStr - Date Time format used
	// for file names and directory names
	FmtDateTimeSecondStr = "20060102150405"

	// FmtDateTimeNanoSecondStr - Custom Date Time Format
	FmtDateTimeNanoSecondStr = "2006-01-02 15:04:05.000000000"

	// FmtDateTimeSecText - Custom Date Time Format
	FmtDateTimeSecText = "2006-01-02 15:04:05"

	// FmtDateTimeDMYNanoTz - Outputs date time to nano seconds with associated time zone
	FmtDateTimeDMYNanoTz = "01/02/2006 15:04:05.000000000 -0700 MST"

	// FmtDateTimeTzNanoYMD - Outputs date time to nano seconds with Year-Month-Date
	FmtDateTimeTzNanoYMD = "2006-01-02 15:04:05.000000000 -0700 MST"

	// FmtDateTimeYMDHMSTz - Outputs Date time to seconds with Year-Month-Date and Time Zone.
	FmtDateTimeYMDHMSTz = "2006-01-02 15:04:05 -0700 MST"

	// FmtDateTimeTzNanoDowYMD - Output date time to nano seconds with Year-Month-Date
	// prefixed by day of the week
	FmtDateTimeTzNanoDowYMD = "Monday 2006-01-02 15:04:05.000000000 -0700 MST"

	// FmtDateTimeTzNanoYMDDow - Output date time to nano seconds with Year-Month-Date
	// prefixed by day of the week
	FmtDateTimeTzNanoYMDDow = "2006-01-02 Monday 15:04:05.000000000 -0700 MST"

	// FmtDateTimeYMDAbbrvDowNano - Output date time to nano seconds with abbreviated
	// day of week.
	FmtDateTimeYMDAbbrvDowNano = "2006-01-02 Mon 15:04:05.000000000 -0700 MST"

	// FmtDateTimeTzSec - Outputs date time to seconds with associated time zone
	FmtDateTimeTzSec = "01/02/2006 15:04:05 -0700 MST"

	// FmtDateTimeEverything - Custom Date Time Format showing virtually
	// all elements of a date time string.
	FmtDateTimeEverything = "Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	// FmtDateTimeNeutralDateFmt - Neutral Date Format without Time Zone
	FmtDateTimeNeutralDateFmt = "2006-01-02 15:04:05.000000000"

	// FmtDateTimeMDYrFmtStr - Month Day Year Date Format String
	FmtDateTimeMDYrFmtStr = "01/02/2006 15:04:05.000000000 -0700 MST"

	// FmtDateTimeUsMilitaryDate2DYr
	FmtDateTimeUsMilitaryDate2DYr = "021504Z 06"

	// FmtDateTimeYrMDayFmtStr - Year Month Day Date Format String
	FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
)

var lockDefaultDateTimeFormat sync.Mutex

const (
	DEFAULTDATETIMEFORMAT = "2006-01-02 15:04:05.000000000 -0700 MST"
)

