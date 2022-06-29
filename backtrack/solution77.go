package backtrack

func combine(n int, k int) [][]int {
	res := make([][]int, 0)

	var backtrack func(in, out, nums, times int, arr *[]int)
	backtrack = func(in, out, nums, times int, arr *[]int) {
		if out == times {
			res = append(res, *arr)
			return
		}
		for i := in; i <= nums; i++ {
			*arr = append(*arr, i)
			backtrack(i+1, out+1, nums, times, arr)
			old := *arr
			oldLen := len(old) - 1
			old = old[0:oldLen]
			arr = &old
		}

	}
	array := make([]int, 0)
	backtrack(1, 0, n, k, &array)
	return res
}
