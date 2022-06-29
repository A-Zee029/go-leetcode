package string

func reverseStr(s string, k int) string {
	i := 0
	n := len(s)
	sb := []byte(s)
	for ; i < n; i += 2 * k {
		x, y := i, i+k-1
		if y >= n {
			y = n - 1
		}
		for ; x < y; x, y = x+1, y-1 {
			sb[x], sb[y] = sb[y], sb[x]
		}
	}

	return string(sb)
}
