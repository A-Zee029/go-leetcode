package others

func firstUniqChar(s string) int {
	counter := make(map[byte]int)
	size := len(s)
	for i := 0; i < size; i++ {
		last, ok := counter[s[i]]
		if ok {
			counter[s[i]] = last + 1
		} else {
			counter[s[i]] = 1
		}
	}
	for i := 0; i < size; i++ {
		last, ok := counter[s[i]]
		if ok && last == 1 {
			return i
		}
	}
	return -1
}
