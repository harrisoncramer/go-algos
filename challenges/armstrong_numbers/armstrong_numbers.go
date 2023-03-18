package armstrongnumbers

import (
	"math"
	"strconv"
)

type MyInt int

/* The Atoi (ASCII to integer) and Itoa (Integer to ASCII) are used to convert numbers to characters */
func (n MyInt) IsArmstrong() (bool, error) {

	remainder := float64(n)
	str := strconv.Itoa(int(n))

	for _, r := range str {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			return false, err
		}
		res := math.Pow(float64(digit), 3)
		remainder -= res
		if remainder < 0 {
			return false, nil
		}
	}

	return remainder == 0, nil
}
