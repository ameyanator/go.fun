package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying namste to people", func(t *testing.T) {
		got := Hello("Ameya", "")
		want := "Namaste Ameya!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Namaste World when empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Namaste World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in English", func(t *testing.T) {
		got := Hello("Soumya", "English")
		want := "Hello Soumya!"

		assertCorrectMessage(t, got, want)
	})
}
