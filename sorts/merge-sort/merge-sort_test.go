package main

import (
	"reflect"
	"testing"

	"github.com/harrisoncramer/golang-algos/sorts/specs"
)

/*
As you proceed through the elements of an array you continuously
swap the target back with it's neighbor until it is in
the correct place. When it's moved, move your index up.
*/
func TestMergeSort(t *testing.T) {

	t.Run("Should merge two sorted arrays", func(t *testing.T) {
		start := []int{3, 6, 9, 10}
		start2 := []int{1, 5, 11, 12}
		want := []int{1, 3, 5, 6, 9, 10, 11, 12}
		got := Merge(start, start2)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %v but wanted %v", got, want)
		}

	})

	t.Run("Should sort the items", func(t *testing.T) {
		for _, test := range specs.MakeTests() {
			got := mergeSort(test.Input)
			if !reflect.DeepEqual(got, test.Want) {
				t.Errorf("Got %v but wanted %v", got, test.Want)
			}
		}
	})

	t.Run("Should handle an empty array", func(t *testing.T) {
		got := mergeSort([]int{})
		if !reflect.DeepEqual(got, []int{}) {
			t.Error("Could not handle empty array")
		}
	})
}
