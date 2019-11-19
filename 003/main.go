package main

import "fmt"

func Sort(arr []int, ch chan int) {
	defer close(ch)
	if len(arr) <= 1 {
		if len(arr) == 1 {
			ch <- arr[0]
		}
		return
	}
	mid := len(arr) / 2
	s1 := make(chan int, mid)
	s2 := make(chan int, len(arr)-mid)
	go Sort(arr[:mid], s1)
	go Sort(arr[mid:], s2)
	Merge(s1, s2, ch)
}

func update(s chan int, ch chan int, c *int, ok *bool) {
	ch <- *c
	*c, *ok = <-s
}

func Merge(s1, s2, ch chan int) {
	v1, ok1 := <-s1
	v2, ok2 := <-s2
	for ok1 && ok2 {
		if v1 < v2 {
			ch <- v1
			v1, ok1 = <-s1
		} else {
			ch <- v2
			v2, ok2 = <-s2
		}
	}
	for ok1 {
		ch <- v1
		v1, ok1 = <-s1
	}
	for ok2 {
		ch <- v2
		v2, ok2 = <-s2
	}
}

func main() {
	a := []int{1, 23, 123, 432, 13, 42, 435, 64, 75, 456345, 24, 3, 13, 234, 35}
	arr := make(chan int, len(a))
	Sort(a, arr)
	for i := 0; i < len(a); i++ {
		num := <-arr
		fmt.Println(num)
	}
}
