package dp

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			left := j
			if left < dp[j] {
				left = dp[j]
			}
			right := i - j
			if right < dp[i-j] {
				right = dp[i-j]
			}

			if dp[i] < left*right {
				dp[i] = left * right
			}
		}
	}
	return dp[n]
}
