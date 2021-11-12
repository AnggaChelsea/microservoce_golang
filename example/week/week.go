package week

import (
	"time"
)

//weekday
func Weekday(year , month, day int) time.Weekday{
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	return date.Weekday()
}