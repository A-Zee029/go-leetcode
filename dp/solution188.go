package dp

func maxProfit2(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2*k+1)
	}

	for i := 1; i < 2*k; i += 2 {
		dp[0][i] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j < 2*k; j += 2 {
			dp[i][j] = dp[i-1][j-1] - prices[i]
			dp[i][j+1] = dp[i-1][j] + prices[i]

			if dp[i][j] < dp[i-1][j] {
				dp[i][j] = dp[i-1][j]
			}

			if dp[i][j+1] < dp[i-1][j+1] {
				dp[i][j+1] = dp[i-1][j+1]
			}

		}
	}

	return dp[len(prices)-1][2*k]

}
