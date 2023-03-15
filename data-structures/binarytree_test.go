package datastructures

import (
	"reflect"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	t.Run("Should build a binary tree", func(t *testing.T) {

		tree := BinaryTree{
			value: 5,
		}

		tree.Add(12)
		tree.Add(2)
		tree.Add(10)
		tree.Add(1)
		tree.Add(3)

		got := tree.DepthFirstSearchPrint([]int{})
		want := []int{1, 2, 3, 5, 10, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but wanted %v", got, want)
		}
	})
}
