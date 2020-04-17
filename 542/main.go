package main

import "fmt"

func main() {
	a := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	fmt.Println(updateMatrix(a))
}

func updateMatrix(matrix [][]int) [][]int {
	step := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	zero := [][2]int{}
	xlong := len(matrix[0])
	ylong := len(matrix)
	for i := 0; i < ylong; i++ {
		for j := 0; j < xlong; j++ {
			if matrix[i][j] == 0 {
				zero = append(zero, [2]int{i, j})
			} else {
				matrix[i][j] = -1
			}
		}
	}

	for len(zero) != 0 {
		a := zero[0]
		fmt.Println(a)
		zero = zero[1:]
		for _, st := range step {
			x := a[1] + st[1]
			y := a[0] + st[0]
			//fmt.Println(x, y)
			if x >= 0 && y >= 0 && x < xlong && y < ylong && matrix[y][x] == -1 {
				matrix[y][x] = matrix[a[0]][a[1]] + 1
				zero = append(zero, [2]int{y, x})
			}

		}
		fmt.Println(matrix)

	}
	return matrix
}
