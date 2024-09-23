package subarrays

//https://leetcode.com/problems/continuous-subarray-sum/description/

func checkSubarraySum(nums []int, k int) bool {
	prefixSum := 0
	prefixModSeen := make(map[int]int)
	prefixModSeen[0] = -1

	for i := range nums {
		prefixSum += nums[i]
		mod := prefixSum % k

		if index, found := prefixModSeen[mod]; found {
			if i-index > 1 {
				return true
			}
		} else {
			prefixModSeen[mod] = i
		}
	}
	return false
}
