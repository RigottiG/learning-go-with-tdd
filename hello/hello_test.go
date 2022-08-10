package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to the people", func(t *testing.T) {
		got := Hello("Rigotti", "")
		want := "Hello, Rigotti"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Rigotti", "Portuguese")
		want := "Olá, Rigotti"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Rigotti", "French")
		want := "Bonjour, Rigotti"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
