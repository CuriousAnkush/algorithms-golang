package sliding_window

import (
	"fmt"
	"math"
)

//Given two strings s and t of lengths m and n respectively, return the minimum window
//substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".
//
//s = "ADOBECODEBANC", t = "ABC"

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	dict := make(map[byte]int)
	var startIndex, EndIndex = 0, 0
	var windowStart, windowEnd int
	minWindLenth := math.MaxInt

	// Store the character count in the dictionary
	for i := range t {
		dict[t[i]]++
	}
	dictLen := len(dict)

	for EndIndex < len(s) {
		// Desired window, we should shrink it by moving startIndex
		if dictLen == 0 {
			length := EndIndex - startIndex + 1

			if length < minWindLenth {
				minWindLenth = length
				windowStart = startIndex
				windowEnd = EndIndex
			}

			if _, found := dict[s[startIndex]]; found {
				if dict[s[startIndex]] == 0 {
					dictLen++
				}
				dict[s[startIndex]]++
			}
			startIndex++
			//	We need to expand the window until the desired window is available
		} else {
			if _, found := dict[s[EndIndex]]; found {
				dict[s[EndIndex]]--
				if dict[s[EndIndex]] == 0 {
					dictLen--
				}
				EndIndex++
				for dictLen == 0 {
					length := EndIndex - startIndex + 1

					if length < minWindLenth {
						minWindLenth = length
						windowStart = startIndex
						windowEnd = EndIndex
					}

					if _, found := dict[s[startIndex]]; found {
						if dict[s[startIndex]] == 0 {
							dictLen++
						}
						dict[s[startIndex]]++
					}
					startIndex++
				}
			} else {
				EndIndex++
			}
		}
	}

	fmt.Println(windowEnd)
	fmt.Println(windowStart)
	return s[windowStart:windowEnd]
}

func minWindow2(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	dict := make(map[byte]int)
	var startIndex, EndIndex = 0, 0
	var windowStart, windowEnd int
	minWindLenth := math.MaxInt

	// Store the character count in the dictionary
	for i := range t {
		dict[t[i]]++
	}
	dictLen := len(dict)

	for EndIndex < len(s) {
		ch := s[EndIndex]

		if _, found := dict[ch]; found {
			dict[ch]--

			if dict[ch] == 0 {
				dictLen--
			}
		}
		EndIndex++

		for dictLen == 0 {
			ch := s[startIndex]

			if _, found := dict[ch]; found {
				dict[ch]++

				if dict[ch] > 0 {
					dictLen++
				}
			}

			if (EndIndex - startIndex) < minWindLenth {
				minWindLenth = EndIndex - startIndex
				windowStart = startIndex
				windowEnd = EndIndex
			}
			startIndex++
		}
	}

	if minWindLenth == math.MaxInt {
		return ""
	}

	return s[windowStart:windowEnd]
}

// Using template
func minWindow3(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	dict := make(map[byte]int)
	var startIndex, EndIndex = 0, 0
	var windowStart, windowEnd int
	minWindLenth := math.MaxInt

	// Store the character count in the dictionary
	for i := range t {
		dict[t[i]]++
	}
	dictLen := len(dict)

	for EndIndex < len(s) {
		ch := s[EndIndex]

		if _, found := dict[ch]; found {
			dict[ch]--

			if dict[ch] == 0 {
				dictLen--
			}
		}
		EndIndex++

		for dictLen == 0 {
			ch := s[startIndex]

			if _, found := dict[ch]; found {
				dict[ch]++

				if dict[ch] > 0 {
					dictLen++
				}
			}

			if (EndIndex - startIndex) < minWindLenth {
				minWindLenth = EndIndex - startIndex
				windowStart = startIndex
				windowEnd = EndIndex
			}
			startIndex++
		}
	}

	if minWindLenth == math.MaxInt {
		return ""
	}

	return s[windowStart:windowEnd]
}
