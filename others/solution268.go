package others

func missingNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	size := len(nums)
	return size*(size+1)/2 - sum
}
