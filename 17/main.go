package main

import (
	"fmt"
	"strconv"
)

func main() {

}

var a2 = []byte{'a', 'b', 'c'}
var a3 = []byte{'d', 'e', 'f'}
var a4 = []byte{'g', 'h', 'i'}
var a5 = []byte{'j', 'k', 'l'}
var a6 = []byte{'m', 'n', 'o'}
var a7 = []byte{'p', 'q', 'r', 's'}
var a8 = []byte{'t', 'u', 'v'}
var a9 = []byte{'w', 'x', 'y', 'z'}
var aa = [][]byte{a2, a3, a4, a5, a6, a7, a8, a9}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	if len(digits) == 1 {
		ret := []string{}
		n, _ := strconv.Atoi(string(c))
		for _, c := range aa[n-2] {
			ret = append(ret, string(c))
		}
		return ret
	}
	ret := []string{}
	retlast := letterCombinations(digits[1:])
	n, _ := strconv.Atoi(string(digits[0]))
	for _, c1 := range retlast {
		for _, c2 := range aa[n-2] {
			ret = append(ret, fmt.Sprintf("%c%s", c2, c1))
		}
	}
	return ret
}
