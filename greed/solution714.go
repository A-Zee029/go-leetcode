package greed

func maxProfit2(prices []int, fee int) int {
	res := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}
		if prices[i] > min+fee {
			res += prices[i] - min - fee
			min = prices[i] - fee
		}
	}
	return res
}
