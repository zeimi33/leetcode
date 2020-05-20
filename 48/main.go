package main

import "fmt"

func main() {
	a := [][]int{{5, 1, 9, 11}, {2, 4, 8, 1}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(a)
	fmt.Println(a)
}

func rotate(matrix [][]int) {
	width := len(matrix)
	k := 0
	for i := 0; i <= width/2; i++ {
		length1 := width - 2*i
		k = 0
		length := width - i - 1
		fmt.Println(i)
		for j := i; k < length1-1; j++ {
			fmt.Println(i, j, k)
			matrix[i][j], matrix[length-k][i], matrix[length][length-k], matrix[j][length] = matrix[length-k][i], matrix[length][length-k], matrix[j][length], matrix[i][j]
			k++
		}
	}
}
