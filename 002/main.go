package main

import "fmt"

//在一个排序好的数组里面删除重复的元素
//Do not allocate extra space for another array, you must do this in place with constant memory.
func main() {

	count := 0
	array := []int{1, 2, 2, 2, 3, 4, 4, 5, 5, 6, 6, 7, 8, 8, 8, 9, 11, 22, 33, 33, 33, 44, 44, 55}
	index := array[0] - 1
	num := len(array)
	for i := 0; i < num; i++ {
		if index == array[count] {
			array = append(array[:count], array[count+1:]...)
		} else {
			index = array[count]
			count++
		}

	}
	fmt.Println(array)
}
