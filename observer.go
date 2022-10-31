package sgp4

import "github.com/SharkEzz/sgp4/cppsgp4"

func NewObserver(coords cppsgp4.CoordGeodetic) (o cppsgp4.Observer, err error) {
	defer catch(&err)

	return cppsgp4.NewObserver(coords), err
}
