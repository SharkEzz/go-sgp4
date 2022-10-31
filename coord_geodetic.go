package sgp4

import "github.com/SharkEzz/sgp4/internal/cppsgp4"

type CoordGeodetic struct {
	ccoordGeodetic cppsgp4.CoordGeodetic
}

func NewCoordGeodetic(lat, lon, alt float64) (c *CoordGeodetic, err error) {
	defer catch(&err)

	coords := cppsgp4.NewCoordGeodetic()
	coords.SetAltitude(alt)
	coords.SetLatitude(lat)
	coords.SetLongitude(lon)

	return &CoordGeodetic{ccoordGeodetic: coords}, err
}

func (c *CoordGeodetic) CoordGeodetic() cppsgp4.CoordGeodetic {
	return c.ccoordGeodetic
}

func (c *CoordGeodetic) Altitude() float64 {
	return c.ccoordGeodetic.GetAltitude()
}

func (c *CoordGeodetic) Latitude() float64 {
	return c.ccoordGeodetic.GetLatitude()
}

func (c *CoordGeodetic) Longitude() float64 {
	return c.ccoordGeodetic.GetLongitude()
}
