package greed

func jump(nums []int) int {
	cur := 0
	next := 0
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > next {
			next = nums[i] + i
		}
		if i == cur {
			cur = next
			res++
		}
	}
	return res
}
