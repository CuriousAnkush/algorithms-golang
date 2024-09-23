package dynamic_programming

import (
	"fmt"
)

//You are given two 0-indexed integer arrays nums and multipliers of size n and m respectively, where n >= m.
//
//You begin with a score of 0. You want to perform exactly m operations. On the ith operation (0-indexed) you will:
//
//Choose one integer x from either the start or the end of the array nums.
//Add multipliers[i] * x to your score.
//Note that multipliers[0] corresponds to the first operation, multipliers[1] to the second operation, and so on.
//Remove x from nums.
//Return the maximum score after performing m operations.
//
//
//
//Example 1:
//
//Input: nums = [1,2,3], multipliers = [3,2,1]
//Output: 14
//Explanation: An optimal solution is as follows:
//- Choose from the end, [1,2,3], adding 3 * 3 = 9 to the score.
//- Choose from the end, [1,2], adding 2 * 2 = 4 to the score.
//- Choose from the end, [1], adding 1 * 1 = 1 to the score.
//The total score is 9 + 4 + 1 = 14.
//Example 2:
//
//Input: nums = [-5,-3,-3,-2,7,1], multipliers = [-10,-5,3,4,6]
//Output: 102
//Explanation: An optimal solution is as follows:
//- Choose from the start, [-5,-3,-3,-2,7,1], adding -5 * -10 = 50 to the score.
//- Choose from the start, [-3,-3,-2,7,1], adding -3 * -5 = 15 to the score.
//- Choose from the start, [-3,-2,7,1], adding -3 * 3 = -9 to the score.
//- Choose from the end, [-2,7,1], adding 1 * 4 = 4 to the score.
//- Choose from the end, [-2,7], adding 7 * 6 = 42 to the score.
//The total score is 50 + 15 - 9 + 4 + 42 = 102.

func maximumScore(nums []int, multipliers []int) int {
	dp := make([][]int, len(multipliers)+1)
	for i := 0; i <= len(multipliers); i++ {
		dp[i] = make([]int, len(multipliers)+1)
	}

	for mIdx := len(multipliers) - 1; mIdx >= 0; mIdx-- {
		for left := 0; left <= mIdx; left++ {
			lSum := nums[left]*multipliers[mIdx] + dp[mIdx+1][left+1]
			rIdx := len(nums) - 1 - (mIdx - left)
			rSum := nums[rIdx]*multipliers[mIdx] + dp[mIdx+1][left]

			dp[mIdx][left] = Max(lSum, rSum)
		}
	}
	return dp[0][0]
}

func maximumScore2(nums []int, multipliers []int) int {
	memo := make(map[string]int)
	return maxScore(0, 0, nums, multipliers, memo)
}

func maxScore(mIdx, lIdx int, nums, mul []int, memo map[string]int) int {
	if mIdx == len(mul) {
		return 0
	}

	rIdx := len(nums) - 1 - (mIdx - lIdx)

	if lIdx == rIdx {
		return nums[lIdx] * mul[mIdx]
	}

	if r, ok := memo[key(mIdx, lIdx, rIdx)]; !ok {
		multiplier := mul[mIdx]
		lMul := multiplier * nums[lIdx]
		rMul := multiplier * nums[rIdx]

		memoKey := key(mIdx, lIdx, rIdx)
		memo[memoKey] = Max(lMul+maxScore(mIdx+1, lIdx+1, nums, mul, memo),
			rMul+maxScore(mIdx+1, lIdx, nums, mul, memo))
		return memo[memoKey]
	} else {
		return r
	}
}

func key(mIdx, lIdx, rIdx int) string {
	return fmt.Sprintf("m_%d_l_%d_r_%d", mIdx, lIdx, rIdx)
}

func Max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
