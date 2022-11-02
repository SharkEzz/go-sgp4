package sgp4

import "github.com/SharkEzz/go-sgp4/internal/cppsgp4"

type Tle struct {
	_tle cppsgp4.Tle
}

func NewTle(name, line1, line2 string) (t *Tle, err error) {
	defer catch(&err)

	return &Tle{cppsgp4.NewTle(name, line1, line2)}, err
}

func (t *Tle) Name() string {
	return t._tle.Name()
}

func (t *Tle) Line1() string {
	return t._tle.Line1()
}

func (t *Tle) Line2() string {
	return t._tle.Line2()
}

func (t *Tle) Epoch() *DateTime {
	return &DateTime{t._tle.Epoch()}
}

func (t *Tle) Inclination(in_degrees bool) float64 {
	return t._tle.Inclination(in_degrees)
}

func (t *Tle) RightAscensionNode(in_degrees bool) float64 {
	return t._tle.RightAscendingNode(in_degrees)
}

func (t *Tle) Eccentricity() float64 {
	return t._tle.Eccentricity()
}

func (t *Tle) ArgumentPerigee(in_degrees bool) float64 {
	return t._tle.ArgumentPerigee(in_degrees)
}

func (t *Tle) MeanAnomaly(in_degrees bool) float64 {
	return t._tle.MeanAnomaly(in_degrees)
}

func (t *Tle) MeanMotion() float64 {
	return t._tle.MeanMotion()
}

func (t *Tle) BStar() float64 {
	return t._tle.BStar()
}

func (t *Tle) Close() {
	cppsgp4.DeleteTle(t._tle)
}
