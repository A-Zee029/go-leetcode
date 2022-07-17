package greed

func PartitionLabels(s string) []int {
	var dict [26]int
	for i, b := range s {
		dict[b-'a'] = i
	}

	res := make([]int, 0)

	max := -1
	last := 0
	for i, b := range s {
		if max == -1 {
			max = dict[b-'a']
			last = i
			continue
		}
		if i > max {
			res = append(res, i-last)
			max = dict[b-'a']
			last = i
		} else {
			if dict[b-'a'] > max {
				max = dict[b-'a']
			}
		}
	}
	res = append(res, len(s)-last)
	return res
}
