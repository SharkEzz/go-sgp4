package sgp4

import "github.com/SharkEzz/sgp4/cppsgp4"

func NewCoordTopocentric(azimuth, elevation, xrange, range_rate float64) (c cppsgp4.CoordTopocentric, err error) {
	defer catch(&err)

	c = cppsgp4.NewCoordTopocentric()
	c.SetAzimuth(azimuth)
	c.SetElevation(elevation)
	c.SetXrange(xrange)
	c.SetRange_rate(range_rate)

	return c, err
}
