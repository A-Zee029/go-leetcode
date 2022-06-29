package backtrack

func letterCombinations(digits string) []string {
	keys := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := make([]string, 0)

	var backtrack func(digits string, path []byte)

	backtrack = func(digits string, path []byte) {
		if len(digits) == 0 && len(path) > 0 {
			res = append(res, string(path))
			return
		}
		digit := digits[0]
		s := keys[digit]
		for i := 0; i < len(s); i++ {
			path = append(path, s[i])
			backtrack(digits[1:], path)
			path = path[:len(path)-1]
		}
	}
	path := make([]byte, 0)

	backtrack(digits, path)
	return res
}
