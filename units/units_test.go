package units

import "testing"

func TestInfo(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"nam", "Celsius"},
		{"arg", "C"},
		{"sym", "°C"},
		{"qty", "temperature"},
	}
	for _, c := range cases {
		got := Info("cel", c.in)
		if got != c.want {
			t.Errorf("Info(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
