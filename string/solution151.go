package string

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	res := ""
	n := len(s)
	i := 0
	last := 0
	for i < n {
		if s[i] == ' ' {
			if len(res) > 0 {
				res = " " + res
			}
			res = s[last:i] + res
			for i+1 < n && s[i] == s[i+1] {
				i++
			}
			last = i + 1
		}
		i++
	}
	if len(res) > 0 {
		res = " " + res
	}
	res = s[last:n] + res
	return res
}
