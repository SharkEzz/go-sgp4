package sgp4

import (
	"github.com/SharkEzz/sgp4/internal/cppsgp4"
	"github.com/SharkEzz/sgp4/utils"
)

type CoordGeodetic struct {
	_coordGeodetic cppsgp4.CoordGeodetic
}

func NewCoordGeodetic(lat, lon, alt float64) (c *CoordGeodetic, err error) {
	defer catch(&err)

	coords := cppsgp4.NewCoordGeodetic()
	coords.SetAltitude(alt)
	coords.SetLatitude(lat)
	coords.SetLongitude(lon)

	return &CoordGeodetic{_coordGeodetic: coords}, err
}

func (c *CoordGeodetic) Altitude() float64 {
	return c._coordGeodetic.GetAltitude()
}

func (c *CoordGeodetic) Latitude() float64 {
	return c._coordGeodetic.GetLatitude()
}

func (c *CoordGeodetic) Longitude() float64 {
	return c._coordGeodetic.GetLongitude()
}

func (c *CoordGeodetic) ToString() string {
	return c._coordGeodetic.ToString()
}

func (c *CoordGeodetic) GetCoords(in_degrees bool) (lat float64, lng float64, alt float64, err error) {
	defer catch(&err)

	if in_degrees {
		lat = utils.Rad2Deg(c.Latitude())
		lng = utils.Rad2Deg(c.Longitude())
	} else {
		lat = c.Latitude()
		lng = c.Longitude()
	}
	alt = c.Altitude()

	return lat, lng, alt, err
}
