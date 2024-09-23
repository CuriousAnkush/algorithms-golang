package algorithms

func relativeSortArray(arr1 []int, arr2 []int) []int {
	freqArray := make([]int, 1001)

	for idx := range arr1 {
		freqArray[arr1[idx]]++
	}

	cursor := 0
	for idx := range arr2 {
		el := arr2[idx]
		freq := freqArray[el]
		for i := 0; i < freq; i++ {
			arr1[cursor] = el
			cursor++
			freqArray[el]--
		}
	}
	for i := 0; i <= 1000; i++ {
		freq := freqArray[i]
		if freq == 0 {
			continue
		}
		for j := 0; j < freq; j++ {
			arr1[cursor] = i
			cursor++
		}
	}

	return arr1
}
