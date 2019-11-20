package main

import "fmt"

func main(){
 	//Given two sorted integer arrays A and B, merge B into A as one sorted array.
 	// Note: You may assume that A has enough space (size that is greater or equal to m + n)
 	// to hold additional elements from B. The number of elements initialized in A and B are m and n respectively.
 	indexa :=0
 	indexb :=0
 	a := []int{1,2,3,4,5,5,6,7,8,9}
 	b := []int{2,3,4,6,8,9,10,11}
 	c := []int{}
 	numa := len(a)
 	numb := len(b)
 	num := len(a)+len(b)
 	for i:=0;i<num;i++{

		if a[indexa] <= b[indexb]{
			c = append(c,a[indexa])
			indexa++
		}else {
			c= append(c,b[indexb])
			indexb++
		}
		if indexa==numa{
			c = append(c,b[indexb:]...)
			break
		}
		if indexb==numb{
			c = append(c,a[indexa:]...)
			break
		}
	}
	fmt.Println(c,"len",len(c))
 }
