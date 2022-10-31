package sgp4

import (
	"time"

	"github.com/SharkEzz/sgp4/internal/cppsgp4"
)

type DateTime struct {
	cdateTime cppsgp4.DateTime
}

func NewDateTimeFromTime(t time.Time) (dt *DateTime, err error) {
	defer catch(&err)

	tUTC := t.UTC()

	return &DateTime{
		cppsgp4.NewDateTime(tUTC.Year(), int(tUTC.Month()), tUTC.Day(), tUTC.Hour(), tUTC.Minute(), tUTC.Second()),
	}, err
}

func (dt *DateTime) DateTime() cppsgp4.DateTime {
	return dt.cdateTime
}

func NewDateTimeNow() (*DateTime, error) {
	return NewDateTimeFromTime(time.Now())
}
