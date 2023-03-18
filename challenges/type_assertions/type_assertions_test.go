package typeassertions

import (
	"reflect"
	"testing"
)

func TestGetDeveloper(t *testing.T) {
	t.Run("Should assign properties and return developer", func(t *testing.T) {

		var name interface{} = "Elliot"
		var age interface{} = 26

		got, err := GetDeveloper(name, age)

		if err != nil {
			t.Errorf("Got error: %e", err)
		}

		want := Developer{
			"Elliot",
			26,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but wanted %v", got, want)
		}

	})

	t.Run("Should return an error if passing an invalid interface", func(t *testing.T) {
		var name interface{} = 22
		var age interface{} = "Sam"

		got, err := GetDeveloper(name, age)

		if err == nil {
			t.Errorf("Did not get error")
		}

		want := Developer{}
		if got != want {
			t.Errorf("Got %v but wanted %v", got, want)
		}
	})
}
