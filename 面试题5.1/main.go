package main

import "fmt"

func main(){
	reversePairs([]int{5,4,3,2,1})
}
func reversePairs(nums []int) int {
	if len(nums)==0{
		return 0
	}
	c,ret := merge(nums)
	return ret
	
}


func merge (a []int)([]int,int){
	fmt.Println(a)
	if len(a) ==1{
		return a,0
	}
	ret :=0
	l := len(a)
	a1Sort,a1Count := merge(a[:l/2])
	a2Sort,a2Count :=merge(a[l/2:])
	a1l := len(a1Sort)
	a2l:= len(a2Sort)
	a3 := []int{}
	a3l := a1l+a2l
	for i:=0;i<a3l;i++{
		fmt.Println(a1Sort,a2Sort,a3)
		if len(a1Sort) == 0{
			a3 = append(a3,a2Sort...)
			break
		}
		if len(a2Sort)==0{
			a3  = append(a3,a1Sort...)
			break
		}
		if a1Sort[0] > a2Sort[0]{
			ret = ret + len(a2Sort)
			a3 = append(a3,a1Sort[0])
			a1Sort = a1Sort[1:]
		}else{
			a3 = append(a3,a2Sort[0])
			a2Sort = a2Sort[1:]
		}
	}
	return a3,ret+a1Count+a2Count
}