package sgp4

import (
	"github.com/SharkEzz/sgp4/internal/cppsgp4"
)

type Eci struct {
	ceci cppsgp4.Eci
}

func NewEci(dt cppsgp4.DateTime, coords cppsgp4.CoordGeodetic) (e *Eci, err error) {
	defer catch(&err)

	return &Eci{cppsgp4.NewEci(dt, coords)}, err
}

func (e *Eci) ToGeodetic() (c *CoordGeodetic, err error) {
	defer catch(&err)

	coords := e.ceci.ToGeodetic()

	geodetic, err := NewCoordGeodetic(coords.GetLatitude(), coords.GetLongitude(), coords.GetAltitude())
	if err != nil {
		return nil, err
	}

	return geodetic, err
}
