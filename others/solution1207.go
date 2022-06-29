package others

func uniqueOccurrences(arr []int) bool {
	cnts := make(map[int]int)
	for _, i := range arr {
		cnt := cnts[i]
		cnt++
		cnts[i] = cnt
	}
	set := make(map[int]int)
	for _, val := range cnts {
		i := set[val]
		if i == 1 {
			return false
		}
		set[val] = 1
	}
	return true
}
