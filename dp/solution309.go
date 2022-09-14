package dp

func maxProfit3(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		if dp[i][0] < dp[i-1][2]-prices[i] {
			dp[i][0] = dp[i-1][2] - prices[i]
		}
		if dp[i][0] < dp[i-1][3]-prices[i] {
			dp[i][0] = dp[i-1][3] - prices[i]
		}

		dp[i][1] = dp[i-1][0] + prices[i]

		dp[i][2] = dp[i-1][2]
		if dp[i][2] < dp[i-1][3] {
			dp[i][2] = dp[i-1][3]
		}

		dp[i][3] = dp[i-1][1]
	}

	res := dp[len(prices)-1][1]
	if res < dp[len(prices)-1][2] {
		res = dp[len(prices)-1][2]
	}
	if res < dp[len(prices)-1][3] {
		res = dp[len(prices)-1][3]
	}
	return res

}
