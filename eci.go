package sgp4

import (
	"github.com/SharkEzz/sgp4/internal/cppsgp4"
	"github.com/SharkEzz/sgp4/utils"
)

type Eci struct {
	ceci cppsgp4.Eci
}

func NewEci(dt cppsgp4.DateTime, coords cppsgp4.CoordGeodetic) (e *Eci, err error) {
	defer catch(&err)

	return &Eci{cppsgp4.NewEci(dt, coords)}, err
}

func (e *Eci) Eci() cppsgp4.Eci {
	return e.ceci
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

func GetGeodeticCoords(eci *Eci, in_degrees bool) (lat float64, lng float64, alt float64, err error) {
	defer catch(&err)

	coords, err := eci.ToGeodetic()
	if err != nil {
		return 0, 0, 0, err
	}

	if in_degrees {
		lat = utils.Rad2Deg(coords.Latitude())
		lng = utils.Rad2Deg(coords.Longitude())
	} else {
		lat = coords.Latitude()
		lng = coords.Longitude()
	}
	alt = coords.Altitude()

	return lat, lng, alt, err
}
