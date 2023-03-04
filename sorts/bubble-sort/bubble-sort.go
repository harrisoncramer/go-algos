package sorts

/*
In a nested loop, compare the value at i to
the value at j, and swap. Continue until you've
looped through the outer loop.
*/
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
