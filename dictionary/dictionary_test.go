package dictionary

import "testing"

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "just a test"}

	t.Run("key in dict", func(t *testing.T) {
		got, _ := dictionary.Search("test")

		want := "just a test"

		assertStrings(t, got, want)
	})

	t.Run("key not in dict", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		want := "could not find the word"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}
