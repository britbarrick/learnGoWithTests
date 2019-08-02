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

	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Area()
		if actual != expected {
			t.Errorf("Expected %.2f, but got %.2f.", expected, actual)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{5.0, 10.0}
		checkArea(t, rectangle, 50.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}
