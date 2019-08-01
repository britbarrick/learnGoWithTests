package iteration

import "testing"

func TestRepeater(t *testing.T) {
	actual := Repeater("e", 3)
	expected := "eee"

	if actual != expected {
		t.Errorf("Expected %q, instead got %q.", expected, actual)
	}
}

func BenchmarkRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeater("e,", 5)
	}
}
