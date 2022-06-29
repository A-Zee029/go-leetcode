package others

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	cnt := 0
	for i, item := range grid {
		for j, b := range item {
			if b == '1' {
				flood(grid, m, n, i, j)
				cnt++
			}
		}
	}
	return cnt
}

func flood(grid [][]byte, m int, n int, i int, j int) {
	if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == '1' {
		grid[i][j] = '0'
		flood(grid, m, n, i+1, j)
		flood(grid, m, n, i, j+1)
		flood(grid, m, n, i-1, j)
		flood(grid, m, n, i, j-1)
	}
}
