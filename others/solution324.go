package others

func wiggleSort(nums []int) {
	n := len(nums)
	quickSelect(nums, 0, n, n/2)
	mid := nums[n/2]
	for i, j, k := 0, 0, n-1; j <= k; {
		ni, nj, nk := (1+2*i)%(n|1), (1+2*j)%(n|1), (1+2*k)%(n|1)
		if nums[nj] > mid {
			nums[ni], nums[nj] = nums[nj], nums[ni]
			i++
			j++
		} else if nums[nj] < mid {
			nums[nk], nums[nj] = nums[nj], nums[nk]
			k--
		} else {
			j++
		}
	}
}

func quickSelect(nums []int, begin, end, n int) {
	target := nums[end-1]
	i, j := begin, begin
	for ; j < end; j++ {
		if nums[j] <= target {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	//nums[i], nums[j] = nums[j], nums[i]
	if i-1 > n {
		quickSelect(nums, begin, i-1, n)
	} else if i <= n {
		quickSelect(nums, i, end, n)
	}
}

//func main() {
//	wiggleSort([]int{3, 2, 1, 1, 3, 2})
//}
