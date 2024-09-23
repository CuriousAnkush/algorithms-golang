package subarrays

func subarraySum2(nums []int, k int) int {
	var totalSubarrayWithCountK int
	prefixSumSeen := make(map[int]int)
	prefixSumSeen[0] = 1
	initPrefixSum := 0
	for i := range nums {
		initPrefixSum += nums[i]

		totalSubarrayWithCountK += prefixSumSeen[initPrefixSum-k]

		prefixSumSeen[initPrefixSum]++
	}
	return totalSubarrayWithCountK
}

// {1,1,1}
func subarraySum(nums []int, k int) int {
	var totalSubarrayWithCountK int
	for i := range nums {
		if nums[i] == k {
			totalSubarrayWithCountK++
		}
		runningSum := nums[i]

		for _, el := range nums[i+1:] {
			runningSum += el
			if runningSum == k {
				totalSubarrayWithCountK++
			}
		}
	}

	return totalSubarrayWithCountK
}
