package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	actual := dictionary.Search("test")
	expected := "this is just a test"

	assertStrings(t, actual, expected)
}

func assertStrings(t *testing.T, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Given: %q, expected %q, got %q.", "test", expected, actual)
	}
}
