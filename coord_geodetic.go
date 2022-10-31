package sgp4

import "github.com/SharkEzz/sgp4/cppsgp4"

func NewCoordGeodetic(lat, lon, alt float64) (c cppsgp4.CoordGeodetic, err error) {
	defer catch(&err)

	c = cppsgp4.NewCoordGeodetic()
	c.SetAltitude(alt)
	c.SetLatitude(lat)
	c.SetLongitude(lon)

	return c, nil
}
