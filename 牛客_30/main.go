package  main

import (
	"fmt"
	"strings"
)

func main(){
	SQL := "explain select * from a"
	fmt.Println( strings.Trim(SQL[strings.Index(SQL,"explain")+7:]," "))
}

func minNumberdisappered( arr []int ) int {
	// write code here
	m := make(map[int]int)
	for _,num := range arr{
		m[num] = -1
	}
	for i:= 1 ;;i++{
		if _,ok := m[i];ok{
			continue
		}
		return i
	}
	return 0
}