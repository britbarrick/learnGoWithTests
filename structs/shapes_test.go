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
		shape    Shape
		expected float64
	}{
		{Rectangle{5.0, 10.0}, 50.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		actual := tt.shape.Area()
		if actual != tt.expected {
			t.Errorf("Expected %.2f, but got %.2f.", tt.expected, actual)
		}
	}

}
