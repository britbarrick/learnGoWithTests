package arraysandslices

// Sum takes an array of 5 integers and returns the sum
func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// SumAllTails takes a varying number of slices and returns a new
// slice containing the totals for the tail of each slice passed in.
// A tail is every item apart from the first one (the "head")
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
