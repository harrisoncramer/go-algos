package setsandsubsets

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func IsSubset(first, second []Flight) bool {
	if len(first) == 0 || len(second) == 0 {
		return false
	}

	flights := map[Flight]int{}
	for _, flight := range second {
		flights[flight] += 1 /* Safe to use += with non-existent key */
	}

	for _, flight := range first {
		if count, found := flights[flight]; !found { /* Combined if statement on a single line */
			return false
		} else {
			if count == 0 {
				return false
			} else {
				flights[flight] -= 1
			}

		}
	}
	return true
}
