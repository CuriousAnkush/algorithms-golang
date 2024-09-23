package algorithms

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	result := 0
	oldHeights := make([]int, len(heights))
	copy(oldHeights, heights)
	sort.Ints(heights)

	for idx, element := range heights {
		if element == oldHeights[idx] {
			continue
		}

		result++
	}

	return result
}

func main() {
	arr := []int{20, 28, 29, 54, 56}
	i, found := sort.Find(len(arr), func(i int) int {
		if 27 == arr[i] {
			return 0
		} else if 27 < arr[i] {
			return 1
		} else {
			return -1
		}
	})

	if found {
		fmt.Println(i)
	}

	//reverseParentheses("(u(love)i)")

}
