package main

import (
	"reflect"
	"testing"

	"github.com/harrisoncramer/golang-algos/sorts/specs"
)

func TestInsertionSort(t *testing.T) {
	t.Run("Should sort the items", func(t *testing.T) {
		for _, test := range specs.MakeTests() {
			got := insertionSort(test.Input)
			if !reflect.DeepEqual(got, test.Want) {
				t.Errorf("Got %v but wanted %v", got, test.Want)
			}
		}
	})

}
