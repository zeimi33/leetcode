package main

import "fmt"

func main(){
a := [][]int {{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
fmt.Println(searchMatrix(a,5))

}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix)==0{
		return false
	}
	x :=0
	y := len(matrix[0])-1
	for x < len(matrix) && y >= 0{

		if matrix[x][y] ==target{
			return true
		}
		if matrix[x][y]> target{
			y--
		}else {
			x++
		}
	}
	return false
}
