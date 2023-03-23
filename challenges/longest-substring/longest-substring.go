package challenges

func LongestSubstring(str string) int {
	myMap, max, left := make(map[rune]int), 0, 0
	for idx, myRune := range str {
		if _, okay := myMap[myRune]; okay == true && myMap[myRune] >= left {
			/* If it's in the map and the index is greater than the left index, then it's a duplicate in the current span of characters */
			/* Then, set the new max (if the current string is larger) and move our left pointer to the right of the repeating character */
			if idx-left > max {
				max = idx - left
			}
			left = myMap[myRune] + 1
		}
		myMap[myRune] = idx
	}

	/* Do one last check after the loop to see if the last string is the longest */
	if len(str)-left > max {
		max = len(str) - left
	}
	return max
}
