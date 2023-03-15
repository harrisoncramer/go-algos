package datastructures

import (
	"reflect"
	"testing"
)

func createTree() BinaryTree {
	tree := BinaryTree{
		value: 5,
	}

	tree.Add(12)
	tree.Add(2)
	tree.Add(10)
	tree.Add(1)
	tree.Add(3)

	return tree
}

func TestBinaryTree(t *testing.T) {
	t.Run("Should print out the a binary tree in ascending order", func(t *testing.T) {

		tree := createTree()

		got := tree.DFSInOrder([]int{})
		want := []int{1, 2, 3, 5, 10, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but wanted %v", got, want)
		}
	})

	t.Run("Should print out a binary tree in descending order", func(t *testing.T) {
		tree := createTree()

		got := tree.DFSReverseOrder([]int{})
		want := []int{12, 10, 5, 3, 2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but wanted %v", got, want)
		}

	})

	t.Run("Should run a DFS search for a value", func(t *testing.T) {

		tree := createTree()

		target := 3
		got := tree.DFS(target)
		want := true

		if got != want {
			t.Errorf("Got %t but wanted %t for target %d", got, want, target)
		}

		target = 12
		got = tree.DFS(target)

		if got != want {
			t.Errorf("Got %t but wanted %t for target %d", got, want, target)
		}

		target = 5
		got = tree.DFS(target)

		if got != want {
			t.Errorf("Got %t but wanted %t for target %d", got, want, target)
		}

		target = 1000
		got = tree.DFS(target)
		want = false

		if got != want {
			t.Errorf("Got %t but wanted %t for target %d", got, want, target)
		}

	})

	t.Run("Should run a BFS search for a value", func(t *testing.T) {
		tree := createTree()

		target := 3
		got := tree.BFS(target)
		want := true

		if got != want {
			t.Errorf("Got %t but wanted %t for target %d", got, want, target)
		}

		target = 12
		got = tree.BFS(target)

		if got != want {
			t.Errorf("Got %t but wanted %t for target %d", got, want, target)
		}

		target = 5
		got = tree.BFS(target)

		if got != want {
			t.Errorf("Got %t but wanted %t for target %d", got, want, target)
		}

		target = 1000
		got = tree.BFS(target)
		want = false

		if got != want {
			t.Errorf("Got %t but wanted %t for target %d", got, want, target)
		}

	})
}
