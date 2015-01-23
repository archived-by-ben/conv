package units

import "testing"

func TestInfo(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"nam", "Celsius"},
		{"sym", "°C"},
	}
	for _, c := range cases {
		got := Info("c", c.in)
		if got != c.want {
			t.Errorf("Info(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
