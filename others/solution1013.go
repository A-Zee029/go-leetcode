package others

func canThreePartsEqualSum(arr []int) bool {
	n := len(arr)
	right := make(map[int][]int)
	sum := 0
	for i := n - 1; i >= 0; i-- {
		sum += arr[i]
		list := right[sum]
		list = append(list, i)
		right[sum] = list
	}

	cur := 0
	for i := 0; i < n-2; i++ {
		cur += arr[i]
		list := right[cur]
		for _, j := range list {
			if j > i+1 && sum-3*cur == 0 {
				return true
			}
		}
	}
	return false
}
