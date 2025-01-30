package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	var dates [4]time.Time

	dates[0], _ = time.Parse("2006-01-02 15:04:05.000000000 MST -07:00", "1609-09-12 19:02:35.123456789 PDT +03:00")
	dates[1], _ = time.Parse("2006-01-02 03:04:05 PM -0700", "1995-11-07 04:29:43 AM -0209")
	dates[2], _ = time.Parse("PM -0700 01/02/2006 03:04:05", "AM -0209 11/07/1995 04:29:43")
	dates[3], _ = time.Parse("Time:Z07:00T15:04:05 Date:2006-01-02 ", "Time:-03:30T19:18:35 Date:2119-10-29")

	defaultFormat := "2006-01-02 15:04:05 PM -07:00 Jan Mon MST"

	formats := []map[string]string{
		{"format": "2006", "description": "Year"},
		{"format": "06", "description": "Year"},

		{"format": "01", "description": "Month"},
		{"format": "1", "description": "Month"},
		{"format": "Jan", "description": "Month"},
		{"format": "January", "description": "Month"},

		{"format": "02", "description": "Day"},
		{"format": "2", "description": "Day"},

		{"format": "Mon", "description": "Week day"},
		{"format": "Monday", "description": "Week day"},

		{"format": "03", "description": "Hours"},
		{"format": "3", "description": "Hours"},
		{"format": "15", "description": "Hours"},

		{"format": "04", "description": "Minutes"},
		{"format": "4", "description": "Minutes"},

		{"format": "05", "description": "Seconds"},
		{"format": "5", "description": "Seconds"},

		{"format": "PM", "description": "AM or PM"},

		{"format": ".000", "description": "Miliseconds"},
		{"format": ".000000", "description": "Microseconds"},
		{"format": ".000000000", "description": "Nanoseconds"},

		{"format": "-0700", "description": "Timezone offset"},
		{"format": "-07:00", "description": "Timezone offset"},
		{"format": "Z0700", "description": "Timezone offset"},
		{"format": "Z07:00", "description": "Timezone offset"},

		{"format": "MST", "description": "Timezone"}}

	for _, date := range dates {
		fmt.Printf("\n\n %s \n", date.Format(defaultFormat))
		fmt.Printf("%-15s + %-12s + %12s \n", strings.Repeat("-", 15), strings.Repeat("-", 12), strings.Repeat("-", 12))
		fmt.Printf("%-15s | %-12s | %12s \n", "Type", "Placeholder", "Value")
		fmt.Printf("%-15s + %-12s + %12s \n", strings.Repeat("-", 15), strings.Repeat("-", 12), strings.Repeat("-", 12))

		for _, f := range formats {
			fmt.Printf("%-15s | %-12s | %-12s \n", f["description"], f["format"], date.Format(f["format"]))
		}
		fmt.Printf("%-15s + %-12s + %12s \n", strings.Repeat("-", 15), strings.Repeat("-", 12), strings.Repeat("-", 12))
	}
}

