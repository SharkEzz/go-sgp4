package sgp4_test

import (
	"testing"

	"github.com/SharkEzz/sgp4"
)

func TestSGP4_Tle(t *testing.T) {
	tle, err := sgp4.NewTle("ISS (ZARYA)", "1 25544U 98067A   22304.51317130  .00016748  00000-0  30494-3 0  9991", "2 25544  51.6438  21.2332 0006312  23.9726 155.8350 15.49597428366324")
	if err != nil {
		t.Fatal(err)
	}

	if tle.Name() != "ISS (ZARYA)" {
		t.Fatal("tle.Name() != \"ISS (ZARYA)\"")
	}

	if tle.Line1() != "1 25544U 98067A   22304.51317130  .00016748  00000-0  30494-3 0  9991" {
		t.Fatal("tle.Line1() != \"1 25544U 98067A   22304.51317130  .00016748  00000-0  30494-3 0  9991\"")
	}

	if tle.Line2() != "2 25544  51.6438  21.2332 0006312  23.9726 155.8350 15.49597428366324" {
		t.Fatal("tle.Line2() != \"2 25544  51.6438  21.2332 0006312  23.9726 155.8350 15.49597428366324\"")
	}

	if tle.Epoch().Year() != 2022 {
		t.Fatal("tle.Epoch().Year() != 2022")
	}

	if tle.Epoch().Month() != 10 {
		t.Fatal("tle.Epoch().Month() != 10")
	}

	if tle.Epoch().Day() != 31 {
		t.Fatal("tle.Epoch().Day() != 4")
	}
}
