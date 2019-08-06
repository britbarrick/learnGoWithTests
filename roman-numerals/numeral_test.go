package numerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Name     string
		Arabic   int
		Expected string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			actual := ConvertToRoman(test.Arabic)
			if actual != test.Expected {
				t.Errorf("Expected %q but got %q", test.Expected, actual)
			}
		})
	}
}
