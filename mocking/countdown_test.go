package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	actual := buffer.String()
	expected := `3
2
1
Go!`

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
