package common

import "testing"

func TestFahrenheit(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{100, 310.9277777777778},
		{0, 255.3722222222222},
		{-540, -44.627777777777766},
	}
	for _, c := range cases {
		got := Fahrenheit(c.in)
		if got != c.want {
			t.Errorf("Fahrenheit(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCelsius(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{300, 573.15},
		{0, 273.15},
		{-494, -220.85000000000002},
	}
	for _, c := range cases {
		got := Celsius(c.in)
		if got != c.want {
			t.Errorf("Celsius(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestHorsepower(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 746},
		{1000.1, 746074.6},
	}
	for _, c := range cases {
		got := Horsepower(c.in)
		if got != c.want {
			t.Errorf("Horsepower(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestWatt(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 1},
	}
	for _, c := range cases {
		got := Watt(c.in)
		if got != c.want {
			t.Errorf("Watt(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestKmh(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{100, 27.777777777778002},
	}
	for _, c := range cases {
		got := Kmh(c.in)
		if got != c.want {
			t.Errorf("Kmh(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMph(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{100, 44.704},
	}
	for _, c := range cases {
		got := Mph(c.in)
		if got != c.want {
			t.Errorf("Mph(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestKn(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{100, 51.4444},
	}
	for _, c := range cases {
		got := Kn(c.in)
		if got != c.want {
			t.Errorf("Kn(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMps(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 1},
	}
	for _, c := range cases {
		got := Mps(c.in)
		if got != c.want {
			t.Errorf("Mps(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCarat(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{10, 2},
	}
	for _, c := range cases {
		got := Carat(c.in)
		if got != c.want {
			t.Errorf("Carat(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestOunce(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 28},
	}
	for _, c := range cases {
		got := Ounce(c.in)
		if got != c.want {
			t.Errorf("Ounce(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestPound(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 453.59237},
	}
	for _, c := range cases {
		got := Pound(c.in)
		if got != c.want {
			t.Errorf("Pound(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestStone(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 6350.29318},
	}
	for _, c := range cases {
		got := Stone(c.in)
		if got != c.want {
			t.Errorf("Stone(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCentimetre(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 0.01},
		{10.5, 0.105},
		{100, 1},
	}
	for _, c := range cases {
		got := Centimetre(c.in)
		if got != c.want {
			t.Errorf("Centimetre(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestInch(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 0.0254},
		{100, 2.54},
		{100.5, 2.5526999999999997},
	}
	for _, c := range cases {
		got := Inch(c.in)
		if got != c.want {
			t.Errorf("Inch(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestYard(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 0.9144},
		{10, 9.144},
		{11.1, 10.14984},
	}
	for _, c := range cases {
		got :=
			Yard(c.in)
		if got != c.want {
			t.Errorf("Yard(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFoot(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 0.3048},
		{10, 3.048},
		{11.1, 3.38328},
	}
	for _, c := range cases {
		got :=
			Foot(c.in)
		if got != c.want {
			t.Errorf("Foot(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMetre(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 1},
		{10, 10},
	}
	for _, c := range cases {
		got :=
			Metre(c.in)
		if got != c.want {
			t.Errorf("Metre(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestKilometre(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 1000},
		{5.5, 5500},
	}
	for _, c := range cases {
		got :=
			Kilometre(c.in)
		if got != c.want {
			t.Errorf("Kilometre(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMile(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 1609.344},
		{1.1, 1770.2784000000001},
		{100, 160934.4},
	}
	for _, c := range cases {
		got :=
			Mile(c.in)
		if got != c.want {
			t.Errorf("Mile(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNautical(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1, 1852},
		{1.1, 2037.2000000000003},
	}
	for _, c := range cases {
		got :=
			Nautical(c.in)
		if got != c.want {
			t.Errorf("Nautical(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
