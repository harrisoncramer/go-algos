package datastructures

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("Should add values to a linked list", func(t *testing.T) {
		ll := LinkedList{}
		ll.Append(5)
		ll.Append(6)
		ll.Append(7)

		want := []int{5, 6, 7}
		got := ll.Traverse()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but wanted %v", got, want)
		}
	})

	t.Run("Should add value at specific index to a linked list", func(t *testing.T) {
		ll := LinkedList{}
		ll.Append(5)
		ll.Append(7)
		ll.Append(8)

		ll.Insert(6, 1)

		if ll.head.val != 5 {
			t.Errorf("Got %d for head but wanted %d", ll.head.val, 5)
		}

		want := []int{5, 6, 7, 8}
		got := ll.Traverse()
		if reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but wanted %v", got, want)
		}

	})

	t.Run("Should 'insert' value at the beginning", func(t *testing.T) {
		ll := LinkedList{}
		ll.Append(5)
		ll.Append(7)
		ll.Append(8)

		ll.Insert(4, 0)

		want := []int{4, 5, 6, 7}
		got := ll.Traverse()
		if reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but wanted %v", got, want)
		}

		if ll.head.val != 4 {
			t.Errorf("Got %d for head but wanted %d", ll.head.val, 4)
		}
	})

	t.Run("Should error when adding value out of bounds", func(t *testing.T) {
		ll := LinkedList{}
		ll.Append(4)
		ll.Append(5)
		ll.Append(6)
		ll.Append(7)

		err := ll.Insert(8, 5)

		if err.Error() != "That index is out of bounds" {
			t.Error("Should have errored with out of bounds insertion")
		}

		want := []int{4, 5, 6, 7}
		got := ll.Traverse()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but wanted %v", got, want)
		}

	})

	t.Run("Should get the index of a value in the linked list", func(t *testing.T) {
		ll := LinkedList{}
		ll.Append(4)
		ll.Append(5)
		ll.Append(6)
		ll.Append(7)

		got := ll.Search(6)
		want := 2
		if got != want {
			t.Errorf("Got %d but wanted %d", got, want)
		}

	})

	t.Run("Should return -1 for a non-existent value", func(t *testing.T) {
		ll := LinkedList{}
		ll.Append(4)
		ll.Append(5)
		ll.Append(6)
		ll.Append(7)

		got := ll.Search(9)
		want := -1
		if got != want {
			t.Errorf("Got %d but wanted %d", got, want)
		}
	})

	t.Run("Should delete the head", func(t *testing.T) {
		ll := LinkedList{}
		ll.Append(4)
		ll.Append(5)
		ll.Append(6)

		ll.Delete(4)

		got := ll.Traverse()
		want := []int{5, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but wanted %v", got, want)
		}
	})

	t.Run("Should delete a value from the middle", func(t *testing.T) {
		ll := LinkedList{}
		ll.Append(4)
		ll.Append(5)
		ll.Append(6)

		ll.Delete(5)
		got := ll.Traverse()
		want := []int{4, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but wanted %v", got, want)
		}

	})

	t.Run("Should delete the tail", func(t *testing.T) {
		ll := LinkedList{}
		ll.Append(4)
		ll.Append(5)
		ll.Append(6)

		ll.Delete(6)
		got := ll.Traverse()
		want := []int{4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but wanted %v", got, want)
		}
	})

	t.Run("Should error if deleting non-existent value", func(t *testing.T) {
		ll := LinkedList{}
		ll.Append(4)
		ll.Append(5)
		ll.Append(6)

		err := ll.Delete(9)
		if err == nil {
			t.Errorf("Did not error")
		}

		got := ll.Traverse()
		want := []int{4, 5, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but wanted %v", got, want)
		}

	})
}
