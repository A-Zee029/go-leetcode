package greed

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	idx := 0
	for _, i := range s {
		if idx < len(g) && i >= g[idx] {
			res++
			idx++
		}
	}
	return res
}
