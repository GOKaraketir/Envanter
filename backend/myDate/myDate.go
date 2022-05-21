package myDate

import "time"

type MyDate struct {
	year                int
	month               time.Month
	day, hour, min, sec int
}

func ToTime(myDate *MyDate) (date time.Time) {
	return time.Date(myDate.year, myDate.month, myDate.day, myDate.hour, myDate.min, myDate.sec, 0, time.Local)
}

func FromTime(date time.Time) (myDate MyDate) {
	myDate.year = date.Year()
	myDate.month = date.Month()
	myDate.day = date.Day()
	myDate.hour = date.Hour()
	myDate.min = date.Minute()
	myDate.sec = date.Second()
	return
}
