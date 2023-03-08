package datastructures

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Should add items to the stack", func(t *testing.T) {

		stack := Stack{}
		stack.Push(3)
		stack.Push(4)

		want := Stack{3, 4}

		if !reflect.DeepEqual(stack, want) {
			t.Errorf("Got %v but wanted %v", stack, want)
		}

	})

	t.Run("Should remove an item from the stack", func(t *testing.T) {
		stack := Stack{}
		stack.Push(3)
		stack.Push(2)
		stack.Push(6)
		stack.Push(3)
		stack.Push(4)
		got, err := stack.Pop()

		if err != nil {
			t.Errorf("Error when popping from stack: %v", err)
		}

		want := 4
		if got != want {
			t.Errorf("Got %d but wanted %d", got, want)
		}

		got, err = stack.Pop()

		if err != nil {
			t.Errorf("Error when popping from stack: %v", err)
		}

		want = 3
		if got != want {
			t.Errorf("Got %d but wanted %d", got, want)
		}

	})

	t.Run("Should error if no items and removing value", func(t *testing.T) {
		stack := Stack{}
		_, err := stack.Pop()

		if err == nil {
			t.Error("Did not receive error when popping from empty stack")
		}

		if err.Error() != "This stack is empty" {
			t.Errorf("Incorrect message when popping from empty stack: %s", err.Error())
		}

	})
}
