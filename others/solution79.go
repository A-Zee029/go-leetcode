package others

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(start, size int)
	backtrack = func(start, size int) {
		if size == len(path) {
			dst := make([]int, size)
			copy(dst, path)
			res = append(res, dst)
			return
		}
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i+1, size)
			path = path[:len(path)-1]
		}
	}

	res = append(res, make([]int, 0))

	for i := 1; i <= len(nums); i++ {
		backtrack(0, i)
	}
	return res
}
