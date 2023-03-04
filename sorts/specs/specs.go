package specs

type Test struct {
	Input []int
	Want  []int
}

func MakeTests() []Test {
	return []Test{
		{
			[]int{3, 1, 6, -10, 54, 3, 2},
			[]int{-10, 1, 2, 3, 3, 6, 54},
		},
	}
}
