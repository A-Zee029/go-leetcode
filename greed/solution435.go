package greed

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})
	last := intervals[0]
	res := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < last[1] {
			res++
			if intervals[i][1] < last[1] {
				last = intervals[i]
			}
		} else {
			last = intervals[i]
		}
	}
	return res
}
