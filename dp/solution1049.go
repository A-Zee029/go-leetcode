package dp

func lastStoneWeightII(stones []int) int {
	sum, target := 0, 0
	for _, stone := range stones {
		sum += stone
	}
	target = sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			if dp[j] < dp[j-stones[i]]+stones[i] {
				dp[j] = dp[j-stones[i]] + stones[i]
			}
		}
	}
	return sum - 2*dp[target]
}
