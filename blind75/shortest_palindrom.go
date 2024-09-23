package blind75

import (
	"math"
	"strings"
)

func shortestPalindrome(s string) string {
	//	Find the tail index of longest prefix of S which is palindrome
	index := longestPrefixRollingHash(s)
	return reverseString(s[index+1:]) + s
}

func reverseString(s string) string {
	sb := strings.Builder{}

	for idx := len(s) - 1; idx > -1; idx-- {
		sb.WriteByte(s[idx])
	}

	return sb.String()
}

func longestPrefix(s string) int {
	for idx := len(s) - 1; idx >= 0; idx-- {
		if isPalindrom(s[:idx]) {
			return idx
		}
	}

	return -1
}

func isPalindrom(s string) bool {
	leftIdx, rightIDx := 0, len(s)-1

	for leftIdx < rightIDx {
		if s[leftIdx] != s[rightIDx] {
			return false
		}
		leftIdx++
		rightIDx--
	}

	return true
}

// Bases on hash

func longestPrefixRollingHash(s string) int {
	var prefixIndex int
	rollingHash, reverseRollingHash, base, mod := 0, 0, 29, 10000000007

	for idx, ch := range s {
		rollingHash = (rollingHash*base + int(ch) - 'a') % mod
		reverseRollingHash = ((int(ch)-'a')*int(math.Pow(float64(base), float64(idx))) + reverseRollingHash) % mod

		//fmt.Println(idx, rollingHash, reverseRollingHash)
		if rollingHash == reverseRollingHash {
			prefixIndex = idx
		}
	}

	return prefixIndex
}
