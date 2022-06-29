package hash_table

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	res := make([]int, 0)
	for _, num := range nums1 {
		if _, ok := set[num]; !ok {
			set[num] = struct{}{}
		}
	}
	for _, num := range nums2 {
		if _, ok := set[num]; ok {
			res = append(res, num)
			delete(set, num)
		}
	}
	return res
}
