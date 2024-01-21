package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is a dictionary",
	}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a dictionary"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFoundWord

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want.Error())
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}
