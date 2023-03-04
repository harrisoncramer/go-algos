package main

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	input := []int{3, 1, 6, -10, 54, 3, 2}
	want := []int{-10, 1, 2, 3, 3, 6, 54}
	t.Run("Should sort the items", func(t *testing.T) {
		got := insertionSort(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but wanted %v", got, want)
		}
	})
}
