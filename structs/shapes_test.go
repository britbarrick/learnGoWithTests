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

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{5.0, 10.0}
		actual := rectangle.Area()
		expected := 50.0

		if actual != expected {
			t.Errorf("Expected %.2v, but got %.2v.", expected, actual)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		actual := circle.Area()
		expected := 314.1592653589793

		if actual != expected {
			t.Errorf("Expected %.2f, but got %.2f.", expected, actual)
		}
	})

}
