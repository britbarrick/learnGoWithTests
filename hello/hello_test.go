package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, actual, expected string) {
		t.Helper()
		if actual != expected {
			t.Errorf("Got %q, but expected %q.", actual, expected)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello("Chris", "")
		expected := "Hello, Chris"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, World"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		actual := Hello("Elodie", "Spanish")
		expected := "Hola, Elodie"
		assertCorrectMessage(t, actual, expected)
	})
}
