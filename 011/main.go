package main

import (
	"fmt"
	"sort"
)

func main (){
	a := []int{-494,-487,-471,-470,-465,-462,-447,-445,-441,-432,-429,-422,-406,-398,-397,-364,-344,-333,-328,-307,-302,-293,-291,-279,-269,-269,-268,-254,-198,-181,-134,-127,-115,-112,-96,-94,-89,-58,-58,-58,-44,-2,-1,43,89,92,100,101,106,106,110,116,143,156,168,173,192,231,248,256,281,316,321,327,346,352,353,355,358,365,371,410,413,414,447,473,473,475,476,481,491,498}
	target :=8511
	c := fourSum(a,target)
	fmt.Println(c)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int
	for i :=0;i<length-3;i++{
		if i>0&&nums[i]==nums[i-1]{
			continue
		}
		min := nums[i]+nums[i+1]+nums[i+2]+nums[i+3]
		if min >target {
			break
		}
		for j :=i+1;j<length-2;j++{
			if  j>i+1&&nums[j]==nums[j-1]{
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2]>target{
				break
			}
			l :=j+1
			r :=length-1
			for ;l<r;{
				if nums[i]+nums[j]+nums[l]+nums[r] <target{
					l++
				}else if nums[i]+nums[j]+nums[l]+nums[r] >target{
					r--
				}else{
					result = append(result,[]int{nums[i],nums[j],nums[l],nums[r]})
					l++
					r--
					for;nums[l]==nums[l-1]&&l<r;{
						l++
					}
				}
			}
		}
	}
	return result
}