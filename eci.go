package sgp4

import (
	"github.com/SharkEzz/go-sgp4/internal/cppsgp4"
)

type Eci struct {
	ceci cppsgp4.Eci
}

func NewEci(dt *DateTime, coords *CoordGeodetic) (e *Eci, err error) {
	defer catch(&err)

	return &Eci{cppsgp4.NewEci(dt._dateTime, coords._coordGeodetic)}, err
}

func (e *Eci) ToGeodetic() (c *CoordGeodetic, err error) {
	defer catch(&err)

	coords := e.ceci.ToGeodetic()

	geodetic, err := NewCoordGeodetic(coords.GetLatitude(), coords.GetLongitude(), coords.GetAltitude(), true)
	if err != nil {
		return nil, err
	}

	return geodetic, err
}
