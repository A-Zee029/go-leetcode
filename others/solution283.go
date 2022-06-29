package others

func moveZeroes(nums []int) {
	n := len(nums)
	for i, j := 0, 0; j < n; j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
