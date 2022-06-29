package hash_table

import "sort"

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < n-3; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				n3, n4 := nums[l], nums[r]
				sum := n1 + n2 + n3 + n4
				if sum == target {
					res = append(res, []int{n1, n2, n3, n4})
					for l < r && n3 == nums[l+1] {
						l++
					}
					for l < r && n4 == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return res
}
