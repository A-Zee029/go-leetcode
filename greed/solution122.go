package greed

func maxProfit(prices []int) int {
	last := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			res = res + prices[i] - last
			last = prices[i]
		}
	}
	diff := prices[len(prices)-1] - last
	if diff > 0 {
		res += diff
	}
	return res
}
