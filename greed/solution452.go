package greed

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	res := make([][]int, 0)
	res = append(res, points[0])
	for i := 1; i < len(points); i++ {
		if points[i][0] > res[len(res)-1][1] {
			res = append(res, points[i])
		} else {
			res[len(res)-1][0] = points[i][0]
			if points[i][1] < res[len(res)-1][1] {
				res[len(res)-1][1] = points[i][1]
			}
		}
	}
	return len(res)
}
