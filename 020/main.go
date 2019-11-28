package main

func missingNumber(nums []int) int {
	a := 0;
	for i:=0;i<len(nums);i++{
		a ^=nums[i]
	}
	for i:=1;i<=len(nums);i++{
		a ^=i
	}
	return a

}
