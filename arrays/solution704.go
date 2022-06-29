package arrays

func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1

	res := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			res = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
