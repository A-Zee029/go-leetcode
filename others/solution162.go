package others

func findPeakElement(nums []int) int {
	i := 1
	for ; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			break
		}
	}
	return i - 1
}
