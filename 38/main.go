package main

import "fmt"

func main() {
	fmt.Println(countAndSay(6))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	ret := countAndSay(n - 1)
	fmt.Println(n, ret)
	retReal := ""
	var c rune
	index := 0
	for _, v := range ret {
		if v == c {
			index++
		} else {
			if index > 0 {
				retReal = fmt.Sprintf("%s%d%c", retReal, index, c)
			}
			index = 1
			c = v
		}
	}
	retReal = fmt.Sprintf("%s%d%c", retReal, index, c)
	return retReal
}
