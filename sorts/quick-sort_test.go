package sorts

import (
	"reflect"
	"testing"

	specs "github.com/harrisoncramer/golang-algos/sorts/specs"
)

func TestQuickSort(t *testing.T) {
	t.Run("Should sort the items", func(t *testing.T) {
		for _, test := range specs.MakeTests() {
			got := QuickSort(test.Input)
			if !reflect.DeepEqual(got, test.Want) {
				t.Errorf("Got %v but wanted %v", got, test.Want)
			}
		}
	})

}
