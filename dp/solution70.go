package dp

func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	first, second := 1, 1
	for i := 2; i <= n; i++ {
		tmp := first + second
		first = second
		second = tmp
	}
	return second
}
