package main

import "fmt"

func main() {
	fmt.Println(sumNums(100))
}

func sumNums(n int) int {
	var f func(res *int, num int) bool
	f = func(res *int, num int) bool {
		*res += num
		return num != 0 && f(res, num-1)
	}
	res := 0
	f(&res, n)
	return res
}
