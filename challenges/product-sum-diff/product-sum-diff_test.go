package productsumdiff

import "testing"

/*
Given an integer number n, return the difference between
the product of its digits and the sum of its digits.
234
*/

type Test struct {
	Input int
	Want  int
}

func TestProductSumDiff(t *testing.T) {

	tests := []Test{
		{234, 15},
		{4421, 21},
	}
	for _, test := range tests {
		got := ProductSumDiff(test.Input)
		if got != test.Want {
			t.Errorf("Got %d but wanted %d", got, test.Want)
		}

	}

}
