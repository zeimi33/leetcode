package main

import "fmt"

//考虑是10进位
func main() {
	array := []int{9, 9, 9, 9, 9, 9, 9, 9}
	num := len(array)
	add := true
	for i := num - 1; i >= 0; i-- {
		if add {
			if array[i] == 9 {
				if i == 0 {
					array[0] = 0
					array = append([]int{1}, array...)
					break
				}
				array[i] = 0
				add = true
			} else {
				array[i] = array[i] + 1
				add = false
			}
		}
	}
	fmt.Println(array)
}
