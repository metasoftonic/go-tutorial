package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := SayHello("Chris", "")
		expected := "Hello, Chris"
		assertCorrectMessage(t, got, expected)

	})
	t.Run("Saying hello to with empty string", func(t *testing.T) {
		got := SayHello("", "")
		expected := "Hello, world"
		assertCorrectMessage(t, got, expected)
	})
	t.Run("In spanish", func(t *testing.T) {
		got := SayHello("Eldoie", "Spanish")
		expected := "Hola, Eldoie"
		assertCorrectMessage(t, got, expected)
	})
}
func assertCorrectMessage(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
