package challenges

import "testing"

type Test struct {
	input string
	want  int
}

/* Given a string s, find the length of the longest substring without repeating characters. */
func TestLongestSubstring(t *testing.T) {

	tests := []Test{
		{
			input: "abcabcbb",
			want:  3,
		},
		{
			input: "aa",
			want:  1,
		},
		{
			input: "bbbbb",
			want:  1,
		},
		{
			input: "",
			want:  0,
		},
	}

	t.Run("Should return the longest substrings", func(t *testing.T) {
		for i, test := range tests {
			got := LongestSubstring(test.input)
			want := test.want
			if got != want {
				t.Errorf("Test %d: Got %d but wanted %d", i+1, got, want)
			}
		}
	})
}
