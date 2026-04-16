package process

import "testing"

func TestWithInheritsMergesParentChain(t *testing.T) {
	original := systemProcessesRaw
	systemProcessesRaw = map[string]map[string]any{
		"Parent": {
			"name":         "Parent",
			"inherits":     "GrandParent",
			"travel_speed": "200",
			"bridge_speed": "50",
		},
		"GrandParent": {
			"name":                "GrandParent",
			"top_surface_pattern": "monotonicline",
		},
	}
	t.Cleanup(func() {
		systemProcessesRaw = original
	})

	p := Process{
		Name:        "Child",
		Inherits:    "Parent",
		TravelSpeed: "300",
		From:        "User",
	}

	if err := withInherits(&p); err != nil {
		t.Fatalf("withInherits returned error: %v", err)
	}

	if p.TravelSpeed != "300" {
		t.Fatalf("expected child travel_speed to win, got %q", p.TravelSpeed)
	}

	if p.BridgeSpeed != "50" {
		t.Fatalf("expected inherited bridge_speed=50, got %q", p.BridgeSpeed)
	}

	if p.TopSurfacePattern != "monotonicline" {
		t.Fatalf("expected inherited top_surface_pattern=monotonicline, got %q", p.TopSurfacePattern)
	}
}

func TestWithInheritsReturnsErrorWhenParentMissing(t *testing.T) {
	original := systemProcessesRaw
	systemProcessesRaw = map[string]map[string]any{}
	t.Cleanup(func() {
		systemProcessesRaw = original
	})

	p := Process{
		Name:     "Child",
		Inherits: "Missing",
	}

	if err := withInherits(&p); err == nil {
		t.Fatal("expected error for missing parent, got nil")
	}
}
