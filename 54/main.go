package main

import "fmt"

func main(){
	fmt.Println(spiralOrder([][]int{{ 7 },{ 9},{ 6 }}))
}

func spiralOrder(matrix [][]int) []int {
	fmt.Println(matrix)
	ret := []int{}
	if len(matrix)==0 || len(matrix[0])==0{
		return ret
	}
	ret = append(ret,matrix[0]...)
	matrix = matrix[1:]
	if len(matrix)==0 || len(matrix[0])==0{
		return ret
	}
	w := len(matrix[0])
	for i := 0 ; i < len(matrix); i++{
		ret = append(ret,matrix[i][w-1])
		matrix[i] = matrix[i][:w-1]
	}
	if len(matrix)==0 || len(matrix[0])==0{
		return ret
	}
	h := len(matrix)-1
	for i:= len(matrix[0])-1;i >=0 ;i--{
		ret =append(ret,matrix[h][i])
	}
	matrix = matrix[:len(matrix)-1]
	if len(matrix)==0 || len(matrix[0])==0{
		return ret
	}
	fmt.Println(matrix)
	for i := len(matrix)-1 ; i  >= 0 ; i--{
		ret = append(ret,matrix[i][0])
		matrix[i] = matrix[i][1:]
	}
	fmt.Println(matrix)
	if len(matrix)==0 || len(matrix[0])==0{
		return ret
	}
	return append(ret,spiralOrder(matrix)...)
}