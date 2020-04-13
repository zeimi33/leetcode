package main

import "fmt"

func main(){
	a := []int{1,23,12,4,345,543,2,21,13,32}
	
		fmt.Println(append(a[:9],a[10:]...))
c := 1
fmt.Println(c+(4/3))
}