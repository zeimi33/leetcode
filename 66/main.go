package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 9, 9, 9}))
}

func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			continue
		}
		digits[i] = digits[i] + 1
		return digits
	}
	a := []int{1}
	a = append(a, digits...)
	return a
}
