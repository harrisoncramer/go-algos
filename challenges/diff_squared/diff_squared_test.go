package diffsquared

import "testing"

type Test struct {
	a    int
	b    int
	want int
}

/*
Calculate the difference between two numbers squared.
*/
func TestDiffSquared(t *testing.T) {
	tests := []Test{
		{2, 4, 12},
		{1, 5, 24},
		{0, 3, 9},
		{-1, -2, 3},
		{1, -2, 3},
		{0, -2, 4},
		{5, 4, 9},
	}
	t.Run("Should return the difference between two squares", func(t *testing.T) {
		for _, test := range tests {
			got := DiffSquared(test.a, test.b)
			want := test.want

			if got != want {
				t.Errorf("Got %d but wanted %d", got, want)
			}
		}
	})
}
