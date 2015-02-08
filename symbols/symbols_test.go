package symbols

import "testing"

func TestProper(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"f", "Fahrenheit"},
		{"c", "Celsius"},
		{"guk", "imperial gallon"},
	}
	for _, c := range cases {
		got := Proper(c.in)
		if got != c.want {
			t.Errorf("Proper(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGlyph(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"f", "°F"},
		{"c", "°C"},
		{"guk", "gal"},
	}
	for _, c := range cases {
		got := Glyph(c.in)
		if got != c.want {
			t.Errorf("Glyph(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
