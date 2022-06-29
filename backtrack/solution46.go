package backtrack

func permute(nums []int) [][]int {
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
			if set[nums[i]] {
				continue
			}
			set[nums[i]] = true
			path = append(path, nums[i])
			backtrack(nums)
			path = path[:len(path)-1]
			set[nums[i]] = false
		}
	}
	backtrack(nums)
	return res
}
