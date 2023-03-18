package doubleletters

import "testing"

type Implementations func(s string) string

/* Write a function that doubles the letters in a string */
func TestDoubleLetters(t *testing.T) {
	tests := []Implementations{DoubleLetters, DoubleLettersTwo}
	t.Run("Should double letters", func(t *testing.T) {
		input := "gopher"

		for _, Implementations := range tests {
			got := Implementations(input)
			want := "ggoopphheerr"
			if got != want {
				t.Errorf("Got %s but wanted %s", got, want)
			}
		}

	})

	t.Run("Should handle non-ASCII characters", func(t *testing.T) {
		input := "🎿"

		for _, Implementations := range tests {
			got := Implementations(input)
			want := "🎿🎿"
			if got != want {
				t.Errorf("Got %s but wanted %s", got, want)
			}
		}
	})

	t.Run("Should not double nothing", func(t *testing.T) {
		input := ""

		for _, Implementations := range tests {
			got := Implementations(input)
			want := ""
			if got != want {
				t.Errorf("Got %s but wanted %s", got, want)
			}
		}
	})

	t.Run("Should double empty string", func(t *testing.T) {
		input := " "

		for _, Implementations := range tests {
			got := Implementations(input)
			want := "  "
			if got != want {
				t.Errorf("Got %s but wanted %s", got, want)
			}
		}
	})
}
