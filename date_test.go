package dutils

import (
	"testing"
	"time"
)

func TestIsEqualDay(t *testing.T) {
	day1, _ := time.Parse(time.RFC3339, "2006-01-03T15:04:05Z")
	day2, _ := time.Parse(time.RFC3339, "1997-01-03T15:04:05Z")
	if IsEqualDay(day1, day2) != true {
		t.Errorf("Incorrent %v %v", day1.Day(), day2.Day())
	}
}

func TestIsEqualYearMonthDay(t *testing.T) {
	day1, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	day2, _ := time.Parse(time.RFC3339, "2006-01-03T15:04:05Z")
	if IsEqualYearMonthDate(day1, day2) == true {
		t.Errorf("%v %v", day1, day2)
	}

	day1, _ = time.Parse(time.RFC3339, "2006-01-02T01:04:05Z")
	day2, _ = time.Parse(time.RFC3339, "2006-01-02T23:04:05Z")
	if IsEqualYearMonthDate(day1, day2) != true {
		t.Errorf("%v %v", day1, day2)
	}
}
