package sgp4

import "github.com/SharkEzz/go-sgp4/internal/cppsgp4"

type Observer struct {
	cobserver cppsgp4.Observer
}

func NewObserver(coords *CoordGeodetic) (o *Observer, err error) {
	defer catch(&err)

	return &Observer{cppsgp4.NewObserver(coords._coordGeodetic)}, err
}

func (o *Observer) GetLocation() (c *CoordGeodetic, err error) {
	defer catch(&err)

	return &CoordGeodetic{o.cobserver.GetLocation()}, err
}

func (o *Observer) GetLookAngle(sat *Eci) (c *CoordTopocentric, err error) {
	defer catch(&err)

	return &CoordTopocentric{o.cobserver.GetLookAngle(sat.ceci)}, err
}
