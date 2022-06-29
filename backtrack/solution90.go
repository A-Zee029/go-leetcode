package backtrack

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		dst := make([]int, len(path))
		copy(dst, path)
		res = append(res, dst)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtrack(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	sort.Ints(nums)
	backtrack(nums, 0)
	return res
}
