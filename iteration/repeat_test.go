package iteration

import (
	"fmt"
	"testing"
)

func TestRepeater(t *testing.T) {
	actual := Repeater("e", 12)
	expected := "eeeeeeeeeeee"

	if actual != expected {
		t.Errorf("Expected %q, instead got %q.", expected, actual)
	}
}

func BenchmarkRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeater("e,", 5)
	}
}

func ExampleRepeater() {
	repeat := Repeater("f", 11)
	fmt.Println(repeat)
	// Output: fffffffffff
}
