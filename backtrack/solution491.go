package backtrack

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		if len(path) >= 2 {
			dst := make([]int, len(path))
			copy(dst, path)
			res = append(res, dst)
		}
		set := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if i > start && set[nums[i]] {
				continue
			}
			size := len(path)
			if size == 0 || nums[i] >= path[size-1] {
				set[nums[i]] = true
				path = append(path, nums[i])
				backtrack(nums, i+1)
				path = path[:size]
			}
		}
	}
	backtrack(nums, 0)
	return res
}
