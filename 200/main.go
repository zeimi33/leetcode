package main

import "fmt"

func main() {
a:
	fmt.Println()
}

func numIslands(grid [][]byte) int {
	xlong := len(grid[0])
	ylong := len(grid)
	num := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '1' {
				num++
				modify(x, y, &grid, xlong, ylong)
			}
		}
	}
	return num
}

func modify(x, y int, grid *[][]byte, xlong, ylong int) {
	steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, step := range steps {
		if x+step[0] < xlong && y+step[1] < ylong && x+step[0] >= 0 && y+step[1] >= 0 {
			if (*grid)[y+step[1]][x+step[0]] == '1' {
				(*grid)[y+step[1]][x+step[0]] = '2'
				modify(x+step[0], y+step[1], grid, xlong, ylong)
			}
		}
	}
}
