package greed

func canJump(nums []int) bool {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if i > max {
			return false
		}
		if i+nums[i] > max {
			max = i + nums[i]
			if max >= len(nums)-1 {
				return true
			}
		}
	}
	return true
}
