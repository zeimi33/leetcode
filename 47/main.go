package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	mid := []int{}
	used := make([]bool, len(nums))
	dfs(nums, &ret, used, mid)
	return ret
}

func dfs(nums []int, ret *[][]int, used []bool, mid []int) {
	if len(mid) == len(nums) {
		r := make([]int, len(mid))
		copy(r, mid)
		(*ret) = append((*ret), r)
		return
	}
	for i, v := range nums {
		if used[i] {
			continue
		}
		if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}

		used[i] = true
		mid = append(mid, v)
		dfs(nums, ret, used, mid)
		mid = mid[:len(mid)-1]
		used[i] = false
	}
}
