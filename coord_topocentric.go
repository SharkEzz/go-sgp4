package sgp4

import "github.com/SharkEzz/sgp4/internal/cppsgp4"

type CoordTopocentric struct {
	_coordTopocentric cppsgp4.CoordTopocentric
}

func NewCoordTopocentric(azimuth, elevation, xrange, range_rate float64) (c *CoordTopocentric, err error) {
	defer catch(&err)

	coords := cppsgp4.NewCoordTopocentric(azimuth, elevation, xrange, range_rate)

	return &CoordTopocentric{coords}, err
}

func (c *CoordTopocentric) Azimuth() float64 {
	return c._coordTopocentric.GetAzimuth()
}

func (c *CoordTopocentric) Elevation() float64 {
	return c._coordTopocentric.GetElevation()
}

func (c *CoordTopocentric) Range() float64 {
	return c._coordTopocentric.GetXrange()
}

func (c *CoordTopocentric) RangeRate() float64 {
	return c._coordTopocentric.GetRange_rate()
}

func (c *CoordTopocentric) ToString() string {
	return c._coordTopocentric.ToString()
}
