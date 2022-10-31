package sgp4

import (
	"time"

	"github.com/SharkEzz/sgp4/internal/cppsgp4"
)

type SGP4 struct {
	csgp4 cppsgp4.SGP4
}

func NewSGP4(tle *Tle) (p *SGP4, err error) {
	defer catch(&err)
	cp := cppsgp4.NewSGP4(tle.ctle)

	return &SGP4{
		csgp4: cp,
	}, nil
}

func (s *SGP4) FindPosition(dt *DateTime) *Eci {
	return &Eci{s.csgp4.FindPosition(dt._dateTime)}
}

func (s *SGP4) Position(lt time.Time) (eci *Eci, err error) {
	defer catch(&err)

	dt, err := NewDateTimeFromTime(lt.UTC())
	if err != nil {
		return nil, err
	}

	return s.FindPosition(dt), err
}
