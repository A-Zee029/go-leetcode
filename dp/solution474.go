package dp

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		one, zero := 0, 0
		for _, i := range str {
			if i == '1' {
				one++
			} else {
				zero++
			}
		}
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				if dp[i-zero][j-one]+1 > dp[i][j] {
					dp[i][j] = dp[i-zero][j-one] + 1
				}
			}
		}
	}
	return dp[m][n]
}
