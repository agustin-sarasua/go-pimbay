package util

import (
	"fmt"
	"strings"
	"time"
)

type DateTime struct {
	time.Time
}

const ctLayout = "2006/01/02|15:04:05"
const ctLayout2 = "2017-06-23T09:50:00-03:00"

func (ct *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(time.RFC3339Nano, s)
	return
}

func (ct *DateTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(time.RFC3339Nano))), nil
}

var nilTime = (time.Time{}).UnixNano()

func (ct *DateTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

type Args struct {
	Time DateTime
}
