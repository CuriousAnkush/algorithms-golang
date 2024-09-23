package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(algorithms.ReverseParentheses("(u(love)i)"))
	//fmt.Println(algorithms.ReverseParentheses("(abcd)"))
	//fmt.Println(algorithms.ReverseParentheses("(ed(et(oc))el)"))
	a := []int{2, 41, 536, 89, 999}

	bu := strings.Builder{}

	for i := range a {
		bu.WriteString(strconv.Itoa(a[i]))
		if i != len(a)-1 {
			bu.WriteString(",")
		}
	}

	fmt.Println(bu.String())

	sort.Slice(a, func(i, j int) bool {
		if freqMap[nums[i]] == freqMap[nums[j]] {
			return nums[i] > nums[j]
		} else {
			return freqMap[nums[i]] == freqMap[nums[j]]
		}
	})
}
