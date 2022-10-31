package sgp4

import (
	"time"

	"github.com/SharkEzz/sgp4/cppsgp4"
)

func NewDateTimeFromTime(t time.Time) (dt cppsgp4.DateTime, err error) {
	defer catch(&err)

	tUTC := t.UTC()

	return cppsgp4.NewDateTime(tUTC.Year(), int(tUTC.Month()), tUTC.Day(), tUTC.Hour(), tUTC.Minute(), tUTC.Second()), err
}
