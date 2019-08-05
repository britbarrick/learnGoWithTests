package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	actual := buffer.String()
	expected := `3
2
1
Go!`

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, expected 4 but got %d", spySleeper.Calls)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
