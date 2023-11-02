package timestamp

import (
	"time"
)

func GetBeginOfADay(year int, month time.Month, day int, location string) int64 {
	loc, _ := time.LoadLocation(location)
	return time.Date(year, month, day, 0, 0, 0, 0, loc).Unix()
}

func GetEndOfADay(year int, month time.Month, day int, location string) int64 {
	loc, _ := time.LoadLocation(location)
	return time.Date(year, month, day, 23, 59, 59, 0, loc).Unix()
}

func GetYearMonthDay() (int, time.Month, int) {
	return time.Now().Date()
}

func ToDateTimeString(ts int64) string {
	tm := time.Unix(ts, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func AddDate(ts int64, years int, months, days int) int64 {
	return time.Unix(ts, 0).AddDate(years, months, days).Unix()
}
