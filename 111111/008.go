package main

import "fmt"

func main(){
	test :=[]int{1,23,4,5,6,7,53,12,31,34321,312}
	a,b:=hash(test,24)
	fmt.Println(a,b)
}

func hash(a []int,c int)(int,int){
	x := len(a)
	dir := map[int]int{}
	for i:=0;i<x;i++{
			dir[a[i]]= i
			if _,ok:=dir[c-a[i]];ok{
				return dir[c-a[i]],i+1
			}
	}
	return 0,0
}