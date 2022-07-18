package dp

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		if dp[i-1]+cost[i-1] < dp[i-2]+cost[i-2] {
			dp[i] = dp[i-1] + cost[i-1]
		} else {
			dp[i] = dp[i-2] + cost[i-2]
		}
	}
	return dp[len(cost)]
}
