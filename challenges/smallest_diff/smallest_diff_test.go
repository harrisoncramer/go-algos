package smallestdiff

import "testing"

type Test struct {
	array1 []int
	array2 []int
	want   int
}

/*
Write a function that returns the smallest difference
between any two integers from two sets of numbers.
Return -1 if the two sets cannot be compared.
*/
func TestCalcSmallestDiff(t *testing.T) {
	t.Run("Should return the smallest difference", func(t *testing.T) {

		tests := []Test{
			{
				[]int{9, 8, 7, 6},
				[]int{7, 3, 2, 5},
				0,
			},
			{
				[]int{1, 2},
				[]int{4, 5},
				2,
			},
			{
				[]int{},
				[]int{4, 5},
				-1,
			},
			{
				[]int{4, 10},
				[]int{},
				-1,
			},
			{
				[]int{4, 10, 100},
				[]int{90},
				10,
			},
		}
		for i, test := range tests {
			got := CalcSmallestDifference(test.array1, test.array2)
			if got != test.want {
				t.Errorf("Test %d: Got %d but wanted %d", i+1, got, test.want)
			}
		}

	})
}
