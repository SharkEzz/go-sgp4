package sgp4

import (
	"time"

	"github.com/SharkEzz/sgp4/cppsgp4"
)

type SGP4 struct {
	csgp4 cppsgp4.SGP4
}

func NewSGP4(tle cppsgp4.Tle) (p *SGP4, err error) {
	defer catch(&err)
	cp := cppsgp4.NewSGP4(tle)

	return &SGP4{
		csgp4: cp,
	}, nil
}

func (s *SGP4) SGP4() cppsgp4.SGP4 {
	return s.csgp4
}

func (s *SGP4) Position(lt time.Time) (lat float64, lng float64, alt float64, err error) {
	defer catch(&err)

	t := lt.UTC()
	dt := cppsgp4.NewDateTime(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())

	pos := s.csgp4.FindPosition(dt)
	geo := pos.ToGeodetic()

	lat = Rad2Deg(geo.GetLatitude())
	lng = Rad2Deg(geo.GetLongitude())
	alt = geo.GetAltitude()

	return lat, lng, alt, nil
}
