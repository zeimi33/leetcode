package main

import (
	"sort"
)

func main(){
	array := []int{}
	target :=0
	threeSumClosest(array,target)
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	clostest := nums[0]+nums[1]+nums[2]
	a := func(num int)int{
		if num <0{
			return  -num
		}
		return num
	}
	for i :=0 ;i<len(nums)-2;i++{//因为三个数相加 不可能循环到倒数第二个作为基数
			l:=i+1
			r :=len(nums)-1
			for;l<r; {

				threesum := nums[i]+nums[l]+nums[r]
				if a(threesum-target)<a(clostest-target){
					clostest = threesum
				}
				if (threesum>target){
					r--
				}else if threesum<target{
					l++
				}else{
					return target
				}
		}
	}
	return clostest
}