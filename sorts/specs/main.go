package sorts

type Test struct {
	Input []int
	Want  []int
}

func MakeTests() []Test {
	return []Test{
		{
			[]int{-1000, 0, 30, 1},
			[]int{-1000, 0, 1, 30},
		},
		{
			[]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{0, 0, 0},
			[]int{0, 0, 0},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{3, 1, 6, -10, 54, 3, 2},
			[]int{-10, 1, 2, 3, 3, 6, 54},
		},
	}
}
