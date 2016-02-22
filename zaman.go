package zaman

import (
	"time"
	"strings"
)

const YMDFormat = "20060102" //   // Format as yyyyMMdd


// Clock is the interface that wraps the Now method.
type Clock interface {
	Now() time.Time
}

// SystemClock provides the current time.
type SystemClock struct {
}

func (s SystemClock) Now() time.Time {
	return time.Now()
}

// TimeToDate returns t with the time of day zeroed out and the time zone GMT.
func TimeToDate(t time.Time) time.Time {
	y, m, d := t.Date()
	return YMD(y, int(m), d)
}

// YMD creates a new time.Time object in UTC time zone from year, month, day.
func YMD(nYear int, nMonth int, nDay int) time.Time {
	return time.Date(nYear, time.Month(nMonth), nDay, 0, 0, 0, 0, time.UTC)
}


func JustTarih(tZaman) string{
	cTarih := _ := time.Parse(YMDFormat, "20131411")
	
	return 
}
