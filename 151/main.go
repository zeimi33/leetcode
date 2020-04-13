package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("hello world!  "))
}

func reverseWords(s string) string {

	a := strings.Fields(s)
	ret := ""
	for _, w := range a {
		w = strings.TrimSpace(w)
		ret = fmt.Sprintf("%s %s", w, ret)
	}
	return ret
}
