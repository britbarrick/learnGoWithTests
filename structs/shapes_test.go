package structs

import "testing"

func TestPerimeter(t *testing.T) {
	actual := Perimeter(10.0, 10.0)
	expected := 40.0

	if actual != expected {
		t.Errorf("Expected %.2f, but got %.2f.", expected, actual)
	}
}
