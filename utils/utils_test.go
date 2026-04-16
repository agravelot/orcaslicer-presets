package utils

import (
	"math"
	"testing"
)

func TestGetLayerHeight(t *testing.T) {
	v := GetLayerHeight("0.20mm Standard @Voron")
	if math.Abs(v-0.20) > 1e-6 {
		t.Fatalf("expected 0.20, got %f", v)
	}
}

func TestGetNozzleSize(t *testing.T) {
	tests := []struct {
		name    string
		inherit string
		want    float64
	}{
		{name: "default 0.4", inherit: "0.20mm Standard @Voron", want: 0.4},
		{name: "0.6", inherit: "0.24mm Optimal 0.6 nozzle @Voron", want: 0.6},
		{name: "0.8", inherit: "0.40mm Standard 0.8 nozzle @Voron", want: 0.8},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetNozzleSize(tc.inherit)
			if math.Abs(got-tc.want) > 1e-6 {
				t.Fatalf("expected %f, got %f", tc.want, got)
			}
		})
	}
}

func TestEllipticalExtrusionRate(t *testing.T) {
	got := EllipticalExtrusionRate(0.4, 0.2, 100)
	want := math.Pi * (0.4 / 2) * (0.2 / 2) * 100
	if math.Abs(got-want) > 1e-9 {
		t.Fatalf("expected %f, got %f", want, got)
	}
}
