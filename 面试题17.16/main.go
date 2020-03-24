package main

import "fmt"

func main() {
	a := []int{2, 7, 9, 3, 1}
	fmt.Println(massage(a))
}

func massage(nums []int) int {
	index := len(nums)
	if index == 1 {
		return nums[0]
	}
	front := 0
	last := 0
	val := 0
	index = 0
	for index < len(nums) {
		val = nums[index] + front
		if val > last {
			front = last
			last = val
		} else {
			front = last
		}
		index++
		fmt.Println("front", front, "last", last)
	}
	if front > last {
		return front
	}
	return last
}
