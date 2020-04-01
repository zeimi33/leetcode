package main

import (
	"fmt"
	"sort"
)

func main(){
	a  := []string{"time","me","bell"}
	c := minimumLengthEncoding(a)
	fmt.Println(c)
}


func minimumLengthEncoding(words []string) int {
	num := 0
	reservewords := []string{}
	for _,w := range words{
		reservewords = append(reservewords,reverseString(w))
	}
	sort.Strings(reservewords)
	index := reservewords[0]
	for _,w := range reservewords{
		if !haveCommonHead(index,w){
			num += len(index)
			num ++
		}
		index = w
	}
	return num +1 + len(reservewords[len(reservewords)-1])
}

func haveCommonHead(index string,word string)bool{
	for i:= 0 ;i < len(index);i++{
		if index[i] != word[i] {
			return false
		}
	}
	return true
}

func reverseString(s string) string {
    runes := []rune(s)
    for from, to := 0, len(runes)-1; from < to; from, to = from + 1, to - 1 {
        runes[from], runes[to] = runes[to], runes[from] 
    } 
    return string(runes)
}