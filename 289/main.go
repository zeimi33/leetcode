package main

import "fmt"

func main() {
	a := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	fmt.Println(a[0][0], a[3][0])
	gameOfLife(a)
	fmt.Println(a)
}
func gameOfLife(board [][]int) {
	steps := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {1, 0}, {1, 1}, {1, -1}, {0, 1}, {0, -1}}
	turnone := 3
	turnzero := 4
	y := len(board)
	x := len(board[0])
	fmt.Println(y)
	alive := 0
	dead := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			alive = 0
			dead = 0
			for _, step := range steps {
				xstep := i + step[0]
				ystep := j + step[1]
				fmt.Println(xstep, ystep)
				if xstep >= 0 && xstep <= x-1 && ystep >= 0 && ystep <= y-1 {
					if board[ystep][xstep] == 1 || board[ystep][xstep] == turnzero {
						alive++
					} else {
						dead++
					}
				}
			}
			if board[j][i] == 1 {
				if alive < 2 || alive > 3 {
					board[j][i] = turnzero
				}
			} else if alive == 3 {
				board[j][i] = turnone
			}
		}
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if board[j][i] == turnone {
				board[j][i] = 1
			}
			if board[j][i] == turnzero {
				board[j][i] = 0
			}
		}
	}

}
