package datastructures

import (
	"errors"
)

type Queue struct {
	pointer int
	index   int
	values  map[int]interface{}
}

func (q *Queue) Init() *Queue {
	q.values = make(map[int]interface{})
	return q
}

func (q *Queue) Enqueue(item interface{}) {
	q.values[q.index] = item
	q.index += 1
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.pointer == q.index {
		return 0, errors.New("This queue is empty")
	}

	value := q.values[q.pointer]
	delete(q.values, q.pointer)
	q.pointer += 1

	return value, nil
}

func (q *Queue) Read() []interface{} {
	var result []interface{}
	for _, value := range q.values {
		result = append(result, value)
	}

	return result
}
