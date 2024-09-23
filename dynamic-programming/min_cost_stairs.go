package dynamic_programming

import "math"

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0

	for idx := 2; idx <= len(cost); idx++ {
		dp[idx] = int(math.Min(float64(dp[idx-1]+cost[idx-1]), float64(dp[idx-2]+cost[idx-2])))
	}

	return dp[len(cost)]
}
