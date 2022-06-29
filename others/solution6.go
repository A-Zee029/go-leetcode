package others

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	max := math.MaxInt

	update := func(cur int) {
		if abs(cur-target) < abs(max-target) {
			max = cur
		}
	}

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			update(sum)
			if sum < target {
				jn := j + 1
				for jn < k && nums[j] == nums[jn] {
					jn++
				}
				j = jn
			} else {
				kn := k - 1
				for j < kn && nums[kn] == nums[k] {
					kn--
				}
				k = kn
			}
		}
	}
	return max
}
