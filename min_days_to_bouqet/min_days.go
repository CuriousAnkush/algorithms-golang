package min_days_to_bouqet

import "slices"

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	maxBloomDay := slices.Max(bloomDay)
	lo := 0
	hi := maxBloomDay

	for lo+1 < hi {
		mid := lo + (hi-lo)/2

		if getMaxBouquests(mid, bloomDay, k) >= m {
			hi = mid
		} else {
			lo = mid
		}
	}

	return hi
}

func getMaxBouquests(daysElapsed int, bloomDay []int, size int) int {
	var maxNumberOfBouquests int
	var count int

	for _, day := range bloomDay {
		if day <= daysElapsed {
			count++
		} else {
			count = 0
		}

		if count == size {
			count = 0
			maxNumberOfBouquests++
		}
	}

	return maxNumberOfBouquests
}
