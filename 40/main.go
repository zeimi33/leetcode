package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(combinationSum2([]int{4,2,5,2,5,3,1,5,2,2},9))
}

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		 return [][]int{}
	}
	sort.Ints(candidates)
	index := candidates[0]
	mnext := make(map[int]int)
	m := make(map[int]int)
	for _, v :=  range candidates {
		if v != index{
			mnext[index] = v
			index =v
		}
		m[v]++
	}
	var dfs func(loc,num int)[][]int
	dfs = func(loc,num int)[][]int{
		ret := [][]int{}
		linshi := []int{}
		poor := target
		for i:= 0 ;i<= m[loc];i++{
			if i > 0{
				linshi = append(linshi,loc)
			}
			poor = num - i * loc
			if poor < loc{
				if poor  == 0  {
					ret = append(ret,linshi)
				}
				break
			}
			if mnext[loc] != 0 {
				r := dfs(mnext[loc],num-i*loc)
				if loc == 2 && num ==9{
					fmt.Println(num- i *loc,num,linshi,i,r)
					}
				for _,v :=range r{
					c := make([]int,len(linshi))
					copy(c,linshi)
					ret = append(ret,append(c,v...))
				}
			}
		}
		return ret
	}
	return dfs(candidates[0],target)
}