package dp

func lengthOfLIS(nums []int) int {
	res := 1
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		max := dp[i]
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if max < dp[j]+1 {
					max = dp[j] + 1
				}
			}
		}
		dp[i] = max
		if max > res {
			res = max
		}
	}
	return res
}
