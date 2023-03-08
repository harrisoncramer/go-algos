package datastructures

import "errors"

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, error) {
	length := len(*s)
	if length == 0 {
		return 0, errors.New("This stack is empty")
	}

	res := (*s)[length-1]
	*s = (*s)[:length-1]

	return res, nil
}
