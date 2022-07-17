package greed

import "sort"

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	last := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > last[1] {
			res = append(res, last)
			last = intervals[i]
		} else {
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		}
	}
	res = append(res, last)
	return res
}
