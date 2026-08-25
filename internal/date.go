package internal

import "time"

func DateDayEqual(a time.Time, b time.Time) bool {
	return a.Year() == b.Year() &&
		a.Month() == b.Month() &&
		a.Day() == b.Day()
}
