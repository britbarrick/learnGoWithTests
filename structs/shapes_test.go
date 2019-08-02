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

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{5.0, 10.0}, expected: 50.0},
		{name: "Circle", shape: Circle{10.0}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12.0, 6.0}, expected: 36.0},
	}

	for _, tt := range areaTests {
		actual := tt.shape.Area()
		if actual != tt.expected {
			t.Errorf("Expected %.2f, but got %.2f.", tt.expected, actual)
		}
	}

}
