package dp

import "math"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if math.Abs(float64(target)) > float64(sum) || (target+sum)%2 == 1 {
		return 0
	}
	target = (target + sum) / 2
	if target < 0 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target]
}
