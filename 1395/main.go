package main

import "fmt"

func main(){
	fmt.Println(numTeams([]int{2,1,3}))
}

func numTeams(rating []int) int {
	ret :=0
    for i :=1 ;i < len(rating)-1 ;i++{
		lbig :=0
		lsmall := 0
		rbig :=0
		rsmall := 0
		for j :=0 ;j < i;j++{
			if rating[j] >rating[i]{
				lbig++
				continue
			}
			lsmall++
		}
		for k:=i+1; k < len(rating);k++{
			if rating[k] >rating[i]{
				rbig++
				continue
			}
			rsmall++
		}
		ret = ret + rbig * lsmall +rsmall *lbig
	}
	return ret
}