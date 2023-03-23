package productsumdiff

import "strconv"

func ProductSumDiff(input int) int {
	product, sum := 1, 0

	ascii := strconv.Itoa(input)
	for _, r := range ascii {
		integer, _ := strconv.Atoi(string(r))
		product *= integer
		sum += integer
	}

	return product - sum
}
