package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	actual := Perimeter(rectangle)
	expected := 40.0

	if actual != expected {
		t.Errorf("Expected %.2f, but got %.2f.", expected, actual)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{5.0, 10.0}
	actual := Area(rectangle)
	expected := 50.0

	if actual != expected {
		t.Errorf("Expected %.2v, but got %.2v.", expected, actual)
	}
}
