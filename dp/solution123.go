package dp

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 5)
	}
	/*********************
	dp[i][0]: do nothing
	dp[i][1]: first buy
	dp[i][2]: first sell
	dp[i][3]: second buy
	dp[i][4]: second sell
	**********************/
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]

	for i := 1; i < len(prices); i++ {
		//dp[i][0] = dp[i-1][0]
		dp[i][1] = dp[i-1][0] - prices[i]
		dp[i][2] = dp[i-1][1] + prices[i]
		dp[i][3] = dp[i-1][2] - prices[i]
		dp[i][4] = dp[i-1][3] + prices[i]

		if dp[i][1] < dp[i-1][1] {
			dp[i][1] = dp[i-1][1]
		}

		if dp[i][2] < dp[i-1][2] {
			dp[i][2] = dp[i-1][2]
		}

		if dp[i][3] < dp[i-1][3] {
			dp[i][3] = dp[i-1][3]
		}

		if dp[i][4] < dp[i-1][4] {
			dp[i][4] = dp[i-1][4]
		}
	}

	return dp[len(prices)-1][4]

}
