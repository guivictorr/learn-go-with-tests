package racer

import "testing"

func TestRacer(t *testing.T) {
	slowUrl := "https://facebook.com"
	fastUrl := "https://guilhermevictor.space"

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
