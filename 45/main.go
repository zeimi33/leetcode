package main

import "fmt"

func main(){
	fmt.Println(jump([]int{2,3,1,1,4}))
}


func jump(nums []int) int {
	if len(nums)== 1 {
		return 0
	}
	step := 0
	index := 0
	max :=0
for {
	index2:= index
	max = 0

	if index >= len(nums){
		return step
	}
	step++
	for i:=1;i<=nums[index];i++{
		if index+i +1 >=len(nums){
			return  step
		}
		if nums[index+i]+i+index > max {
			max = nums[index+i]+index+i
			index2 = i+index
		}
	}
	
	index = index2
	fmt.Println(index)
}
	return 0
}