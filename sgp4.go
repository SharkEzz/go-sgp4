package sgp4

import (
	"time"

	"github.com/SharkEzz/sgp4/cppsgp4"
	"github.com/SharkEzz/sgp4/utils"
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

	dt, err := NewDateTimeFromTime(lt.UTC())
	if err != nil {
		return 0, 0, 0, err
	}

	pos := s.csgp4.FindPosition(dt)
	geo := pos.ToGeodetic()

	lat = utils.Rad2Deg(geo.GetLatitude())
	lng = utils.Rad2Deg(geo.GetLongitude())
	alt = geo.GetAltitude()

	return lat, lng, alt, err
}
