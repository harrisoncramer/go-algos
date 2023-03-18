package countwords

import (
	"sort"
	"strings"
)

type Word struct {
	text       string
	occurences int
}

/*
Create a slice type which implements the sort interface, meaning it has
the Len(), Swap(), and Less() methods. This type can be returned from
functions that expect a []Word type.
*/

type ByFrequency []Word

func (b ByFrequency) Len() int {
	return len(b)
}

func (b ByFrequency) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByFrequency) Less(i, j int) bool {
	if b[i].occurences == b[j].occurences {
		return b[i].text < b[j].text
	}
	return b[i].occurences < b[j].occurences
}

func (b *ByFrequency) Sort() {
	sort.Sort(ByFrequency(*b))
}

func CountWords(text string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Split(text, " ") /* Use the strings.Split method to split on spaces */

	for _, word := range words {
		count, exists := wordMap[word] /* Use two arguments to access potentially missing map values safely */
		if exists {
			wordMap[word] = count + 1
		} else {
			wordMap[word] = 1
		}
	}

	return wordMap
}

func Top5Words(wordmap map[string]int) []Word {
	set := ByFrequency{}
	for newWord, val := range wordmap {
		set = append(set, Word{
			text:       newWord,
			occurences: val,
		})
	}

	sort.Sort(set)
	return set[len(set)-5:]
}
