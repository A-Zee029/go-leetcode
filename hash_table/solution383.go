package hash_table

func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, 26)
	for _, i := range magazine {
		m[i-'a']++
	}
	for _, i := range ransomNote {
		if m[i-'a'] == 0 {
			return false
		}
		m[i-'a']--
	}
	return true
}
