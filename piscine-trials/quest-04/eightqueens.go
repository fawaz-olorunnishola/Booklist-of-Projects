package piscinego

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(board, 0)
}

func solve(board [8]int, col int) {
	if col == 8 {
		printSolution(board)
		return
	}

	for row := 1; row <= 8; row++ {
		if isSafe(board, row, col) {
			board[col] = row
			solve(board, col+1)
		}
	}
}

func isSafe(board [8]int, row, col int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row || board[i]-i == row-col || board[i]+i == row+col {
			return false
		}
	}
	return true
}

func printSolution(board [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + '0'))
	}
	z01.PrintRune('\n')
}
