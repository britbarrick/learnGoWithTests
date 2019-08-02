package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		actual, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertStrings(t, actual, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, actual := dictionary.Search("unknown")

		assertError(t, actual, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	expected := "this is just a test"
	actual, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if expected != actual {
		t.Errorf("Expected %q, but got %q.", expected, actual)
	}
}

func assertStrings(t *testing.T, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Given: %q, expected %q, got %q.", "test", expected, actual)
	}
}

func assertError(t *testing.T, actual, expected error) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
