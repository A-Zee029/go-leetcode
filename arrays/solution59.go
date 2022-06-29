package arrays

func generateMatrix(n int) [][]int {

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	on := n
	n--
	cur, oi, oj := 1, 0, 0
	for ; n > 0; n -= 2 {
		for j := oj; j < oj+n; j++ {
			res[oi][j] = cur
			cur++
		}
		for i := oi; i < oi+n; i++ {
			res[i][oj+n] = cur
			cur++
		}
		for j := oj + n; j > oj; j-- {
			res[oi+n][j] = cur
			cur++
		}
		for i := oi + n; i > oi; i-- {
			res[i][oj] = cur
			cur++
		}
		oi++
		oj++
	}
	if on%2 == 1 {
		res[oi][oj] = cur
	}
	return res
}
