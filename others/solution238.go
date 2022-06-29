package others

func productExceptSelf(nums []int) []int {
	size := len(nums)
	left := make([]int, size)
	right := make([]int, size)
	for i := 0; i < size; i++ {
		r := size - i - 1
		if i == 0 {
			left[i] = nums[i]
			right[r] = nums[r]
		} else {
			left[i] = left[i-1] * nums[i]
			right[r] = right[r+1] * nums[r]
		}
	}
	res := make([]int, size)
	for i := 1; i < size-1; i++ {
		res[i] = left[i-1] * right[i+1]
	}
	res[0] = right[1]
	res[size-1] = left[size-2]
	return res
}
