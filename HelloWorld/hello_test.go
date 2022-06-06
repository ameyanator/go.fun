package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ameya")
	want := "Namaste Ameya!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}