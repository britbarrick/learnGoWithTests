package main

import "testing"

func TestHello(t *testing.T) {
	actual := Hello("Chris")
	expected := "Hello, Chris"

	if actual != expected {
		t.Errorf("Got %q, but expected %q.", actual, expected)
	}
}
