package main

import "fmt"

func main(){
	fmt.Println(1)
}

func reverse(x int) int {
	max := 2147483647
	min := -2147483648
	index := 1
	if x <= 0{
		index =-1
		num*=index
	}
	m := []int{}
	for x >0{
		m = append(m,x%10)
		x = x/10
	}
	for i,num := range m{
		if i==0&&num ==0{
			continue
		}
		x = x*10 + num
	}
	x*=index
	if x>max  || x<min{
		return 0
	}
	return x

}

