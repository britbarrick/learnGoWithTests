package arraysandslices

// Sum takes an array of 5 integers and returns the sum
func Sum(numbers [5]int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
