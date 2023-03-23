package averagesalary

import (
	"sort"
)

func Average(salary []int) (result float64) {
	sort.Sort(sort.IntSlice(salary))

	if len(salary) <= 2 {
		return -1
	}

	noOutliers := salary[1 : len(salary)-1]
	res := 0
	for _, val := range noOutliers {
		res += val
	}

	return float64(float64(res) / float64(len(noOutliers)))
}

/* 
This solution is better because it only requires a single
loop over the input array.
*/
func AverageTwo(salary []int) float64 {
	var maxNum, minNum = salary[0], salary[0]
	var sumArr int

	for _, val := range salary {
		sumArr += val

		if val > maxNum {
			maxNum = val
		}

		if val < minNum {
			minNum = val
		}
	}

	return float64(sumArr-minNum-maxNum) / float64(len(salary)-2)
}
