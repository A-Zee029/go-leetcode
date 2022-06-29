package others

func getDescentPeriods(prices []int) int64 {
	n := len(prices)
	sum, pre := n, 0

	for i := 1; i < n; i++ {
		if prices[i] == prices[i-1]-1 {
			pre++
			sum += pre
		} else {
			pre = 0
		}
	}

	return int64(sum)
}
