package algorithms

import "fmt"

//Given a string array words, return an array of all characters that show up in all strings within the words (including duplicates). You may return the answer in any order.
//
//
//
//Example 1:
//
//Input: words = ["bella","label","roller"]
//Output: ["e","l","l"]
//Example 2:
//
//Input: words = ["cool","lock","cook"]
//Output: ["c","o"]

func commonChars(words []string) []string {
	var coChars []string
	freqCounts := make([]map[rune]int, len(words))

	for i, word := range words {
		freqCount := make(map[rune]int)
		for _, char := range word {
			freqCount[char]++
			freqCounts[i] = freqCount
		}
	}

	// For each character in the first word,
	// see if it exists in all the maps.
	for _, char := range words[0] {
		foundInAll := true
		for _, countsMap := range freqCounts {
			foundInAll = foundInAll && countsMap[char] > 0
			countsMap[char]--
		}
		if foundInAll {
			coChars = append(coChars, fmt.Sprintf("%c", char))
		}
	}
	return coChars
}
