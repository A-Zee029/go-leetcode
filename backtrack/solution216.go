package backtrack

func combinationSum3(k int, n int) [][]int {

	res := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(start, n, k int, path []int)
	backtrack = func(start, n, k int, path []int) {
		if n < 0 {
			return
		}
		if len(path) == k {
			if n == 0 {
				path2 := make([]int, len(path))
				copy(path2, path)
				res = append(res, path2)
			}
			return
		}
		for i := start; i <= 9; i++ {
			path = append(path, i)
			n -= i
			backtrack(i+1, n, k, path)
			n += i
			path = path[:len(path)-1]
		}
	}

	backtrack(1, n, k, path)
	return res
}
