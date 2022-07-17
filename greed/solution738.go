package greed

import "strconv"

func monotoneIncreasingDigits(n int) int {
	if n < 10 {
		return n
	}
	bs := []byte(strconv.Itoa(n))
	last := 0
	for i := 1; i < len(bs); i++ {
		if bs[i] == bs[i-1] {
			continue
		} else if bs[i] > bs[i-1] {
			last = i
		} else {
			bs[last] = bs[last] - 1
			for j := last + 1; j < len(bs); j++ {
				bs[j] = '9'
			}

			res, _ := strconv.Atoi(string(bs))
			return res
		}
	}
	return n
}
