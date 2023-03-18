package highestdividend

import (
	"testing"
)

/* Write a function that takes in JSON stock data and returns the stock name with the highest dividend */
func TestHighestDividend(t *testing.T) {

	t.Run("Should return the highest dividend", func(t *testing.T) {
		got, err := HighestDividend(`[{"ticker":"APPL","dividend":0.5},{"ticker":"MSFT","dividend":0.2}]`)

		if err != nil {
			t.Errorf("Returned error while parsing highest dividend: %s", err)
		}

		want := "APPL"

		if got != want {
			t.Errorf("Got %s but wanted %s", got, want)
		}
	})

	t.Run("Should error with invalid JSON", func(t *testing.T) {
		res, err := HighestDividend("{{{")

		if err == nil {
			t.Errorf("Did not error with invalid JSON")
		}

		if res != "" {
			t.Errorf("Expected zero value but got %s", res)
		}
	})

	t.Run("Should return first value if all are the same", func(t *testing.T) {
		got, err := HighestDividend(`[{"ticker":"APPL","dividend":0.2},{"ticker":"MSFT","dividend":0.2}]`)

		if err != nil {
			t.Errorf("Returned error while parsing highest dividend: %s", err)
		}

		want := "APPL"

		if got != want {
			t.Errorf("Got %s but wanted %s", got, want)
		}
	})

	t.Run("Should handle empty ticker", func(t *testing.T) {
		got, err := HighestDividend(`[]`)

		if err != nil {
			t.Errorf("Returned error while parsing highest dividend: %s", err)
		}

		want := ""

		if got != want {
			t.Errorf("Got %s but wanted %s", got, want)
		}
	})
}
