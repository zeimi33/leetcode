package main

import "fmt"

func main() {
	a := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(a))
}

func trap(height []int) int {
	index := 0
	jishui := map[int]int{}
	ret := 0
	for _, a := range height {

		for i := 0; i <= a; i++ {
			if _, ok := jishui[i]; ok {
				ret += jishui[i]
				jishui[i] = 0
			}
		}
		if a < index {
			for i := 1; i <= index-a; i++ {
				jishui[i+a]++
			}
		} else {
			for i := 1; i < index; i++ {
				if _, ok := jishui[i]; ok {
					ret += jishui[i]
					jishui[i] = 0
				}
			}
			index = a
		}
	}
	return ret
}
