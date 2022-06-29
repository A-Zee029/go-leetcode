package hash_table

func isAnagram(s string, t string) bool {
	len1 := len(s)
	len2 := len(t)
	if len1 != len2 {
		return false
	}
	arr := make([]int, 26)
	for i := 0; i < len1; i++ {
		arr[s[i]-'a']++
	}
	for i := 0; i < len1; i++ {
		arr[t[i]-'a']--
	}
	for _, val := range arr {
		if val != 0 {
			return false
		}
	}
	return true
}
