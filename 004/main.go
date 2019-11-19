package main

//移除重复的元素，但是可以允许最多两次重复元素存在。
import "fmt"

func main() {

	count := 0
	array := []int{1, 2, 2, 2, 3, 4, 4, 5, 5, 6, 6, 7, 8, 8, 8, 9, 11, 22, 33, 33, 33, 44, 44, 55}
	index := array[0] - 1
	num := len(array)
	countRepeat := 0
	for i := 0; i < num; i++ {
		if index == array[count] {
			if countRepeat < 1 {
				countRepeat = +1
				count++
				continue
			} else {
				array = append(array[:count], array[count+1:]...)
			}
		} else {
			index = array[count]
			countRepeat = 0
			count++
		}

	}
	fmt.Println(array)
}
