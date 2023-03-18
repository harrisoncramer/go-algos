package triangularnums

func TriangularRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return n + TriangularRecursive(n-1)
}

func Triangular(n int) int {
	result := 0
	for n > 0 {
		result += n
		n -= 1
	}

	return result
}

/* Imagine that each triangle forms half of a rectangle... */
func TriangularShortcut(n int) int {
	return (n * (n + 1)) / 2
}
