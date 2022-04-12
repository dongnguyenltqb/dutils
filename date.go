package dutils

import "time"

func IsEqualDay(d1 time.Time, d2 time.Time) bool {
	return d1.Day() == d2.Day()
}

func IsEqualYearMonthDate(d1 time.Time, d2 time.Time) bool {
	year1, month1, day1 := d1.Date()
	year2, month2, day2 := d2.Date()
	return year1 == year2 && month1 == month2 && day1 == day2
}
