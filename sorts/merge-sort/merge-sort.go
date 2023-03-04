package sorts

/*
Merge sort is a divide and conquer algorithm. First, recursively
divide the array until you get an array of just 1 length.
Implement a function "merge" that merges two sorted arrays.
Then merge these arrays and return them.

This algorithm requires multiple recursive calls at the same time,
which each sort one half of the array at a time.
*/

func Merge(a []int, b []int) []int {
	res := []int{}
	i, j := 0, 0

	/* Compare, adding lower value from either array */
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	/* Append the remainder */
	if i < len(a) {
		res = append(res, a[i:]...)
	} else {
		res = append(res, b[j:]...)
	}
	return res
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	/* Using the same index splits the array evenly
	without losing any values */
	a := mergeSort(arr[0 : len(arr)/2])
	b := mergeSort(arr[len(arr)/2:])
	return Merge(a, b)
}
