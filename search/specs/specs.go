package search

type Test struct {
	Input  []int
	Target int
	Want   int
}

func MakeTests() []Test {
	return []Test{
		{
			[]int{2, 5, 19, 110, 153, 189, 200, 290},
			200,
			6,
		},
		{
			[]int{2, 5, 19, 110, 153, 189, 200, 290},
			2,
			0,
		},
		{
			[]int{2, 5, 19, 110, 153, 189, 200, 290},
			290,
			7,
		},
		{
			[]int{},
			80,
			-1,
		},
		{
			[]int{4, 6, 100},
			90,
			-1,
		},
		{
			[]int{2, 5, 19, 110, 153, 189, 200, 290},
			5,
			1,
		},
	}
}
