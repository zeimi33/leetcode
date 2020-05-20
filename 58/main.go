package main

import (
	"fmt"
	"strings"
)

func main(){
fmt.Println(lengthOfLastWord("wo cao ni mal  le a "))
}

func lengthOfLastWord(s string) int {
	s = strings.Trim(s," ")
	sl := strings.Split(s," ")
	fmt.Println(sl)
	fmt.Println(sl[len(sl)-1])
	return  len(strings.Trim(sl[len(sl)-1]," "))
}