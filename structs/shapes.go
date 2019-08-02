package structs

import "math"

// Shape groups functionality based upon Area()
type Shape interface {
	Area() float64
}

// Rectangle gives the dimensions of the shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle gives dimensions of the shape
type Circle struct {
	Radius float64
}

// Perimeter takes width and height of a rectangle as float64 and
// returns the perimeter.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area takes the width and height of a rectangle as float64 and
// returns the area.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area takes the width and height of a circle as float64 and
// returns the area.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
