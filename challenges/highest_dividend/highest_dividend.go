package highestdividend

import (
	"encoding/json"
	"errors"
)

type Stock struct {
	Ticker   string  `json:"ticker"`
	Dividend float64 `json:"dividend"`
}

func HighestDividend(s string) (string, error) {

	/* First, we unmarshal the JSON into a struct */
	var data []Stock
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		return "", errors.New("Could not unmarshal JSON")
	}

	/* Then we compare all of the values in the struct */
	var largest float64
	var nameLargest string
	for _, val := range data {
		if val.Dividend > largest {
			largest = val.Dividend
			nameLargest = val.Ticker
		}
	}

	return nameLargest, nil
}
