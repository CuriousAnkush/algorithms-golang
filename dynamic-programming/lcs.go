package dynamic_programming

func longestCommonSubsequence(text1 string, text2 string) int {
	lcs := make([][]int, len(text1)+1)

	for j := 0; j <= len(text1); j++ {
		lcs[j] = make([]int, len(text2)+1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				lcs[i][j] = 1 + lcs[i+1][j+1]
			} else {
				lcs[i][j] = max(lcs[i][j+1], lcs[i+1][j])
			}
		}
	}

	return lcs[0][0]
}
