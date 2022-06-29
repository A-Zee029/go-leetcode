package backtrack

func solveSudoku(board [][]byte) {
	var backtrack func(row, col int) bool

	rows := [9][9]bool{}
	cols := [9][9]bool{}
	zones := [9][9]bool{}

	for i, bytes := range board {
		for j, b := range bytes {
			if b != '.' {
				idx := int(b - '1')
				rows[i][idx] = true
				cols[j][idx] = true
				zones[3*(i/3)+j/3][idx] = true
			}
		}
	}

	backtrack = func(row, col int) bool {
		if row == 9 {
			return true
		}

		if col == 9 {
			return backtrack(row+1, 0)
		}

		if board[row][col] == '.' {
			for k := 1; k < 10; k++ {
				if rows[row][k-1] || cols[col][k-1] || zones[3*(row/3)+col/3][k-1] {
					continue
				}

				board[row][col] = byte('0' + k)
				rows[row][k-1] = true
				cols[col][k-1] = true
				zones[3*(row/3)+col/3][k-1] = true

				b := backtrack(row, col+1)
				if b {
					return true
				}

				board[row][col] = '.'
				rows[row][k-1] = false
				cols[col][k-1] = false
				zones[3*(row/3)+col/3][k-1] = false
			}
			return false
		} else {
			return backtrack(row, col+1)
		}
	}

	backtrack(0, 0)
}

//func main() {
//	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
//	solveSudoku(board)
//	fmt.Println(board)
//	//b := byte(9)
//	//fmt.Printf("%v %T", b, b)
//}
