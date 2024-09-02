package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying Hello To people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying Hello World when empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying hello in spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in French ", func(t *testing.T) {
		got := Hello("Napoleon", "French")
		want := "Bon, Napoleon"

		assertCorrectMessage(t, got, want)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // redirects error line to the point of failure rather showing here

	if got != want {
		t.Errorf("Got: %q , Wanted: %q ", got, want)

	}
}
