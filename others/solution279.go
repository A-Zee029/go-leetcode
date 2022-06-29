package others

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		lmin := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			if dp[i-j*j] < lmin {
				lmin = dp[i-j*j]
			}
		}
		dp[i] = lmin + 1
	}
	return dp[n]
}
