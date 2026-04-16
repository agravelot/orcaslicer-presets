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

func TestWithInheritsRecordsMetadata(t *testing.T) {
	original := systemProcessesRaw
	systemProcessesRaw = map[string]map[string]any{
		"Parent": {
			"name":                "Parent",
			"inherits":            "GrandParent",
			"bridge_speed":        "50",
			"top_surface_pattern": "monotonicline",
		},
		"GrandParent": {
			"name":                      "GrandParent",
			"support_interface_spacing": "0.2",
		},
	}
	t.Cleanup(func() {
		systemProcessesRaw = original
	})

	p := Process{
		Name:        "Child",
		Inherits:    "Parent",
		TravelSpeed: "300", // Code-defined, should not appear in InheritedFrom
		From:        "User",
	}

	if err := withInherits(&p); err != nil {
		t.Fatalf("withInherits returned error: %v", err)
	}

	// Inherited keys should have metadata
	if p.InheritedFrom == nil {
		t.Fatal("expected InheritedFrom to be populated")
	}

	if source, ok := p.InheritedFrom["bridge_speed"]; !ok || source != "Parent" {
		t.Fatalf("expected bridge_speed inherited from Parent, got %q", source)
	}

	if source, ok := p.InheritedFrom["top_surface_pattern"]; !ok || source != "Parent" {
		t.Fatalf("expected top_surface_pattern inherited from Parent, got %q", source)
	}

	if source, ok := p.InheritedFrom["support_interface_spacing"]; !ok || source != "GrandParent" {
		t.Fatalf("expected support_interface_spacing inherited from GrandParent, got %q", source)
	}

	// Child-defined key should not be in InheritedFrom
	if _, ok := p.InheritedFrom["travel_speed"]; ok {
		t.Fatal("expected travel_speed (child-defined) to not be in InheritedFrom")
	}

	// Excluded keys should not be present
	if _, ok := p.InheritedFrom["name"]; ok {
		t.Fatal("expected name to be excluded from InheritedFrom")
	}
}
