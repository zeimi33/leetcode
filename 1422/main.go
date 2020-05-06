package main

import "fmt"

func main(){
	fmt.Println(maxScore("00"))
}

func maxScore(s string) int {
	num1 := 0
	num0 :=0
	index := []int{}
	max1 :=0
	for _,v := range s{
		if v == '0'{
			num0++

		}else{
			num1++
		}
		index = append(index,num0)
	}
	fmt.Println(index)
	for i := range s{
		if i ==len(s)-1{
			break
		}
		fmt.Println(index[i],num1,i)
		c := index[i] + num1 - (i+1-index[i])
		fmt.Println(c)
		if c > max1 {
			max1 =c
		}
	}
	return max1
}