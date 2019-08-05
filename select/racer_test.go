package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.uk"

	expected := fastURL
	actual := Racer(slowURL, fastURL)

	if expected != actual {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
