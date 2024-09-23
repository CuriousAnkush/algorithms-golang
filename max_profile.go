package algorithms

import (
	"math"
	"sort"
)

type ProfitProfile struct {
	difficulty int
	profit     int
}

type Profiles []*ProfitProfile

func (p Profiles) Len() int {
	return len(p)
}

func (p Profiles) Less(i, j int) bool {
	return p[i].difficulty <= p[j].difficulty
}

func (p Profiles) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	var maxProfit int
	var profitProfile Profiles
	seenDifficulty := make(map[int]*ProfitProfile)
	for idx, diff := range difficulty {
		if profile, found := seenDifficulty[diff]; !found {
			p := &ProfitProfile{
				difficulty: diff,
				profit:     profit[idx],
			}
			profitProfile = append(profitProfile, p)
			seenDifficulty[diff] = p
		} else {
			profile.profit = int(math.Max(float64(profile.profit), float64(profit[idx])))
		}

	}

	sort.Sort(profitProfile)

	maxP := math.MinInt
	for i, profile := range profitProfile {
		maxP = int(math.Max(float64(maxP), float64(profile.profit)))
		profitProfile[i].profit = maxP
	}

	for _, capability := range worker {
		matchingDifficultIndex := sort.Search(len(profitProfile), func(i int) bool {
			return profitProfile[i].difficulty >= capability
		})
		if len(profitProfile) == matchingDifficultIndex || profitProfile[matchingDifficultIndex].difficulty != capability {
			matchingDifficultIndex--
		}
		if matchingDifficultIndex == -1 {
			continue
		}
		maxProfit += profitProfile[matchingDifficultIndex].profit
	}

	return maxProfit
}
