package main

import (
	"fmt"
	"sort"
)

func main() {

	a := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merge(a)
	b := [][]int{{0, 0}}
	fmt.Println(merge(b))
}

type sortBy [][]int

func (a sortBy) Len() int           { return len(a) }
func (a sortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortBy) Less(i, j int) bool { return a[i][0] < a[j][0] }

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Sort(sortBy(intervals))
	ret := [][]int{}
	index := intervals[0]
	for len(intervals) != 0 {
		head := intervals[0]
		if index[1] >= head[0] {
			index[1] = max(index[1], head[1])
		} else {
			ret = append(ret, index)
			index = head
		}
		intervals = intervals[1:]
	}
	ret = append(ret, index)
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
