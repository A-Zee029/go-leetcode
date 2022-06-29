package hash_table

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	m := make(map[int]int)
	for i, num := range nums {
		if v, ok := m[target-num]; ok {
			res[0] = i
			res[1] = v
			break
		}
		m[num] = i
	}
	return res
}
