package iteration

import "testing"

func TestRepeater(t *testing.T) {
	actual := Repeater("e")
	expected := "eeeee"

	if actual != expected {
		t.Errorf("Expected %q, instead got %q.", expected, actual)
	}
}
