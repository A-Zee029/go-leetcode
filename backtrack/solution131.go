package backtrack

func partition(s string) [][]string {

	res := make([][]string, 0)

	var backtrack func(s string, start int, path []string)
	backtrack = func(s string, start int, path []string) {
		if start == len(s) {
			dst := make([]string, len(path))
			copy(dst, path)
			res = append(res, dst)
		}
		for i := start + 1; i <= len(s); i++ {
			if !isParadise(s[start:i]) {
				continue
			}
			path = append(path, s[start:i])
			backtrack(s, i, path)
			path = path[:len(path)-1]
		}
	}
	path := make([]string, 0)
	backtrack(s, 0, path)
	return res
}

func isParadise(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
