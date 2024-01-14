package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Guilherme")
		want := "Hello, Guilherme"

		assertCorrectMessage(t, got, want)
	})
	t.Run("empty strings defaults to 'World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
