package backtrack

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	set := make(map[int]bool)
	var backtrack func(nums []int)
	backtrack = func(nums []int) {
		if len(path) == len(nums) {
			dst := make([]int, len(path))
			copy(dst, path)
			res = append(res, dst)
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !set[i-1] {
				continue
			}
			set[i] = true
			path = append(path, nums[i])
			backtrack(nums)
			path = path[:len(path)-1]
			set[i] = false
		}
	}
	backtrack(nums)
	return res
}
