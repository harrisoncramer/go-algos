package datastructures

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {

	t.Run("Should enqueue elements correctly", func(t *testing.T) {
		q := new(Queue[int])
		q.Init()

		q.Enqueue(3)
		q.Enqueue(4)
		q.Enqueue(6)
		q.Enqueue(1)

		got := q.values
		want := map[int]int{0: 3, 1: 4, 2: 6, 3: 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but wanted %v", got, want)
		}
	})

	t.Run("Should dequeue elements correctly ", func(t *testing.T) {

		q := Queue[int]{}
		q.Init()
		q.Enqueue(3)
		q.Enqueue(4)
		q.Enqueue(6)

		got, err := q.Dequeue()

		if err != nil {
			t.Errorf("Errored during dequeue: %v", err)
		}

		want := 3

		if got != want {
			t.Errorf("Got %d but wanted %d", got, want)
		}

		remainder := q.Read()
		wantRemainder := []int{4, 6}

		if !reflect.DeepEqual(remainder, wantRemainder) {
			t.Errorf("Got %v but wanted %v", remainder, wantRemainder)
		}
	})

	t.Run("Should error when attempting to dequeue non-existent items", func(t *testing.T) {
		q := Queue[int]{}
		q.Init()

		q.Enqueue(1)

		q.Dequeue()
		got, err := q.Dequeue()

		if got != 0 {
			t.Errorf("Dequeue failure returned non-zero value: %d", got)
		}

		if err == nil {
			t.Errorf("Dequeue failure did not return error: %v", err)
		}

		if err.Error() != "This queue is empty" {
			t.Errorf("Dequeue failure error message incorrect: %s", err.Error())
		}

	})

}
