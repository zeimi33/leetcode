package main

import (
	"fmt"
	"sync"
)

type student struct {
	ch  chan int
	run bool
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("len", len(a), "cap", cap(a))
	b := a[:0]
	fmt.Println("len", len(b), "cap", cap(b))
	s := student{}
	w := sync.WaitGroup{}
	s.run = true
	go s.a()
	w.Add(1)
	s.ch = make(chan int)
	for i := 0; i < 10000; i++ {
		s.ch <- i
	}
}

func (stu *student) a() {
	var ok bool
	ok = true
	num := 0
	for stu.run && ok {
		num, ok = <-stu.ch
		fmt.Println(num)
	}
	stu.run = false
}
