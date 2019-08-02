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
		_, err := dictionary.Search("unknown")
		expected := "cannot find the word you're looking for"

		if err == nil {
			t.Fatal("expected an error.")
		}

		assertStrings(t, err.Error(), expected)
	})

}

func assertStrings(t *testing.T, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Given: %q, expected %q, got %q.", "test", expected, actual)
	}
}
