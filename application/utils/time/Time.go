package timeUtils

import (
	"time"
)

func GetMonthStartEnd(t string) (int64, int64) {
	var times time.Time
	var types string
	switch len(t) {
	case 7:
		types = "2006-01"
	case 10:
		types = "2006-01-02"

	}
	times, _ = time.ParseInLocation(types, t, time.Local)

	monthStartDay := times.AddDate(0, 0, -times.Day()+1)
	monthStartTime := time.Date(monthStartDay.Year(), monthStartDay.Month(), monthStartDay.Day(), 0, 0, 0, 0, times.Location())
	monthEndDay := monthStartTime.AddDate(0, 1, -1)
	monthEndTime := time.Date(monthEndDay.Year(), monthEndDay.Month(), monthEndDay.Day(), 23, 59, 59, 0, times.Location())

	return FmtTimeToTimestamp(monthStartTime), FmtTimeToTimestamp(monthEndTime)
}
func FmtTimeToTimestamp(t time.Time) int64 {
	times, _ := time.ParseInLocation("2006-01-02", t.Format("2006-01-02"), time.Local)
	return times.Unix()
}

func GetDate(v int64) string {
	return time.Unix(v, 0).Format("2006-01-02")

}
