package main

import "fmt"

func main() {
	fmt.Println("wodiu")
}

func subarraySum(nums []int, k int) int {
	ret := 0
	index := 0
	m := make(map[int]int)
	m[0] = 1
	length := len(nums)
	for i := 0; i < length; i++ {
		index += nums[i]
		if _, ok := m[index-k]; ok {
			ret += m[index-k]
		}
	}
	m[index]++
}
