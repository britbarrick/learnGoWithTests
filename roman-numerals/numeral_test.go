package numerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	actual := ConvertToRoman(1)
	expected := "I"

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
