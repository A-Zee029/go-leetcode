package arrays

func minSubArrayLen(target int, nums []int) int {
	i, sum, res := 0, 0, 100000
	n := len(nums)
	for j := 0; j < n; j++ {
		sum += nums[j]
		for sum >= target {
			if j-i+1 < res {
				res = j - i + 1
			}
			sum -= nums[i]
			i++
		}
	}
	if res == 100000 {
		res = 0
	}
	return res
}
