package hash_table

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		m[n] = true
		n = getSum2(n)
	}
	return n == 1
}

func getSum2(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
