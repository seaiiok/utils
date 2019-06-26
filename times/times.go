package times

import (
	"time"
)

type times struct {
}

const TimeLayout = "2006-01-02 15:04:05"

func New() *times {
	return &times{}
}

func (t *times) TimeString2Unix(timeString string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TimeLayout, timeString, loc)
	return theTime.Unix()
}

func (t *times) Unix2TimeString(unix int64) string {
	return time.Unix(unix, 0).Format(TimeLayout)
}
