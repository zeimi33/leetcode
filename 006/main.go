package main

import "fmt"

//得到一个这样的数组 Pascal's Triangle
//[ [1],
// [1,1],
// [1,2,1],
// [1,3,3,1],
// [1,4,6,4,1] ]
func main(){

	num := 0
	fmt.Scanf("%d",&num)
	pascal:= make([][]int,num)
	for i:=0 ;i<num;i++{
		a := []int{}
		for n :=0;n<=i;n++{
			if n==0 || n==i{
				a =append(a,1)
			}else {
				a = append(a,pascal[i-1][n-1]+pascal[i-1][n])
			}
		}
		pascal[i] = a
	}
	fmt.Println(pascal)
}
