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

// Triangle gives dimentsions of the shape
type Triangle struct {
	Height float64
	Base   float64
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

// Area takes the radius of a circle as float64 and
// returns the area.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area takes the base and height of a triangle as a float 64
// and returns the area
func (t Triangle) Area() float64 {
	return .5 * t.Base * t.Height
}
