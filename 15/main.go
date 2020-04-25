package main

import (
	"fmt"
	"sort"
)

func main() {
	a := threeSum([]int{3, 0, -2, -1, 1, 2})
	fmt.Println(a)
}

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }
func threeSum(nums []int) [][]int {
	sort.Sort(SortBy(nums))
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	ret := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		fmt.Println(nums[i])
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j] > 0 {
				break
			}
			if k, ok := m[-nums[i]-nums[j]]; ok && k > j {
				ret = append(ret, []int{nums[i], nums[j], -nums[i] - nums[j]})
			}
		}
	}
	return ret
}
