package sorts

/*
As you proceed through the elements of an array you continuously
swap the target back with it's neighbor until it is in
the correct place. When it's moved, move your index up.
*/
func InsertionSort(arr []int) []int {
	i := 1
	var j int
	for i < len(arr) {
		j = i
		for j >= 1 {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
			j--
		}
		i++
	}
	return arr
}
