package main

import "fmt"

func main() {
	c := "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring(c))
}

func lengthOfLongestSubstring(s string) int {
	freq := make([]int, 128)
	var res = 0
	start, end := 0, -1
	for start < len(s) {
		if end+1 < len(s) && freq[s[end+1]] == 0 {
			end++
			freq[s[end]]++
		} else {
			freq[s[start]]--
			start++
		}
		res = max(res, end-start+1)
	}
	return res
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
