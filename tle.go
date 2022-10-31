package sgp4

import "github.com/SharkEzz/sgp4/cppsgp4"

func NewTle(name, line1, line2 string) (t cppsgp4.Tle, err error) {
	defer catch(&err)

	return cppsgp4.NewTle(name, line1, line2), err
}
