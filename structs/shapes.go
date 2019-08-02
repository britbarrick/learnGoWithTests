package structs

// Rectangle gives the dimensions of the shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter takes width and height of a rectangle as float64 and
// returns the perimeter.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area takes the width and height of a rectangle as float64 and
// returns the area.
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
