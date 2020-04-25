package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"aa", "a"}))
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	max := len(strs[0])
	sIndex := 0
	for i, s := range strs {
		if len(s) == 0 {
			return ""
		}
		if max > len(s) {
			max = len(s)
			sIndex = i
		}
	}
	fmt.Println(max, sIndex)
	ret := ""
	for i, c := range strs[sIndex] {
		for _, s := range strs {
			if s[i] != byte(c) {
				return ret
			}
		}
		ret = fmt.Sprintf("%s%c", ret, c)
	}
	return ret
}
