package sgp4

import "github.com/SharkEzz/sgp4/cppsgp4"

func NewEci(dt cppsgp4.DateTime, coords cppsgp4.CoordGeodetic) (e cppsgp4.Eci, err error) {
	defer catch(&err)

	return cppsgp4.NewEci(dt, coords), err
}
