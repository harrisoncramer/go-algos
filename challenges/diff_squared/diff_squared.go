package diffsquared

import "math"

func DiffSquared(a int, b int) (res int) {
	aSquared := math.Pow(float64(a), 2)
	bSquared := math.Pow(float64(b), 2)

	return int(math.Abs(aSquared - bSquared))
}
