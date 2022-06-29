package others

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				jn := j + 1
				for jn < k && nums[jn] == nums[j] {
					jn++
				}
				j = jn

				kn := k - 1
				for j < kn && nums[kn] == nums[k] {
					kn--
				}
				k = kn
			} else if sum < 0 {
				jn := j + 1
				for jn < k && nums[jn] == nums[j] {
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
	return res
}

//func main() {
//	sum := threeSum([]int{-1, 0, 1, 2, -1, -4})
//	fmt.Println(sum)
//}
