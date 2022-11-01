package sgp4

import "github.com/SharkEzz/go-sgp4/internal/cppsgp4"

type Tle struct {
	ctle cppsgp4.Tle
}

func NewTle(name, line1, line2 string) (t *Tle, err error) {
	defer catch(&err)

	return &Tle{cppsgp4.NewTle(name, line1, line2)}, err
}

func (t *Tle) Name() string {
	return t.ctle.Name()
}

func (t *Tle) Line1() string {
	return t.ctle.Line1()
}

func (t *Tle) Line2() string {
	return t.ctle.Line2()
}

func (t *Tle) Epoch() *DateTime {
	return &DateTime{t.ctle.Epoch()}
}

func (t *Tle) Inclination(in_degrees bool) float64 {
	return t.ctle.Inclination(in_degrees)
}

func (t *Tle) RightAscensionNode(in_degrees bool) float64 {
	return t.ctle.RightAscendingNode(in_degrees)
}

func (t *Tle) Eccentricity() float64 {
	return t.ctle.Eccentricity()
}

func (t *Tle) ArgumentPerigee(in_degrees bool) float64 {
	return t.ctle.ArgumentPerigee(in_degrees)
}

func (t *Tle) MeanAnomaly(in_degrees bool) float64 {
	return t.ctle.MeanAnomaly(in_degrees)
}

func (t *Tle) MeanMotion() float64 {
	return t.ctle.MeanMotion()
}

func (t *Tle) BStar() float64 {
	return t.ctle.BStar()
}
