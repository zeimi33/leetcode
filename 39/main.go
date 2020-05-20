package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 5, 8, 4}, 10))
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combinationSum1(candidates, target)
}

func combinationSum1(candidates []int, target int) [][]int {
	ret := [][]int{}
	tmp := 0
	for i, v := range candidates {
		//	fmt.Println("v:", v, "target", target)
		tmp = target
		tmpret := []int{}
		for tmp-v >= 0 {
			tmp = tmp - v
			tmpret = append(tmpret, v)
			fmt.Println(tmpret, ret)
			if tmp == 0 {
				ret = append(ret, tmpret)
				continue
			}
			tmpret1 := combinationSum(candidates[i+1:], tmp)
			if len(tmpret1) > 0 {
				for _, retv := range tmpret1 {
					fmt.Println("retv", retv)
					c := make([]int, len(tmpret))
					copy(c, tmpret)
					c = append(c, retv...)
					fmt.Println(c)
					ret = append(ret, c)
					fmt.Println(ret)
				}
			}
			//fmt.Println(ret)
		}
	}
	return ret

}
