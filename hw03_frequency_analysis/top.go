package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	// Place your code here.
	top := 10
	words := strings.Fields(text)
	wordsCount := make(map[string]int, len(words))

	for _, word := range words {
		_, ok := wordsCount[word]
		if !ok {
			wordsCount[word] = 0
		}
		wordsCount[word]++
	}

	if len(wordsCount) < top {
		top = len(wordsCount)
	}

	keys := make([]string, 0, len(wordsCount))
	for k := range wordsCount {
		keys = append(keys, k)
	}

	values := make([]int, 0, len(wordsCount))
	for _, val := range wordsCount {
		values = append(values, val)
	}

	// Sort all words lexographically
	sort.Strings(keys)

	// Get top 10 values
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	// Make slice with answer
	answer := make([]string, 0, top)

	for _, val := range values[0:top] {
		for _, k := range keys {
			if wordsCount[k] == val {
				// fmt.Println(k, wordsCount[k])
				answer = append(answer, k)
				delete(wordsCount, k) // Delete word which already added to answer
				break
			}
		}
	}

	return answer
}
