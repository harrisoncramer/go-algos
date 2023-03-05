package search

import (
	"testing"

	specs "github.com/harrisoncramer/golang-algos/search/specs"
)

func TestBinarySearch(t *testing.T) {
	t.Run("Should find the result", func(t *testing.T) {
		tests := specs.MakeTests()
		for _, test := range tests {
			got := BinarySearch(test.Input, test.Target)
			want := test.Want
			if got != want {
				t.Errorf("Got index %d but want %d", got, want)
			}
		}
	})
}
