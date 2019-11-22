package main

import "fmt"

func main(){
	a := []int{3,1,3}
	c := findMin(a)
	fmt.Println(c)
}
func findMin(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}else if len(nums)>=2&&nums[0]<nums[len(nums)-1]{
		return  nums[0]
	}else{

		a := findMin(nums[0:len(nums)/2])
		b :=findMin(nums[len(nums)/2:])
		if a<b{
			return a
		}else {
			return b
		}
	}
}