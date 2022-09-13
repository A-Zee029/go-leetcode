package dp

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	dp := make([]int, 10001)
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			if dp[j-nums[i]]+nums[i] > dp[j] {
				dp[j] = dp[j-nums[i]] + nums[i]
			}
		}
	}
	return dp[sum] == sum
}
