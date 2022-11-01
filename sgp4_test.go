package sgp4_test

import (
	"testing"
	"time"

	"github.com/SharkEzz/go-sgp4"
)

const (
	tle1 = "1 33591U 09005A   22304.34359593  .00000180  00000-0  12210-3 0  9995"
	tle2 = "2 33591  99.1341 341.2512 0014233 162.9432 197.2221 14.12630118707632"
)

func TestSGP4_FindPosition(t *testing.T) {
	tle, err := sgp4.NewTle("NOAA 19", tle1, tle2)
	if err != nil {
		t.Fatal(err)
	}
	p, err := sgp4.NewSGP4(tle)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		tString string
		wantLat float64
		wantLng float64
		wantAlt float64
	}{
		{
			name:    "NOAA 19",
			tString: "2022-10-31T15:15:08Z",
			wantLat: 43.11834312758037,
			wantLng: 64.24461016754081,
			wantAlt: 861.6024972133991,
		},
		{
			name:    "NOAA 19",
			tString: "2022-11-01T15:15:08Z",
			wantLat: 80.22435624311764,
			wantLng: 4.721711641150241,
			wantAlt: 860.6844430232495,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t1, err := time.Parse(time.RFC3339, tt.tString)
			if err != nil {
				t.Fatal(err)
			}

			eci, err := p.PositionFromTime(t1)
			if err != nil {
				t.Fatal(err)
			}

			geodetic, err := eci.ToGeodetic()
			if err != nil {
				t.Fatal(err)
			}
			lat, lng, alt, err := geodetic.GetCoords(true)
			if err != nil {
				t.Fatal(err)
			}

			if lat != tt.wantLat {
				t.Fatalf("Latitude = %v, want %v", lat, tt.wantLat)
			}
			if lng != tt.wantLng {
				t.Fatalf("Longitude = %v, want %v", lng, tt.wantLng)
			}
			if alt != tt.wantAlt {
				t.Fatalf("Altitude = %v, want %v", alt, tt.wantAlt)
			}
		})
	}
}
