package main

import (
	"fmt"
	"sync"
)

func main() {
	a := superEggDrop(4, 2000)
	fmt.Println(a)
}

func superEggDrop(K int, N int) int {
	lock := sync.Mutex{}
	memo := [100][10000]int{}
	res := make(chan int)
	go dp(K, N, &memo, res, lock)
	return <-res
}

func dp(K int, N int, memo *[100][10000]int, res chan<- int, lock sync.Mutex) {

	if K == 1 {
		res <- N
		return
	}
	if N == 0 {
		res <- 0
		return
	}
	if a := (*memo)[K][N]; a != 0 {
		res <- a
		return
	}
	front := make(chan int)
	last := make(chan int)
	resVal := 10000000000000000
	for i := 1; i < N+1; i++ {
		go dp(K, N-i, memo, front, lock)
		go dp(K-1, i-1, memo, last, lock)
		a := <-front
		b := <-last
		if a > b {
			if resVal > a+1 {
				resVal = a + 1
			}
		} else if resVal > b+1 {
			resVal = b + 1
		}
	}
	(*memo)[K][N] = resVal
	fmt.Println(K, N, resVal)
	res <- resVal
}