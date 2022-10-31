package sgp4

import "github.com/SharkEzz/sgp4/internal/cppsgp4"

type Observer struct {
	cobserver cppsgp4.Observer
}

func NewObserver(coords *CoordGeodetic) (o *Observer, err error) {
	defer catch(&err)

	obs := cppsgp4.NewObserver(coords._coordGeodetic)

	return &Observer{obs}, err
}

func (o *Observer) GetLocation() *CoordGeodetic {
	return &CoordGeodetic{o.cobserver.GetLocation()}
}

func (o *Observer) GetLookAngle(sat *Eci) *CoordTopocentric {
	return &CoordTopocentric{o.cobserver.GetLookAngle(sat.ceci)}
}
