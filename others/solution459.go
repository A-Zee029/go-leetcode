package others

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 1 {
		return true
	}
	for i := 1; i < n; i++ {
		if n%i != 0 {
			continue
		}
		temp := s[0:i]
		j := i
		same := true
		for ; j < n; j += i {
			if s[j:j+i] != temp {
				same = false
				break
			}
		}
		if same {
			return true
		}
	}
	return false
}
