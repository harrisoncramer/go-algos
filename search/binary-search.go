package search

/*
Binary search works by finding the middle of the (sorted) input
array, and comparing it against the target. Depending on
whether the target is smaller or larger, the left or right
pointer is moved, shrinking the size of the searchable area.
found, or not.
*/

func BinarySearch(arr []int, target int) int {

	if len(arr) == 0 {
		return -1
	}

	left := 0
	right := len(arr) - 1
	var mid int

	for left <= right {
		mid = (left + right) / 2
		if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			return mid
		}

	}

	return -1
}
