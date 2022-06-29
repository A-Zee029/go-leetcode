package backtrack

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	var backtrack func(s string, start int, path []string)
	backtrack = func(s string, start int, path []string) {
		size := len(path)
		if size == 3 {
			if isValidPart(s[start:]) {
				join := strings.Join(path, ".")
				join = join + "." + s[start:]
				res = append(res, join)
			}
			return
		}
		for i := start + 1; i <= len(s) && isValidPart(s[start:i]) && len(s)-start <= 3*(4-len(path)); i++ {

			path = append(path, s[start:i])
			backtrack(s, i, path)
			path = path[:len(path)-1]
		}
	}
	path := make([]string, 0)
	backtrack(s, 0, path)
	return res
}

func isValidPart(s string) bool {

	atoi, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if atoi > 255 || atoi < 0 || len(strconv.Itoa(atoi)) != len(s) {
		return false
	}
	return true
}
