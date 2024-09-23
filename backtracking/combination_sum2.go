package backtracking

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var subset []int
	sort.Ints(candidates)
	backtrack(candidates, 0, 0, target, subset, &result)
	return result
}

func backtrack(array []int, start, sum, target int, subset []int, result *[][]int) {
	if sum == target {
		*result = append(*result, append([]int{}, subset...))
		return
	}

	if sum > target {
		return
	}

	for j := start; j < len(array); j++ {
		if j != start && array[j] == array[j-1] {
			continue
		}
		subset = append(subset, array[j])
		sum := sum + array[j]
		backtrack(array, j+1, sum, target, subset, result)
		sum = sum - array[j]
		subset = subset[:len(subset)-1]
	}
}
