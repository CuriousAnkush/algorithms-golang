package subarrays

func subarraysDivByK(nums []int, k int) int {
	prefixSum := 0
	prefixModSeen := make(map[int]int)
	totalCount := 0

	for i := range nums {
		if nums[i]%k == 0 && i != 0 {
			totalCount++
		}
		prefixSum += nums[i]
		mod := prefixSum % k
		prefixModSeen[mod]++
	}

	for i, i2 := range prefixModSeen {
		if i2 > 1 {
			totalCount += prefixModSeen[i]
		}
		if i == 0 {
			totalCount++
		}
	}
	return totalCount
}
