package dp

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := robRange(nums, 0, len(nums)-2)
	res2 := robRange(nums, 1, len(nums)-1)
	if res > res2 {
		return res
	} else {
		return res2
	}
}

func robRange(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	dp := make([]int, len(nums))
	dp[start] = nums[start]
	dp[start+1] = nums[start]
	if nums[start+1] > nums[start] {
		dp[start+1] = nums[start+1]
	}
	for i := start + 2; i <= end; i++ {
		dp[i] = dp[i-1]
		if dp[i-2]+nums[i] > dp[i] {
			dp[i] = dp[i-2] + nums[i]
		}
	}
	return dp[end]
}
