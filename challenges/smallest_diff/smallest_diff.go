package smallestdiff

import (
	"math"
	"sort"
)

func CalcSmallestDifference(arr1, arr2 []int) int {

	if len(arr1) == 0 || len(arr2) == 0 {
		return -1
	}

	sort.Ints(arr1) /* We can use sort.Ints() to sort an integer array without creating the Swap(), Len(), and Less() methods */
	sort.Ints(arr2)

	p1, p2 := 0, 0
	minDiff := math.Abs(float64(arr1[p2]) - float64(arr2[p1]))
	for p1 < len(arr1) && p2 < len(arr2) {
		diff := float64(arr1[p1]) - float64(arr2[p2])
		if math.Abs(diff) < math.Abs(minDiff) {
			minDiff = math.Abs(diff)
		}

		if minDiff == 0 {
			return 0
		} else if diff < 0 {
			p1++
		} else {
			p2++
		}
	}

	return int(math.Abs(minDiff))
}
