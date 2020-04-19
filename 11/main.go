package main

import "fmt"

func main(){
	fmt.Println(maxArea([]int{1,2}))
}

func maxArea(height []int) int {
	max := 0
	area := 0
	l :=0
	r :=len(height)-1
	for l<r{
		if height[l]>=height[r]{
			area = height[r]*(r-l)
			r--
		}else{
			area = height[l]*(r-l)
			l++
		}
		if max <  area{
			max = area
		}
	}
	return max
}
