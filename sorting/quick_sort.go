package sorting

func sortArray(nums []int) []int {
	quickSortRec(nums, 0, len(nums)-1)
	return nums
}

func quickSortRec(list []int, start, end int) {
	if start > end {
		return
	}
	newPivotIdx := placePivot(list, start, end)

	if newPivotIdx != start {
		quickSortRec(list, start, newPivotIdx-1)
	}
	quickSortRec(list, newPivotIdx+1, end)
}

func placePivot(list []int, low, high int) int {
	pivot := list[high]

	leftIdx := low
	rightIdx := len(list) - 2

	for leftIdx <= rightIdx {
		for list[leftIdx] < pivot && leftIdx <= rightIdx {
			leftIdx++
		}
		for rightIdx >= low && leftIdx <= rightIdx && list[rightIdx] >= pivot {
			rightIdx--
		}
		if leftIdx <= rightIdx {
			swap(list, leftIdx, rightIdx)
		}

	}
	swap(list, high, rightIdx+1)
	return rightIdx + 1
}

func swap(list []int, fIdx, Lidx int) {
	list[fIdx], list[Lidx] = list[Lidx], list[fIdx]
}
