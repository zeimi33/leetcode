package main

import (
	"fmt"
	"strconv"
)

func main() {
	solveNQueens(8)
}

func solveNQueens(n int) [][]string {
	ret := make([][]string,n)
	for i:= 0 ;i < n ; i ++ {
		ret[i] = make([]string, n)
	}
	place(0,&ret)
	for i := 0 ;i <n ;i++{
		for j:=0; j<n ; j++{
			if ret[i][j] != "Q"{
				ret[i][j] = "."
			}
		}
	}
	p(&ret)
	return ret
}
func place(row int,ret *[][]string)bool{
	if row == len(*ret)-1 {
		for i := 0 ; i <len(*ret) ; i++ {
			if (*ret)[row][i] == ""{
				(*ret)[row][i] ="Q"
				return true
			}
		}
		return false
	}
	p(ret)
	for i := 0 ; i < len(*ret) ;i ++ {
		if (*ret)[row][i] != ""{
			continue
		}
		(*ret)[row][i] = "Q"
		full(row,i,len(*ret),strconv.Itoa(row),ret)
		r := place(row+1,ret)
		if r {
			return r
		}

		drop(row,i,len(*ret),strconv.Itoa(row),ret)
		continue
	}
	return false
}
func p (ret *[][]string){
	for i := 0 ; i < len(*ret) ; i++{
		fmt.Println((*ret)[i])
	}
	fmt.Println("----------------")
}
func full(i, j ,n int,c string, ret *[][]string) {
	//fmt.Println(i,j,n,ret)
	for k:= 0 ; k < n ; k ++{
		if k != i {
			if (*ret)[k][j] == ""{
			(*ret)[k][j] = c
			}
		}
		if k != j {
			if (*ret)[i][k] == ""{
			(*ret)[i][k] = c
			}
		}
	}
	//fmt.Println(i,j,n,ret)
	for k := 1 ; ; k++{
		if i-k >= 0 && j -k >= 0 {
			if (*ret)[i-k][j-k]==""{
			(*ret)[i-k][j-k] = c
			}
			continue
		}
		break
	}
	//fmt.Println(i,j,n,ret)
	for k := 1; ; k++{
		if i+k <n &&j+k <n{
			if (*ret)[i+k][j+k]==""{
			(*ret)[i+k][j+k] = c
			}
			continue
		}
		break
	}
	for k := 1 ; ; k++{
		if i+k <n && j-k >=0 {
			if (*ret)[i+k][j-k] == ""{
			(*ret)[i+k][j-k] = c
			}
			continue
		}
		break
	}
	for k :=1 ; ; k++{
		if i-k >=0 && j +k < n {
			if (*ret)[i-k][j+k] == ""{
			(*ret)[i-k][j+k] = c
			}
			continue
		}
		break
	}
	//fmt.Println(ret)
}
func drop(i, j ,n int, c string ,ret *[][]string) {
	//fmt.Println(i,j,n,ret)
	(*ret)[i][j] =""
	for k:= 0 ; k < n ; k ++{
		if k != i {
			if (*ret)[k][j] == c{
			(*ret)[k][j] = ""
			}
		}
		if k != j {
			if (*ret)[i][k] == c {
			(*ret)[i][k] = ""
			}
		}
	}
	//fmt.Println(i,j,n,ret)
	for k := 1 ; ; k++{
		if i-k >= 0 && j -k >= 0 {
			if (*ret)[i-k][j-k] ==c {
			(*ret)[i-k][j-k] = ""
			}
			continue
		}
		break
	}
	//fmt.Println(i,j,n,ret)
	for k := 1; ; k++{
		if i+k <n &&j+k <n{
			if (*ret)[i+k][j+k] == c{
			(*ret)[i+k][j+k] = ""
			}
			continue
		}
		break
	}
	for k := 1 ; ; k++{
		if i+k <n && j-k >=0 {
			if (*ret)[i+k][j-k] ==c {
			(*ret)[i+k][j-k] = ""
			}
			continue
		}
		break
	}
	for k :=1 ; ; k++{
		if i-k >=0 && j +k < n {
			if (*ret)[i-k][j+k] ==c {
			(*ret)[i-k][j+k] = ""
			}
			continue
		}
		break
	}
	//fmt.Println(ret)
}
