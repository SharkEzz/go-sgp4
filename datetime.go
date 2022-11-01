package sgp4

import (
	"time"

	"github.com/SharkEzz/go-sgp4/internal/cppsgp4"
)

type DateTime struct {
	_dateTime cppsgp4.DateTime
}

func NewDateTimeFromTime(t time.Time) (dt *DateTime, err error) {
	defer catch(&err)

	tUTC := t.UTC()

	return &DateTime{
		cppsgp4.NewDateTime(tUTC.Year(), int(tUTC.Month()), tUTC.Day(), tUTC.Hour(), tUTC.Minute(), tUTC.Second()),
	}, err
}

func NewDateTimeNow() (*DateTime, error) {
	return &DateTime{cppsgp4.DateTimeNow(false)}, nil
}

func (dt *DateTime) Year() int {
	return dt._dateTime.Year()
}

func (dt *DateTime) Month() int {
	return dt._dateTime.Month()
}

func (dt *DateTime) Day() int {
	return dt._dateTime.Day()
}

func (dt *DateTime) Hour() int {
	return dt._dateTime.Hour()
}

func (dt *DateTime) Minute() int {
	return dt._dateTime.Minute()
}

func (dt *DateTime) Second() int {
	return dt._dateTime.Second()
}

func (dt *DateTime) Time() time.Time {
	return time.Date(dt.Year(), time.Month(dt.Month()), dt.Day(), dt.Hour(), dt.Minute(), dt.Second(), 0, time.UTC)
}
