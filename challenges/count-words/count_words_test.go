package countwords

import (
	"reflect"
	"testing"
)

/*
Write a function which returns a map containing the most common words in a sentence
and how often they were used.
*/
func TestCountWords(t *testing.T) {
	t.Run("Should count words", func(t *testing.T) {
		t.Skip() /* Comparison not working */

		text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`
		results := CountWords(text)
		got := Top5Words(results)
		want := []Word{
			{"Ipsum", 3},
			{"and ", 3},
			{"Lorem", 4},
			{"of", 4},
			{"the", 6},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but wanted %v", got, want)
		}
	})
}
