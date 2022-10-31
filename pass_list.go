package sgp4

import (
	"time"

	"github.com/SharkEzz/sgp4/cppsgp4"
)

func GeneratePassList(lat, lng, alt float64, sgp4 cppsgp4.SGP4, start, end time.Time, step int) (p []cppsgp4.PassDetails, err error) {
	defer catch(&err)

	startUtc := start.UTC()
	endUtc := end.UTC()

	startDt := cppsgp4.NewDateTime(startUtc.Year(), int(startUtc.Month()), startUtc.Day(), startUtc.Hour(), startUtc.Minute(), startUtc.Second())
	endDt := cppsgp4.NewDateTime(endUtc.Year(), int(endUtc.Month()), endUtc.Day(), endUtc.Hour(), endUtc.Minute(), endUtc.Second())

	coords, err := NewCoordGeodetic(lat, lng, alt)
	if err != nil {
		return nil, err
	}

	passList := cppsgp4.GeneratePassList(coords, sgp4, startDt, endDt, step)

	passListSlice := make([]cppsgp4.PassDetails, passList.Size())

	for i := int64(0); i < passList.Size(); i++ {
		passListSlice[i] = passList.Get(int(i))
	}

	return passListSlice, nil
}
