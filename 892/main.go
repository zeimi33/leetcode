package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	a := [][]int{{1, 2}, {3, 4}}
	b := surfaceArea(a)
	fmt.Println(b)
}

func surfaceArea(grid [][]int) int {
	if grid == nil {
		return 0
	}
	atomic.CompareAndSwapUint32()
	num := 0
	next := 0
	length_raw := len(grid)
	high_raw := len(grid[0])

	for length := 0; length < len(grid); length++ {
		for high := 0; high < len(grid[0]); high++ {
			num += grid[length][high]
			if length > 0 {
				next += min(grid[length][high], grid[length-1][high])
			}
			if length < length_raw-1 {
				next += min(grid[length][high], grid[length+1][high])
			}
			if high > 0 {
				next += min(grid[length][high], grid[length][high-1])
			}
			if high < high_raw-1 {
				next += min(grid[length][high], grid[length][high+1])
			}
			if grid[length][high] > 1 {
				next += (grid[length][high] - 1) * 2
			}
			fmt.Println(next)
		}
	}
	fmt.Println(num, next)
	return num*6 - next

}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
