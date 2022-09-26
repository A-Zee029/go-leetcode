package greed

func maxProfit3(prices []int) int {
	low, result := prices[0], 0
	for _, price := range prices {
		if price < low {
			low = price
		}
		if result < price-low {
			result = price - low
		}
	}
	return result
}
