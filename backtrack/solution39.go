package backtrack

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var backtrack func(candidates, path []int, target, start int)
	backtrack = func(candidates, path []int, target, start int) {

		if target <= 0 {
			if target == 0 {
				dst := make([]int, len(path))
				copy(dst, path)
				res = append(res, dst)
			}
			return
		}
		for i := start; i < len(candidates) && candidates[i] <= target; i++ {
			path = append(path, candidates[i])
			backtrack(candidates, path, target-candidates[i], i)
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	sort.Ints(candidates)
	backtrack(candidates, path, target, 0)
	return res
}