/*	Output
	$ go run main.go


	 1609-09-12 19:02:35 PM +03:00 Sep Sat PDT
	--------------- + ------------ + ------------
	Type            | Placeholder  |        Value
	--------------- + ------------ + ------------
	Year            | 2006         | 1609
	Year            | 06           | 09
	Month           | 01           | 09
	Month           | 1            | 9
	Month           | Jan          | Sep
	Month           | January      | September
	Day             | 02           | 12
	Day             | 2            | 12
	Week day        | Mon          | Sat
	Week day        | Monday       | Saturday
	Hours           | 03           | 07
	Hours           | 3            | 7
	Hours           | 15           | 19
	Minutes         | 04           | 02
	Minutes         | 4            | 2
	Seconds         | 05           | 35
	Seconds         | 5            | 35
	AM or PM        | PM           | PM
	Miliseconds     | .000         | .123
	Microseconds    | .000000      | .123456
	Nanoseconds     | .000000000   | .123456789
	Timezone offset | -0700        | +0300
	Timezone offset | -07:00       | +03:00
	Timezone offset | Z0700        | +0300
	Timezone offset | Z07:00       | +03:00
	Timezone        | MST          | PDT
	--------------- + ------------ + ------------


	 1995-11-07 04:29:43 AM -02:09 Nov Tue -0209
	--------------- + ------------ + ------------
	Type            | Placeholder  |        Value
	--------------- + ------------ + ------------
	Year            | 2006         | 1995
	Year            | 06           | 95
	Month           | 01           | 11
	Month           | 1            | 11
	Month           | Jan          | Nov
	Month           | January      | November
	Day             | 02           | 07
	Day             | 2            | 7
	Week day        | Mon          | Tue
	Week day        | Monday       | Tuesday
	Hours           | 03           | 04
	Hours           | 3            | 4
	Hours           | 15           | 04
	Minutes         | 04           | 29
	Minutes         | 4            | 29
	Seconds         | 05           | 43
	Seconds         | 5            | 43
	AM or PM        | PM           | AM
	Miliseconds     | .000         | .000
	Microseconds    | .000000      | .000000
	Nanoseconds     | .000000000   | .000000000
	Timezone offset | -0700        | -0209
	Timezone offset | -07:00       | -02:09
	Timezone offset | Z0700        | -0209
	Timezone offset | Z07:00       | -02:09
	Timezone        | MST          | -0209
	--------------- + ------------ + ------------


	 1995-11-07 04:29:43 AM -02:09 Nov Tue -0209
	--------------- + ------------ + ------------
	Type            | Placeholder  |        Value
	--------------- + ------------ + ------------
	Year            | 2006         | 1995
	Year            | 06           | 95
	Month           | 01           | 11
	Month           | 1            | 11
	Month           | Jan          | Nov
	Month           | January      | November
	Day             | 02           | 07
	Day             | 2            | 7
	Week day        | Mon          | Tue
	Week day        | Monday       | Tuesday
	Hours           | 03           | 04
	Hours           | 3            | 4
	Hours           | 15           | 04
	Minutes         | 04           | 29
	Minutes         | 4            | 29
	Seconds         | 05           | 43
	Seconds         | 5            | 43
	AM or PM        | PM           | AM
	Miliseconds     | .000         | .000
	Microseconds    | .000000      | .000000
	Nanoseconds     | .000000000   | .000000000
	Timezone offset | -0700        | -0209
	Timezone offset | -07:00       | -02:09
	Timezone offset | Z0700        | -0209
	Timezone offset | Z07:00       | -02:09
	Timezone        | MST          | -0209
	--------------- + ------------ + ------------


	 2119-10-29 19:18:35 PM -03:30 Oct Sun -0330
	--------------- + ------------ + ------------
	Type            | Placeholder  |        Value
	--------------- + ------------ + ------------
	Year            | 2006         | 2119
	Year            | 06           | 19
	Month           | 01           | 10
	Month           | 1            | 10
	Month           | Jan          | Oct
	Month           | January      | October
	Day             | 02           | 29
	Day             | 2            | 29
	Week day        | Mon          | Sun
	Week day        | Monday       | Sunday
	Hours           | 03           | 07
	Hours           | 3            | 7
	Hours           | 15           | 19
	Minutes         | 04           | 18
	Minutes         | 4            | 18
	Seconds         | 05           | 35
	Seconds         | 5            | 35
	AM or PM        | PM           | PM
	Miliseconds     | .000         | .000
	Microseconds    | .000000      | .000000
	Nanoseconds     | .000000000   | .000000000
	Timezone offset | -0700        | -0330
	Timezone offset | -07:00       | -03:30
	Timezone offset | Z0700        | -0330
	Timezone offset | Z07:00       | -03:30
	Timezone        | MST          | -0330
	--------------- + ------------ + ------------

*/
