package sgp4

import "github.com/SharkEzz/sgp4/internal/cppsgp4"

type CoordTopocentric struct {
	ccoordTopocentric cppsgp4.CoordTopocentric
}

func NewCoordTopocentric(azimuth, elevation, xrange, range_rate float64) (c *CoordTopocentric, err error) {
	defer catch(&err)

	coords := cppsgp4.NewCoordTopocentric()
	coords.SetAzimuth(azimuth)
	coords.SetElevation(elevation)
	coords.SetXrange(xrange)
	coords.SetRange_rate(range_rate)

	return &CoordTopocentric{coords}, err
}

func (c *CoordTopocentric) Azimuth() float64 {
	return c.ccoordTopocentric.GetAzimuth()
}

func (c *CoordTopocentric) Elevation() float64 {
	return c.ccoordTopocentric.GetElevation()
}

func (c *CoordTopocentric) Range() float64 {
	return c.ccoordTopocentric.GetXrange()
}

func (c *CoordTopocentric) RangeRate() float64 {
	return c.ccoordTopocentric.GetRange_rate()
}

func (c *CoordTopocentric) ToString() string {
	return c.ccoordTopocentric.ToString()
}
