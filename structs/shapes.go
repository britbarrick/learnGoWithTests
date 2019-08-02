package structs

// Perimeter takes width and height of a rectangle as float64 and
// returns the perimeter.
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

// Area takes the width and height of a rectangle as float64 and
// returns the area.
func Area(width, height float64) float64 {
	return width * height
}
