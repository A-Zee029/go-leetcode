package others

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	res := 0
	for i := 0; i < n; i++ {
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i+1] = max(dp[i+1], dp[j+1])
			}
		}
		dp[i+1]++
		res = max(res, dp[i+1])
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

//func main() {
//	arr := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
//	lengthOfLIS(arr)
//}
