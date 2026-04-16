package process

import "testing"

func TestGetMode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "silent", input: "STANDARD SILENT", want: "SILENT"},
		{name: "speed", input: "STRUCTURAL SPEED", want: "PERFORMANCE"},
		{name: "default", input: "STANDARD", want: "NORMAL"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := getMode(tc.input)
			if got != tc.want {
				t.Fatalf("expected %q, got %q", tc.want, got)
			}
		})
	}
}

func TestMinSpeed(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{name: "numeric lower a", a: "120", b: "200", want: "120"},
		{name: "numeric lower b", a: "300", b: "200", want: "200"},
		{name: "percent a", a: "80%", b: "200", want: "80%"},
		{name: "percent b", a: "200", b: "80%", want: "80%"},
		{name: "invalid a", a: "x", b: "200", want: "x"},
		{name: "invalid b", a: "200", b: "x", want: "x"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := minSpeed(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("expected %q, got %q", tc.want, got)
			}
		})
	}
}

func TestAvoidNoisySpeeds(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "percent passthrough", input: "150%", want: "150%", wantErr: false},
		{name: "safe speed", input: "80", want: "80", wantErr: false},
		{name: "first noisy range", input: "30", want: "24", wantErr: false},
		{name: "second noisy range", input: "120", want: "94", wantErr: false},
		{name: "invalid speed", input: "abc", want: "abc", wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := avoidNoisySpeeds(tc.input)
			if (err != nil) != tc.wantErr {
				t.Fatalf("expected error=%v, got error=%v", tc.wantErr, err != nil)
			}
			if got != tc.want {
				t.Fatalf("expected %q, got %q", tc.want, got)
			}
		})
	}
}
