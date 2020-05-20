package main

import "fmt"

func main(){
	a := findTheLongestSubstring("eleetminicoworoep")//"leetcodeisgreat"
	fmt.Println(a)
}


func findTheLongestSubstring(s string) int {
	m := make(map[rune]int)
	pos := make(map[int]int)
	pos[0]=0
	for i, v := range []rune{'a','e','i','o','u'}{
		m[v] = 1 << i
	}
	bit := 0
	ret := 0
	for i, v := range s {
		if b,ok := m[v]; ok{
			bit ^=b
		}

		fmt.Println(pos,bit,i+1)
		if v,ok := pos[bit]; ok {
			ret = max(ret,i+1-v)
		}else{
			pos[bit] = i+1
		}
	}
	return ret
}

func max (a,b int)int{
	if a > b {
		return a
	}
	return b
}