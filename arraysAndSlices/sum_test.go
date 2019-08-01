package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	actual := Sum(numbers)
	expected := 15

	if actual != expected {
		t.Errorf("Expected %d, but got %d, %v", expected, actual, numbers)
	}

}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v.", expected, actual)
	}
}
