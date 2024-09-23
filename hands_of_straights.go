package algorithms

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	groups := make([]map[int][]int, 0, len(hand))
	sort.Ints(hand)

	for i := range hand {
		found := false
		grpIdx := -1
		for i2 := range groups {
			groupMap := groups[i2]
			if groupMap[hand[i]-1] != nil && len(groupMap[hand[i]-1]) < groupSize {
				found = true
				grpIdx = i2
			}
		}

		// If the matching group found, then insert it in the group,
		// and then change the root of the group
		if found {
			newGroup := make(map[int][]int)
			groups[grpIdx][hand[i]-1] = append(groups[grpIdx][hand[i]-1], hand[i])
			newGroup[hand[i]] = groups[grpIdx][hand[i]-1]
			groups[grpIdx] = newGroup
		} else {
			group := make(map[int][]int)
			group[hand[i]] = []int{hand[i]}
			groups = append(groups, group)
		}
	}

	for i := range groups {
		for _, values := range groups[i] {
			if len(values) < groupSize {
				return false
			}
		}
	}

	return true
}
