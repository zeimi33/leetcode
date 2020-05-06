package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	long := len(nums)
	find := nums[0] - 1
	for i := 0; i < long; i++ {
		if nums[index] == find {
			nums = append(nums[:index], nums[index+1:]...)
		} else {
			find = nums[index]
			index++
		}
	}
	fmt.Println(nums)
	return len(nums)
}
