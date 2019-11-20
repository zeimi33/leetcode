package main

import "fmt"

//给定一个整型数组num，找出这个数组中满足这个条件的所有数字：
// num[i]+num[j]+num[k] = 0. 并且所有 的答案是要和其他不同的，
// 也就是说两个相同的答案是不被接受的。
func main(){
	array:=[]int{-1,-2,-3,-4,-5,0,1,2,3,4,5}
	num := len(array)
	ret := map[int]map[int]int{}
	for i:=0;i<num;i++{
		for n:=i+1;n<num;n++{
			for m:=n+i;m<num;m++{
				if array[i]+array[n]+array[m]==0{
					if ret[i]==nil{
						ret[i] = map[int]int{}
					}
					ret[i][n]=m
				}
			}
		}
	}
	fmt.Println(ret)
}
