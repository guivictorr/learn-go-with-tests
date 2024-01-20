package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{
		"test": "this is a dictionary",
	}
	got := Search(dictionary, "test")
	want := "this is a dictionary"

	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}
