package dp

import "math"

func rob(nums []int) int {
	n := len(nums)
	a := make([]int, n)
	b := make([]int, n)
	for i, num := range nums {
		if i == 0 {
			a[i] = num
			b[i] = 0
			continue
		}
		a[i] = b[i-1] + num
		b[i] = int(math.Max(float64(a[i-1]), float64(b[i-1])))
	}
	return int(math.Max(float64(a[n-1]), float64(b[n-1])))
}
