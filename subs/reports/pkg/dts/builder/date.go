package builder

import (
	"errors"
	"time"
)

// Builder to wrap time.Date
//
// Takes advantage of zero values so you don't have to set
// values you don't need
type Date struct {
	year  int
	month time.Month
	day   int
	hour  int
	min   int
	sec   int
	nsec  int
	loc   *time.Location
}

func NewDate() *Date {
	return &Date{}
}

// MustBuild builds the time.Time represented by the Date by calling
// time.Date
//
// Date returns the Time corresponding to
//
//	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
//
// in the appropriate zone for that time in the given location.
//
// The month, day, hour, min, sec, and nsec values may be outside
// their usual ranges and will be normalized during the conversion.
// For example, October 32 converts to November 1.
//
// A daylight savings time transition skips or repeats times.
// For example, in the United States, March 13, 2011 2:15am never occurred,
// while November 6, 2011 1:15am occurred twice. In such cases, the
// choice of time zone, and therefore the time, is not well-defined.
// Date returns a time that is correct in one of the two zones involved
// in the transition, but it does not guarantee which.
//
// MustBuild panics if d.loc is nil.
func (d *Date) MustBuild() time.Time {
	return time.Date(
		d.year,
		d.month,
		d.day,
		d.hour,
		d.min,
		d.sec,
		d.nsec,
		d.loc,
	)
}

// MustBuild builds the time.Time represented by the Date by calling
// time.Date
//
// Date returns the Time corresponding to
//
//	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
//
// in the appropriate zone for that time in the given location.
//
// The month, day, hour, min, sec, and nsec values may be outside
// their usual ranges and will be normalized during the conversion.
// For example, October 32 converts to November 1.
//
// A daylight savings time transition skips or repeats times.
// For example, in the United States, March 13, 2011 2:15am never occurred,
// while November 6, 2011 1:15am occurred twice. In such cases, the
// choice of time zone, and therefore the time, is not well-defined.
// Date returns a time that is correct in one of the two zones involved
// in the transition, but it does not guarantee which.
//
// Build will return ErrMissingLocation if d.loc is nil.
func (d *Date) Build() (time.Time, error) {
	if d.loc == nil {
		return time.Time{}, ErrMissingLocation
	}

	return time.Date(
		d.year,
		d.month,
		d.day,
		d.hour,
		d.min,
		d.sec,
		d.nsec,
		d.loc,
	), nil
}

func (d *Date) Year(year int) *Date {
	d.year = year
	return d
}

func (d *Date) Month(month time.Month) *Date {
	d.month = month
	return d
}

func (d *Date) Day(day int) *Date {
	d.day = day
	return d
}

func (d *Date) Hour(hour int) *Date {
	d.hour = hour
	return d
}

func (d *Date) Min(min int) *Date {
	d.min = min
	return d
}

func (d *Date) Sec(sec int) *Date {
	d.sec = sec
	return d
}

func (d *Date) Nsec(nsec int) *Date {
	d.nsec = nsec
	return d
}

func (d *Date) Loc(loc *time.Location) *Date {
	d.loc = loc
	return d
}

var (
	ErrMissingLocation = errors.New("time: missing Location in call to Date")
)
