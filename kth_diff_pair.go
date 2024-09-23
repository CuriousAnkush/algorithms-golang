package algorithms

import "math"

func smallestDistancePair(nums []int, k int) int {
	minNUm := math.MaxInt
	maxNum := math.MinInt
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
		if num < minNUm {
			minNUm = num
		}
	}

	bucket := make([]int, maxNum-minNUm+1)

	for idx := range nums {
		for j := idx + 1; j < len(nums); j++ {
			bucket[Abs(nums[j]-nums[idx])]++
		}
	}

	for diff, diffFreq := range bucket {
		k = k - diffFreq
		if k <= 0 {
			return diff
		}
	}

	return -1
}

func Abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}
