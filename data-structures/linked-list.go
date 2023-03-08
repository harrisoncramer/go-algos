package datastructures

import "errors"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Append(v int) {

	newNode := &Node{
		val:  v,
		next: nil,
	}

	current := ll.head
	if current == nil {
		ll.head = newNode
	} else {
		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}
}

func (ll *LinkedList) Search(v int) int {
	current := ll.head

	var j int
	for current != nil {
		if current.val == v {
			return j
		}
		j++
		current = current.next
	}

	return -1
}

func (ll *LinkedList) Insert(v int, j int) error {
	if ll.head == nil {
		ll.Append(v)
		return nil
	}

	if j == 0 {
		temp := ll.head
		ll.head = &Node{
			val:  v,
			next: temp,
		}
		return nil
	}

	current := ll.head.next
	prev := ll.head

	i := 0
	for i < j && current != nil {
		prev = current
		current = current.next
		i++
	}

	if i < j {
		return errors.New("That index is out of bounds")
	}

	prev.next = &Node{
		val:  v,
		next: current,
	}

	return nil
}

func (ll *LinkedList) Delete(v int) error {
	if ll.head == nil {
		return errors.New("This list is empty")
	}

	/* Deleting head */
	if ll.head.val == v {
		ll.head = ll.head.next
		return nil
	}

	current := ll.head
	prev := ll.head

	for current.val != v && current.next != nil {
		prev = current
		current = current.next
	}

	if current.next == nil {
		/* We reached the end of the list */
		if current.val == v {
			prev.next = nil
			return nil
		} else {
			return errors.New("That value does not exist")
		}
	} else {
		/* We reached the node, rmeove  */
		prev.next = current.next
	}

	return nil
}

func (ll *LinkedList) Traverse() []int {
	current := ll.head
	results := []int{}
	for current != nil {
		results = append(results, current.val)
		current = current.next
	}

	return results
}
