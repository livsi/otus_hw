package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(txt string) []string {
	wf := make(map[string]int)
	for _, w := range strings.Fields(txt) {
		if w == "" {
			continue
		}

		wf[w]++
	}

	// var wordFrs []wordFreq //<-- prealloc linter warning
	wordFrs := make([]wordFreq, 0, len(wf))
	for word, frq := range wf {
		wordFrs = append(wordFrs, wordFreq{word: word, frq: frq})
	}

	sort.Slice(wordFrs, func(i, j int) bool {
		if wordFrs[i].frq == wordFrs[j].frq {
			return wordFrs[i].word < wordFrs[j].word
		}
		return wordFrs[i].frq > wordFrs[j].frq
	})

	var top10 []string
	for i := 0; i < 10; i++ {
		if i > len(wordFrs)-1 {
			return top10
		}
		top10 = append(top10, wordFrs[i].word)
	}

	return top10
}

type wordFreq struct {
	word string
	frq  int
}
