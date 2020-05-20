package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(numRescueBoats([]int{3,5,3,4},5))
}

func numRescueBoats(people []int, limit int) int {
	ret := 0
	sort.Ints(people)
	for len(people) >1{
	//	fmt.Println(people)
		ret++
		if people[len(people)-1] + people[0] <= limit{
			people = people[1:len(people)-1]
		}else{
			people = people[0:len(people)-1]
		}
	}
	if len(people) !=0{
		ret++
	}
	return ret
}