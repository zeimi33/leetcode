package main

import (
	"fmt"
	"sort"
)

func main(){
		a := []int{2,1,5,6,2,3}
		fmt.Println(largestRectangleArea(a))
}

func largestRectangleArea(heights []int) int {
	if len(heights)==0{
		return 0
	}
	if len(heights)==1{
		return heights[0]
	}
	rep :=make([]int,len(heights))
	copy(rep,heights)
	sort.Ints(rep)

	max := heights[0]
	for i:=0;i<len(heights);i++{
		println(i)
		if i==0 {
			max =returnTheArea(heights,1,0,0,heights[0])
			continue
		}
		if i==(len(heights)-1){
			max =returnTheArea(heights,0,1,i,max)
			continue
		}
		max =returnTheArea(heights,1,1,i,max)

	}
	return max
}

func returnTheArea (array []int,right int,left int,index int,max int)int{
	mut := 1
	if left==1 {
		for i:=index-1;i>=0;i--{
			if array[i]>=array[index]{
				mut++
			}else {
				break
			}
		}
	}
	if right==1 {
		for i:=index+1;i<len(array);i++{
			if array[i]>=array[index]{
				mut++
			}else{
				break
			}
		}
	}
	if max > mut*array[index]{
		return max
	}else{
		return mut*array[index]
	}
}