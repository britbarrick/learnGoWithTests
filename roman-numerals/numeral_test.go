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
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			actual := ConvertToRoman(test.Arabic)
			if actual != test.Expected {
				t.Errorf("Expected %q but got %q", test.Expected, actual)
			}
		})
	}

	t.Run("1 gets converted to I", func(t *testing.T) {
		actual := ConvertToRoman(1)
		expected := "I"

		if actual != expected {
			t.Errorf("Expected %q but got %q", expected, actual)
		}
	})

	t.Run("2 gets converted to II", func(t *testing.T) {
		actual := ConvertToRoman(2)
		expected := "II"

		if actual != expected {
			t.Errorf("Expected %q but got %q", expected, actual)
		}
	})
}
