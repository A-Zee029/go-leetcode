package arrays

func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1
	for l <= r {
		if nums[l] == val {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		} else {
			l++
		}
	}
	return r + 1
}
