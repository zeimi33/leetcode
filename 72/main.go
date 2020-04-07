package main

import "fmt"

func main() {
	a := minDistance("horse", "ros")
	fmt.Println(a)
}

func minDistance(word1 string, word2 string) int {
	i := len(word1)
	j := len(word2)
	dp := [][]int{}
	s := []int{}
	for a := 0; a <= j; a++ {
		s = append(s, a)
	}
	dp = append(dp, s)
	for a := 1; a <= i; a++ {
		sli := []int{a}
		dp = append(dp, sli)
	}
	fmt.Println(dp)
	for a, w1 := range word1 {
		for b, w2 := range word2 {

			minChange := min(dp[a][b+1], dp[a+1][b], dp[a][b])
			if w1 == w2 {
				dp[a+1] = append(dp[a+1], dp[a][b])
			} else {
				dp[a+1] = append(dp[a+1], minChange+1)
			}
		}
		fmt.Println(dp)
	}
	return dp[i][j]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c

}
