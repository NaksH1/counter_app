package util

import (
	"time"
)

// converts time from string to ISO 8601 format
func FormatDateToISO(date string) (*time.Time, error) {
	time, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	return &time, nil
}

// compares dates and return true if date1 is less than or equal to date2 else false
func CompareDates(date1 time.Time, date2 time.Time) bool {
	return date1.Before(date2) || date1.Equal(date2)
}

// converts date from format time.Time to string format 2006-12-29
func FormatDate(date time.Time) string {
	return date.Format("2006-01-02")
}
