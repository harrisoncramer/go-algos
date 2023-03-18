package triangularnums

import "testing"

type Implementations []func(i int) int

type Test struct {
	input int
	want  int
}

type Tests []Test

/* Write a function that, when given value n, returns the nth triangular number */
func TestTriangular(t *testing.T) {

	tests := Tests{
		{3, 6},
		{0, 0},
		{1, 1},
		{6, 21},
	}

	Implementations := Implementations{Triangular, TriangularRecursive, TriangularShortcut}
	t.Run("Should return the nth triangular number", func(t *testing.T) {
		for _, f := range Implementations {
			for _, test := range tests {
				got := f(test.input)
				want := test.want
				if got != want {
					t.Errorf("Got %d but wanted %d", got, want)
				}
			}
		}

	})
}
