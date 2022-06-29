package greed

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	for i := 0; i < len(nums) && k > 0; i++ {
		if nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	if k%2 == 1 {
		nums[len(nums)-1] *= -1
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
