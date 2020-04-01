package main

import "fmt"

func main() {
	a := []int{100, 4, 200, 1, 3, 2}
	c := longestConsecutive(a)
	fmt.Println(c)
}

func longestConsecutive(nums []int) int {
	nummap := make(map[int]int)
	for _, i := range nums {
		nummap[i] = 1
	}

	max := 1
	num := 0
	ok := false
	c := 0
	index := 0
	for i := range nummap {
		num = 1
		index = i
		for {
			index++
			c, ok = nummap[index]
			if ok {
				if c > 1 {
					num = num + c
					break
				}
				num++
			} else {
				break
			}
			fmt.Println(index, c)
		}
		nummap[i] = num
		if max < num {
			max = num
		}
	}

	return max
}
