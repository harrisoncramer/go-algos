package armstrongnumbers

import (
	"testing"
)

/*
Create a method which tells whether an integer is an armstrong number.
An armstrong number is a three-digit number where if you cube each of it's digits
they add up to the original number.
*/
func TestArmstrong(t *testing.T) {
	t.Run("Should return true for a valid armstrong number", func(t *testing.T) {
		int := MyInt(371)
		got, err := int.IsArmstrong()

		if err != nil {
			t.Errorf("Got error: %v", err)
		}

		want := true

		if got != want {
			t.Errorf("Got %t but wanted %t", got, want)
		}

	})

	t.Run("Should return false for an invalid armstrong number", func(t *testing.T) {
		int := MyInt(310)

		got, err := int.IsArmstrong()

		if err != nil {
			t.Errorf("Got error: %v", err)
		}

		want := false

		if got != want {
			t.Errorf("Got %t but wanted %t", got, want)
		}
	})
}
