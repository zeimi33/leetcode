package main

import "fmt"

func main(){
	fmt.Println()
}

func firstMissingPositive(nums []int) int {
	m := make(map[int]bool)
	for _,v  := range nums {
		m[v] = true
	}
	index := 0
	for {
		if !m[index]{
			return index
		}
		index++
	}
	return 0
}