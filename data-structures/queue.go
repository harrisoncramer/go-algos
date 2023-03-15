package datastructures

import (
	"errors"
)

type Queue[T any] struct {
	pointer int
	index   int
	values  map[int]T
}

func (q *Queue[T]) Init() *Queue[T] {
	q.values = make(map[int]T)
	return q
}

func (q *Queue[T]) Enqueue(item T) {
	q.values[q.index] = item
	q.index += 1
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.pointer == q.index {
		var empty T
		return empty, errors.New("This queue is empty")
	}

	value := q.values[q.pointer]
	delete(q.values, q.pointer)
	q.pointer += 1

	return value, nil
}

func (q *Queue[T]) Read() []T {
	var result []T
	for _, value := range q.values {
		result = append(result, value)
	}

	return result
}
