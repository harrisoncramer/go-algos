package needlehaystack

import "testing"

/* Given two strings needle and haystack,
return the index of the first occurrence of needle in haystack,
or -1 if needle is not part of haystack. */

type Test struct {
	needle   string
	haystack string
	want     int
}

func TestNeedleHaystacks(t *testing.T) {
	tests := []Test{
		{
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			haystack: "",
			needle:   "",
			want:     -1,
		},
		{
			haystack: "",
			needle:   "ahsdf",
			want:     -1,
		},
		{
			haystack: "abcdef",
			needle:   "abcdefg",
			want:     -1,
		},
	}

	t.Run("NeedleHaystack", func(t *testing.T) {
		for i, test := range tests {
			got := NeedleHaystack(test.needle, test.haystack)
			want := test.want

			if got != want {
				t.Errorf("Test %d: Got %d but wanted %d", i, got, want)
			}
		}
	})

	t.Run("NeedleHaystackTwo", func(t *testing.T) {
		for i, test := range tests {
			got := NeedleHaystackTwo(test.needle, test.haystack)
			want := test.want

			if got != want {
				t.Errorf("Test %d: Got %d but wanted %d", i, got, want)
			}
		}
	})

	t.Run("NeedleHaystackThree", func(t *testing.T) {
		for i, test := range tests {
			got := NeedleHaystackThree(test.needle, test.haystack)
			want := test.want

			if got != want {
				t.Errorf("Test %d: Got %d but wanted %d", i, got, want)
			}
		}
	})

}
