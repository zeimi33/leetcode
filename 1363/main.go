package main

import "fmt"

func main() {
	fmt.Println(largestMultipleOfThree([]int{8, 7, 6, 1, 0}))
}
func largestMultipleOfThree(digits []int) string {
	if len(digits) == 0 {
		return ""
	}
	del := func(index *[10]int, d int) bool {
		for i := d; i < 9; i += 3 {
			if (index)[i] != 0 {
				(index)[i]--
				return true
			}
		}
		return false
	}
	index := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sum := 0
	ret := ""
	for _, i := range digits {
		index[i]++
		sum += i
	}
	if sum%3 == 1 {
		if !del(&index, 1) {
			del(&index, 2)
			del(&index, 2)
		}
	}
	if sum%3 == 2 {
		if !del(&index, 2) {
			del(&index, 1)
			del(&index, 1)
		}
	}
	for i := 9; i >= 0; i-- {
		for index[i] > 0 {
			ret = fmt.Sprintf("%s%d", ret, i)
			index[i]--
		}
	}
	if len(ret) > 0 && ret[0] == '0' {
		return "0"
	}
	return ret
}
