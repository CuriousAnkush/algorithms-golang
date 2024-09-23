package algorithms

import (
	"math"
	"sort"
)

//Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
//
//
//
//Example 1:
//
//Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
//Output: [[1,6],[8,10],[15,18]]
//Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
//Example 2:
//
//Input: intervals = [[1,4],[4,5]]
//Output: [[1,5]]
//Explanation: Intervals [1,4] and [4,5] are considered overlapping.

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}

	var intervalList []Interval
	for _, interval := range intervals {
		intervalList = append(intervalList, Interval{
			start: interval[0],
			end:   interval[1],
		})
	}

	sort.Slice(intervalList, func(i, j int) bool {
		return intervalList[i].Less(intervalList[j])
	})

	var nonOverlappingIntervals = []Interval{intervalList[0]}
	for cursor := 1; cursor < len(intervalList); cursor++ {
		interval2 := intervalList[cursor]
		interval1 := nonOverlappingIntervals[len(nonOverlappingIntervals)-1]

		if interval1.IsOverlapping(interval2) {
			interval := interval1.Merge(interval2)
			nonOverlappingIntervals[len(nonOverlappingIntervals)-1] = interval
		} else {
			nonOverlappingIntervals = append(nonOverlappingIntervals, interval2)
		}
	}

	result := [][]int{}
	for i := range nonOverlappingIntervals {
		inte := nonOverlappingIntervals[i]
		result = append(result, []int{inte.start, inte.end})
	}

	return result
}

type Interval struct {
	start int
	end   int
}

func (i Interval) IsOverlapping(j Interval) bool {
	if j.start >= i.start && j.start <= i.end {
		return true
	}

	return false
}

func (i Interval) Merge(j Interval) Interval {
	start := math.Min(float64(i.start), float64(j.start))
	end := math.Max(float64(i.end), float64(j.end))

	return Interval{
		start: int(start),
		end:   int(end),
	}
}

func (i Interval) Less(j Interval) bool {
	if i.start == j.start {
		return i.end <= j.end
	} else {
		return i.start <= j.start
	}
}
