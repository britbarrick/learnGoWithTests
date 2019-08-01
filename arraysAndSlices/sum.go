package arraysandslices

// Sum takes an array of 5 integers and returns the sum
func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// SumAll takes a varying number of slices and returns a new
// slice containing the totals for each slice passed in.
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNums := len(numbersToSum)
	sums := make([]int, lengthOfNums)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
