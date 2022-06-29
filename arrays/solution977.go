package arrays

func sortedSquares(nums []int) []int {
	n := len(nums)
	arr := make([]int, n)
	left, right, idx := 0, n-1, n-1
	for left <= right {
		left2, right2 := nums[left]*nums[left], nums[right]*nums[right]
		if left2 > right2 {
			arr[idx] = left2
			left++
		} else {
			arr[idx] = right2
			right--
		}
		idx--
	}
	return arr
}
