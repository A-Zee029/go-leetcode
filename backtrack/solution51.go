package backtrack

import "strings"

func solveNQueens(n int) [][]string {
	board := make([][]string, 0)
	for i := 0; i < n; i++ {
		board = append(board, make([]string, n))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	checkAvailable := func(i, j int) bool {
		for idx := 0; idx < i; idx++ {
			j1 := j - (i - idx)
			j2 := j + (i - idx)
			if (j1 >= 0 && j1 < n && board[idx][j1] == "Q") || (j2 >= 0 && j2 < n && board[idx][j2] == "Q") || board[idx][j] == "Q" {
				return false
			}
		}
		return true
	}

	result := make([][]string, 0)

	var backtrack func(now int)
	backtrack = func(now int) {
		if n == now {
			tmp := make([]string, n)
			for i, s := range board {
				tmp[i] = strings.Join(s, "")
			}
			result = append(result, tmp)

			return
		}

		for i := 0; i < n; i++ {
			if checkAvailable(now, i) {
				board[now][i] = "Q"
				backtrack(now + 1)
				board[now][i] = "."
			}
		}
	}

	backtrack(0)
	return result
}
