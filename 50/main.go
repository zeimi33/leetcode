package main

import "fmt"

func main() {
	fmt.Println(myPow(2.1, 3))
}

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	zhishu := 0
	if n < 0 {
		n = -n
		zhishu = -1
	}
	a := 1
	index := 1.0
	if x < 0 && n%2 == 1 {
		x = -x
		index = -1
	}
	ret := 1.0
	temporary := x
	for a <= n {
		if a&n > 0 {
			fmt.Println(a, n, temporary)
			ret = ret * temporary
		}
		temporary = temporary * temporary
		a = a << 1
	}
	if zhishu == -1 {
		ret = 1.0 / ret
	}
	return ret * index
}
