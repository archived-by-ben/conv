package si

import "testing"

func TestRound(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{100, "100"},
		{1.1, "1.1"},
		{1.25, "1.2"},
		{81.0, "81"},
		{81.000000000001, "81"},
		{81.01, "81"},
		{839.87365, "839.9"},
		{-839.87365, "-839.9"},
		{1.39700, "1.4"},
	}
	for _, c := range cases {
		got := Round(c.in)
		if got != c.want {
			t.Errorf("Round(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFahrenheit(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{100, -279.67},
		{0, -459.67},
		{-44, -538.87},
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
		{300, 26.850000000000023},
		{0, -273.15},
		{-44, -317.15},
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
		{746, 1},
		{1000.1, 1.3406166219839142},
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
		{100.1, 100.1},
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
		{10, 35.99999999999971},
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
		{10, 22.369362920544024},
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
		{10, 19.438461717893492},
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
		{2, 10},
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
		{28, 1},
		{500, 17.857142857142858},
		{222.222, 7.9365000000000006},
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
		{1, 0.002204622621848776},
		{500, 1.1023113109243878},
		{444.4, 0.9797342931495959},
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
		{1, 0.00015747304441776972},
		{11111, 1.7496829965258391},
		{5555.5, 0.8748414982629196},
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
		{1, 100},
		{10.5, 1050},
		{100, 10000},
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
		{1, 39.37007874015748},
		{100, 3937.0078740157483},
		{100.5, 3956.692913385827},
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
		{1, 1.0936132983377078},
		{11.1, 12.139107611548557},
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
		{1, 3.280839895013123},
		{11.1, 36.41732283464567},
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
		{1.1, 1.1},
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
		{1.1, 0.0011},
		{1001, 1.001},
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
		{1, 0.0006213711922373339},
		{1.1, 0.0006835083114610674},
		{1609.344, 1},
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
		{1.1, 0.0005939524838012959},
		{15000, 8.099352051835853},
	}
	for _, c := range cases {
		got :=
			Nautical(c.in)
		if got != c.want {
			t.Errorf("Nautical(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
