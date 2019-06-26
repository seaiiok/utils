package utils

import (
	"time"
)

type _times struct {
}

// 格式时间<2006-01-02 15:04:05>转化为时间戳
func (t *_times) TimeString2Unix(timeString string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TimeLayout, timeString, loc)
	return theTime.Unix()
}

// 时间戳转化为格式时间<2006-01-02 15:04:05>
func (t *_times) Unix2TimeString(unix int64) string {
	return time.Unix(unix, 0).Format(TimeLayout)
}
