package sgp4

import "github.com/SharkEzz/sgp4/cppsgp4"

func NewDateTime(year, month, day, hour, minute, second int) (dt cppsgp4.DateTime, err error) {
	defer catch(&err)

	return cppsgp4.NewDateTime(year, month, day, hour, minute, second), nil
}
