package algorithms

import (
	"sort"
)

func countButtonPress(word string) int {
	var result int
	freqCount := make([]int, 26)
	for idx := range word {
		freqCount[word[idx]-'a']++
	}

	sort.Slice(freqCount, func(i, j int) bool {
		return freqCount[i] > freqCount[j]
	})

	for i := range freqCount {
		result += freqCount[i] * (i/8 + 1)
	}
	return result
}
