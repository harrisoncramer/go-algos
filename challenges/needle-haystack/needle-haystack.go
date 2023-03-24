package needlehaystack

import "strings"

func NeedleHaystack(n, h string) int {
	if len(n) == 0 || len(h) == 0 {
		return -1
	}

	var (
		p1    int
		p2    int
		start int
	)

	for p2 < len(n) {

		if p1 == len(h) {
			return -1
		}

		if h[p1] == n[p2] {
			p1++
			p2++
		} else {
			p2 = 0
			p1 = start + 1
			start += 1
		}
	}

	if p2 == len(n) {
		return start
	} else {
		return -1
	}
}

func NeedleHaystackTwo(n, h string) int {
	if len(n) == 0 || len(h) == 0 {
		return -1
	}
	return strings.Index(h, n)
}

func NeedleHaystackThree(needle, haystack string) int {
	h := len(haystack)
	n := len(needle)
	if n == 0 || h == 0 {
		return -1
	}

	/* Check to see if the needle matches a chunk from the haystack,
	if not, increment by one. Do this until you reach the index where
	this slice is impossible because it's too long: len(haystack) - len(needle) + 1 */
	for i := 0; i < h-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
