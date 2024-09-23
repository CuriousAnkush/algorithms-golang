package algorithms

type LongestPalindrom struct {
	s string
}

func (p LongestPalindrom) GetLength() int {
	totalLength := 0
	text := p.s
	freqCount := make(map[rune]int)

	for _, char := range text {
		freqCount[char]++
	}

	hasOdd := false

	for _, count := range freqCount {
		if count%2 == 0 {
			totalLength = totalLength + count
		} else {
			totalLength += count - 1
			hasOdd = true
		}
	}

	if hasOdd {
		totalLength++
	}

	return totalLength
}
