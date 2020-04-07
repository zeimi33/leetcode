package main

import "fmt"

func main() {
	c := maxDepthAfterSplit("(()())")
	fmt.Println(c)
}

func maxDepthAfterSplit(seq string) []int {
	ret := []int{}
	index := 0
	for _, w := range seq {

		if w == '(' {
			index++
			ret = append(ret, index%2)

		} else {
			res[i] = depth % 2
			index--
			
		}
	}
	return ret
}

