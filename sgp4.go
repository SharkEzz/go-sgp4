package sgp4

import (
	"time"

	"github.com/SharkEzz/go-sgp4/internal/cppsgp4"
)

type SGP4 struct {
	_sgp4 cppsgp4.SGP4
}

func NewSGP4(tle *Tle) (p *SGP4, err error) {
	defer catch(&err)
	cp := cppsgp4.NewSGP4(tle._tle)

	return &SGP4{
		_sgp4: cp,
	}, nil
}

func (s *SGP4) PositionFromDateTime(dt *DateTime) (eci *Eci, err error) {
	defer catch(&err)

	return &Eci{s._sgp4.FindPosition(dt._dateTime)}, err
}

func (s *SGP4) PositionFromTime(lt time.Time) (eci *Eci, err error) {
	defer catch(&err)

	dt, err := NewDateTimeFromTime(lt.UTC())
	if err != nil {
		return nil, err
	}

	return s.PositionFromDateTime(dt)
}

func (s *SGP4) Close() {
	cppsgp4.DeleteSGP4(s._sgp4)
}
