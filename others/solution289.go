package others

func gameOfLife(board [][]int) {
	dir := [][]int{{-1, 0}, {-1, 1}, {-1, -1}, {0, 1}, {0, -1}, {1, 1}, {1, 0}, {1, -1}}
	row := len(board)
	col := len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			liveCnt := 0
			for k := 0; k < 8; k++ {
				nx := dir[k][0] + i
				ny := dir[k][1] + j
				if nx < row && nx >= 0 && ny < col && ny >= 0 && abs(board[nx][ny]) == 1 {
					liveCnt++
				}
			}
			if board[i][j] == 1 && (liveCnt < 2 || liveCnt > 3) {
				board[i][j] = -1
			}
			if board[i][j] == 0 && liveCnt == 3 {
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			} else if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
