package main

import (
	"fmt"
	"math"
)

func main() {
	a := superEggDrop(4, 2000)
	fmt.Println(a)
}

func superEggDrop(K int, N int) int {
	memo := make(map[string]int)
	res := dp(K, N, &memo)
	return res
}

func dp(K int, N int, memo *map[string]int) int {
	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}
	var (
		res  = math.MaxInt32
		temp = fmt.Sprintf("%d,%d", K, N)
		l    = 1
		r    = N
	)
	if (*memo)[temp] != 0 {
		return (*memo)[temp]
	}
	for l <= r {
		i := l + (r-l)>>1
		notbroken := dp(K, N-i, memo)
		broken := dp(K-1, i-1, memo)
		res = min(res, max(notbroken, broken)+1)
		if notbroken > broken {
			l = i + 1
		} else {
			r = i - 1
		}
	}
	(*memo)[temp] = res
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
