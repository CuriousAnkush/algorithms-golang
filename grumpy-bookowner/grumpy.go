package grumpy_bookowner

import (
	"math"
)

func MaxSatisfied(customers []int, grumpy []int, minutes int) int {
	sumSoFar := 0
	sumMax := 0
	actualSum := 0
	maxDiff := math.MinInt
	for idx := range customers {
		if idx < minutes {
			if grumpy[idx] == 0 {
				sumSoFar += customers[idx]
				actualSum += customers[idx]
			}
			sumMax += customers[idx]
			if idx == minutes-1 {
				diff := sumMax - sumSoFar
				if diff > maxDiff {
					maxDiff = diff
				}
			}
		} else {
			sumMax = sumMax - customers[idx-minutes] + customers[idx]
			if grumpy[idx-minutes] == 0 {
				sumSoFar = sumSoFar - customers[idx-minutes]
			}
			if grumpy[idx] == 0 {
				actualSum += customers[idx]
				sumSoFar += customers[idx]
			}
			diff := sumMax - sumSoFar
			if diff > maxDiff {
				maxDiff = diff
			}
		}
	}
	return maxDiff + actualSum
}
