package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {

	check := func(t testing.TB, got, want []string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("seperated by comma", func(t *testing.T) {
		got := Split("a,b,c", ",")
		want := []string{"a", "b", "c"}

		check(t, got, want)
	})

	t.Run("separeated by empty string", func(t *testing.T) {
		got := Split(" xyz ", "")
		want := []string{" ", "x", "y", "z", " "}

		check(t, got, want)
	})

	t.Run("seperated by chars", func(t *testing.T) {
		got := Split("a man a plan a canal panama", "a ")
		want := []string{"", "man ", "plan ", "canal panama"}

		check(t, got, want)
	})

	t.Run("split by a char/string not persent in string", func(t *testing.T) {
		got := Split("", "Bernardo O'Higgins")
		want := []string{""}

		check(t, got, want)
	})

	/* If the separator is an empty string, strings.Split splits after each UTF-8 sequence.
	// Therefore, an empty string will be split into its individual UTF-8 characters,
	// which in the case of an empty string will result in an empty slice []string{}. */
	t.Run("split by a empty str not persent in string", func(t *testing.T) {
		got := Split("", "")
		want := []string{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
