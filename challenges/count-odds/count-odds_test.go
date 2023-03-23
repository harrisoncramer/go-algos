package countodds

import "testing"

type Test struct {
	Low  int
	High int
	Want int
}

/*
Write a function that returns the number of odds between
a high and low value, inclusive of those two numbers
*/
func TestCountOdds(t *testing.T) {

	tests := []Test{
		{3, 7, 3},
		{1, 5, 3},
		{0, 4, 2},
		{1, 10, 5},
		{0, 1, 1},
		// {1, 1, 1},
		{0, 0, 0},
		{2, 5, 2},
	}

	for _, test := range tests {
		got := CountOdds(test.Low, test.High)
		if got != test.Want {
			t.Errorf("Got %d but wanted %d for high of %d and low of %d", got, test.Want, test.High, test.Low)
		}

	}

}
