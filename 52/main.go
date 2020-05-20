package main

import (
	"bytes"
	"fmt"
)

func main(){
	fmt.Println(solveNQueens(4))
}

func totalNQueens(n int) int {

}

func totalNQueens(n int) int{
	var result int
	board := make([][]byte, n)
	for i := 0; i < len(board); i++ {
		tmpRow := bytes.Repeat([]byte{'.'}, len(board))
		board[i] = tmpRow
	}

	backtrack(&board, 0, &result)

	return result
}

func backtrack(board *[][]byte, row int, result *int) {
	if row == len(*board) {
		*result++
	}

	for col := 0; col < len(*board); col++ {
		if !isValid(board, row, col) {
			continue
		}

		(*board)[row][col] = 'Q'
		backtrack(board, row+1, result)
		(*board)[row][col] = '.'
	}
}

func isValid(board *[][]byte, row, col int) bool {
	// 检查列
	for i := 0; i < len(*board); i++ {
		if (*board)[i][col] == 'Q' {
			return false
		}
	}

	// 检查右上
	for i, j := row-1, col+1; i >= 0 && j < len(*board); {
		if (*board)[i][j] == 'Q' {
			return false
		}

		i--
		j++
	}

	// 检查左上
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if (*board)[i][j] == 'Q' {
			return false
		}

		i--
		j--
	}

	return true
}