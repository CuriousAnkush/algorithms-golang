package dynamic_programming

import (
	"math"
)

func maximalSquare(matrix [][]byte) int {
	maxDim := 0
	dp := make([][]int, len(matrix)+1)

	for i := 0; i <= len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				continue
			}

			dp[i][j] = 1 + min(dp[i][j+1], dp[i+1][j], dp[i+1][j+1])
			if maxDim < dp[i][j] {
				maxDim = dp[i][j]
			}
		}
	}

	return int(math.Pow(float64(maxDim), 2))
}
